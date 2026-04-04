package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/timschumi/watchtower/cmd"
)

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	cmd.Execute()
}
