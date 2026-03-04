package utils

type Bus interface {
	Send(action string, data ...any)
}

const (
	TransactionLog      = "payment:tr_log"
	APPaidAmountChanged = "payment:ap_paid_amount_changed"
	TransactionInfo     = "payment:tr_info"
	GenCardRating       = "payment:gen_card_rating"
	PspCardRating       = "payment:psp_card_rating"
)
