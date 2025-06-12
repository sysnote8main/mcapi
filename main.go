package main

import (
	"log/slog"

	"github.com/sysnote8main/mcapi/pkg/papermc"
)

func main() {
	allProjects, err := papermc.GetAllAvailableProjects()
	if err != nil {
		panic(err)
	}
	slog.Info("Result", slog.Any("data", allProjects))

	projectInfo, err := papermc.GetProjectInfo("paper")
	if err != nil {
		panic(err)
	}
	slog.Info("project info", slog.Any("data", projectInfo))

	builds, err := papermc.GetProjectBuilds("paper", "1.21.1")
	if err != nil {
		panic(err)
	}
	slog.Info("builds", slog.Any("data", builds))
}
