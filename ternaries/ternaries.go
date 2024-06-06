package ternaries

func Eval[T any](is bool, trueValue, falseValue T) T {
	if is {
		return trueValue
	}
	return falseValue
}
