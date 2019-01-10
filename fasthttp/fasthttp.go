package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type MyHandler struct {
	foobar string
}

//// request handler in net/http style, i.e. method bound to MyHandler struct.
//func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
//	// notice that we may access MyHandler properties here - see h.foobar.
//	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
//		ctx.Path(), h.foobar)
//}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {

	fmt.Printf("ctx.RequestURI(): %s\n", string(ctx.RequestURI()))
	fmt.Printf("ctx.Path(): %s\n", string(ctx.Path()))

	if ctx.Request.Header.IsGet() {
		switch string(ctx.Path()) {
		case "/account/filter/":
			filter(ctx)
		case "/account/group/":
			group(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	} else if ctx.Request.Header.IsPost() {
		switch string(ctx.Path()) {
		case "/account/new/":
			filter(ctx)
		case "/account/likes/":
			group(ctx)
		default:
			uri := ctx.Request.RequestURI()
			if i64, ok := byteSliceToInt64(uri[9:]); ok {
				id := int32(i64)
				fmt.Println(id)
			}
		}
	} else {
		ctx.Error("", fasthttp.StatusNotFound)
	}

	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

func byteSliceToInt64(s []byte) (res int64, ok bool) {
	sign := len(s) > 0 && s[0] == '-'
	if sign {
		s = s[1:]
	}

	ok = true

	res = 0
	for _, c := range s {
		if v := int64(c - '0'); v < 0 || v > 9 {
			ok = false
			break
		} else {
			res = res*10 + v
		}
	}

	if sign {
		res = -res
	}

	return
}

func main() {
	// pass bound struct method to fasthttp
	//myHandler := &MyHandler{
	//	foobar: "foobar",
	//}
	//fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)

	// pass plain function to fasthttp
	fasthttp.ListenAndServe(":80", fastHTTPHandler)
}

func filter(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}

func group(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}

func recommend(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}

func suggest(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}

func add(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}

func update(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}

func addLikes(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}
