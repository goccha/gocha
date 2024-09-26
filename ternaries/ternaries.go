package ternaries

type F[T any] func() T

func Eval[T any](is bool, trueValue, falseValue T) T {
	if is {
		return trueValue
	}
	return falseValue
}

func EvalFunc[T any](is bool, trueFunc F[T], falseFunc F[T]) T {
	if is {
		return trueFunc()
	}
	return falseFunc()
}
