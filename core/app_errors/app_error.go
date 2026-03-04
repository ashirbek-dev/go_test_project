package app_errors

import (
	"fmt"
)

type ApplicationError struct {
	code      int
	errorText string
}

func (e ApplicationError) Error() string {
	return e.errorText
}

func (e ApplicationError) Code() int {
	return e.code
}

func (e ApplicationError) Is(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(ApplicationError)
	if ok {
		return true
	}
	_, ok = err.(*ApplicationError)
	if ok {
		return true
	}
	return false
}

func NewAppErr(code int, errorDetails ...string) *ApplicationError {
	err := &ApplicationError{code: code, errorText: getCodeText(code)}
	for _, detail := range errorDetails {
		err.errorText = fmt.Sprintf("%s. %s", err.errorText, detail)
	}
	return err
}

func NewFromErr(code int, errors ...error) *ApplicationError {
	err := &ApplicationError{code: code, errorText: getCodeText(code)}
	for _, errorItem := range errors {
		(*err).errorText += fmt.Sprintf("%s. %s", err.errorText, errorItem.Error())
	}
	return err
}

func getCodeText(code int) string {
	switch code {
	case ENTITY_NOT_FOUND:
		return "entity not found"
	case PAN_INVALID:
		return "invalid card pan"
	case EXPIRY_INVALID:
		return "invalid card expiry"
	case TR_ALREADY_EXISTS:
		return "transaction is already exists"
	case EMPTY_RESULT:
		return "result is empty"
	case TR_NOT_PAID:
		return "transaction is not paid"
	case TR_ALREADY_IN_REVERSAL_STATE:
		return "transaction is already in reversal state"
	case BALANCE_LOWER_THEN_MIN_AMOUNT:
		return "balance is lower then min amount"
	case AMOUNT_IS_ZERO:
		return "payment amount is zero"
	case ENTITY_ALREADY_EXISTS:
		return "entity already exists"
	case VALUE_IS_NULL_OR_EMPTY:
		return "value is null or empty"
	case UNTRUSTED_CARD:
		return "untrusted card"
	case TR_ERROR:
		return "transaction error"
	case INVALID_FORMAT:
		return "invalid format"
	case CARD_BLOCKED:
		return "card is blocked"
	case CARD_NOT_FOUND:
		return "card is not found"
	case SERVICE_OPERATION_ERROR:
		return "service operation error"
	case CLIENT_NOT_FOUND:
		return "client not found"
	case CARD_NOT_BELONG_TO_CLIENT:
		return "card is not belong to client"
	case EXT_SRV_CONN_ERR:
		return "external service connection error"
	case EXT_SRV_EMPTY_RESULT:
		return "external service empty result"
	case TR_NOT_CANCELLABLE:
		return "transaction can not be cancelled"
	case TR_NOT_APPROVABLE:
		return "transaction can not be approved"
	case TR_NOT_RETURNABLE:
		return "transaction can not be returned"
	case INVALID_STATE:
		return "entity is in invalid state"
	case CARD_EXPIRED:
		return "card is expired"
	case INSUFFICIENT_FUNDS:
		return "insufficient funds"
	case PS_CARD_IS_BLOCKED_BY_PIN:
		return "card is blocked by pin"
	case PS_CARD_IS_EXPIRED_OR_NOT_FOUND:
		return "card is expired or not found"
	case PS_CARD_LIMIT_FULL:
		return "card limit exceed"
	case PS_CARD_NOT_FOR_EPOS:
		return "card is not for this e-pos"
	case PS_CARD_WITHOUT_PIN:
		return "card is without pin"
	case PS_CARD_PIN_COUNT_BLOCK:
		return "pin count block"
	case PS_CARD_LOST:
		return "card is lost"
	case PS_CARD_STOLEN:
		return "card is stolen"
	case PS_CARD_TEMP_BLOCKED:
		return "card is temp. blocked"
	case PS_CARD_TEMP_BLOCKED_2:
		return "card is temp. blocked v2"
	case PS_CARD_LIMIT_EXCEED:
		return "card limit exceed v2"
	case PS_CARD_TEMP_BLOCKED_3:
		return "card is temp. blocked v3"
	case PS_CARD_TEMP_BLOCKED_4:
		return "card is temp. blocked v4"
	case PS_EMITENT_SPECIAL_CONDITIONS:
		return "emitent special conditions"
	case HTTP_REQUEST_ERROR:
		return "http request error"
	case HTTP_REQUEST_ERROR_404:
		return "http request 404 error"
	case HTTP_RESPONSE_PARSING_ERROR:
		return "http response parsing error"
	case HTTP_UNEXPECTED_ERROR:
		return "http unexpected error"
	default:
		return ""
	}
}

const (
	ENTITY_NOT_FOUND                = -1
	PAN_INVALID                     = -2
	EXPIRY_INVALID                  = -3
	TR_ALREADY_EXISTS               = -4
	EMPTY_RESULT                    = -5
	TR_NOT_PAID                     = -6
	TR_ALREADY_IN_REVERSAL_STATE    = -7
	BALANCE_LOWER_THEN_MIN_AMOUNT   = -8
	AMOUNT_IS_ZERO                  = -9
	ENTITY_ALREADY_EXISTS           = -10
	VALUE_IS_NULL_OR_EMPTY          = -11
	UNTRUSTED_CARD                  = -12
	TR_ERROR                        = -13
	INVALID_FORMAT                  = -14
	CARD_BLOCKED                    = -15
	CARD_NOT_FOUND                  = -16
	SERVICE_OPERATION_ERROR         = -17
	CLIENT_NOT_FOUND                = -18
	CARD_NOT_BELONG_TO_CLIENT       = -19
	EXT_SRV_CONN_ERR                = -20
	EXT_SRV_EMPTY_RESULT            = -21
	TR_NOT_CANCELLABLE              = -22
	TR_NOT_APPROVABLE               = -23
	TR_NOT_RETURNABLE               = -24
	INVALID_STATE                   = -25
	CARD_EXPIRED                    = -26
	INSUFFICIENT_FUNDS              = -50
	PS_CARD_IS_BLOCKED_BY_PIN       = -200
	PS_CARD_IS_EXPIRED_OR_NOT_FOUND = -201
	PS_CARD_LIMIT_FULL              = -202
	PS_CARD_NOT_FOR_EPOS            = -203
	PS_CARD_WITHOUT_PIN             = -204
	PS_CARD_PIN_COUNT_BLOCK         = -205
	PS_CARD_LOST                    = -206
	PS_CARD_STOLEN                  = -207
	PS_CARD_TEMP_BLOCKED            = -208
	PS_CARD_TEMP_BLOCKED_2          = -209
	PS_CARD_LIMIT_EXCEED            = -210
	PS_CARD_TEMP_BLOCKED_3          = -211
	PS_CARD_TEMP_BLOCKED_4          = -212
	PS_EMITENT_SPECIAL_CONDITIONS   = -213
	HTTP_REQUEST_ERROR              = -1001
	HTTP_REQUEST_ERROR_404          = -1002
	HTTP_RESPONSE_PARSING_ERROR     = -1003
	HTTP_UNEXPECTED_ERROR           = -1004
	PROVIDER_NOT_FOUND              = -1101
)
