package oticker

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTickerAdvance(t *testing.T) {
	var flag = false
	//advance will be executed once in the go routine
	TickerAdvance(500*time.Millisecond, func() {
		flag = !flag
	})
	assert.Equal(t, flag, true)
	<-time.After(750 * time.Millisecond)
	assert.Equal(t, flag, false)
}

func TestTickerAdvanceCtx(t *testing.T) {
	var flag = false
	ctx, cancelFunc := context.WithCancel(context.Background())
	TickerAdvanceCtx(ctx, 500*time.Millisecond, func() {
		flag = !flag
	})
	assert.Equal(t, flag, true)
	cancelFunc()
	<-time.After(750 * time.Millisecond)
	assert.Equal(t, flag, true)
}
func TestTicker(t *testing.T) {
	var flag = false
	Ticker(500*time.Millisecond, func() {
		flag = !flag
	})
	assert.Equal(t, flag, false)
	<-time.After(750 * time.Millisecond)
	assert.Equal(t, flag, true)
}

func TestTickerCtx(t *testing.T) {
	var flag = false
	ctx, cancelFunc := context.WithCancel(context.Background())
	TickerCtx(ctx, 500*time.Millisecond, func() {
		flag = !flag
	})
	assert.Equal(t, flag, false)
	cancelFunc()
	<-time.After(750 * time.Millisecond)
	assert.Equal(t, flag, false)
}
