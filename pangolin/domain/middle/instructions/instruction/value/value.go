package value

import (
	var_value "github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value"
)

type value struct {
	val var_value.Value
	op  Operation
}

func createValue(val var_value.Value, op Operation) Value {
	out := value{
		val: val,
		op:  op,
	}

	return &out
}

// Value returns the value
func (obj *value) Value() var_value.Value {
	return obj.val
}

// Operation returns the operation
func (obj *value) Operation() Operation {
	return obj.op
}
