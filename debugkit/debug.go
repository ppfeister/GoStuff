package debugkit

import "fmt"

var Severity = map[string]uint8{
	"fatal": 0,
	"warn":  1,
	"debug": 2,
	"info":  3,
}

var stdprefix = map[uint8]string{
	Severity["fatal"]: ":::  fatal  :::",
	Severity["warn"]:  ":::  warn   :::",
	Severity["debug"]: ":::  debug  :::",
	Severity["info"]:  ":::  info   :::",
}

var verbosity_pref uint8 = Severity["warn"]

func SetVerbosity(verbosity uint8) {
	if verbosity > uint8(len(Severity)-1) {
		verbosity = uint8(len(Severity) - 1)
	}
	verbosity_pref = verbosity
}

func Out(severity uint8, message string) {
	if severity > verbosity_pref {
		return
	}
	if severity > uint8(len(Severity)-1) {
		severity = uint8(len(Severity) - 1)
	}
	fmt.Printf("%v   %v\n", stdprefix[severity], message)
}
