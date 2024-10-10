package main

import (
	"bufio"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"

	//	"net/rpc"
	"flag"
	"net/rpc"
	"os"
	//	"bufio"
	//	"os"

	"fmt"
)

func workerReverse(words []string, client *rpc.Client, start int, finish int, responseChan chan []string) {
	finalResponse := []string{}
	for i := start; i < finish; i++ {
		request := stubs.Request{Message: words[i]}
		response := new(stubs.Response)
		client.Call(stubs.PremiumReverseHandler, request, response)
		finalResponse = append(finalResponse, response.Message)
	}
	responseChan <- finalResponse
}

func main() {
	server1 := flag.String("server1", "127.0.0.1:8030", "IP:port string to connect to as server")
	server2 := flag.String("server2", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	client1, _ := rpc.Dial("tcp", *server1)
	client2, _ := rpc.Dial("tcp", *server2)
	clients := []*rpc.Client{client1, client2}
	defer client1.Close()
	defer client2.Close()

	file, err := os.Open("../wordlist")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	responseChans := []chan []string{make(chan []string), make(chan []string)}

	for i := 0; i < 2; i++ {
		go workerReverse(words,, start, finish clients[i], responseChans[i])
	}

	finalReversedWords := []string{}
	for i := 0; i < 2; i++ {
		finalReversedWords = append(finalReversedWords, <-responseChans[i]...)
	}

	for word := range finalReversedWords {
		fmt.Println(word)
	}

	//TODO: connect to the RPC server and send the request(s)
}
