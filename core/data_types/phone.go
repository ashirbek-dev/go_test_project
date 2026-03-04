package data_types

import (
	"fmt"
	"gateway/core/app_errors"
	"gateway/core/utils"
	"strings"
	"unicode/utf8"
)

type Phone struct {
	PhoneNumber     string
	phoneCleared    string
	phoneWithPrefix string
	phoneNoPrefix   string
}

func CreatePhone(phone string) (*Phone, error) {

	phoneCleared := utils.StrLeaveOnlyDigits(phone)

	length := utf8.RuneCountInString(phoneCleared)

	if length < 9 {
		return nil, app_errors.NewAppErr(app_errors.INVALID_FORMAT, "phone length less then 9")
	}

	if length > 12 {
		return nil, app_errors.NewAppErr(app_errors.INVALID_FORMAT, "phone length more then 12")
	}

	var phoneWithPrefix string
	var phoneNoPrefix string

	if length > 9 {
		phoneNoPrefix = phoneCleared[len(phoneCleared)-9:]
		if length == 12 {
			prefix := phoneCleared[:3]
			if prefix == "000" {
				phoneWithPrefix = fmt.Sprintf(`998%s`, phoneNoPrefix)
			} else {
				phoneWithPrefix = phoneCleared
			}
		} else {
			return nil, app_errors.NewAppErr(app_errors.INVALID_FORMAT, "phone is invalid")
		}
	} else {
		phoneWithPrefix = fmt.Sprintf(`998%s`, phoneCleared)
		phoneNoPrefix = phoneCleared
	}

	return &Phone{
		PhoneNumber:     phoneWithPrefix,
		phoneCleared:    phoneCleared,
		phoneWithPrefix: phoneWithPrefix,
		phoneNoPrefix:   phoneNoPrefix,
	}, nil
}

func (ct *Phone) GetFullNumber() string {
	return ct.phoneWithPrefix
}

func (ct *Phone) GetShortNumber() string {
	return ct.phoneNoPrefix
}

func (ct *Phone) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), " ")
	s = strings.Trim(s, "+")
	if s == "null" {
		ct.PhoneNumber = ""
		return
	}
	return
}
func (ct *Phone) MarshalJSON() ([]byte, error) {
	if ct.PhoneNumber == "" {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.PhoneNumber)), nil
}
