package main

import (
	"github.com/abc7468/roycoin/cli"
	"github.com/abc7468/roycoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
