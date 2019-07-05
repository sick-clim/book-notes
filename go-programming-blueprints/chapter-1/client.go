package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行っている1人のユーザを表す
type client struct {
	// このクライアントのためのWebSocket
	socket *websocket.Conn
	// メッセージが送られるチャネル
	send chan []byte
	// このクライアントが参加しているチャットルーム
	room *room
}
