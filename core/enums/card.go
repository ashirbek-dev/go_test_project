package enums

const (
	CardStateUnknown         = 0
	CardStateActive          = 1
	CardStateExpired         = 2
	CardStateBlocked         = 3
	CardStateInvalid         = 4
	CardStateNotFound        = 5
	CardStateInactiveAccount = 6
	CardStateInactiveIssuer  = 7
	CardStateInactiveUser    = 8
)

const (
	CardSearchStateNew          = 0
	CardSearchStatePending      = 1
	CardSearchStateSuccess      = 2
	CardSearchStateError        = 3
	CardSearchStateNothingFound = 4
)

const (
	CardSourceUnknown                  = 0
	CardSourcePhone                    = 1
	CardSourcePinfl                    = 2
	CardSourceCustomerId               = 3
	CardSourceUnknownCheckedOwns       = 4
	CardSourcePhoneCheckedOwns         = 5
	CardSourceCustomerIdCheckedOwns    = 6
	CardSourceUnknownCheckedNotOwns    = 7
	CardSourcePhoneCheckedNotOwns      = 8
	CardSourceCustomerIdCheckedNotOwns = 9
	CardSourceGenesys                  = 10
	CardSourceUnknownCheckedNoPinfl    = 11
	CardSourcePhoneCheckedNoPinfl      = 12
	CardSourceCustomerIdCheckedNoPinfl = 13
)

const (
	CardPinflSearchResultStateUnknown      = 0
	CardPinflSearchResultStateNew          = 1
	CardPinflSearchResultStateAlreadyFound = 2
)
