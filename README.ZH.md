# golang-library

[![Tests And Coverage](https://github.com/baishan-development-guizhou/golang-library/actions/workflows/test.yml/badge.svg)](https://github.com/baishan-development-guizhou/golang-library/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/baishan-development-guizhou/golang-library/branch/master/graph/badge.svg?token=MBPD4JCBSL)](https://codecov.io/gh/baishan-development-guizhou/golang-library)

[英文/English](README.md)

## 描述

`golang` 常用公共库.

| 模块 | 描述 | 备注 |
| ------ | ----------- | ------ |
| [echo_ext](https://github.com/baishan-development-guizhou/golang-library/tree/master/echo_ext) | [echo](https://github.com/labstack/echo) 扩展. 包括路由熔断、错误处理等... |
| [log](https://github.com/baishan-development-guizhou/golang-library/tree/master/log) | [zap](https://github.com/uber-go/zap) 扩展. 它支持从上下文获取 `Logger` 并在 `goroutine` 中使用. |
| [ocommon](https://github.com/baishan-development-guizhou/golang-library/tree/master/ocommon) | 开放的 common 模块，支持`string`, `slice`, 或者一些其他的公共操作。  |
| [report](https://github.com/baishan-development-guizhou/golang-library/blob/master/report/README.md) | 为你的应用自动上报信息到监控平台. | 只支持 `v2` 端点.

## 开源协议

遵循 [MIT](https://opensource.org/licenses/MIT) 开源协议 (MIT)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fbaishan-development-guizhou%2Fgolang-library.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fbaishan-development-guizhou%2Fgolang-library?ref=badge_large)
