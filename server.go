package main

import (
    "net/rpc/jsonrpc"
    "net/rpc"
    "net"
    "log"
    "strings"
    "flag"
    "os/exec"
)

type RPCFunc uint8

func (*RPCFunc) Echo(arg *string, result *string) error {
    log.Print("Arg passed: " + *arg)
    *result = ">" + *arg + "<"
    return nil
}

func (*RPCFunc) Build(arg *string, result *string) error {
    log.Print("BUILD SYSTEM INVOKED")
    log.Print("Runnig command: " + *arg)

    commands := strings.Fields(*arg)
    caller := string(commands[0])
    cmd := exec.Command(caller, strings.Join(commands[1:], " "))
    err := cmd.Start()
    if err != nil {
        log.Print("Uh oh: " + err.Error())
    }
    err = cmd.Wait()
    log.Printf("Command finished with error: %v", err)

    *result = ">" + *arg + "<"
    return nil
}

func main() {
    var socket = flag.String(
        "socket", "/tmp/build.sock", "The socket to bind to")
    flag.Parse()

    log.Print("Starting Server...")
    l, err := net.Listen("unix", *socket)
    defer l.Close()
    if err != nil {
        log.Fatal(err)
    }
    log.Print("listening on: ", l.Addr())
    rpc.Register(new (RPCFunc))
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
