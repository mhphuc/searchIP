package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"log"
	"math/rand"
	"os"
	"time"
)

var filename = "compromised-ips.txt"

type ipHashTable map[byte]ipHashTable

var rootHashTable ipHashTable
var hashTablesCount int = 0
var logger *log.Logger
var min = 1
var max = 255
var trie = NewTrie()

func main() {
	fmt.Println("Search IP from list")
	rand.Seed(time.Now().UnixNano())

	connections, err := net.Connections("all")
	if err != nil {
		panic(err)
	}

	// Using Trie
	loadIPFromFileTrie(filename)

	/*
		Remind time for search ip
	*/
	//t:=time.Now()
	//for _, connection := range connections {
	//	if connection.Status == "ESTABLISHED" {
	//		//fmt.Println(connection.Raddr.IP) //view list connection
	//		if searchIpInTrie(connection.Raddr.IP){
	//			fmt.Println("Block")
	//			fmt.Println(connection.Raddr.IP)
	//		}
	//	}
	//
	//}
	//fmt.Print("End Time   ")
	//fmt.Println(time.Now().Sub(t))

	for true {
		for _, connection := range connections {
			if connection.Status == "ESTABLISHED" {
				//fmt.Println(connection.Raddr.IP) //view list connection
				if searchIpInTrie(connection.Raddr.IP) {
					fmt.Println(connection.Raddr.IP)
				}
			}
		}
		fmt.Println("------------------------------------------")
		time.Sleep(500 * time.Millisecond)
	}

}

func loadIPFromFileTrie(filename string) {
	file, err := os.Open(filename)
	var ips = 0
	if err != nil {
		logger.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		logger.Fatal(err)
	}
	t := time.Now()
	fmt.Println("	Loading.")
	for scanner.Scan() {
		str := scanner.Text()
		trie.Insert([]byte(str), []byte("true"))
		ips++
	}
	fmt.Print(time.Now().Sub(t))
	fmt.Println("	IP in list: ", ips)
}

func searchIpInTrie(ip string) bool {
	_, ok := trie.Search([]byte(ip))
	if !ok {
		return false
	}
	return true
}
