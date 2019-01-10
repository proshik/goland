package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name/", Hello)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

//
//r.GET("/account/filter", filter)                    // GET: /accounts/filter - для поиска аккаунтов по чётким критериям;
////r.GET("/account/group", group)                      // GET: /accounts/group - для подсчёта пользователей по группам;
//r.GET("/account/*id/recommend/", recommend) // GET: /accounts/<id>/recommend - для поиска подходящей "второй половинки";
////r.GET("/account/:id/suggest/", suggest)       // GET: /accounts/<id>/suggest/ - для поиска возможных симпатий по симпатиям других;
////r.POST("/account/new", add)                         // POST: /accounts/new/ - для добавления новых пользователей;
////r.POST("/account/:id", update)                       // POST: /accounts/<id>/ - для обновления данных;
////r.POST("/account/likes", addLikes)                  // POST: /accounts/likes/ - для добавления новых симпатий;
