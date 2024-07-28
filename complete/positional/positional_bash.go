package positional

import (
	"github.com/atticus-sullivan/go-complete/internal"
	"strings"
)

func (p *Positional) GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash) {
	add := make([]*internal.AddFuncBash, 0)
	ifs := make([]*internal.AddIfBash, 0)

	p.Arg = p.Arg.SetIdxS(p.Idx)
	a, i := p.Arg.GenerateBash(builderArguments, indent, progName)
	add = append(add, a...)
	ifs = append(ifs, i...)

	return add, ifs
}
