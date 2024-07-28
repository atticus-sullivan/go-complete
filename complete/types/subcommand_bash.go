package types

import (
	"github.com/atticus-sullivan/go-complete/internal"
	"fmt"
	"strings"
	sh"mvdan.cc/sh/v3/syntax"
)

func (ct CTsubcommands) GenerateBash(builder *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash) {
	add := make([]*internal.AddFuncBash, 0)
	ifs := make([]*internal.AddIfBash, 0)

	fmt.Fprintf(builder, `%[1]sif [[ ${#positionals[@]} -eq %[2]d ]] ; then
%[1]s    case "${COMP_WORDS[i]}" in
`, indent, ct.idx)

	var cmds []string
	for _, c := range ct.Cmds {
		name_s, err := sh.Quote(c.Name, sh.LangBash)
		if err != nil {
			continue
		}
		fmt.Fprintf(builder, `%[1]s        %[2]s)
%[1]s            _%[3]s_%[2]s $i
%[1]s            return 0
%[1]s            ;;
`, indent, name_s, progName)
		cmds = append(cmds, c.Name)
		add = append(add, &internal.AddFuncBash{
			Fun: func(b *strings.Builder, i string) []*internal.AddFuncBash {
				return c.GenerateBash(b, i, progName+"_"+c.Name)
			},
		})
	}

	cmd := strings.Join(cmds, " ")
	ifs = append(ifs, &internal.AddIfBash{
		Fun: func(b *strings.Builder, i string) {
			fmt.Fprintf(builder, `%[1]sif [[ ${#positionals[@]} -eq %[2]d ]] ; then
%[1]s    COMPREPLY=( $(compgen -W "%[3]s ${opts}" -- ${cur}) )
%[1]s    return 0
%[1]sfi
`, i, ct.idx, cmd)
		},
	})

	builder.WriteString(indent)
	builder.WriteString("    esac\n")
	builder.WriteString(indent)
	builder.WriteString("fi\n")

	return add, ifs
}
