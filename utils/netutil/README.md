# Net Utils

- provide some network utils. eg: `InternalIPv4`
- sub package: 
  - `httpctypes` - provide some commonly http content types.
  - `httpheader` - provide some commonly http header names.
  - `httpreq` - provide some http request utils

## Install

```bash
go get github.com/yoas0bi/micro-toolkit/utils/netutil
```

## Go docs

- [Go docs](https://pkg.go.dev/github.com/yoas0bi/micro-toolkit/utils/netutil)

## Usage

```go
import "github.com/yoas0bi/micro-toolkit/utils/netutil"
```

```go
// Get internal IPv4 address
netutil.InternalIPv4()
```

## Testings

```shell
go test -v ./netutil/...
```

Test limit by regexp:

```shell
go test -v -run ^TestSetByKeys ./netutil/...
```
