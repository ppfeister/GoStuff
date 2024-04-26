package main

import (
	"flag"

	. "github.com/ppfeister/GoStuff/utils"
)

func main() {
	var num_workers uint
	var manifest_uri_simple string
	var manifest_uri_batteringram string
	var debug bool
	var verbosity uint
	flag.UintVar(&num_workers, "t", 20, "set number of concurrent threads")
	flag.StringVar(&manifest_uri_simple, "s", "", "use a simple manifest")
	flag.StringVar(&manifest_uri_batteringram, "b", "", "use a grouped manifest (not recommended)")
	flag.BoolVar(&debug, "d", false, "dump debugging info to stdout")
	flag.UintVar(&verbosity, "verbosity", 1, "")
	flag.Parse()
	var single_target []string = flag.Args()

	var debugkit DebugKit = BuildDebugKit(uint8(verbosity))

	/*if len(single_target) > 0 {
		fmt
	}*/

	debugkit.Print(debugkit, Severity["info"])
}
