package flag

import (
	"github.com/atticus-sullivan/go-complete/internal"
	"fmt"
	"strings"
)

func (f *Flag) GenerateBash(builder *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash) {
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
				builder.WriteString(id)
				builder.WriteString("if [[ ")
				if f.Short != 0 {
					builder.WriteString(`"${prev}" == "-`)
					builder.WriteRune(f.Short)
					if f.Long != "" {
						builder.WriteString(`" || `)
					} else {
						builder.WriteString(`" `)
					}
				}
				if f.Long != "" {
					builder.WriteString(`"${prev}" == "--`)
					builder.WriteString(f.Long)
					builder.WriteString(`" `)
				}
				builder.WriteString("]] ; then\n")
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
