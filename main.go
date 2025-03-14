package main

import (
	"fmt"
	"log"
	"os"
	"rank/infra"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Try: go run main.go <*.csv>")
	}

	processor := infra.NewRepoProcessor()
	dataChannel := make(chan []string, 100)

	go func() {
		var err = infra.LoadCSV(os.Args[1], dataChannel)
		if err != nil {
			fmt.Println("Erro ao carregar CSV:", err)
		}
	}()

	processor.ProcessData(dataChannel)

	topRepos := processor.GetTopRepos(10)
	topRepos.PrintTable()
}
