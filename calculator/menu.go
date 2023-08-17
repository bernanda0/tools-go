package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Knetic/govaluate"
)

func start() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Welcome, input basic math operation. 'q' to exit!")

	for {
		fmt.Print("\n# ")
		input, _ := reader.ReadString('\n')

		// Remove trailing newline character
		input = strings.TrimSpace(input)

		if input == "q" {
			fmt.Print("Exiting calculator...")
			break
		}

		expression, err := govaluate.NewEvaluableExpression(input)
		if err != nil {
			fmt.Print("Invalid expression:", err)
			continue
		}

		result, err := expression.Evaluate(nil)
		if err != nil {
			fmt.Print("Error evaluating expression:", err)
			continue
		}

		fmt.Print("  = ", result)
	}

	fmt.Println()
}
