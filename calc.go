package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/Knetic/govaluate"
)

func calc() {
	functions := map[string]govaluate.ExpressionFunction{
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
	}

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(os.Args[1], functions)

	parameters := make(map[string]interface{}, 8)
	parameters["epoch"] = time.Now().UnixNano()

	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < 3; i++ {
			parameters["i"] = i

			result, err := expression.Evaluate(parameters)

			fmt.Println(result, err)
		}
	}
}
