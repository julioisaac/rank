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
			name: "Valid data processing",
			records: [][]string{
				{"1", "user1", "repo1", "3", "10", "5"},
				{"2", "user2", "repo1", "2", "5", "2"},
				{"3", "user3", "repo2", "1", "1", "1"},
			},
			expected: map[string]domain.RepoStats{
				"repo1": {Commits: 2, Files: 5, Additions: 15, Deletions: 7},
				"repo2": {Commits: 1, Files: 1, Additions: 1, Deletions: 1},
			},
		},
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

func TestGetTopRepos(t *testing.T) {
	rp := NewRepoProcessor()
	rp.stats["repo1"] = domain.RepoStats{Commits: 3, Files: 5, Additions: 10, Deletions: 2}
	rp.stats["repo2"] = domain.RepoStats{Commits: 5, Files: 1, Additions: 2, Deletions: 1}
	rp.stats["repo3"] = domain.RepoStats{Commits: 2, Files: 4, Additions: 5, Deletions: 3}

	expected := Results{
		{Repo: "repo1", Stats: rp.stats["repo1"]},
		{Repo: "repo3", Stats: rp.stats["repo3"]},
		{Repo: "repo2", Stats: rp.stats["repo2"]},
	}

	topRepos := rp.GetTopRepos(3)

	if !reflect.DeepEqual(topRepos, expected) {
		t.Errorf("GetTopRepos() failed, got %v, expected %v", topRepos, expected)
	}
}
