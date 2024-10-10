package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	//	"net/rpc"
	//	"fmt"
	//	"time"
	//	"net"
)

var nextAddr string
var registered = false
var nextRound *rpc.Client

type BottlesOfBeer struct{}

func (b *BottlesOfBeer) Round(inToken Token, outToken *Token) (err error) {
	bottles := inToken.NumOfBottles
	bottlesOfBeer(bottles)
	if bottles > 0 {
		passBottle(bottles - 1)
	}
	return
}

type Token struct {
	NumOfBottles int
}

func bottlesOfBeer(numOfBottles int) {
	if numOfBottles > 1 {
		fmt.Println(numOfBottles, " Bottles of Beer on the wall, ", numOfBottles, " bottles of beer. Take one down, pass it around...")
	} else if numOfBottles == 1 {
		fmt.Println("1 Bottle of Beer on the wall, 1 bottle of beer. Take one down, pass it around...")
	} else {
		fmt.Println("NO MORE BEER :(")
	}

}

func passBottle(bottles int) {
	reqs := Token{NumOfBottles: bottles}
	resp := new(Token)
	if !registered {
		nextRound, _ = rpc.Dial("tcp", nextAddr)
		registered = true
	}
	nextRound.Go("BottlesOfBeer.Round", reqs, resp, nil)
}

func main() {
	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	bottles := flag.Int("n", 0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.
	listener, _ := net.Listen("tcp", ":"+*thisPort)
	defer listener.Close()
	rpc.Register(&BottlesOfBeer{})
	if *bottles > 0 {
		bottlesOfBeer(int(*bottles))
		go passBottle(*bottles - 1)
	}
	rpc.Accept(listener)
}
