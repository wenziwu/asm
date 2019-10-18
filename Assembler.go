package asm

import (
	"bytes"
	"encoding/binary"

	"github.com/akyoto/asm/stringtable"
	"github.com/akyoto/asm/syscall"
	"github.com/akyoto/asm/utils"
)

type Assembler struct {
	bytes.Buffer
	StringTable     *stringtable.StringTable
	SectionPointers []utils.Pointer
}

func New() *Assembler {
	return &Assembler{
		StringTable: stringtable.New(),
	}
}

func (a *Assembler) AddString(msg string) int64 {
	address := a.StringTable.Add(msg)

	a.SectionPointers = append(a.SectionPointers, utils.Pointer{
		Address:  address,
		Position: a.Len(),
	})

	return address
}

func (a *Assembler) WriteBytes(someBytes ...byte) {
	for _, b := range someBytes {
		a.WriteByte(b)
	}
}

func (a *Assembler) Mov(registerName string, num interface{}) {
	baseCode := byte(0xb8)
	registerID, exists := registerIDs[registerName]

	if !exists {
		panic("Unknown register name: " + registerName)
	}

	switch num.(type) {
	case string, int64:
		a.WriteByte(REX(1, 0, 0, 0))
	}

	a.WriteByte(baseCode + registerID)

	switch v := num.(type) {
	case string:
		_ = binary.Write(a, binary.LittleEndian, a.AddString(v))
	default:
		_ = binary.Write(a, binary.LittleEndian, num)
	}
}

func (a *Assembler) Syscall(parameters ...interface{}) {
	for count, parameter := range parameters {
		a.Mov(syscall.Registers[count], parameter)
	}

	a.WriteBytes(0x0f, 0x05)
}

func (a *Assembler) Println(msg string) {
	a.Print(msg + "\n")
}

func (a *Assembler) Print(msg string) {
	a.Syscall(syscall.Write, int32(1), msg, int32(len(msg)))
}

func (a *Assembler) Exit(code int32) {
	a.Syscall(syscall.Exit, code)
}
