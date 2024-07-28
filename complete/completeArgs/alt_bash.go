package completeargs

import (
	"github.com/atticus-sullivan/go-complete/internal"
	"strings"
)

func (ca CTalternatives) GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash) {
	builderArguments.WriteString(indent)
	builderArguments.WriteString(`    COMPREPLY+=( $(compgen -W "`)
	first := true
	for _,a := range ca.Alts {
		if !first {
			builderArguments.WriteRune(' ')
		}
		first = false
		builderArguments.WriteString(a)
	}
	builderArguments.WriteString(`" -- "${cur}") )`)
	builderArguments.WriteRune('\n')
	return nil,nil
}
