package main

import (
	"github.com/abiosoft/ishell"
	"strconv"
	"fmt"
	"ara/color"
)

func initShell(hsv *color.Hsv) *ishell.Shell {
	shell := ishell.New()

	shell.AddCmd(&ishell.Cmd {
		Name: "hue",
		Help: "Set hue channel. Usage: hue [float 0-360...]",
		Func: func(c *ishell.Context) {
			if result, err := strconv.ParseFloat(c.Args[0], 64); err == nil {
				hsv.Hue = result
			} else {
				fmt.Println(err)
			}
		},
	})

	shell.AddCmd(&ishell.Cmd {
		Name: "saturation",
		Help: "Set saturation channel. Usage: hue [float 0-360...]",
		Func: func(c *ishell.Context) {
			if result, err := strconv.ParseFloat(c.Args[0], 64); err == nil {
				hsv.Saturation = result
			} else {
				fmt.Println(err)
			}
		},
	})

	shell.AddCmd(&ishell.Cmd {
		Name: "value",
		Help: "Set value channel. Usage: hue [float 0-360...]",
		Func: func(c *ishell.Context) {
			if result, err := strconv.ParseFloat(c.Args[0], 64); err == nil {
				hsv.Value = result
			} else {
				fmt.Println(err)
			}
		},
	})

	shell.Run()

	return shell
}