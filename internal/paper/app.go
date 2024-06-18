package paper

import (
	"goodnode/internal/correlation"
	"goodnode/internal/graph"
)

type Id string

type Text string

type AddInput struct {
	name graph.NodeName
	text Text
	// source はこのpaperの参照を持つnode
	// targetは追加されるこのpaper自身
	source graph.NodeId
}

func NewAddInput(name graph.NodeName, text Text, source graph.NodeId) *AddInput {
	return &AddInput{name: name, text: text, source: source}
}

type App struct {
	repo Repo
}

func (a *App) Add(input *AddInput) correlation.Id {
	return a.repo.Add(input)
}

func NewApp(repo Repo) *App {
	return &App{repo: repo}
}
