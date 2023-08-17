package calculator

import (
	abstract "bernanda/learn/abstraction"
)

type CalculatorResponse struct {
	baseResponse abstract.BaseResponse
	result       interface{}
}

func RunCalculator() {
	start()

	// div := CalculatorResponse{}
	// div.result, div.baseResponse.Err, div.baseResponse.PanicMsg = Division(10, 0)

	// fmt.Println(div)
}
