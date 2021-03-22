# `ocommon`

[中文/Chinese](README.ZH.md)

## Description

Open common module. Support `string`, `slice`, or others common operation. 

|           | description         |
| --------- | ------------------- |
| `oid`    | Get `goroutine` id. |
| `omap`    | `map` operate.      |
| `oslice`  | `slice` operate.    |
| `ostring` | `string` operate.    |
| `oticker` | `ticker` operate.    |

## Index

Methods index.

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

## Usage

Please see tests.
