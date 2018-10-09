package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
)

// RPCService doc
type RPCService struct{}

// Call doc
func (h *RPCService) Call(req *http.Request, args *struct{ Arg0 string }, response *struct{ Message string }) error {
	log.Printf("%v", args)
	response.Message = "Hello " + args.Arg0 + "!"
	return nil
}

func main() {
	PORT := 8800
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	RPC.RegisterService(new(RPCService), "")
	http.Handle("/rpc", RPC)

	log.Printf("Starting XML-RPC server on localhost:%d/rpc\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
