package JJJCarrier

type Sender interface {
	Send(Msg) error
}

type Client interface {
	PushTheard() error
	PullTheard() (interface{}, error)
	Sender
	String() string
	Close() error
}
