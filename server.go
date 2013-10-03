package dangerlib

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os/exec"
	"strings"
)

func (*RPCFunc) Runner(request *Request, result *string) error {
	log.Print("Running command: " + strings.Join(request.Args, " "))

	// The command is running the first string
	command := request.Args[0]
	// The args are getting passed in
	args := request.Args[1:]

	out, err := exec.Command(command, args...).CombinedOutput()

	if err != nil {
		log.Print("Failure: " + err.Error())
	}

	if len(out) == 0 {
		*result = fmt.Sprintf("Silent success")
	} else {
		*result = fmt.Sprintf("%s", out)
	}
	return nil
}

func Serve() {
	log.Print("Starting Server...")
	l, err := net.Listen("unix", "/tmp/build.sock")
	defer l.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("listening on: ", l.Addr())
	rpc.Register(new(RPCFunc))
	for {
		log.Print("waiting for connections ...")
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %s", conn)
			continue
		}
		log.Printf("connection started: %v", conn.RemoteAddr())
		go jsonrpc.ServeConn(conn)
	}
}
