package silver

import (
	"errors"
	"fmt"

	"code.google.com/p/go.net/websocket"
	"github.com/jjjabc/JJJCarrier"
)

type silverClient struct {
	pipe     chan Silver
	pullpipe chan Silver
	order    chan string
	wsConn   *websocket.Conn
}

func (c *silverClient) Send(msg JJJCarrier.Msg) error {
	json := JsonSilver{Time: msg.Time().String(), Price: float32(msg.GetContext().(float32))}
	return websocket.JSON.Send(c.wsConn, json)
}
func (c *silverClient) PushTheard() error {
	for {
		select {
		case msg := <-c.pipe:
			err := c.Send(msg)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		case ord := <-c.order:
			if ord == "quit" {
				return errors.New("client: order to quit")
			}

		}
	}
	for msg := range c.pipe {
		err := c.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *silverClient) PullTheard() (interface{}, error) {
	return JJJCarrier.ErrNotSupport, nil
}
func (c *silverClient) Close() error {
	c.order <- "quit"
	return nil
}
func (c *silverClient) String() string {
	return c.wsConn.RemoteAddr().String() + c.wsConn.Request().RemoteAddr
}
