package types

import (
	"fmt"
	"strings"

	"github.com/atticus-sullivan/go-complete/internal"
	sh "mvdan.cc/sh/v3/syntax"
)

func (ct CTsubcommands) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	add := make([]*internal.AddFuncZsh, 0)

	builderArguments.WriteRune('(')

	first := true
	for _, c := range ct.Cmds {
		name_s, err := sh.Quote(c.Name, sh.LangBash)
		if err != nil {
			continue
		}
		if !first {
			builderArguments.WriteRune(' ')
		} else {
			first = false
		}
		builderArguments.WriteString(c.Name)
		add = append(add, &internal.AddFuncZsh{
			Cas: func(b *strings.Builder, i string) {
				fmt.Fprintf(b, `%[1]s%[2]s)
%[1]s    _%[3]s_%[2]s
%[1]s    ;;
`, i, name_s, progName)
			},
			Fun: func(b *strings.Builder, i string) []*internal.AddFuncZsh {
				return c.GenerateZsh(b, i, progName+"_"+c.Name)
			},
		})
	}

	builderArguments.WriteRune(')')

	return add
}
