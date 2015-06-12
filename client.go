package JJJCarrier

import (
	"strconv"
	"time"

	"code.google.com/p/go.net/websocket"
	"github.com/jjjabc/JJJCarrier/silver"
)

//Pusher接口
type Pusher interface {
	Push() error
}

//Puller接口
type Puller interface {
	Pull() error
}
type Msg interface {
	String() string
	Time() time.Time
	GetContext() interface{}
}

type message struct {
	s silver.Silver
}

func (c message) String() string {
	return strconv.FormatFloat(float64(c.s.Price), 'f', 3, 32)
}
func (c message) Time() time.Time {
	return c.s.T
}
func (c message) GetContext() interface{} {
	return c.s.Price
}

type Client interface {
	Puller
	Pusher
	Send(Msg) error
}
type client struct {
	pipe   chan message
	wsConn *websocket.Conn
}

func (c *client) Send(msg Msg) error {
	json := JsonSilver{Time: msg.Time().String(), Price: float32(msg.GetContext().(float32))}
	return websocket.JSON.Send(c.wsConn, json)
}
func (c *client) Push() error {
	for msg := range c.pipe {

		err := c.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil

}

type JsonSilver struct {
	Time  string  `json:"time"`
	Price float32 `json:"price"`
}
