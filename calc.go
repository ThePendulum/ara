package main

import (
	"github.com/yuin/gopher-lua"
	"fmt"
)

var L = lua.NewState()

func compile(expression string) (*lua.LFunction, error) {
	err := L.DoString(fmt.Sprintf(`
		for k, v in pairs (math) do _G[k] = v end

		return function (beat, index)
			return (%s)
		end
	`, expression))

	if err != nil {
		fmt.Println(err)
	}

	fn, _ := L.Get(-1).(*lua.LFunction)

	L.Pop(1)

	if fn == nil {
		return nil, fmt.Errorf("Express did not result in a function: %q", expression)
	}

	return fn, nil
}

func call(fn *lua.LFunction, context map[string]int) float64 {
	defer L.SetTop(L.GetTop())

	L.Push(fn)

	for _, value := range context {
		L.Push(lua.LNumber(value))
	}

	if err := L.PCall(3, 1, nil); err != nil {
		return 0
	}

	num, _ := L.Get(-1).(lua.LNumber)

	return float64(num)
}
