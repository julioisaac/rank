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

	for i := 0; i < len(topRepos); i++ {
		r := topRepos[i]
		rank := fmt.Sprintf("%v", i+1)
		score := fmt.Sprintf("%f", r.Stats.CalcActivityScore())
		fmt.Printf("Rank: %s, Repo: %s, Score: %s\n", rank, r.Repo, score)
	}
}
