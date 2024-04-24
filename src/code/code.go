package code

import (
	"encoding/binary"
	"fmt"
)

// Instructions is a slice of bytes
type Instructions []byte

// Opcode is a byte
type Opcode byte

// OpConstant has one operand
// iota generate increasing byte values for us
const (
	OpConstant Opcode = iota
)

// Definition defines an Opcode
type Definition struct {
	Name          string // helps to an Opcode readable
	OperandWidths []int  //OperandWidths contains the number of bytes each operand takes up
}

// definition for OpConstant is that its only operand is two bytes wide
// which makes it an uint16 and limits its maximum value to 65535
var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

// Make easily build up bytecode instructions without having to check for errors for each call
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}
	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}
