package infra

import (
	"rank/domain"
	"reflect"
	"testing"
)

func TestProcessData(t *testing.T) {
	tests := []struct {
		name     string
		records  [][]string
		expected map[string]domain.RepoStats
	}{
		{
			name: "Ignores invalid records",
			records: [][]string{
				{"1", "user1", "repo1", "3", "10", "5"},
				{"2", "user2", "repo1"},
				{"3", "user3", "repo2", "a", "b", "c"},
			},
			expected: map[string]domain.RepoStats{
				"repo1": {Commits: 1, Files: 3, Additions: 10, Deletions: 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := NewRepoProcessor()
			ch := make(chan []string)

			go func() {
				for _, record := range tt.records {
					ch <- record
				}
				close(ch)
			}()

			rp.ProcessData(ch)

			if !reflect.DeepEqual(rp.stats, tt.expected) {
				t.Errorf("ProcessData() failed, got %v, expected %v", rp.stats, tt.expected)
			}
		})
	}
}
