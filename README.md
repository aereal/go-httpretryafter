![CI][ci-badge]
[![PkgGoDev][pkggodev-badge]][pkggodev]

# go-httpretryafter

Provides functions to parse [Retry-After][retry-after] header.

## Synopsis

```go
import (
  "net/http"

  httpretryafter "github.com/aereal/go-httpretryafter"
)

var req *http.Request
httpretryafter.Parse(req.Header.Get("Retry-After"))
```

## Installation

```
go get github.com/aereal/go-httpretryafter
```

[ci-badge]: https://github.com/aereal/go-httpretryafter/workflows/CI/badge.svg?branch=main
[pkggodev-badge]: https://pkg.go.dev/badge/aereal/go-httpretryafter
[pkggodev]: https://pkg.go.dev/github.com/aereal/go-httpretryafter
[retry-after]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After
