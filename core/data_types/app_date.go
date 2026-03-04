package data_types

import (
	"fmt"
	"strings"
	"time"
)

type AppDate struct {
	time.Time
}

const dLayout = "02.01.2006"

func (ct *AppDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(dLayout, s)
	return
}

func (ct *AppDate) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(dLayout))), nil
}

func (ct *AppDate) IsSet() bool {
	return !ct.IsZero()
}
