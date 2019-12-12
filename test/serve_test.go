package test

import (
	"log"
	"net/http"
	"testing"

	"github.com/mszsgo/hgraph"
)

func TestServeHandel(t *testing.T) {
	log.Print("启动服务： http://localhost:9990")
	http.ListenAndServe(":9990", hgraph.GraphqlHttpHandler(Query{}, Mutation{}))
}
