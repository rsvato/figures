package main

import (
	"fmt"

	"github.com/rsvato/figures/figures"
)

func main() {
	desk := figures.NewDesk(8, 8)
	desk.Print()
	for t := range 64 {
		p, ok := desk.PlaceFigure(&figures.Knight{})
		if !ok {
			fmt.Printf("Total of %d figures placed\n", t)
			break
		}
		fmt.Printf("Placed figure %s\n", *p)
		desk.Print()
	}
}
