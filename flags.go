package main

import (
	"flag"
	"fmt"
	"os/user"
)

type stringArrayFlag []string

func (s *stringArrayFlag) String() string {
	return fmt.Sprintf("%#v", s)
}

func (s *stringArrayFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func initFlags() {
	config = &configType{}
	thisUser, _ := user.Current()
	flag.Var(&config.inventoryPath, "inv", "Path to ansible inventory")
	flag.StringVar(&config.user, "user", thisUser.Username, "User to invoke ssh with")
	flag.StringVar(&config.identity, "identity", "", "Path to identity file")
	flag.Parse()
}
