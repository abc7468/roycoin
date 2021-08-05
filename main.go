package main

import (
	"github.com/abc7468/roycoin/blockchain"
	"github.com/abc7468/roycoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
