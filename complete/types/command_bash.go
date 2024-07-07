package types

import (
	"complete/internal"
	"fmt"
	"slices"
	"strings"
)

func (mc *Completer) GenerateBash(builder *strings.Builder, indent string, progName string) []*internal.AddFuncBash {
	addFuncs := make([]*internal.AddFuncBash, 0)
	addIfs := make([]*internal.AddIfBash, 0)

	var flags []string
	for _, f := range mc.Flags {
		flags = append(flags, "-"+string(f.Short))
		flags = append(flags, "--"+f.Long)
	}

	fmt.Fprintf(builder, `%[1]sfunction _%[2]s {
%[1]s    local cur opts
%[1]s    cur="${COMP_WORDS[COMP_CWORD]}"
%[1]s    opts="%[3]s"

%[1]s    local positionals=()
%[1]s    for ((i=$1+1; i<COMP_CWORD; i++)); do
`, indent, progName, strings.Join(flags, " "))

	for _, p := range mc.Positionals {
		a, i := p.GenerateBash(builder, indent+"        ", mc.Name)
		addFuncs = append(addFuncs, a...)
		addIfs = append(addIfs, i...)
	}

	builder.WriteString(indent)
	builder.WriteString(`        case "${COMP_WORDS[i]}" in
`)
	for _, f := range mc.Flags {
		addFuncs = append(addFuncs, f.GenerateBash(builder, indent+"        ", mc.Name)...)
	}
	fmt.Fprintf(builder, `%[1]s        *)
%[1]s            positionals+=( "${COMP_WORDS[i]}" )
%[1]s            ;;
%[1]s        esac
%[1]s    done
`, indent)

	if slices.ContainsFunc(addIfs, func(x *internal.AddIfBash) bool { return x != nil }) {
		for _, i := range addIfs {
			if i == nil {
				continue
			}
			i.Fun(builder, indent+"    ")
		}
	}

	builder.WriteRune('}')
	builder.WriteRune('\n')

	return addFuncs
}
