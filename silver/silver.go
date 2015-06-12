package silver

import (
	"time"
)

type Silver struct {
	T     time.Time `json:"time"`
	Price float32   `json:"price"`
}