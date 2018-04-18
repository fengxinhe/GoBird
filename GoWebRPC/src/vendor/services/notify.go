package services

import (
    "golang.org/x/net/websocket"
    "fmt"
)

type connection struct {
    // The websocket connection.
    ws *websocket.Conn

    // Buffered channel of outbound messages.
    send chan string
}
type message struct {
    method   string  `json:"method"`
    like     int    `json:"like"`
}
func (c *connection) reader() {
    for {
        var msg string
        fmt.Println("wwwrrrrr")
        err := websocket.Message.Receive(c.ws, &msg)
        if err != nil {
            break
        }
        H.broadcast <- msg
    }
    c.ws.Close()
}

func (c *connection) writer() {
    for message := range c.send {
        fmt.Println("wwwwwwwwwwww")
        err := websocket.Message.Send(c.ws, message)
        if err != nil {
            break
        }
    }
    c.ws.Close()
}

// func (c *connection) writetoall() {
//     for message := range c.send {
//         for client := range h.connections {
//
//         err := websocket.Message.Send(client.ws, message)
//         if err != nil {
//             break
//         }
//     }
//     }
//     c.ws.Close()
// }

func NotifyHandler(ws *websocket.Conn) {
    c := &connection{send: make(chan string, 256), ws: ws}
    fmt.Println("dddddddddd")
    H.register <- c
    //H.broadcast <- "sss"
    defer func() { H.unregister <- c }()
    go c.writer()
    c.reader()
}