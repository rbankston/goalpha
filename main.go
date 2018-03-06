package main

import (
	"flag"
	"fmt"
)

func main() {

	phasePtr := flag.Int("phase", 1, "Phase of the program you are on: Choose 1, 2, 3")
	weightPtr := flag.Int("weight", 250, "Your current weight in whole numbers")
	fatPtr := flag.Int("fat", 20, "Current Body Fat percentage in whole numbers")

	flag.Parse()

	fmt.Println("Phase:", *phasePtr)
	fmt.Println("Weight:", *weightPtr)
	fmt.Println("Fat Percentage:", *fatPtr)

}
