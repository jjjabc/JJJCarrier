package transfer

type Notifyer interface{
	Notify(x interface{})error
}