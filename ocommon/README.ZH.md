# `ocommon`

[英文/English](README.md)

## Description

开放的 common 模块，支持`string`, `slice`, 或者一些其他的公共操作。 

|           | 描述         |
| --------- | ------------------- |
| `oid`    | 获取 `goroutine` id. |
| `omap`    | `map` 操作.      |
| `oslice`  | `slice` 操作.    |
| `ostring` | `string` 操作.    |
| `oticker` | `ticker` 操作.    |

## Index

方法索引

|  package  | method  |  return |
| --------- | ------- | ------- |
| `oid`    | `ID`    | `uint64` |
| `omap`    | `MapStrStrToString`    | `string` |
| `oslice`    | `IntSliceToStringSlice`    | `[]string` |
| `oslice`    | `IntSliceContains`    | `bool` |
| `oslice`    | `IntSliceContainsAny`    | `bool` |
| `oslice`    | `IntSliceContainsAll`    | `bool` |
| `oslice`    | `IntSliceIsEmpty`    | `bool` |
| `oslice`    | `StringSliceContains`    | `bool` |
| `oslice`    | `StringSliceContainsAny`    | `bool` |
| `oslice`    | `StringSliceContainsAll`    | `bool` |
| `oslice`    | `StringSliceIsEmpty`    | `bool` |
| `ostring`    | `IsEmpty`    | `bool` |
| `ostring`    | `IsNotEmpty`    | `bool` |
| `ostring`    | `IsAllEmpty`    | `bool` |
| `ostring`    | `IsAnyEmpty`    | `bool` |
| `ostring`    | `IsBlank`    | `bool` |
| `ostring`    | `IsNotBlank`    | `bool` |
| `ostring`    | `IsAnyBlank`    | `bool` |
| `ostring`    | `DefaultIfEmpty`    | `string` |
| `ostring`    | `DefaultIfBlank`    | `string` |
| `ostring`    | `FirstNotEmpty`    | `string` |
| `ostring`    | `FirstNotBlank`    | `string` |
| `ostring`    | `FromBytes`    | `string` |
| `ostring`    | `ToBytes`    | `[]byte` |

## 使用

请参考测试用例。
