package completer

import (
	"complete/internal"
	"fmt"
	"slices"
	"strings"
)

func (mc *MainCompleter) GenerateZsh(builder *strings.Builder, indent string) {
	addFuncs := make([]*internal.AddFuncZsh, 0)

	fmt.Fprintf(builder, `%[1]sfunction _%[2]s {
%[1]s    local context state state_descr line opt_args
%[1]s    _arguments -C :`, indent, mc.Name)

	argIndent := indent + "        "
	for _, p := range mc.Positionals {
		addFuncs = append(addFuncs, p.GenerateZsh(builder, argIndent, mc.Name)...)
	}
	for _, f := range mc.Flags {
		addFuncs = append(addFuncs, f.GenerateZsh(builder, argIndent, mc.Name)...)
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
	builder.WriteRune('\n')
	fmt.Fprintf(builder, "%[1]scompdef _%[2]s %[2]s\n", indent, mc.Name)

	if slices.ContainsFunc(addFuncs, func(x *internal.AddFuncZsh) bool { return x != nil }) {
		builder.WriteRune('\n')
		first := true

		addFuncsProcess := addFuncs
		for slices.ContainsFunc(addFuncsProcess, func(x *internal.AddFuncZsh) bool { return x != nil }) {
			addFuncs = nil

			for _, i := range addFuncsProcess {
				if i == nil {
					continue
				}
				if !first {
					builder.WriteRune('\n')
					builder.WriteRune('\n')
				} else {
					first = false
				}
				addFuncs = append(addFuncs, i.Fun(builder, "")...)
			}

			addFuncsProcess = addFuncs
		}
	}
}
