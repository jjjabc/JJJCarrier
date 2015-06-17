package JJJCarrier

import (
	"sort"
	"errors"
	"time"
)

var ErrNotSupport = errors.New("JJJCarrier: function not support")
var ErrNilPointer = errors.New("JJJCarrier: fointer is nill")

type Msg interface {
	String() string
	Time() time.Time
	GetContext() interface{}
}
type MsgSlice []Msg

func (this MsgSlice) Len() int           { return len(this) }
func (this MsgSlice) Less(i, j int) bool { return this[i].Time().Before(this[j].Time()) }
func (this MsgSlice) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }

func SortByTimeInc(messages []Msg) {
	sort.Sort(sort.Reverse(MsgSlice(messages)))
}
func GetNewMsg(curs []Msg, olds []Msg) (r []Msg, err error) {
	if olds == nil {
		return curs, nil
	}
	if curs==nil{
		return nil,ErrNilPointer;
	}
	for _,cur:=range curs{
		isNew:=true
		for _,old:=range olds{
			if cur==old{
				isNew=false
				break
			}
		}
		if isNew{
			r=append(r,cur)
		}
	}
	SortByTimeInc(r)
	return
}