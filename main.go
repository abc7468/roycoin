package main

import (
	"github.com/abc7468/roycoin/explorer"
	"github.com/abc7468/roycoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
