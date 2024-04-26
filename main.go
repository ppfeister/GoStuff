package main

import (
	"flag"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/ppfeister/GoStuff/messages"
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

	//debugkit.SetVerbosity(uint8(verbosity))

	if len(single_target) > 0 {
		if len(single_target) != 2 {
			log.Fatal(messages.Fetch(i18n.LocalizeConfig{
				MessageID: "Err_SinglePairCredQty",
			}))
		}
		if manifest_uri_simple != "" || manifest_uri_batteringram != "" {
			log.Fatal("Manifest can not be provided alongside single-pair creds")
		}
	}

	//debugkit.Out(debugkit.Severity["warn"], "Something fucked up!")
}
