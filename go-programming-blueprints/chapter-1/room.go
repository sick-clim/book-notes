package main

type room struct {
	// 他のクライアントに転送するためのメッセージを保持するチャネル
	forward chan []byte
	// チャットルームに参加しようとしているクライアントのためのチャネル
	join chan *client
	// チャットルームから退出しようとしているクライアントのためのチャネル
	leave chan *client
	// 在室しているすべてのクライアントが保持されます
	clients map[*client]bool
}
