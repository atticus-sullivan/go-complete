package types

import (
	"complete/internal"
	"fmt"
	"slices"
	"strings"
)

func (compl *Completer) GenerateZsh(builder *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	addFuncs := make([]*internal.AddFuncZsh, 0)

	fmt.Fprintf(builder, `%[1]sfunction _%[2]s {
%[1]s    local context state state_descr line opt_args
%[1]s    _arguments -C :`, indent, progName)

	argIndent := indent + "        "
	for _, p := range compl.Positionals {
		addFuncs = append(addFuncs, p.GenerateZsh(builder, argIndent, progName+"_"+compl.Name)...)
	}
	for _, f := range compl.Flags {
		addFuncs = append(addFuncs, f.GenerateZsh(builder, argIndent, progName+"_"+compl.Name)...)
	}

	if slices.ContainsFunc(addFuncs, func(x *internal.AddFuncZsh) bool { return x != nil }) {
		builder.WriteString(" \\\n")
		builder.WriteString(argIndent)
		builder.WriteString(`"*::arg:->args"`)
		builder.WriteRune('\n')
	}

	builder.WriteRune('\n')

	if slices.ContainsFunc(addFuncs, func(x *internal.AddFuncZsh) bool { return x != nil }) {
		builder.WriteString("    case ${line[1]} in\n")
		for _, i := range addFuncs {
			if i == nil {
				continue
			}
			i.Cas(builder, argIndent)
		}
		builder.WriteString("    esac\n")
	}

	builder.WriteRune('}')

	return addFuncs
}
