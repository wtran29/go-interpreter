package vm

import (
	"github.com/wtran29/go-interpreter/src/code"
	"github.com/wtran29/go-interpreter/src/object"
)

type Frame struct {
	fn          *object.CompiledFunction // points to the compiled function reference by frame
	ip          int                      //instruction pointer in this frame, for this function
	basePointer int                      // used to know how many loacals a function is going to use
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	f := &Frame{
		fn:          fn,
		ip:          -1,
		basePointer: basePointer,
	}
	return f
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
