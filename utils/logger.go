package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintInfo(to_print string) {
	fmt.Printf("%v %v\n", color.CyanString("►"), to_print)
}

func PrintSuccess(to_print string) {
	fmt.Printf("%v %v\n", color.CyanString("✔"), to_print)
}

func PrintError(to_print string) {
	fmt.Printf("%v %v\n", color.RedString("✘"), to_print)
}
