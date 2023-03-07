package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Params.ByName("name")
		ctx.JSON(200, gin.H{
			"msg": "hello " + name,
		})
	})

	r.Run()
}
