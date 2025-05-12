package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	auth := r.Group("/auth")
	auth.POST("/login", func(ctx *gin.Context) {})
	auth.POST("/register", func(ctx *gin.Context) {})

	r.POST("/push", func(ctx *gin.Context) {
		//check if the user has the perms to push to a given repo
		//if so then save the pushed file
		//bad file names should probably return a 400?

		repo := ctx.GetHeader("Repo")
		filename := ctx.GetHeader("Filename")

		_, err := os.Stat(path.Join(repo, filename))
		if err == nil {
			ctx.String(http.StatusBadRequest, "File already exists")
			return
		}

		_, err = os.Stat(repo)
		if err != nil {
			if err = os.Mkdir(repo, os.ModePerm); err != nil {
				//fyi i know that panics will kill the whole program
				panic(err)
			}
		}

		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(path.Join(repo, filename), body, os.ModePerm)
		if err != nil {
			panic(err)
		}
		ctx.String(http.StatusOK, "ok")
	})

	r.Run()
}
