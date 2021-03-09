package report

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

type testHandler struct {
	name     string
	before   func(options *Options) bool
	after    func(options *Options) bool
	validate func(data map[string]interface{}) bool
	times    int
}

var wg sync.WaitGroup

func mockHttpServer(t *testing.T, tt testHandler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		// Test request url
		if req.URL.String() != "/some/path" {
			t.Errorf("Get error path %v, wanted %v.", req.URL.String(), "/some/path")
		}
		// Test request parameters
		all, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Errorf("Get body err: %s", err.Error())
		}
		var arr []map[string]interface{}
		err = json.Unmarshal(all, &arr)
		if err != nil {
			t.Errorf("Get body struct err: %s", err.Error())
		}
		if tt.validate != nil && !tt.validate(arr[0]) {
			t.Errorf("Validate false.")
		}
		// Send response to be tested
		_, _ = rw.Write([]byte(`OK`))
	}
}

func testHandlerCases(t *testing.T, tests []testHandler) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg.Add(tt.times)
			server := httptest.NewServer(mockHttpServer(t, tt))
			// Close the server when test finishes
			options := Configure().
				WithClient(server.Client()).
				WithUrl(server.URL + "/some/path").
				WithAfter(func(point Point, success bool) {
					wg.Done()
				})
			if tt.before != nil && !tt.before(options) {
				t.Errorf("Validate before failed.")
			}
			options.Start()
			if tt.after != nil && !tt.after(options) {
				t.Errorf("Validate after failed.")
			}
			wg.Wait()
		})
	}
}

func TestOptions_AddReporter(t *testing.T) {
	tests := []testHandler{
		{
			name: "Test AddReporter Succeed",
			before: func(options *Options) bool {
				options.AddReporter(func() Point {
					return Point{}
				}, time.Second*2)
				return true
			},
			times: 1,
		},
	}
	testHandlerCases(t, tests)
}

func TestOptions_Send(t *testing.T) {
	tests := []testHandler{
		{
			name: "Test SendReporter Succeed",
			after: func(options *Options) bool {
				options.SendReporter(Reporter{
					Point: func() Point {
						return Point{Value: 64}
					},
					Interval: time.Second,
				})
				return true
			},
			validate: func(data map[string]interface{}) bool {
				return data["value"].(float64) == 64
			},
			times: 2,
		},
		{
			name: "Test SendReporterPayLoad Succeed",
			after: func(options *Options) bool {
				options.SendReporterPayLoad(
					func() Point {
						return Point{Value: 64}
					},
					time.Second,
				)
				return true
			},
			validate: func(data map[string]interface{}) bool {
				return data["value"].(float64) == 64
			},
			times: 1,
		},
		{
			name: "Test Send Succeed",
			after: func(options *Options) bool {
				options.Send(Point{Value: 64})
				return true
			},
			validate: func(data map[string]interface{}) bool {
				return data["value"].(float64) == 64
			},
			times: 1,
		},
		{
			name: "Test SendPayLoad Succeed",
			after: func(options *Options) bool {
				options.SendPayLoad(Fields{
					"a": "!",
					"b": "a",
				}, Tags{}, 1)
				return true
			},
			validate: func(data map[string]interface{}) bool {
				fields := data["fields"].(map[string]interface{})
				return data["value"].(float64) == 1 && fields["a"] == "!" && fields["b"] == "a"
			},
			times: 1,
		},
		{
			name: "Test SendPayLoadWithPoint Succeed",
			after: func(options *Options) bool {
				options.SendPayLoadWithPoint(Point{
					Name:     "1",
					Endpoint: "2",
					Value:    1,
					Step:     2,
					Fields:   Fields{"b": "a"},
					Tags:     Tags{"a": "1"},
					Time:     3,
				})
				return true
			},
			validate: func(data map[string]interface{}) bool {
				return data["endpoint"].(string) == "2" && data["fields"].(string) == "b=a" &&
					data["metric"].(string) == "1" && data["step"].(float64) == 2 &&
					data["tags"].(string) == "a=1" && data["time_stamp"].(float64) == 3
			},
			times: 1,
		},
	}
	testHandlerCases(t, tests)
}
