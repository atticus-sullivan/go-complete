package completer

import (
	"complete/internal"
	"fmt"
	"slices"
	"strings"
)

func (mc *MainCompleter) GenerateBash(builder *strings.Builder, indent string) {
	addFuncs := make([]*internal.AddFuncBash, 0)
	addIfs := make([]*internal.AddIfBash, 0)

	var flags []string
	for _, f := range mc.Flags {
		flags = append(flags, "-"+string(f.Short))
		flags = append(flags, "--"+f.Long)
	}

	fmt.Fprintf(builder, `%[1]sfunction _%[2]s {
%[1]s    local cur opts
%[1]s    COMPREPLY=()
%[1]s    cur="${COMP_WORDS[COMP_CWORD]}"
%[1]s    opts="%[3]s"

%[1]s    local positionals=()
%[1]s    for ((i=1; i<COMP_CWORD; i++)); do
`, indent, mc.Name, strings.Join(flags, " "))

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

	builder.WriteString(indent)
	builder.WriteString(`    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
`)

	builder.WriteRune('}')
	builder.WriteRune('\n')
	fmt.Fprintf(builder, "%[1]scomplete -F _%[2]s %[2]s\n", indent, mc.Name)

	if slices.ContainsFunc(addFuncs, func(x *internal.AddFuncBash) bool { return x != nil }) {
		builder.WriteRune('\n')
		first := true

		addFuncsProcess := addFuncs
		for slices.ContainsFunc(addFuncsProcess, func(x *internal.AddFuncBash) bool { return x != nil }) {
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
