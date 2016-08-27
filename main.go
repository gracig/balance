package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	//Retrieve file name parameters from command line
	var (
		master string
		tx     string
	)
	flag.StringVar(&master, "master", "", "The Master file")
	flag.StringVar(&tx, "tx", "", "The Transaction File to be processed")
	flag.Parse()
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	if !(seen["master"] && seen["tx"]) {
		flag.PrintDefaults()
		os.Exit(2)
	}

	chMaster := openFileChannel(master)
	chTx := openFileChannel(tx)

	//Main loop
	mLine, isMasterActive := <-chMaster
	txLine, isTransactionActive := <-chTx
	for isMasterActive || isTransactionActive {
		var action, value string
		var nextMaster, nextTx bool
		if !isTransactionActive {
			action, value, nextMaster, nextTx = "del", mLine, true, false
		} else if !isMasterActive {
			action, value, nextMaster, nextTx = "new", txLine, false, true
		} else {
			if txLine == mLine {
				action, value, nextMaster, nextTx = "upt", txLine, true, true
			} else if txLine > mLine {
				action, value, nextMaster, nextTx = "del", mLine, true, false
			} else if txLine < mLine {
				action, value, nextMaster, nextTx = "new", txLine, false, true
			}
		}
		if nextMaster {
			mLine, isMasterActive = <-chMaster
		}
		if nextTx {
			txLine, isTransactionActive = <-chTx
		}
		fmt.Println(value + "," + action)
	}
}

func openFileChannel(file string) <-chan string {
	ch := make(chan string)
	go func(ch chan string) {
		defer close(ch)
		f, err := os.Open(file)
		if err != nil {
			log.Printf("Could not open file: %v. %v", file, err)
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Printf("Error while scanning file: %v. %v", file, err)
		}
	}(ch)
	return ch
}
