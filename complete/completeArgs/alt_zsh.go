package completeargs

import (
	"github.com/atticus-sullivan/go-complete/internal"
	"strings"
)

func (ca CTalternatives) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	builderArguments.WriteString(": :(")
	first := true
	for _,a := range ca.Alts {
		if !first {
			builderArguments.WriteRune(' ')
		}
		first = false
		builderArguments.WriteString(a)
	}
	builderArguments.WriteRune(')')
	return nil
}
