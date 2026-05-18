package targs

import "slices"

// Option represents a CLI flag with its associated handler.
type Option struct {
	Options     []string
	Handler     func(*string)
	HasExtraArg bool
}

// HandleArgs processes os.Args against the provided options.
func HandleArgs(args []string, options []Option, invalidArg func()) {
	seen := []string{}

	fail := func() {
		if invalidArg != nil {
			invalidArg()
		}
	}

	for i := 0; i < len(args); i++ {
		matched := false

		if slices.Contains(seen, args[i]) {
			fail()
			return
		}

		for _, option := range options {
			if !slices.Contains(option.Options, args[i]) {
				continue
			}

			matched = true
			var extraArg *string

			if option.HasExtraArg {
				i++
				if i >= len(args) {
					return
				}
				extraArg = &args[i]
			}

			option.Handler(extraArg)
			seen = append(seen, option.Options...)

			break
		}

		if !matched {
			fail()
			return
		}
	}
}
