package types

import (
	"fmt"
	"slices"
	"strings"

	"github.com/atticus-sullivan/go-complete/internal"
	sh "mvdan.cc/sh/v3/syntax"
)

func (compl *Completer) GenerateZsh(builder *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	addFuncs := make([]*internal.AddFuncZsh, 0)

	fmt.Fprintf(builder, `%[1]sfunction _%[2]s {
%[1]s    local context state state_descr line opt_args
%[1]s    _arguments -C :`, indent, progName)

	name_s, err := sh.Quote(compl.Name, sh.LangBash)
	if err != nil {
		// TODO
		return nil
	}

	argIndent := indent + "        "
	for _, p := range compl.Positionals {
		addFuncs = append(addFuncs, p.GenerateZsh(builder, argIndent, progName+"_"+name_s)...)
	}
	for _, f := range compl.Flags {
		addFuncs = append(addFuncs, f.GenerateZsh(builder, argIndent, progName+"_"+name_s)...)
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
