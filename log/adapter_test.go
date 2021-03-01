package log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZapLoggerAdapter_With(t *testing.T) {
	logger := G()
	withLogger := logger.With("123", "123")
	assert.NotNil(t, withLogger)
	assert.NotEqual(t, fmt.Sprintf("%v", logger), fmt.Sprintf("%v", withLogger))
}
