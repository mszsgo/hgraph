package hgraph

import (
	"net/http"

	"github.com/graphql-go/handler"
)

// Http Handler
// h := hgraph.GraphqlHttpHandler(&Query{}, &Mutation{})
func GraphqlHttpHandler(query, mutation interface{}) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   GraphqlSchema(query, mutation),
		Pretty:   true,
		GraphiQL: true,
	})
}

func HttpHandle(query, mutation interface{}) {
	h := GraphqlHttpHandler(query, mutation)
	// Graphql服务
	/*http.Handle("/graphql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ctx := context.WithValue(r.Context(), "token", r.Header.Get("MS-token"))
		// token、用户id、用户名等信息存储header，用于记录操作日志
		//h.ContextHandler(ctx, w, r)
	}))*/
	http.Handle("/graphql", h)
}
