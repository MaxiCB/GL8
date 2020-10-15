package cpu

// CPU Struct Definition
type CPU struct {
	// Initialize RAM as 256 bytes
	RAM [256]int
	// # 8 general-purpose 8-bit numeric registers R0-R7.
	//     # * R5 is reserved as the interrupt mask (IM)
	//     # * R6 is reserved as the interrupt status (IS)
	//     # * R7 is reserved as the stack pointer (SP)
	reg [8]int
	// Memory Address Register
	// Current address being written/read
	mar int
	// Memory Data Register
	// Data from current read or write register
	mdr int
	// Program Counter
	pc int
	ir int
	// Flag Register
	fl [8]int
	// Instruction Iterator Pointer
	iter int
	// Stack Pointer
	sp int
	// CALL Instruction Pointe
	cache int
}

// CreateCPU Initializes a CPU instance
func CreateCPU() CPU {
	return CPU{[256]int{}, [8]int{}, 0, 0, 0, 0, [8]int{}, 0, 255, 0}
}

// ReadRAM Takes in a address integer && returns the data at that address in RAM
func (c *CPU) ReadRAM(address int) int {
	data := c.RAM[address]
	return data
}

// WriteRAM Takes in a address integer, and data value and stores it in the corresponding RAM address
func (c *CPU) WriteRAM(address int, data int) {
	c.RAM[address] = data
	return
}
