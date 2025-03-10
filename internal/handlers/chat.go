package handlers

import (
	"videochat-project/pkg/chat"
	w "videochart/pkg/webrtc"
	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket/v2"
)

func RoomChat(c *fiber.Ctx)error{
	return c.Render("chat",fiber.Map{}, "layouts/main")
}

func RoomChatWebsocket(c *websocket.Conn){
	uuid := c.Params("uuid")
	if uuid == ""{
		return
	}
	w.RoomsLock.Lock()
	room := w.Rooms[uuid]
	w.Roomslock.unlock()
	if room == nil {
		return
	}
	if room.Hub == nil {
		return
	}
	chat.PeerChatConn(c.conn,room.Hub)
}

func StreamChatWebSocket(c *websocket.Conn){
	suuid := c.Params("suuid")
	if suuid == ""{
		return
	}
	w.RoomsLock.Lock()
	if stream, ok := w.Streams[suuid]; ok{
		w.RoomsLock.Unlock()
		if stream.Hub == nil {
			hub := chat.NewHub()
			stream.Hub = hub
			go hub.Run()
		}
		chat.PeerChatConn(c.Conn, stream.Hub)
		return
	}
	w.RoomsLock.Unlock()
}