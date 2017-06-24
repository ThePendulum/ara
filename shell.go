package main

import (
	"github.com/abiosoft/ishell"
	"ara/color"
	"log/syslog"
	"strings"
)

func initShell(logger *syslog.Writer, hsv *color.Hsv) {
	shell := ishell.New()

	shell.AddCmd(&ishell.Cmd {
		Name: "hue",
		Help: "Set hue channel. Usage: hue [float 0-360...]",
		Func: func(c *ishell.Context) {
			expression := strings.Join(c.Args, " ")

			if hueFn, err := compile(expression); err == nil {
				hsv.Hue = hueFn
			}
		},
	})

	shell.AddCmd(&ishell.Cmd {
		Name: "saturation",
		Help: "Set saturation channel. Usage: hue [float 0-360...]",
		Func: func(c *ishell.Context) {
			expression := strings.Join(c.Args, " ")

			if saturationFn, err := compile(expression); err == nil {
				hsv.Saturation = saturationFn
			}
		},
	})

	shell.AddCmd(&ishell.Cmd {
		Name: "value",
		Help: "Set value channel. Usage: hue [float 0-360...]",
		Func: func(c *ishell.Context) {
			expression := strings.Join(c.Args, " ")

			if valueFn, err := compile(expression); err == nil {
				hsv.Value = valueFn
			}
		},
	})

	shell.Println("Welcome to Ara!")
	shell.Run()
}
