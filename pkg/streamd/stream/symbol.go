package stream_manager

import (
	"fmt"
	"strings"
)

const (
	SYMBOL_PREFIX = "metathings.streamd"
)

type Symbol interface {
	Id() string
	Component() string
	Alias() string
	String() string
}

type symbol struct {
	id        string
	component string
	alias     string
}

func (self *symbol) Id() string        { return self.id }
func (self *symbol) Component() string { return self.component }
func (self *symbol) Alias() string     { return self.alias }
func (self *symbol) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", SYMBOL_PREFIX, self.component, self.id, self.alias)
}

func NewSymbol(id, component, alias string) Symbol {
	return &symbol{
		id:        id,
		component: component,
		alias:     alias,
	}
}

func FromString(x string) (Symbol, error) {
	if !strings.HasPrefix(x, SYMBOL_PREFIX) {
		return nil, ErrBadSymbolString
	}

	buf := strings.Split(x[len(SYMBOL_PREFIX)+1:], ".")
	if len(buf) != 3 {
		return nil, ErrBadSymbolString
	}

	return NewSymbol(buf[0], buf[1], buf[2]), nil
}

type SymbolTable interface {
	Lookup(id_or_alias string) Symbol
}

type symbolTable struct {
	id_idx    map[string]Symbol
	alias_idx map[string]Symbol
}

func (self *symbolTable) Lookup(id_or_alias string) Symbol {
	if sym, ok := self.id_idx[id_or_alias]; ok {
		return sym
	}

	if sym, ok := self.alias_idx[id_or_alias]; ok {
		return sym
	}

	return nil
}

func NewSymbolTable(syms []Symbol) SymbolTable {
	sym_tbl := &symbolTable{
		id_idx:    make(map[string]Symbol),
		alias_idx: make(map[string]Symbol),
	}
	for _, sym := range syms {
		sym_tbl.id_idx[sym.Id()] = sym
		sym_tbl.alias_idx[sym.Alias()] = sym
	}
	return sym_tbl
}
