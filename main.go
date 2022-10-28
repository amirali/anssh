package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"gopkg.in/ini.v1"
)

type configType struct {
	inventoryPath stringArrayFlag
	targetHost    string
	user          string
	identity      string
}

var config *configType

func executeSSH() {
	sshString := fmt.Sprintf("%s@%s", config.user, config.targetHost)

	if config.identity != "" {
		sshString += fmt.Sprintf(" -i %s", config.identity)
	}

	cmd := exec.Command("ssh", sshString)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	initFlags()

	var content []byte
	for _, filename := range config.inventoryPath {
		thisContent, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		content = append(content, thisContent...)
	}

	inventory, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys: true,
		AllowShadows:     true,
	}, content)
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(initModel(inventory))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	if config.targetHost != "" {
		executeSSH()
	}
}

func extarctHost(keys ...string) []string {
	var hosts []string
	for _, key := range keys {
		host := strings.Split(key, " ")[0]
		hosts = append(hosts, host)
	}
	return hosts
}
