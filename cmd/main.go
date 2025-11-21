package main

import (
	"fmt"
	"os"

	"github.com/MaksimBBB/maks-GoJunior-cli/converter"
	"github.com/MaksimBBB/maks-GoJunior-cli/ui"
)

func main() {
	fmt.Println("Ласкаво просимо до PDF Converter!")

	conv := converter.NewConverter()
	if err := ui.RunInteractive(conv); err != nil {
		fmt.Fprintln(os.Stderr, "Помилка:", err)
		os.Exit(1)
	}
}
