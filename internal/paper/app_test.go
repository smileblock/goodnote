package paper

import (
	"goodnode/internal/correlation"
	"testing"
)

type GraphRepo struct {
	t *testing.T
}

func (g *GraphRepo) Add(input *AddInput) correlation.Id {
	g.t.Log(input)
	return "1"
}

func NewGraphRepo(t *testing.T) Repo {
	return &GraphRepo{t: t}
}

func TestApp(t *testing.T) {
	repo := NewGraphRepo(t)
	app := NewApp(repo)
	input := NewAddInput("A", "AAA", "AAA")
	app.Add(input)
}
