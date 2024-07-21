package flag

import (
	"complete/internal"
	"fmt"
	"strings"
)

func (f *Flag) GenerateBash(builder *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash) {
	if len(f.Args) > 1 {
		// TODO not nice that this generates an error
		// maybe we can let this silently fail instead and just always only use the first argument
		panic("only flags with one argument are supported when generating bash completion")
	}

	addFuncs := make([]*internal.AddFuncBash, 0)
	addIfs := make([]*internal.AddIfBash, 0)

	if f.Short != 0 {
		fmt.Fprintf(builder, `%[1]s-%[2]c)
%[1]s    (( i+=%[3]d ))
%[1]s    ;;
`, indent, f.Short, len(f.Args))
	}

	if f.Long != "" {
		fmt.Fprintf(builder, `%[1]s--%[2]s)
%[1]s    (( i+=%[3]d ))
%[1]s    ;;
`, indent, f.Long, len(f.Args))
	}

	if len(f.Args) > 0 {
		addIfs = append(addIfs, &internal.AddIfBash{
			Fun: func(builder *strings.Builder, id string) {
				fmt.Fprintf(builder, `%[1]sif [[ "${prev}" == "-%[2]c" || "${prev}" == "--%[3]s" ]] ; then
`, id, f.Short, f.Long)
				builder.WriteString(id)
				builder.WriteString("    COMPREPLY=()\n")
				f.Args[0].GenerateBash(builder, id, progName)
				builder.WriteString(id)
				builder.WriteString("fi\n")
				// %[1]s    COMPREPLY=( $(compgen -W "%[3]s ${opts}" -- ${cur}) )
				// %[1]s    return 0
				// %[1]sfi
			},
		})
	}

	return addFuncs, addIfs
}
