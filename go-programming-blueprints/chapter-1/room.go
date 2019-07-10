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

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// 参加
			r.clients[client] = true
		case client := <-r.leav:
			// 退出
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// 全てのクライアントにメッセージを転送
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージ送信
				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
