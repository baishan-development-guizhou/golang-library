<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

# Table of Content

- [Name](#name)
- [Description](#description)
- [Author](#author)
- [Copyright and License](#copyright-and-license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Name

golang-library-log

a golang log including `Debug` `Debugf` `Info` `Infof` `Warn` `Warnf` `Error` `Errorf` and `With` function.

# Description

a simple encapsulation of zap logger is made,it supports getting `ILogger` from context to using in `goroute`

`Global` returns global ILogger extends from zap.Logger.

`ReplaceGlobal` replace the global ILogger with the one passed by the parameter.

`Context` is getting `context.Context` and `ILogger` from `context.Context`,if the context does not contain a logger,
the context and logger are generated and returned.

`Associate` returns a copy of context.Context in which the ILogger associated.

`Routine` returns the ILogger bound to the current goroutine,you'd better not use this func.

`BindRoutine` bind the ILogger to the current goroutine

## Synopsis

```go
package main

import (
	"context"
	"github/baishan-development-guizhou/golang-library/log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//config
	logger := log.Configure().WithCaller(true).WithStacktrace(true).WithLevel(log.InfoLevel).Init()
	logger = logger.With("X-Request-ID", "balabala")
	logger.Error("first")
	//associate
	ctx, _ := log.Associate(context.Background(), logger)
	go func() {
		//use
		_, iLogger := log.Context(ctx)
		iLogger.Info("second")
		wg.Done()
	}()
	wg.Wait()
}

```

# Author

RuiFG (樊国睿) <guorui.fan@baishancloud.com>

# Copyright and License

The MIT License (MIT)

Copyright (c) 2021 RuiFG (樊国睿) <guorui.fan@baishancloud.com>