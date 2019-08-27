package util

type Repository struct {
	Title string
	Path  string
}

type Config struct {
	Author       string
	Repositories []Repository
	Exclude      []string
}
