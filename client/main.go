package main

import (
    "log"
    "net/rpc"
    "client/utils"
)

func main() {
    client, err := rpc.Dial("tcp", "localhost:9000")
    if err != nil {
        log.Fatal("Dialing error:", err)
    }

    users, err := utils.ReadUsersFromFile("users.json")
    if err != nil {
        log.Fatal("Error reading users:", err)
    }

    for _, user := range users {
        var reply string
        err = client.Call("UserHandler.CreateUser", user, &reply)
        if err != nil {
            log.Println("Error creating user:", err)
        } else {
            log.Println(reply)
        }
    }
}
