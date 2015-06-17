package silver

import(
		"code.google.com/p/go.net/websocket"
	"github.com/jjjabc/JJJCarrier"
)
type WsServer struct{
	client *silverClient
	station *JJJCarrier.Station
}
func (ws *WsServer)Init(sta *JJJCarrier.Station,cli *silverClient){
	ws.station=sta
	ws.client=cli	
}
func (ws *WsServer)WebsocketHandler(c *websocket.Conn){
	ws.client.wsConn=c
	ws.client.PushTheard()
	defer c.Close()
}