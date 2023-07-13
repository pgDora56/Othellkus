package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.SetFuncMap(template.FuncMap{
	// 	"safe": noescape,
	// })
	// js,css,faviconなどを読み込むためのasstes設定
	router.LoadHTMLGlob("view/*.tmpl")
	router.Static("/resource", "./resource")
	router.StaticFile("/favicon.ico", "./resource/favicon.ico")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "board.tmpl", gin.H{})
	})

	router.GET("/controller", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "controller.tmpl", gin.H{})
	})

	router.Run(":52417")

	log.Println("End Introquiz Portal Square Azure")
}
