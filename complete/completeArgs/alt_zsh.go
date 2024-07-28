package completeargs

import (
	"strings"

	"github.com/atticus-sullivan/go-complete/internal"
	sh "mvdan.cc/sh/v3/syntax"
)

func (ca CTalternatives) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	builderArguments.WriteString(": :(")
	first := true
	for _,a := range ca.Alts {
		if !first {
			builderArguments.WriteRune(' ')
		}
		first = false
		s, err := sh.Quote(a, sh.LangBash)
		if err != nil {
			// TODO
			continue
		}
		builderArguments.WriteString(s)
	}
	builderArguments.WriteRune(')')
	return nil
}
