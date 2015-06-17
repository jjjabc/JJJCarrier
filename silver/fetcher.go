package silver

import (
	"time"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/jjjabc/JJJCarrier"
)

type SilverFetcher struct {
	sPrices []JJJCarrier.Msg
	addr    url.URL
	marshal JJJCarrier.Marshaler
	intervalMS time.Duration
}
type silverMarshal struct {
}
func (this *SilverFetcher) GetIntMs() time.Duration {
	return this.intervalMS
}
func (this *SilverFetcher) SetIntMs(ms time.Duration) {
	this.intervalMS = ms
}
func (this *SilverFetcher) GetPrices() []JJJCarrier.Msg {
	return this.sPrices
}
func (this *SilverFetcher) GetNew(marshal JJJCarrier.Marshaler) ([]JJJCarrier.Msg, error) {
	if this.addr.Scheme != "http" {
		return nil, errors.New("SliverFetcher: URL isn't http")
	}
	this.marshal = marshal
	for {
		slice, err := this.fetchMsg(this.addr.String())
		if err != nil {
			fmt.Printf("%s-continue", err.Error())
			continue
		}
		newMsgs, err := JJJCarrier.GetNewMsg(slice, this.sPrices)
		if err != nil {
			return nil, err
		}
		if newMsgs != nil {
			return newMsgs, nil
		}

	}
}
func (this *SilverFetcher) fetchMsg(url string) ([]JJJCarrier.Msg, error) {
	code, state := getHtml(url)
	if state != 200 {
		return nil, errors.New("HttpFetch: http state return " + strconv.Itoa(state))
	}
	slice, err := this.marshal.Marshal(code)
	if err != nil {
		return nil, err
	} else {
		return slice, nil
	}
}
func (this silverMarshal) Marshal(v interface{}) ([]JJJCarrier.Msg, error) {
	rawRegxpString := `(dataCell.cell0.*)dataObjs\[`
	rawdate := silverRaw2slice(rawRegxpString, v.(string))
	if rawdate == nil {
		return nil, errors.New("Marshal error:Raw2Slice return nil!")
	}
	date := silverDataTidy(rawdate)
	ss := silverParseSlice(date)
	r := make([]JJJCarrier.Msg, len(ss))
	for i, v := range ss {
		r[i] = v
	}
	return r, nil
}
