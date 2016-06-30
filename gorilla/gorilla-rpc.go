package main

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

func main() {
	/*
		POST {"method":"HelloService.Say","params":[{"Who":"Test"}], "id":"1"} to localhost:8080/rpc
	*/

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(HelloService), "")
	http.Handle("/rpc", s)
	http.ListenAndServe(":8080", nil)
}

type HelloArgs struct {
	Who string
}

type HelloReply struct {
	Message string
}

type HelloService struct{}

func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	reply.Message = "Hello " + args.Who
	return nil
}
