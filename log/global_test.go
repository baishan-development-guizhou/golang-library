package log

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"sync"
	"testing"
)

func TestLoggerContextKey(t *testing.T) {
	ctx := context.Background()
	logger := G()
	withValue := context.WithValue(ctx, loggerKey{}, logger)
	value := ctx.Value(loggerKey{})
	assert.NotEqual(t, fmt.Sprintf("%T", value), fmt.Sprintf("%T", logger))
	assert.Equal(t, withValue.Value(loggerKey{}), logger)
	assert.Equal(t, context.WithValue(withValue, "", "").Value(loggerKey{}), logger)

}
func TestLoggerFunc(t *testing.T) {
	parent := context.Background()
	ctx, logger := AssociateC(parent, G())
	assert.NotNil(t, ctx)
	assert.NotEqual(t, ctx, parent)
	assert.NotNil(t, logger)
	loggerEq := C(ctx)
	assert.Equal(t, logger, loggerEq)
}

func TestContext(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, logger := AssociateC(context.TODO(), G())
	iLogger := C(ctx)
	assert.Equal(t, logger, iLogger)

}

func BenchmarkLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("")
	}
}
func BenchmarkZapLogger(b *testing.B) {
	zap.ReplaceGlobals(zap.NewNop())
	for i := 0; i < b.N; i++ {
		zap.S().Info("")
	}
}
