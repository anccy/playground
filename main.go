package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Checker struct {
	c *ethclient.Client
}

func NewChecker(rawURL string) (*Checker, error) {
	var err error
	c := &Checker{}
	c.c, err = ethclient.Dial(rawURL)
	return c, err
}

func (c *Checker) checkBalanceEth(addr string) (*big.Float, error) {
	account := common.HexToAddress(addr)
	o, err := c.c.BalanceAt(context.Background(), account, nil)

	fbalance := new(big.Float)
	fbalance.SetString(o.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue, err
}

func main() {
	checker, err := NewChecker("https://cloudflare-eth.com")
	if err != nil {
		fmt.Printf("invalid net: %v\n", err)
		return
	}

	if !common.IsHexAddress(os.Args[1]) {
		fmt.Println("invalid address")
		return
	}

	o, err := checker.checkBalanceEth(os.Args[1])
	if err != nil {
		fmt.Printf("balance check error:%v\n", err)
	}
	fmt.Printf("balance is %v eth\n", o)
}