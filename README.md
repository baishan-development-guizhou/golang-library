# golang-library

[![release](https://img.shields.io/github/v/release/baishan-development-guizhou/golang-library?color=orange)](https://github.com/baishan-development-guizhou/golang-library/releases)
[![Go mod](https://img.shields.io/github/go-mod/go-version/baishan-development-guizhou/golang-library?style=plastic)](https://golang.org/dl/)
[![Tests And Coverage](https://github.com/baishan-development-guizhou/golang-library/actions/workflows/test.yml/badge.svg)](https://github.com/baishan-development-guizhou/golang-library/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/baishan-development-guizhou/golang-library/branch/master/graph/badge.svg?token=MBPD4JCBSL)](https://codecov.io/gh/baishan-development-guizhou/golang-library)

[中文/Chinese](README.ZH.md)

## Description

Common public library by `golang`.

| module | description | remark |
| ------ | ----------- | ------ |
| [echo_ext](echo_ext) | [echo](https://github.com/labstack/echo) Extension. Add circuit breaker and error handler etc... |
| [log](log) | [zap](https://github.com/uber-go/zap) Extension. It supports getting Logger from context to using in goroutine. |
| [ocommon](ocommon) | Open common module. Support `string`, `slice`, or others common operation. |
| [report](report) | Report our application information to server metric. | Only support `v2` endpoint.

## License

[![MIT](https://img.shields.io/github/license/baishan-development-guizhou/golang-library)](https://opensource.org/licenses/MIT)

The [MIT](https://opensource.org/licenses/MIT) License (MIT)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fbaishan-development-guizhou%2Fgolang-library.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fbaishan-development-guizhou%2Fgolang-library?ref=badge_large)
