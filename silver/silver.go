package silver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type Silver struct {
	T     time.Time `json:"time"`
	Price float32   `json:"price"`
}
type JsonSilver struct {
	Time  string  `json:"time"`
	Price float32 `json:"price"`
}

func (c Silver) String() string {
	return strconv.FormatFloat(float64(c.Price), 'f', 3, 32)
}
func (c Silver) Time() time.Time {
	return c.T
}
func (c Silver) GetContext() interface{} {
	return c.Price
}
func silverParse(raw string) (index Silver) {
	var err error
	var price float64
	stime := silverRaw2slice(`dataCell.cell0 = "(.*)";dataCell.cell1 =`, raw)
	sprice := silverRaw2slice(`dataCell.cell1 = "(.*)";dataCell.cell2 =`, raw)
	index.T, err = time.Parse("2006-01-02 15:04:05", stime[0])
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	price, err = strconv.ParseFloat(sprice[0], 32)
	if err != nil {
		fmt.Printf(err.Error())
	}
	index.Price = float32(price)
	return
}
func silverParseSlice(s []string) []Silver {
	var ss []Silver
	for _, key := range s {
		ss = append(ss, silverParse(key))
	}
	return ss
}
func silverRegexpSubmatch(s string, content string) [][]string {
	rx := regexp.MustCompile(s)
	return rx.FindAllStringSubmatch(content, -1)

}
func silverRaw2slice(regxp string, content string) []string {
	s := silverRegexpSubmatch(regxp, content)
	if s == nil {
		return nil
	}
	var slice []string
	for _, key := range s {
		slice = append(slice, key[1])
	}
	return slice
}
func silverDataTidy(date []string) []string {
	var tidyeddate []string
	for _, s := range date {
		isOnly := true
		for _, t := range tidyeddate {
			if s == t {
				isOnly = false
				break
			}
		}
		if isOnly {
			tidyeddate = append(tidyeddate, s)
		}
	}
	return tidyeddate
}
func getHtml(url string) (content string, statusCode int) {
	resp, err := http.Get(url)
	if err != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}
