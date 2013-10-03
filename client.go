package dangerlib

// Connect to JSONRPC Server and send command-line args to build

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"strings"
)

func Connect(args []string) {
	conn, e := net.Dial("unix", "/tmp/build.sock")
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not connect: %s\n", e)
		os.Exit(1)
	}

	client := jsonrpc.NewClient(conn)
	var reply string

	request := &Request{"RPCFunc.Runner", args}

	fmt.Printf("Sending: %s\n", strings.Join(args, " "))
	client.Call(request.Command, request, &reply)
	fmt.Printf("Reply: %s\n", reply)
}
