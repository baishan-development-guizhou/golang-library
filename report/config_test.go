package report

import (
	"context"
	"fmt"
	"github.com/baishan-development-guizhou/golang-library/log"
	"net/http"
	"reflect"
	"testing"
)

type test struct {
	name string
	init *Options
	args *Options
	want *Options
	test func(this test) bool
}

func testCasesWithCondition(t *testing.T, tests []test, condition func(test) bool) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if condition != nil && !condition(tt) {
				t.Errorf("condittion error, want %v", tt.want)
			}
			if tt.test != nil && !tt.test(tt) {
				t.Errorf("condittion error, want %v", tt.want)
			}
		})
	}
}

func testCase(t *testing.T, tests []test) {
	testCasesWithCondition(t, tests, nil)
}

func TestOptions_Configure(t *testing.T) {
	tests := []test{
		{
			name: "Test Configure Succeed.",
			want: &Options{
				url: "http://127.0.0.1:10699/v2/push",
			},
			test: func(this test) bool {
				return this.want.url == Configure().url
			},
		},
	}
	testCase(t, tests)
}

func TestOptions_WithConfigure(t *testing.T) {
	tests := []test{
		{
			name: "Test WithConfigure Succeed",
			args: &Options{
				url: "test",
			},
			want: &Options{
				url: "test",
			},
			test: func(this test) bool {
				return reflect.DeepEqual(this.want, WithConfigure(this.args))
			},
		},
	}
	testCase(t, tests)
}

func TestOptions_WithConfigureSelf(t *testing.T) {
	tests := []test{
		{
			name: "Test WithConfigureSelf Succeed",
			init: &Options{},
			args: &Options{
				url: "test",
			},
			want: &Options{
				url: "test",
			},
			test: func(this test) bool {
				return reflect.DeepEqual(this.want, this.init.WithConfigure(*this.args))
			},
		},
	}
	testCase(t, tests)
}

func TestOptions_WithParameter(t *testing.T) {
	defaultValue := Point{}
	beforeFunc := func(point Point) Point {
		return point
	}
	afterFunc := func(point Point, success bool) {
		fmt.Println(true)
	}
	c := context.Background()
	client := http.Client{}
	errorFunc := func(point Point, status int) {
		fmt.Println(true)
	}
	logInstance := log.Configure().Init()
	tests := []test{
		{
			name: "Test WithUrl Succeed.",
			init: &Options{},
			args: &Options{url: "Test Url"},
			want: &Options{url: "Test Url"},
			test: func(this test) bool {
				return reflect.DeepEqual(this.init.WithUrl(this.args.url), this.want)
			},
		},
		{
			name: "Test WithDefaultValue Succeed.",
			init: &Options{},
			args: &Options{defaultValue: &defaultValue},
			want: &Options{defaultValue: &defaultValue},
			test: func(this test) bool {
				return reflect.DeepEqual(this.init.WithDefaultValue(this.args.defaultValue), this.want)
			},
		},
		{
			name: "Test WithBefore Succeed.",
			init: &Options{},
			args: &Options{before: beforeFunc},
			want: &Options{before: beforeFunc},
			test: func(this test) bool {
				got := this.init.WithBefore(this.args.before)
				return reflect.ValueOf(got.before) == reflect.ValueOf(this.want.before)
			},
		},
		{
			name: "Test WithAfter Succeed.",
			init: &Options{},
			args: &Options{after: afterFunc},
			want: &Options{after: afterFunc},
			test: func(this test) bool {
				got := this.init.WithAfter(this.args.after)
				// For the same parameter, the result is the same
				return reflect.ValueOf(got.after) == reflect.ValueOf(this.want.after)
			},
		},
		{
			name: "Test WithContext Succeed.",
			init: &Options{},
			args: &Options{context: c},
			want: &Options{context: c},
			test: func(this test) bool {
				got := this.init.WithContext(this.args.context)
				return reflect.ValueOf(got.context) == reflect.ValueOf(this.want.context)
			},
		},
		{
			name: "Test WithLog Succeed.",
			init: &Options{},
			args: &Options{log: logInstance},
			want: &Options{log: logInstance},
			test: func(this test) bool {
				got := this.init.WithLog(this.args.log)
				return reflect.ValueOf(got.log) == reflect.ValueOf(this.want.log)
			},
		},
		{
			name: "Test WithClient Succeed.",
			init: &Options{},
			args: &Options{client: &client},
			want: &Options{client: &client},
			test: func(this test) bool {
				return reflect.DeepEqual(this.init.WithClient(this.args.client), this.want)
			},
		},
		{
			name: "Test WithErrorHandler Succeed.",
			init: &Options{},
			args: &Options{errorHandler: errorFunc},
			want: &Options{errorHandler: errorFunc},
			test: func(this test) bool {
				got := this.init.WithErrorHandler(this.args.errorHandler)
				return reflect.ValueOf(got.errorHandler) == reflect.ValueOf(this.want.errorHandler)
			},
		},
	}
	testCase(t, tests)
}
