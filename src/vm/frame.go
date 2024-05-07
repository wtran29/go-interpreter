package vm

import (
	"github.com/wtran29/go-interpreter/src/code"
	"github.com/wtran29/go-interpreter/src/object"
)

type Frame struct {
	fn *object.CompiledFunction // points to the compiled function reference by frame
	ip int //instruction pointer in this frame, for this function
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
