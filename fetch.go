package JJJCarrier

import (
	"fmt"
)
type Marshaler interface{
	Marshal(interface{}) ([]Msg, error)
}

/*Fetcher接口描述了通过URL返回一个Msg的管道，管道在有新内容时回返回输送出新的内容。*/
type Fetcher interface {
	GetNew(Marshaler) ([]Msg, error)
}
type Fetch struct {
	order   chan string
}
func (this *Fetch) FetchThread(f Fetcher,m *chan Msg,marshal Marshaler) {
	if m == nil {
		panic("FetchThread: *chan Msg is nil pointer")
	}
	for {
		tempChan := make(chan Msg)
		go func() {
			news, err := f.GetNew(marshal)
			if err != nil {
				panic(err)
			}
			for _, msg := range news {
				tempChan <- msg
			}
		}()
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic:",err)
				fmt.Printf("Restart FetchThread")
				this.FetchThread(f,m,marshal)
			}
		}()
		select {
		case msg := <-tempChan:
			*m <- msg
		case ord := <-this.order:
			if ord == "quit" {
				return
			}
		}
	}
}
