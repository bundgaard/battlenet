package types

import "time"

type Image struct {
	Name      string
	Data      []byte
	CreatedAt time.Time
}
