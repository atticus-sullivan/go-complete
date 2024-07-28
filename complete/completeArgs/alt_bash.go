package completeargs

import (
	"github.com/atticus-sullivan/go-complete/internal"
	sh "mvdan.cc/sh/v3/syntax"
	"strings"
)

func (ca CTalternatives) GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash) {
	for _,a := range ca.Alts {
		s, err := sh.Quote(a, sh.LangBash)
		if err != nil || strings.HasPrefix(s, "\"") {
			// TODO
			continue
		}
		builderArguments.WriteString(indent)
		builderArguments.WriteString(`    COMPREPLY+=( $(compgen -W "`)
		builderArguments.WriteString(s)
		builderArguments.WriteString(`" -- "${cur}") )`)
		builderArguments.WriteRune('\n')
	}
	return nil,nil
}
