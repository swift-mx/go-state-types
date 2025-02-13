package init

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
)

var Methods = map[uint64]builtin.MethodMeta{
	1: {"Constructor", *new(func(*ConstructorParams) *abi.EmptyValue)}, // Constructor
	2: {"Exec", *new(func(*ExecParams) *ExecReturn)},                   // Exec
	3: {"Exec4", *new(func(*Exec4Params) *ExecReturn)},                 // Exec4
}
