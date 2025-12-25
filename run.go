package main

import (
	"fmt"

	"github.com/rsvato/figures/figures"
)

func main() {
	desk := figures.NewDesk(8, 8)
	desk.Print()
	for t := range desk.X * desk.Y {
		p, ok := desk.PlaceFigure(&figures.Queen{})
		if !ok {
			fmt.Printf("Total of %d figures placed\n", t)
			break
		}
		fmt.Printf("Placed figure %s\n", *p)
		desk.Print()
	}
}
