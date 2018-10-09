package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"
)

// XMLRPCCall doc
func XMLRPCCall(method string, args struct{ Arg0 string }) (response struct{ Message string }, err error) {
	buf, _ := xml.EncodeClientRequest(method, &args)
	rpcURI := "http://localhost:8800/rpc"
	rpcType := "text/xml"
	rpcPayload := bytes.NewBuffer(buf)

	resp, err := http.Post(rpcURI, rpcType, rpcPayload)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = xml.DecodeClientResponse(resp.Body, &response)
	return
}

func main() {
	response, err := XMLRPCCall("RPCService.Call", struct{ Arg0 string }{"Arg0 Value"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %s\n", response.Message)
}
