# `log`

[英文/English](README.md)

对 [zap](https://github.com/uber-go/zap) 做了一个简单的封装，它支持从上下文获取 `logger` 并可以在 `goroutine` 中使用。

## Description

包括 `Debug` `Debugf` `Info` `Infof` `Warn` `Warnf` `Error` `Errorf` `Fatal` `Fatalf` `Panic` `Panicf`
和 `With` `Named` `Sync` `StdLogger`
等函数.

|      |   描述   |
|  --- | -------- |
| `G`  | 返回一个 继承了 `zap.Logger` 的全局日志 |
| `ReplaceG` | 通过参数将全局日志 `G` 替换 |
|`C` |  直接从 `context.Context` 获取到 `Logger`. 如果上下文中没有包含一个 `logger`, 将会返回 `G()`.
| `AssociateC` | 返回 `Logger` 关联的 `context.Context` 的副本。

## 使用

```go
package main

import (
	"context"
	"github.com/baishan-development-guizhou/golang-library/log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	logger := log.Configure().WithOutputEncoder(log.ConsoleOutputEncoder).WithCallerEncoder(log.ShortRoutineCallerEncoder).
		WithStacktrace(false).WithLevel(log.DebugLevel).
		Init()
	logger = logger.With("dsd", "sdd")
	ctx, _ := log.AssociateC(context.Background(), logger)
	go func() {
		iLogger := log.C(ctx)
		iLogger.Info("second")
		wg.Done()
	}()
	wg.Wait()
}
```
