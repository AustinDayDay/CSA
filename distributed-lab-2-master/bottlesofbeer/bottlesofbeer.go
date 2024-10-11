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

func bottlesOfBeer(numOfBottles int) {
	if numOfBottles > 1 {
		fmt.Printf("%v Bottles of Beer on the wall, %v bottles of beer. Take one down, pass it around...\n", numOfBottles, numOfBottles)
	} else if numOfBottles == 1 {
		fmt.Printf("1 Bottle of Beer on the wall, 1 bottle of beer. Take one down, pass it around...\n")
	} else {
		fmt.Printf("NO MORE BEER :(\n")
	}

}

func passBottle(bottles int) {
	reqs := Token{NumOfBottles: bottles}
	resp := new(Token)
	if registered == false {
		var err error
		nextRound, err = rpc.Dial("tcp", nextAddr)
		if err != nil {
			fmt.Println(err)
		}
		registered = true
	}
	nextRound.Go("BottlesOfBeer.Round", reqs, resp, nil)
}

type BottlesOfBeer struct{}
type Token struct {
	NumOfBottles int
}

func (b *BottlesOfBeer) Round(inToken Token, outToken *Token) (err error) {
	bottles := inToken.NumOfBottles
	bottlesOfBeer(bottles)
	if bottles > 0 {
		passBottle(bottles - 1)
	}
	return
}

func main() {
	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	bottles := flag.Int("n", 0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.
	rpc.Register(&BottlesOfBeer{})
	listener, _ := net.Listen("tcp", ":"+*thisPort)
	defer listener.Close()
	if *bottles > 0 {
		bottlesOfBeer(int(*bottles))
		go passBottle(*bottles - 1)
	}
	rpc.Accept(listener)
}
