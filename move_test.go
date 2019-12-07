package asm_test

import (
	"math"
	"testing"

	"github.com/akyoto/asm"
	"github.com/akyoto/assert"
)

func TestMoveRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Register string
		Number   int64
		Code     []byte
	}{
		{"rax", 1, []byte{0x48, 0xb8, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"eax", 1, []byte{0xb8, 0x01, 0x00, 0x00, 0x00}},
		{"ax", 1, []byte{0x66, 0xb8, 0x01, 0x00}},
		{"al", 1, []byte{0xb0, 0x01}},
		{"rcx", 1, []byte{0x48, 0xb9, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"ecx", 1, []byte{0xb9, 0x01, 0x00, 0x00, 0x00}},
		{"cx", 1, []byte{0x66, 0xb9, 0x01, 0x00}},
		{"cl", 1, []byte{0xb1, 0x01}},
		{"rdx", 1, []byte{0x48, 0xba, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"edx", 1, []byte{0xba, 0x01, 0x00, 0x00, 0x00}},
		{"dx", 1, []byte{0x66, 0xba, 0x01, 0x00}},
		{"dl", 1, []byte{0xb2, 0x01}},
		{"rbx", 1, []byte{0x48, 0xbb, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"ebx", 1, []byte{0xbb, 0x01, 0x00, 0x00, 0x00}},
		{"bx", 1, []byte{0x66, 0xbb, 0x01, 0x00}},
		{"bl", 1, []byte{0xb3, 0x01}},
		{"rsi", 1, []byte{0x48, 0xbe, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"esi", 1, []byte{0xbe, 0x01, 0x00, 0x00, 0x00}},
		{"si", 1, []byte{0x66, 0xbe, 0x01, 0x00}},
		{"sil", 1, []byte{0x40, 0xb6, 0x01}},
		{"rdi", 1, []byte{0x48, 0xbf, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"edi", 1, []byte{0xbf, 0x01, 0x00, 0x00, 0x00}},
		{"di", 1, []byte{0x66, 0xbf, 0x01, 0x00}},
		{"dil", 1, []byte{0x40, 0xb7, 0x01}},
		{"rsp", 1, []byte{0x48, 0xbc, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"esp", 1, []byte{0xbc, 0x01, 0x00, 0x00, 0x00}},
		{"sp", 1, []byte{0x66, 0xbc, 0x01, 0x00}},
		{"spl", 1, []byte{0x40, 0xb4, 0x01}},
		{"rbp", 1, []byte{0x48, 0xbd, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"ebp", 1, []byte{0xbd, 0x01, 0x00, 0x00, 0x00}},
		{"bp", 1, []byte{0x66, 0xbd, 0x01, 0x00}},
		{"bpl", 1, []byte{0x40, 0xb5, 0x01}},
		{"r8", 1, []byte{0x49, 0xb8, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"r8d", 1, []byte{0x41, 0xb8, 0x01, 0x00, 0x00, 0x00}},
		{"r8w", 1, []byte{0x66, 0x41, 0xb8, 0x01, 0x00}},
		{"r8b", 1, []byte{0x41, 0xb0, 0x01}},
		{"r9", 1, []byte{0x49, 0xb9, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"r9d", 1, []byte{0x41, 0xb9, 0x01, 0x00, 0x00, 0x00}},
		{"r9w", 1, []byte{0x66, 0x41, 0xb9, 0x01, 0x00}},
		{"r9b", 1, []byte{0x41, 0xb1, 0x01}},
		{"r10", 1, []byte{0x49, 0xba, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"r10d", 1, []byte{0x41, 0xba, 0x01, 0x00, 0x00, 0x00}},
		{"r10w", 1, []byte{0x66, 0x41, 0xba, 0x01, 0x00}},
		{"r10b", 1, []byte{0x41, 0xb2, 0x01}},
		{"r11", 1, []byte{0x49, 0xbb, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"r11d", 1, []byte{0x41, 0xbb, 0x01, 0x00, 0x00, 0x00}},
		{"r11w", 1, []byte{0x66, 0x41, 0xbb, 0x01, 0x00}},
		{"r11b", 1, []byte{0x41, 0xb3, 0x01}},
		{"r12", 1, []byte{0x49, 0xbc, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"r12d", 1, []byte{0x41, 0xbc, 0x01, 0x00, 0x00, 0x00}},
		{"r12w", 1, []byte{0x66, 0x41, 0xbc, 0x01, 0x00}},
		{"r12b", 1, []byte{0x41, 0xb4, 0x01}},
		{"r13", 1, []byte{0x49, 0xbd, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"r13d", 1, []byte{0x41, 0xbd, 0x01, 0x00, 0x00, 0x00}},
		{"r13w", 1, []byte{0x66, 0x41, 0xbd, 0x01, 0x00}},
		{"r13b", 1, []byte{0x41, 0xb5, 0x01}},
		{"r14", 1, []byte{0x49, 0xbe, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"r14d", 1, []byte{0x41, 0xbe, 0x01, 0x00, 0x00, 0x00}},
		{"r14w", 1, []byte{0x66, 0x41, 0xbe, 0x01, 0x00}},
		{"r14b", 1, []byte{0x41, 0xb6, 0x01}},
		{"r15", 1, []byte{0x49, 0xbf, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"r15d", 1, []byte{0x41, 0xbf, 0x01, 0x00, 0x00, 0x00}},
		{"r15w", 1, []byte{0x66, 0x41, 0xbf, 0x01, 0x00}},
		{"r15b", 1, []byte{0x41, 0xb7, 0x01}},

		// Conversion tests
		{"rax", 0xff, []byte{0x48, 0xb8, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},               // 8 bit
		{"rax", 0xffff, []byte{0x48, 0xb8, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},             // 16 bit
		{"rax", 0xffffffff, []byte{0x48, 0xb8, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00}},         // 32 bit
		{"rax", 0x7fffffffffffffff, []byte{0x48, 0xb8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}}, // 64 bit
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		a.EnableOptimizer = false
		t.Logf("mov %s, %d", pattern.Register, pattern.Number)
		a.MoveRegisterNumber(pattern.Register, uint64(pattern.Number))
		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}

func TestMoveRegisterNumberOptimized(t *testing.T) {
	usagePatterns := []struct {
		Register string
		Number   int64
		Code     []byte
	}{
		{"rax", 1, []byte{0xb8, 0x01, 0x00, 0x00, 0x00}},
		{"rcx", 1, []byte{0xb9, 0x01, 0x00, 0x00, 0x00}},
		{"rdx", 1, []byte{0xba, 0x01, 0x00, 0x00, 0x00}},
		{"rbx", 1, []byte{0xbb, 0x01, 0x00, 0x00, 0x00}},
		{"rsi", 1, []byte{0xbe, 0x01, 0x00, 0x00, 0x00}},
		{"rdi", 1, []byte{0xbf, 0x01, 0x00, 0x00, 0x00}},
		{"rsp", 1, []byte{0xbc, 0x01, 0x00, 0x00, 0x00}},
		{"rbp", 1, []byte{0xbd, 0x01, 0x00, 0x00, 0x00}},
		{"r8", 1, []byte{0x41, 0xb8, 0x01, 0x00, 0x00, 0x00}},
		{"r9", 1, []byte{0x41, 0xb9, 0x01, 0x00, 0x00, 0x00}},
		{"r10", 1, []byte{0x41, 0xba, 0x01, 0x00, 0x00, 0x00}},
		{"r11", 1, []byte{0x41, 0xbb, 0x01, 0x00, 0x00, 0x00}},
		{"r12", 1, []byte{0x41, 0xbc, 0x01, 0x00, 0x00, 0x00}},
		{"r13", 1, []byte{0x41, 0xbd, 0x01, 0x00, 0x00, 0x00}},
		{"r14", 1, []byte{0x41, 0xbe, 0x01, 0x00, 0x00, 0x00}},
		{"r15", 1, []byte{0x41, 0xbf, 0x01, 0x00, 0x00, 0x00}},

		// Conversion tests
		{"rax", math.MaxInt8, []byte{0xb8, 0x7f, 0x00, 0x00, 0x00}},                                // 8 bit
		{"rax", math.MaxInt16, []byte{0xb8, 0xff, 0x7f, 0x00, 0x00}},                               // 16 bit
		{"rax", math.MaxInt32, []byte{0xb8, 0xff, 0xff, 0xff, 0x7f}},                               // 32 bit
		{"rax", math.MaxInt64, []byte{0x48, 0xb8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}}, // 64 bit
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		t.Logf("mov %s, %d", pattern.Register, pattern.Number)
		a.MoveRegisterNumber(pattern.Register, uint64(pattern.Number))
		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}

func TestMoveRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination string
		Source      string
		Code        []byte
	}{
		{"rax", "rax", []byte{0x48, 0x89, 0xc0}},
		{"eax", "eax", []byte{0x89, 0xc0}},
		{"ax", "ax", []byte{0x66, 0x89, 0xc0}},
		{"al", "al", []byte{0x88, 0xc0}},
		{"rcx", "rcx", []byte{0x48, 0x89, 0xc9}},
		{"ecx", "ecx", []byte{0x89, 0xc9}},
		{"cx", "cx", []byte{0x66, 0x89, 0xc9}},
		{"cl", "cl", []byte{0x88, 0xc9}},
		{"rdx", "rdx", []byte{0x48, 0x89, 0xd2}},
		{"edx", "edx", []byte{0x89, 0xd2}},
		{"dx", "dx", []byte{0x66, 0x89, 0xd2}},
		{"dl", "dl", []byte{0x88, 0xd2}},
		{"rbx", "rbx", []byte{0x48, 0x89, 0xdb}},
		{"ebx", "ebx", []byte{0x89, 0xdb}},
		{"bx", "bx", []byte{0x66, 0x89, 0xdb}},
		{"bl", "bl", []byte{0x88, 0xdb}},
		{"rsi", "rsi", []byte{0x48, 0x89, 0xf6}},
		{"esi", "esi", []byte{0x89, 0xf6}},
		{"si", "si", []byte{0x66, 0x89, 0xf6}},
		{"sil", "sil", []byte{0x40, 0x88, 0xf6}},
		{"rdi", "rdi", []byte{0x48, 0x89, 0xff}},
		{"edi", "edi", []byte{0x89, 0xff}},
		{"di", "di", []byte{0x66, 0x89, 0xff}},
		{"dil", "dil", []byte{0x40, 0x88, 0xff}},
		{"rsp", "rsp", []byte{0x48, 0x89, 0xe4}},
		{"esp", "esp", []byte{0x89, 0xe4}},
		{"sp", "sp", []byte{0x66, 0x89, 0xe4}},
		{"spl", "spl", []byte{0x40, 0x88, 0xe4}},
		{"rbp", "rbp", []byte{0x48, 0x89, 0xed}},
		{"ebp", "ebp", []byte{0x89, 0xed}},
		{"bp", "bp", []byte{0x66, 0x89, 0xed}},
		{"bpl", "bpl", []byte{0x40, 0x88, 0xed}},
		{"r8", "r8", []byte{0x4d, 0x89, 0xc0}},
		{"r8d", "r8d", []byte{0x45, 0x89, 0xc0}},
		{"r8w", "r8w", []byte{0x66, 0x45, 0x89, 0xc0}},
		{"r8b", "r8b", []byte{0x45, 0x88, 0xc0}},
		{"r9", "r9", []byte{0x4d, 0x89, 0xc9}},
		{"r9d", "r9d", []byte{0x45, 0x89, 0xc9}},
		{"r9w", "r9w", []byte{0x66, 0x45, 0x89, 0xc9}},
		{"r9b", "r9b", []byte{0x45, 0x88, 0xc9}},
		{"r10", "r10", []byte{0x4d, 0x89, 0xd2}},
		{"r10d", "r10d", []byte{0x45, 0x89, 0xd2}},
		{"r10w", "r10w", []byte{0x66, 0x45, 0x89, 0xd2}},
		{"r10b", "r10b", []byte{0x45, 0x88, 0xd2}},
		{"r11", "r11", []byte{0x4d, 0x89, 0xdb}},
		{"r11d", "r11d", []byte{0x45, 0x89, 0xdb}},
		{"r11w", "r11w", []byte{0x66, 0x45, 0x89, 0xdb}},
		{"r11b", "r11b", []byte{0x45, 0x88, 0xdb}},
		{"r12", "r12", []byte{0x4d, 0x89, 0xe4}},
		{"r12d", "r12d", []byte{0x45, 0x89, 0xe4}},
		{"r12w", "r12w", []byte{0x66, 0x45, 0x89, 0xe4}},
		{"r12b", "r12b", []byte{0x45, 0x88, 0xe4}},
		{"r13", "r13", []byte{0x4d, 0x89, 0xed}},
		{"r13d", "r13d", []byte{0x45, 0x89, 0xed}},
		{"r13w", "r13w", []byte{0x66, 0x45, 0x89, 0xed}},
		{"r13b", "r13b", []byte{0x45, 0x88, 0xed}},
		{"r14", "r14", []byte{0x4d, 0x89, 0xf6}},
		{"r14d", "r14d", []byte{0x45, 0x89, 0xf6}},
		{"r14w", "r14w", []byte{0x66, 0x45, 0x89, 0xf6}},
		{"r14b", "r14b", []byte{0x45, 0x88, 0xf6}},
		{"r15", "r15", []byte{0x4d, 0x89, 0xff}},
		{"r15d", "r15d", []byte{0x45, 0x89, 0xff}},
		{"r15w", "r15w", []byte{0x66, 0x45, 0x89, 0xff}},
		{"r15b", "r15b", []byte{0x45, 0x88, 0xff}},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		t.Logf("mov %s, %s", pattern.Destination, pattern.Source)
		a.MoveRegisterRegister(pattern.Destination, pattern.Source)
		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}

func TestMoveMemoryNumber(t *testing.T) {
	usagePatterns := []struct {
		Register  string
		ByteCount byte
		Number    int64
		Code      []byte
	}{
		{"rax", 8, 0x7f, []byte{0x48, 0xc7, 0x00, 0x7f, 0x00, 0x00, 0x00}},
		{"rax", 4, 0x7f, []byte{0xc7, 0x00, 0x7f, 0x00, 0x00, 0x00}},
		{"rax", 2, 0x7f, []byte{0x66, 0xc7, 0x00, 0x7f, 0x00}},
		{"rax", 1, 0x7f, []byte{0xc6, 0x00, 0x7f}},
		{"rcx", 8, 0x7f, []byte{0x48, 0xc7, 0x01, 0x7f, 0x00, 0x00, 0x00}},
		{"rcx", 4, 0x7f, []byte{0xc7, 0x01, 0x7f, 0x00, 0x00, 0x00}},
		{"rcx", 2, 0x7f, []byte{0x66, 0xc7, 0x01, 0x7f, 0x00}},
		{"rcx", 1, 0x7f, []byte{0xc6, 0x01, 0x7f}},
		{"rdx", 8, 0x7f, []byte{0x48, 0xc7, 0x02, 0x7f, 0x00, 0x00, 0x00}},
		{"rdx", 4, 0x7f, []byte{0xc7, 0x02, 0x7f, 0x00, 0x00, 0x00}},
		{"rdx", 2, 0x7f, []byte{0x66, 0xc7, 0x02, 0x7f, 0x00}},
		{"rdx", 1, 0x7f, []byte{0xc6, 0x02, 0x7f}},
		{"rbx", 8, 0x7f, []byte{0x48, 0xc7, 0x03, 0x7f, 0x00, 0x00, 0x00}},
		{"rbx", 4, 0x7f, []byte{0xc7, 0x03, 0x7f, 0x00, 0x00, 0x00}},
		{"rbx", 2, 0x7f, []byte{0x66, 0xc7, 0x03, 0x7f, 0x00}},
		{"rbx", 1, 0x7f, []byte{0xc6, 0x03, 0x7f}},
		{"rdi", 8, 0x7f, []byte{0x48, 0xc7, 0x07, 0x7f, 0x00, 0x00, 0x00}},
		{"rdi", 4, 0x7f, []byte{0xc7, 0x07, 0x7f, 0x00, 0x00, 0x00}},
		{"rdi", 2, 0x7f, []byte{0x66, 0xc7, 0x07, 0x7f, 0x00}},
		{"rdi", 1, 0x7f, []byte{0xc6, 0x07, 0x7f}},
		{"rsi", 8, 0x7f, []byte{0x48, 0xc7, 0x06, 0x7f, 0x00, 0x00, 0x00}},
		{"rsi", 4, 0x7f, []byte{0xc7, 0x06, 0x7f, 0x00, 0x00, 0x00}},
		{"rsi", 2, 0x7f, []byte{0x66, 0xc7, 0x06, 0x7f, 0x00}},
		{"rsi", 1, 0x7f, []byte{0xc6, 0x06, 0x7f}},
		{"rbp", 8, 0x7f, []byte{0x48, 0xc7, 0x45, 0x00, 0x7f, 0x00, 0x00, 0x00}},
		{"rbp", 4, 0x7f, []byte{0xc7, 0x45, 0x00, 0x7f, 0x00, 0x00, 0x00}},
		{"rbp", 2, 0x7f, []byte{0x66, 0xc7, 0x45, 0x00, 0x7f, 0x00}},
		{"rbp", 1, 0x7f, []byte{0xc6, 0x45, 0x00, 0x7f}},
		{"rsp", 8, 0x7f, []byte{0x48, 0xc7, 0x04, 0x24, 0x7f, 0x00, 0x00, 0x00}},
		{"rsp", 4, 0x7f, []byte{0xc7, 0x04, 0x24, 0x7f, 0x00, 0x00, 0x00}},
		{"rsp", 2, 0x7f, []byte{0x66, 0xc7, 0x04, 0x24, 0x7f, 0x00}},
		{"rsp", 1, 0x7f, []byte{0xc6, 0x04, 0x24, 0x7f}},
		{"r8", 8, 0x7f, []byte{0x49, 0xc7, 0x00, 0x7f, 0x00, 0x00, 0x00}},
		{"r8", 4, 0x7f, []byte{0x41, 0xc7, 0x00, 0x7f, 0x00, 0x00, 0x00}},
		{"r8", 2, 0x7f, []byte{0x66, 0x41, 0xc7, 0x00, 0x7f, 0x00}},
		{"r8", 1, 0x7f, []byte{0x41, 0xc6, 0x00, 0x7f}},
		{"r9", 8, 0x7f, []byte{0x49, 0xc7, 0x01, 0x7f, 0x00, 0x00, 0x00}},
		{"r9", 4, 0x7f, []byte{0x41, 0xc7, 0x01, 0x7f, 0x00, 0x00, 0x00}},
		{"r9", 2, 0x7f, []byte{0x66, 0x41, 0xc7, 0x01, 0x7f, 0x00}},
		{"r9", 1, 0x7f, []byte{0x41, 0xc6, 0x01, 0x7f}},
		{"r10", 8, 0x7f, []byte{0x49, 0xc7, 0x02, 0x7f, 0x00, 0x00, 0x00}},
		{"r10", 4, 0x7f, []byte{0x41, 0xc7, 0x02, 0x7f, 0x00, 0x00, 0x00}},
		{"r10", 2, 0x7f, []byte{0x66, 0x41, 0xc7, 0x02, 0x7f, 0x00}},
		{"r10", 1, 0x7f, []byte{0x41, 0xc6, 0x02, 0x7f}},
		{"r11", 8, 0x7f, []byte{0x49, 0xc7, 0x03, 0x7f, 0x00, 0x00, 0x00}},
		{"r11", 4, 0x7f, []byte{0x41, 0xc7, 0x03, 0x7f, 0x00, 0x00, 0x00}},
		{"r11", 2, 0x7f, []byte{0x66, 0x41, 0xc7, 0x03, 0x7f, 0x00}},
		{"r11", 1, 0x7f, []byte{0x41, 0xc6, 0x03, 0x7f}},
		{"r12", 8, 0x7f, []byte{0x49, 0xc7, 0x04, 0x24, 0x7f, 0x00, 0x00, 0x00}},
		{"r12", 4, 0x7f, []byte{0x41, 0xc7, 0x04, 0x24, 0x7f, 0x00, 0x00, 0x00}},
		{"r12", 2, 0x7f, []byte{0x66, 0x41, 0xc7, 0x04, 0x24, 0x7f, 0x00}},
		{"r12", 1, 0x7f, []byte{0x41, 0xc6, 0x04, 0x24, 0x7f}},
		{"r13", 8, 0x7f, []byte{0x49, 0xc7, 0x45, 0x00, 0x7f, 0x00, 0x00, 0x00}},
		{"r13", 4, 0x7f, []byte{0x41, 0xc7, 0x45, 0x00, 0x7f, 0x00, 0x00, 0x00}},
		{"r13", 2, 0x7f, []byte{0x66, 0x41, 0xc7, 0x45, 0x00, 0x7f, 0x00}},
		{"r13", 1, 0x7f, []byte{0x41, 0xc6, 0x45, 0x00, 0x7f}},
		{"r14", 8, 0x7f, []byte{0x49, 0xc7, 0x06, 0x7f, 0x00, 0x00, 0x00}},
		{"r14", 4, 0x7f, []byte{0x41, 0xc7, 0x06, 0x7f, 0x00, 0x00, 0x00}},
		{"r14", 2, 0x7f, []byte{0x66, 0x41, 0xc7, 0x06, 0x7f, 0x00}},
		{"r14", 1, 0x7f, []byte{0x41, 0xc6, 0x06, 0x7f}},
		{"r15", 8, 0x7f, []byte{0x49, 0xc7, 0x07, 0x7f, 0x00, 0x00, 0x00}},
		{"r15", 4, 0x7f, []byte{0x41, 0xc7, 0x07, 0x7f, 0x00, 0x00, 0x00}},
		{"r15", 2, 0x7f, []byte{0x66, 0x41, 0xc7, 0x07, 0x7f, 0x00}},
		{"r15", 1, 0x7f, []byte{0x41, 0xc6, 0x07, 0x7f}},
	}

	for _, pattern := range usagePatterns {
		a := asm.New()
		t.Logf("mov %db [%s], %d", pattern.ByteCount, pattern.Register, pattern.Number)
		a.MoveMemoryNumber(pattern.Register, pattern.ByteCount, uint64(pattern.Number))
		assert.DeepEqual(t, a.Bytes(), pattern.Code)
	}
}
