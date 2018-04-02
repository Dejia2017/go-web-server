package controllers

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"go-web-server/models"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.Message)
)

func init() {
	go handleMessages()
}

//广播发送至页面
func handleMessages() {
	for {
		msg := <-broadcast
		fmt.Println("clients len ", len(clients))
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}