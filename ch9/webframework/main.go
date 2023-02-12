package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	var connHandler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		ctx.Value(BirdnestCtxKey).(*ContextInfo).Params["LocalAddrContextKey"] = r.Context().Value(http.LocalAddrContextKey)
	}
	var authHandler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == "123456" {
			ctx.Value(BirdnestCtxKey).(*ContextInfo).Params["Valid"] = true
		}
	}

	var b = New()
	// 添加中间件
	b.UseFunc(connHandler)
	b.UseFunc(authHandler)

	// 添加handler
	mux := httprouter.New()
	mux.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx := r.Context()
		params := ctx.Value(BirdnestCtxKey).(*ContextInfo).Params

		localAddr := params["LocalAddrContextKey"].(*net.TCPAddr)
		valid := params["Valid"]
		w.Write([]byte(fmt.Sprintf("hello world. localAddr: %s, valid: %v\n", localAddr, valid)))
	})
	b.SetMux(mux)

	b.Run(":8080")
}

// ContextInfo 是一个保存上下文信息的对象.
type ContextInfo struct {
	Params map[string]any
}

// 在上下文中的key的类型
type contextKey struct {
	name string
}

// BirdnestCtxKey 是一个上下文中的key.
var BirdnestCtxKey = contextKey{"BirdnestCtxKey"}

// NewContext 创建一个新的上下文.
func NewContext() context.Context {
	return context.WithValue(context.Background(), BirdnestCtxKey, &ContextInfo{
		Params: make(map[string]any),
	})
}

// Handler 是中间件类型，定义了插件的方法签名.
type Handler interface {
	ServeHTTP(c context.Context, w http.ResponseWriter, r *http.Request)
}

// HandlerFunc 也是一个中间件类型，以接口的形式提供.
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)

func (fn HandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fn(ctx, w, r)
}

func Wrap(handler http.Handler) Handler {
	return HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

// Birdnest 是一个http服务的中间件框架.
type Birdnest struct {
	middleware []Handler
	mux        http.Handler
}

// New 创建一个新的Birdnest实例.
func New() *Birdnest {
	return &Birdnest{
		mux: http.DefaultServeMux,
	}
}

// Use 增加一个中间件.
// 运行之后不能再增加了，否则会出现非预期的现象.
func (b *Birdnest) Use(handler Handler) {
	b.middleware = append(b.middleware, handler)
}

// UseFunc 增加一个中间件.
func (b *Birdnest) UseFunc(handleFunc HandlerFunc) {
	b.Use(HandlerFunc(handleFunc))
}

// SetMux 使用定制化的mux, 比如httprouter.
func (b *Birdnest) SetMux(mux http.Handler) {
	b.mux = mux
}

// ServeHTTP 实现http.Handler接口.
func (b *Birdnest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext()

	for _, handler := range b.middleware {
		handler.ServeHTTP(ctx, w, r)
	}

	b.mux.ServeHTTP(w, r.WithContext(ctx))
}

// Run 运行http服务.
func (b *Birdnest) Run(addr string) error {
	return http.ListenAndServe(addr, b)
}
