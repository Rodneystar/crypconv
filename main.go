package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rodneystar/crypconv/cc"
)

func main() {
	params := parseParams()
	result, err := cc.Convert(params.Amount, params.Coin)
	if err != nil {
		usage()
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%f %s -> %f USD\n", params.Amount, params.Coin, result)
}

func parseParams() convertParms {
	args := os.Args[1:]

	if len(args) != 2 {
		usage()
		os.Exit(1)
	}
	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		usage()
		fmt.Println(err)
		os.Exit(1)
	}
	return convertParms{
		Coin:   args[1],
		Amount: amount,
	}
}

type convertParms struct {
	Coin   string
	Amount float64
}

func usage() {
	fmt.Println()
	fmt.Println("usage: crypconv <amount> <ticker>")
	fmt.Println()
}
