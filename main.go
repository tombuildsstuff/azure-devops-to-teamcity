package main

import (
	"flag"
	"os"

	"github.com/tombuildsstuff/azure-devops-to-teamcity/ado"
)

func main() {
	org := os.Getenv("ADO_ORGANIZATION")
	token := os.Getenv("ADO_TOKEN")
	project := os.Getenv("ADO_PROJECT")
	buildStepNumber := flag.Int("build-number", -1, "-build-number=12")

	flag.Parse()

	client := ado.NewClient("https://dev.azure.com", org, project, token)

	builds, err := client.GetAvailableBuilds()
	if err != nil {
		panic(err)
	}

	firstBuildId := (*builds)[0].Id
	buildLog, err := client.GetBuildLog(firstBuildId, *buildStepNumber)
	if err != nil {
		panic(err)
	}

	// write that out locally
	fileName := "build.log"
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(*buildLog)
	if err != nil {
		panic(err)
	}

	outputs, err := ParseBuildLog(fileName)
	if err != nil {
		panic(err)
	}

	err = FormatForTeamCity(*outputs)
	if err != nil {
		panic(err)
	}
}
