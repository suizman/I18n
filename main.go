package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/prometheus/common/log"
)

const version = "0.1"

var flagVersion = false
var flagHelp = false

var printVersion = flag.CommandLine.Name() + " version: " + version
var printUsage = "usage: " + flag.CommandLine.Name() + " keyword ..."

func i18n(s string) (int, error) {

	if len(s) < 3 {
		fmt.Printf("Keyword should have at least 3 letters and %s has: %d\n", s, len(s))
		os.Exit(1)
	}

	b := []byte(s)
	resp, err := fmt.Printf("%s%d%s", string(b[0]), len(b)-2, string(b[len(b)-1]))
	if err != nil {
		log.Errorf("Error: %v:\n", err)
	}
	return resp, nil
}

func init() {
	flag.BoolVar(&flagVersion, "version", false, "display program version")
	flag.BoolVar(&flagHelp, "help", false, printUsage)
}

func main() {

	flag.Parse()
	switch {
	case flagVersion:
		fmt.Println(printVersion)
	case flagHelp:
		flag.Usage()
	case len(os.Args) <= 1:
		flag.Usage()
		os.Exit(1)
	case len(os.Args) > 1:
		i18n(os.Args[1])
	}

}
