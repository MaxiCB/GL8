package cpu

import (
	"testing"
)

// TestCreateCPU calls CreateCPU, checking for a valid return value
func TestCreateCPU(t *testing.T) {
	want := CPU{[256]int{}, [8]int{}, 0, 0, 0, 0, [8]int{}, 0, 255, 0}
	created := CreateCPU()
	if want != created {
		t.Fatalf(`CreateCPU() = %q, want match for %v`, created, want)
	}
}

// TestReadRAM calls ReadRAM with a address, checking for the correct return value.
func TestReadRAM(t *testing.T) {
	cpu := CPU{[256]int{}, [8]int{}, 0, 0, 0, 0, [8]int{}, 0, 255, 0}
	want := 0
	created := cpu.ReadRAM(0)
	if want != created {
		t.Fatalf(`ReamRAM(0) = %q, want match for %v`, created, want)
	}
}

// TestWriteRAM calls WriteRAM with a address and data, writing to the corresponding address with the data provided
func TestWriteRAM(t *testing.T) {
	cpu := CPU{[256]int{}, [8]int{}, 0, 0, 0, 0, [8]int{}, 0, 255, 0}
	want := 2
	cpu.WriteRAM(0, 2)
	created := cpu.ReadRAM(0)
	if want != created {
		t.Fatalf(`ReamRAM(0) = %q, want match for %v`, created, want)
	}
}

// TestLoadProgram calls LoadProgram with the value "test", and check that the corresponding data is placed in memory
func TestLoadProgram(t *testing.T) {
	cpu := CPU{[256]int{}, [8]int{}, 0, 0, 0, 0, [8]int{}, 0, 255, 0}
	want := 10000010
	cpu.LoadProgram("test")
	created := cpu.ReadRAM(0)
	if want != created {
		t.Fatalf(`ReamRAM(0) = %q, want match for %v`, created, want)
	}
}

func TestLoadFullProgram(t *testing.T) {
	cpu := CPU{[256]int{}, [8]int{}, 0, 0, 0, 0, [8]int{}, 0, 255, 0}
	cpu.LoadProgram("test")
	for index, value := range cpu.RAM[0:6] {
		if cpu.ReadRAM(index) != int(value) {
			t.Fatalf(`Address %q = %v, want match for %#q`, index, value, cpu.ReadRAM(index))
		}
	}
}

// TestCPURun calls CPURun() and verifies the program was loaded into memory
func TestCPURun(t *testing.T) {
	cpu := CPU{[256]int{}, [8]int{}, 0, 0, 0, 0, [8]int{}, 0, 255, 0}
	cpu.LoadProgram("mult")
	cpu.CPURun()
}
