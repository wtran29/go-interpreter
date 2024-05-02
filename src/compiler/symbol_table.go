package compiler

// set as an unique string to differentiate between different scopes
type SymbolScope string

const (
	GlobalScope SymbolScope = "GLOBAL"
)

// Symbol is a struct that holds all the necessary info
// about a symbol we encounter in Funckey - Name, Scope and Index
type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

// SymbolTable associates strings with Symbols in its store
// and keeps track of numDefinitions it has.
type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
}

// NewSymbolTable returns an instance of SymbolTable
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

// Define takes an identifier as an argument, create definition and return a Symbol
func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions, Scope: GlobalScope}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	return obj, ok
}
