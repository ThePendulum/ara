package main

import (
	"github.com/abiosoft/ishell"
	"strconv"
	"fmt"
)

func initShell() *ishell.Shell {
	shell := ishell.New()

	shell.AddCmd(&ishell.Cmd {
		Name: "hue",
		Help: "Set hue channel. Usage: hue [float 0-360...]",
		Func: setHue,
	})

	shell.AddCmd(&ishell.Cmd {
		Name: "saturation",
		Help: "Set saturation channel. Usage: hue [float 0-360...]",
		Func: setSaturation,
	})

	shell.AddCmd(&ishell.Cmd {
		Name: "value",
		Help: "Set value channel. Usage: hue [float 0-360...]",
		Func: setValue,
	})

	shell.Run()

	return shell
}

func setHue(c *ishell.Context) {
	if result, err := strconv.ParseFloat(c.Args[0], 64); err == nil {
		hsv.Hue = result
	} else {
		fmt.Println(err)
	}
}

func setSaturation(c *ishell.Context) {
	if result, err := strconv.ParseFloat(c.Args[0], 64); err == nil {
		hsv.Saturation = result
	} else {
		fmt.Println(err)
	}
}

func setValue(c *ishell.Context) {
	if result, err := strconv.ParseFloat(c.Args[0], 64); err == nil {
		hsv.Value = result
	} else {
		fmt.Println(err)
	}
}