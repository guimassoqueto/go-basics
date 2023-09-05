package main

import (
	"gos/config"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	logger.Infof("Galera, sou gay -> %s", "você tambem")
}