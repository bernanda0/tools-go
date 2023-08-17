package main

import (
	abstract "bernanda/learn/abstraction"
	"bernanda/learn/calculator"
	"fmt"
)

type CalculatorResponse struct {
	baseResponse abstract.BaseResponse
	result       interface{}
}

func runCalculator() {
	fmt.Println(calculator.SumAll(12, 12, 12, 12))

	div := CalculatorResponse{}
	div.result, div.baseResponse.Err, div.baseResponse.PanicMsg = calculator.Division(10, 0)

	fmt.Println(div)
}
