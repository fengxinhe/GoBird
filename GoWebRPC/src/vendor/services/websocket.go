package services

import (
    "golang.org/x/net/websocket"
    "net/rpc/jsonrpc"
    "log"
    "fmt"
)

func JsonrpcHandler(ws *websocket.Conn) {
    jsonrpc.ServeConn(ws)
}

func PushHandler(ws *websocket.Conn) {
    var id int
    c := jsonrpc.NewClient(ws)

    err := c.Call("User.Getpsd", nil, &id)
    if err != nil {
        log.Print("User.Getid error:", err)
        return
    } else {
        fmt.Println("websocket")
        NotifyHandler(ws)
        User.Client[id] = c;
        User.Mutex[id].Lock()
    }

}