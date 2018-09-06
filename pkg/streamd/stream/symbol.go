package stream_manager

import "fmt"

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
	return fmt.Sprintf("metathings.streamd.%v.%v.%v", self.component, self.id, self.alias)
}

func NewSymbol(id, component, alias string) Symbol {
	return &symbol{
		id:        id,
		component: component,
		alias:     alias,
	}
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
