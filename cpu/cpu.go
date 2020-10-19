package cpu

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// CPU Struct Definition
type CPU struct {
	// Initialize RAM as 256 bytes
	// This need's converted into a map
	// A Map would allow for on the fly conversion to the decimal representation
	// As well as allow us to refer to address' in a similar manner as before
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
	// CALL Instruction Pointer
	cache int
}

// CreateCPU Initializes a CPU instance
func CreateCPU() CPU {
	return CPU{[256]int{}, [8]int{}, 0, 0, 0, 0, [8]int{}, 0, 255, 0}
}

func convertBinaryToDecimal(number int) int {
	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2, counter))
		number = number / 10
		counter++
	}
	return decimal
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

// LoadProgram takes in a file-name and loads the file content into memory
func (c *CPU) LoadProgram(name string) {
	address := 0
	file, err := os.Open(`programs/` + name + `.gl8`)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cleaned := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(cleaned, "#") || strings.HasPrefix(cleaned, "//") {
		} else {
			// Need to store the raw string in RAM
			// This would allow for accessing, as well as being able to reference other address'
			conv, _ := strconv.Atoi(cleaned)
			c.WriteRAM(address, conv)
			address++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

// LDI Stores the next two instruction inside memory
func (c *CPU) LDI() {
	// Require's Next 2 Address
	inst0 := c.RAM[c.ir+1]
	inst1 := c.RAM[c.ir+2]
	c.reg[inst0] = inst1
	c.iter = 3
}

// PRINT Print's the data at a given register
func (c *CPU) PRINT() {
	inst1 := c.RAM[c.ir+1]
	c.iter = 2
	converted := convertBinaryToDecimal(c.reg[inst1])
	fmt.Println(converted)
}

// MULT Multiply's the data in two register's together
func (c *CPU) MULT() {
	inst1 := c.RAM[c.ir+1]
	inst2 := c.RAM[c.ir+2]
	fmt.Println(inst1)
	fmt.Println(inst2)
	c.reg[inst1] = convertBinaryToDecimal(inst1) * convertBinaryToDecimal(inst2)
	c.iter = 3
}

// OPSwitch Branches based on the instruction value given
func (c *CPU) OPSwitch(instruction int) {
	switch instruction {
	case 10000010:
		c.LDI()
	case 1000111:
		c.PRINT()
	case 10100010:
		c.MULT()
	case 10100000:
		fmt.Println("Add")
	case 1000101:
		fmt.Println("Push")
	case 1000110:
		fmt.Println("Pop")
	case 1010000:
		fmt.Println("Call")
	case 10001:
		fmt.Println("Return")
	case 10100111:
		fmt.Println("Compare")
	case 1010100:
		fmt.Println("Jump")
	case 1010101:
		fmt.Println("Jump Equal")
	case 1010110:
		fmt.Println("Jump Not Equal")
	case 10000100:
		fmt.Println("ST")
	case 11111111:
		fmt.Println("Histo")
	}
}

// CPURun Run's the internally stored program
func (c *CPU) CPURun() {
	for _, value := range c.RAM {
		if value == 1 && c.iter <= 0 {
			fmt.Println("TERMINATING")
			break
		} else {
			// fmt.Println(value)
			c.OPSwitch(value)
			c.iter--
		}
		c.ir++
	}
}
