package domain

import "testing"

func TestCalcActivityScore(t *testing.T) {
	tests := []struct {
		name     string
		stats    RepoStats
		expected float64
	}{
		{
			name:     "All zero values",
			stats:    RepoStats{Commits: 0, Files: 0, Additions: 0, Deletions: 0},
			expected: 0.0,
		},
		{
			name:     "Only commits",
			stats:    RepoStats{Commits: 10, Files: 0, Additions: 0, Deletions: 0},
			expected: 10.0,
		},
		{
			name:     "Only files changed",
			stats:    RepoStats{Commits: 0, Files: 4, Additions: 0, Deletions: 0},
			expected: 8.0,
		},
		{
			name:     "Only additions",
			stats:    RepoStats{Commits: 0, Files: 0, Additions: 20, Deletions: 0},
			expected: 10.0,
		},
		{
			name:     "Only deletions",
			stats:    RepoStats{Commits: 0, Files: 0, Additions: 0, Deletions: 10},
			expected: 5.0,
		},
		{
			name:     "Mixed values",
			stats:    RepoStats{Commits: 3, Files: 4, Additions: 10, Deletions: 6},
			expected: 3*1.0 + 4*2.0 + 10*0.5 + 6*0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stats.CalcActivityScore()
			if got != tt.expected {
				t.Errorf("CalcActivityScore() = %.2f, expected %.2f", got, tt.expected)
			}
		})
	}
}
