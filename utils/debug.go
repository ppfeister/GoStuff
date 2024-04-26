package utils

import "fmt"

var Severity = map[string]int{
	"fatal": 0,
	"warn":  1,
	"debug": 2,
	"info":  3,
}

type DebugKit struct {
	stdprefix []string
	verbosity uint8
	Print     func(DebugKit, uint8, string)
}

func BuildDebugKit(verbosity uint8) DebugKit {
	return DebugKit{
		stdprefix: []string{": fatal :", ": warn : ", ": debug :", ": inform :"},
		verbosity: verbosity,
		Print: func(self DebugKit, sev uint8, msg string) {
			// Prevent severity boundary from being overrun
			if sev >= uint8(len(self.stdprefix)) {
				sev = uint8(len(self.stdprefix) - 1)
			}
			// Present messages to console only if verbosity level is high enough
			if sev < self.verbosity {
				fmt.Printf("%v %v", self.stdprefix[sev], msg)
			}
			fmt.Printf("%v %v", self.stdprefix, msg)
		},
	}
}
