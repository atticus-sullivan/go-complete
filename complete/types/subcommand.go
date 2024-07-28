package types

import completeargs "github.com/atticus-sullivan/go-complete/complete/completeArgs"

type CTsubcommands struct {
	Cmds []Completer
	idx uint
}

func (cs CTsubcommands) SetIdxS(idx uint) completeargs.CompleteTypeSub {
	cs.idx = idx
	return cs
}
