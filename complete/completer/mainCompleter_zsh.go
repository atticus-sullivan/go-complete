package completer

import (
	"fmt"
	"slices"
	"strings"

	"github.com/atticus-sullivan/go-complete/internal"
	sh "mvdan.cc/sh/v3/syntax"
)

func (mc *MainCompleter) GenerateZsh(builder *strings.Builder, indent string) {
	addFuncs := make([]*internal.AddFuncZsh, 0)

	name_s, err := sh.Quote(mc.Name, sh.LangBash)
	if err != nil {
		// TODO
		return
	}

	if mc.Opts.ToFile {
		fmt.Fprintf(builder, "#compdef _%[1]s %[1]s\n\n", name_s)
	}

	fmt.Fprintf(builder, `%[1]sfunction _%[2]s {
%[1]s    local context state state_descr line opt_args
%[1]s    _arguments -C :`, indent, name_s)

	argIndent := indent + "        "
	for _, p := range mc.Positionals {
		addFuncs = append(addFuncs, p.GenerateZsh(builder, argIndent, name_s)...)
	}
	for _, f := range mc.Flags {
		addFuncs = append(addFuncs, f.GenerateZsh(builder, argIndent, name_s)...)
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
	if !mc.Opts.ToFile {
		fmt.Fprintf(builder, "%[1]scompdef _%[2]s %[2]s\n", indent, name_s)
	}

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
