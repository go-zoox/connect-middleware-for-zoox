# Connect Middleware for Zoox

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/connect-middleware-for-zoox)](https://pkg.go.dev/github.com/go-zoox/connect-middleware-for-zoox)
[![Build Status](https://github.com/go-zoox/connect-middleware-for-zoox/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/connect-middleware-for-zoox/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/connect-middleware-for-zoox)](https://goreportcard.com/report/github.com/go-zoox/connect-middleware-for-zoox)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/connect-middleware-for-gin/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/connect-middleware-for-gin?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/connect-middleware-for-gin.svg)](https://github.com/go-zoox/connect-middleware-for-zoox/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/connect-middleware-for-gin.svg?label=Release)](https://github.com/go-zoox/connect-middleware-for-zoox/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/connect-middleware-for-zoox
```

## Getting Started

```go
package main

import (
	"github.com/go-zoox/connect-middleware-for-zoox"
	"github.com/go-zoox/logger"
	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
)

func main() {
	r := defaults.Default()

	r.Use(connect.Create("YOUR_SECRET_KEY"))

	r.Get("/user", func(c *zoox.Context) {
		user, err := connect.GetUser(c)
		if err != nil {
			c.JSON(401, zoox.H{
				"message": "unauthorized",
			})
			return
		}

		c.JSON(200, zoox.H{
			"user": user,
		})
	})

	r.Get("/", func(c *zoox.Context) {
		c.JSON(200, zoox.H{
			"message": "helloworld",
		})
	})

	if err := r.Run(); err != nil {
		logger.Fatal("error running server: %s", err)
	}
}

```

## Related Projects
* [go-zoox/connect](https://github.com/go-zoox/connect) - The Auth Connect.
* [go-zoox/connect-middleware-for-gin](https://github.com/go-zoox/connect-middleware-for-gin) - The Auth Connect Middleware for Gin.

## License
GoZoox is released under the [MIT License](./LICENSE).
