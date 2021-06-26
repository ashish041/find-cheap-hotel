package peon

import (
	"log"
	"os"
)

func QuitIfRoot() {
	if isRoot() {
		log.Fatal("this program can not be run as root.")
	}
}

func isRoot() bool {
	return os.Getuid() == 0
}
