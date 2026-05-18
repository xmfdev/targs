package targs

import (
	"slices"
	"strings"
)

type Option struct {
	Options     []string
	Handler     func(string)
	HasExtraArg bool
}

func stripEqualSign(arg string) string {
	if idx := strings.IndexRune(arg, '='); idx != -1 {
		return arg[:idx+1]
	}

	return arg
}

func isEqualSignArg(arg string) bool {
	return strings.ContainsRune(arg, '=')
}

func getEqualSignExtraArg(arg string) string {
	_, after, _ := strings.Cut(arg, "=")
	return after
}

func HandleArgs(args []string, options []Option, invalidArg func()) {
	seen := []string{}

	fail := func() {
		if invalidArg != nil {
			invalidArg()
		}
	}

	for i := 0; i < len(args); i++ {
		matched := false
		argKey := stripEqualSign(args[i])

		if slices.Contains(seen, argKey) {
			fail()
			return
		}

		for _, option := range options {
			if !slices.Contains(option.Options, argKey) {
				continue
			}

			matched = true
			var extraArg string

			if option.HasExtraArg {
				if isEqualSignArg(args[i]) {
					extraArg = getEqualSignExtraArg(args[i])
				} else {
					i++

					if i >= len(args) {
						return
					}

					extraArg = args[i]
				}

			}

			option.Handler(extraArg)

			for _, v := range option.Options {
				seen = append(seen, stripEqualSign(v))
			}

			break
		}

		if !matched {
			fail()
			return
		}
	}
}
