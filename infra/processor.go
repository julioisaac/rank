package infra

import (
	"encoding/csv"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"rank/domain"
	"sort"
	"strconv"
)

type RepoProcessor struct {
	stats map[string]domain.RepoStats
}

type Results []Result

type Result struct {
	Repo  string
	Stats domain.RepoStats
}

func NewRepoProcessor() *RepoProcessor {
	return &RepoProcessor{stats: make(map[string]domain.RepoStats)}
}

func LoadCSV(filePath string, ch chan<- []string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("error reading CSV header: %w", err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			close(ch)
			break
		}
		ch <- record
	}
	return nil
}

func (rp *RepoProcessor) ProcessData(ch <-chan []string) {
	for record := range ch {
		if len(record) < 6 {
			log.Printf("Invalid register ignored: %v", record)
			continue
		}

		files, err := strconv.Atoi(record[3])
		if err != nil {
			log.Printf("Error converting files field: %v", record[3])
			continue
		}

		additions, err := strconv.Atoi(record[4])
		if err != nil {
			log.Printf("Error converting addings field: %v", record[4])
			continue
		}

		deletions, err := strconv.Atoi(record[5])
		if err != nil {
			log.Printf("Error converting deletions field: %v", record[5])
			continue
		}

		repo := record[2]
		stats := rp.stats[repo]
		stats.Commits++
		stats.Files += files
		stats.Additions += additions
		stats.Deletions += deletions
		rp.stats[repo] = stats
	}
}

func (rp *RepoProcessor) GetTopRepos(n int) Results {
	results := make(Results, 0, len(rp.stats))

	for repo, stats := range rp.stats {
		results = append(results, Result{Repo: repo, Stats: stats})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Stats.CalcActivityScore() > results[j].Stats.CalcActivityScore()
	})

	if len(results) > n {
		return results[:n]
	}
	return results
}

func (r Results) PrintTable() {
	fmt.Printf("\n--- Top %v Repositories ---\n", len(r))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Repository", "Score"})
	for i := 0; i < len(r); i++ {
		r := r[i]
		rank := fmt.Sprintf("%v", i+1)
		score := fmt.Sprintf("%f", r.Stats.CalcActivityScore())
		row := []string{rank, r.Repo, score}
		table.Append(row)
	}
	table.Render()
}
