package calculator

import (
	"fmt"
)

func chainOperation(opsType string, numbers []float64) (result float64) {
	for i := 0; i < len(numbers); i++ {
		switch opsType {
		case "sum":
			result += numbers[i]
		case "product":
			result *= numbers[i]
		case "reduce":
			result -= numbers[i]
		case "divide":
			result /= numbers[i]
		default:
			return
		}

	}
	return
}

func sum(numbers ...float64) {
	chainOperation("sum", numbers)
}

func product(numbers ...float64) {
	chainOperation("product", numbers)
}

func reduce(numbers ...float64) {
	chainOperation("reduce", numbers)
}

func divide(numbers ...float64) {
	defer divisionByZero()
	chainOperation("divide", numbers)
}

func divisionByZero() {
	msg := recover()
	if msg != nil {
		fmt.Println(msg)
	}
}

// func isDivisionByZero(err *error, pm *interface{}) {
// 	msg := recover()
// 	if msg != nil {
// 		*err = errors.New("Oops Division by Zero")
// 		*pm = msg
// 	}
// }

// func Division(a, b int) (result int, err error, panic_msg interface{}) {
// 	defer isDivisionByZero(&err, &panic_msg)
// 	result = a / b
// 	return
// }
