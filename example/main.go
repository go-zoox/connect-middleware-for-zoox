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
