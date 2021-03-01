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

a golang logger
including `Debug` `Debugf` `Info` `Infof` `Warn` `Warnf` `Error` `Errorf` `Fatal` `Fatalf` `Panic` `Panicf`
and `With` `Named` `Sync`
function.

# Description

a simple encapsulation of zap logger is made,it supports getting `Logger` from context to using in `goroute`

`G` returns global Logger extends from zap.Logger.

`ReplaceG` replace the global Logger with the one passed by the parameter.

`C` is getting `Logger` from `context.Context`,if the context does not contain a logger, the `G()` logger will return.

`AssociateC` returns a copy of context.Context in which the Logger associated.

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

# Author

RuiFG (樊国睿) <guorui.fan@baishancloud.com>

# Copyright and License

The MIT License (MIT)

Copyright (c) 2021 RuiFG (樊国睿) <guorui.fan@baishancloud.com>