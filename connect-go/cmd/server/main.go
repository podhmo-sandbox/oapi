package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"

	greetv1 "github.com/podhmo-sandbox/oapi/connect-go/gen/greet/v1"        // generated by protoc-gen-go
	"github.com/podhmo-sandbox/oapi/connect-go/gen/greet/v1/greetv1connect" // generated by protoc-gen-connect-go
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err // automatically coded correctly
	}
	log.Println("Request headers: ", req.Header())
	// if err := validateGreetRequest(req.Msg); err != nil {
	// 	return nil, connect.NewError(connect.CodeInvalidArgument, err)
	// }
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	// // http2
	// http.ListenAndServe(
	// 	"localhost:8080",
	// 	// Use h2c so we can serve HTTP/2 without TLS.
	// 	h2c.NewHandler(mux, &http2.Server{}),
	// )

	if err := http.ListenAndServe(":33333", mux); err != nil {
		panic(err)
	}
}
