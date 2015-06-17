package JJJCarrier

import (
	"errors"
)

type Notifyer interface {
	Notify(v interface{}) error
}
type Station struct {
	clients map[string]Client
	msg     []Msg
	in      chan Msg
	order   chan string
}

func (this *Station) ClientReg(c Client) error {
	this.clients[c.String()] = c
	return nil
}
func (this *Station) ClientUnreg(c Client) error {
	err := c.Close()
	if err != nil {
		return err
	}
	delete(this.clients, c.String())
	return nil
}
func (this *Station) TransferTheard() error {
	for {
		select {
		case m := <-this.in:
			this.receive(m)
		case o := <-this.order:
			if o == "quit" {
				return errors.New("Station: order to quit")
			}
		}
	}
}
func (this *Station) Broadcast(m Msg) {
	for _, client := range this.clients {
		client.Send(m)
	}
}
func (this *Station) receive(m Msg) {
	this.Broadcast(m)
}
func (this *Station) Close() {
	for _, v := range this.clients {
		this.ClientUnreg(v)
	}
	this.order<-"quit"
}
func (this *Station) Fetchloop(f Fetch){
	
}
