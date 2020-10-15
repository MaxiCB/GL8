package main

import (
	"fmt"

	"aaroncb.com/cpu"
)

func main() {
	// # Initialization of the CPU
	newCPU := cpu.CreateCPU()
	fmt.Println(newCPU.RAM)
	fmt.Println(newCPU.ReadRAM(2))
}
