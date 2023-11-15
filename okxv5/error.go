package okxv5

import (
	"fmt"

	"github.com/msw-x/moon/ujson"
	"golang.org/x/exp/slices"
)

type Error struct {
	Code ujson.Int64
	Text string
}

func (o *Error) Std() error {
	if o.Empty() {
		return nil
	} else {
		return o
	}
}

func (o *Error) Empty() bool {
	// 1 added, as PlaceOrder response contains info inside
	return o.Code == 0 || o.Code == 1
}

func (o *Error) Error() string {
	return fmt.Sprintf("code[%d]: %s", o.Code.Value(), o.Text)
}

// func (o *Error) RequestParameterError() bool {
// 	return o.Code ==
// }

func (o *Error) ApiKeyInvalid() bool {
	codes := []ujson.Int64{
		50100, // API frozen, please contact customer service.
		50101, // APIKey does not match current environment.
		50111, // Invalid OK-ACCESS-KEY.
		50119, // API key doesn't exist

	}
	return slices.Contains(codes, o.Code)
}

// func (o *Error) ApiKeyExpired() bool {
// 	return o.Code ==
// }

// func (o *Error) TooManyVisits() bool {
// 	return o.Code ==
// }

func (o *Error) UnmatchedIp() bool {
	return o.Code == 50110
}

func (o *Error) SymbolIsNotWhitelisted() bool {
	return o.Code == 51001 // Instrument ID does not exist.
}

func (o *Error) KycNeeded() bool {
	codes := []ujson.Int64{
		51732, // Required user KYC level not met
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) Timeout() bool {
	codes := []ujson.Int64{
		50004, // API endpoint request timeout (does not mean that the request was successful or failed, please check the request result).
		52912, // Server timeout
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) InsufficientBalance() bool {
	codes := []ujson.Int64{
		51008, // Order failed. Insufficient {param0}...
		51131, // Insufficient balance
		51502, // Order failed. Insufficient {param0}...
		51736, // Insufficient {ccy} balance
		52916, // Insufficient balance in funding account
		58229, // Insufficient funding account balance to pay fees {fee} USDT
		58350, // Insufficient balance.
		58372, // Insufficient small assets.
		59103, // Account margin is insufficient and leverage is too low. Please increase the leverage.
		59108, // Your account leverage is too low and has insufficient margins. Please increase the leverage.
		59200, // Insufficient account balance.
		59303, // Insufficient available margin, add margin or reduce the borrowing amount
		59304, // Insufficient equity for borrowing. Keep enough funds to pay interest for at least one day.
	}
	return slices.Contains(codes, o.Code)
}
