package account

import (
	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
)

var Methods = map[uint64]builtin.MethodMeta{
	1: {"Constructor", *new(func(*address.Address) *abi.EmptyValue)},                   // Constructor
	2: {"PubkeyAddress", *new(func(*abi.EmptyValue) *address.Address)},                 // PubkeyAddress
	3: {"AuthenticateMessage", *new(func(*AuthenticateMessageParams) *abi.EmptyValue)}, // AuthenticateMessage
	uint64(builtin.UniversalReceiverHookMethodNum): {"UniversalReceiverHook", *new(func(*abi.CborBytesTransparent) *abi.EmptyValue)}, // UniversalReceiverHook
}
