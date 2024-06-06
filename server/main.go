package main

import (
    "log"
    "net"
    "net/rpc"
    "server/handlers"
    "server/db"
)

func main() {
    listener, err := net.Listen("tcp", ":9000")
    if err != nil {
        log.Fatal("Listener error:", err)
    }
    defer listener.Close()

    userHandler := new(handlers.UserHandler)
    rpc.Register(userHandler)

    db.InitDB()

    log.Println("RPC server is running on port 9000")
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Connection error:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
