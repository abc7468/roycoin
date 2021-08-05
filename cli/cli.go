package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/abc7468/roycoin/explorer"
	"github.com/abc7468/roycoin/rest"
)

func usage() {
	fmt.Printf("Welcome to roy coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000: 		Set the Port of server\n")
	fmt.Printf("-mode=rest:		Choose between 'html' and 'rest'\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()
	fmt.Println(*port, *mode)
	switch *mode {
	case "rest":
		//start rest api
		rest.Start(*port)
	case "html":
		//start html explorer
		explorer.Start(*port)

	default:
		usage()
	}
}
