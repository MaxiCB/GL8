package main

import (
	"aaroncb.com/cpu"
)

func main() {
	// # Initialization of the CPU
	newCPU := cpu.CreateCPU()
	newCPU.LoadProgram("test")
}
