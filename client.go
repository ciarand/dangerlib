package main
// Connect to JSONRPC Server and send command-line args to Echo

import (
  "net/rpc/jsonrpc"
  "os"
  "net"
  "fmt"
  "flag"
)

func main() {
    var socket = flag.String(
        "socket", "/tmp/build.sock", "The socket to bind to")

    flag.Parse()
    conn, e := net.Dial("unix", *socket)
    if e != nil {
        fmt.Fprintf(os.Stderr, "Could not connect: %s\n", e)
        os.Exit(1)
    }
    client := jsonrpc.NewClient(conn)
    var reply string
    for i, arg := range os.Args {
        if i == 0 { continue } // Ignore program name
        fmt.Printf("Sending: %s\n", arg)
        client.Call("RPCFunc.Build", arg, &reply)
        fmt.Printf("Reply: %s\n", reply)
    }
}
