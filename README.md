# 🏆 Activity Ranking

## Approach
The approach for reading and processing the file involved implementing a Go routine that reads and processes the data concurrently, using the functions provided as inputs.
To calculate the score for each repository, a weight was assigned to the relevant fields.
```domain/score.go```
```go
const (
    wCommit = 1.0  // weight for commits
    wFiles  = 2.0  // weight for files changed
    wAdds   = 0.5  // weight for additions
    wDels   = 0.5  // weight for deletions
)
```

## Running

* **How to Test**
```bash
$ make tests
```

* **How to Run**
```bash
$ make run path=data/commits.csv
```