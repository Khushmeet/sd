package main

import "errors"

/*
Scope defines the map that stores variable value
*/
type Scope struct {
	S map[string]int
}

/*
Add creates a new scope
*/
func Add(x string, val int) Scope {
	m := make(map[string]int)
	m[x] = val

	s := Scope{
		S: m,
	}
	return s
}

/*
Get retrives value from scope
*/
func (s *Scope) Get(x string) int {
	return s.S[x]
}

/*
SymbolTable keep tracks of the varibales in the program
*/
type SymbolTable struct {
	Table []Scope
}

/*
Push append the scope into the table
*/
func (st *SymbolTable) Push(scope Scope) {
	st.Table = append(st.Table, scope)
}

/*
Lookup retrives the value for x
*/
func (st *SymbolTable) Lookup(x string) (int, error) {
	for i := len(st.Table); i > 0; i-- {
		val := st.Table[i].Get(x)
		if val != -1 {
			return val, nil
		}
	}
	return -1, errors.New("Undefined")
}
