package enums

const (
	TrTypePayment  = 0
	TrTypeReversal = 1
)

const (
	TrStateNew                           = 0
	TrStateHoldingPayment                = 1
	TrStateSuccessHold                   = 2
	TrStateHoldConfirming                = 3
	TrStateHoldPaid                      = 4
	TrStateHoldingError                  = 5
	TrStatePaymentApproveServiceError    = 6
	TrStateHoldServiceError              = 7
	TrStateHoldInternalError             = 8
	TrStateHoldEmptyResponse             = 9
	TrStateInsufficientFunds             = 10
	TrStateHoldConnectionError           = 11
	TrStateHoldCancelled                 = 12
	TrStatePaymentApproveInternalError   = 13
	TrStatePaymentApproveEmptyResponse   = 14
	TrStatePaymentApproveConnectionError = 15
)

// ReversalState
//  0 - no reversal
// -1 - pending reversal
// -2 - success reversal
// -3 - reversal error

const (
	TrRevStateDefault         = 0
	TrRevStatePending         = -1
	TrRevStateSuccess         = -2
	TrRevStateError           = -3
	TrRevStateServiceError    = -4
	TrRevStateConnectionError = -5
	TrRevStateEmptyResponse   = -6
)
