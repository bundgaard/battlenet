// Copyright (C) 2021 David Bundgaard
package types

import (
	"encoding/json"
	"strconv"
	"time"
)

// ProfileTime ...
type ProfileTime struct {
	time.Time
}

// UnmarshalJSON ...
func (pt *ProfileTime) UnmarshalJSON(data []byte) error {
	m, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	t := time.Unix(m, 0)
	pt.Time = t
	return nil
}

// MarshalJSON back to epoc
func (pt ProfileTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(pt.Unix())
}
