package evaluator

import (
	"github.com/wtran29/go-interpreter/src/object"
)

var builtins = map[string]*object.Builtin{
	"len": object.GetBuiltinByName("len"),
	// first will return the first element of the given arary
	"first": object.GetBuiltinByName("first"),
	// last will return the last element of the given array
	"last": object.GetBuiltinByName("last"),
	// tail returns a new array containing all elements of the array except the first one.
	"tail": object.GetBuiltinByName("tail"),
	// push adds a new element to the end of array but does not modify the array.
	// it allocates a new array with the same elements as the old one plus the new
	"push": object.GetBuiltinByName("push"),
	"show": object.GetBuiltinByName("show"),
}
