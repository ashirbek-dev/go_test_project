package data_types

import (
	"fmt"
	"strings"
	"time"
)

type AppTime struct {
	time.Time
}

const tLayout = "02.01.2006 15:04:05"

func (ct *AppTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(tLayout, s)
	return
}

func (ct *AppTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(tLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *AppTime) IsSet() bool {
	return !ct.IsZero()
}
