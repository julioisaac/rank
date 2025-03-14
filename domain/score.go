package domain

const (
	wCommit = 1.0
	wFiles  = 2.0
	wAdds   = 0.5
	wDels   = 0.5
)

type RepoStats struct {
	Commits   int
	Files     int
	Additions int
	Deletions int
}

func (r RepoStats) CalcActivityScore() float64 {
	return float64(r.Commits)*wCommit + float64(r.Files)*wFiles + float64(r.Additions)*wAdds + float64(r.Deletions)*wDels
}
