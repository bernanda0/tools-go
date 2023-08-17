package calculator

import (
	"errors"
)

func SumAll(numbers ...int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return
}

func isDivisionByZero(err *error, pm *interface{}) {
	msg := recover()
	if msg != nil {
		*err = errors.New("Oops Division by Zero")
		*pm = msg
	}
}

func Division(a, b int) (result int, err error, panic_msg interface{}) {
	defer isDivisionByZero(&err, &panic_msg)
	result = a / b
	return
}
