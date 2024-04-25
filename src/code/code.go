package code

import (
	"bytes"
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
	// allocate the instruction []byte and
	// add the Opcode as its first byte – by casting it into one
	instructionLen := 1
	// iterate over the defined OperandWidths
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	// take the matching element from operands and put it in the instruction
	for i, o := range operands {
		width := def.OperandWidths[i]
		// switch statement with a different method for each operand,
		// depending on how wide the operand is
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}

func (ins Instructions) String() string {
	var out bytes.Buffer
	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}
		operands, read := ReadOperands(def, ins[i+1:])
		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))
		i += 1 + read
	}
	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n", len(operands), operandCount)
	}
	switch operandCount {
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}
	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0
	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}
		offset += width
	}

	return operands, offset
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
