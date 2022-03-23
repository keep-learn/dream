package service

import (
	"dream/pkg/conf"
	"dream/pkg/log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	conf.Init()
	log.Init()

	exitCode := m.Run()
	os.Exit(exitCode)
}
