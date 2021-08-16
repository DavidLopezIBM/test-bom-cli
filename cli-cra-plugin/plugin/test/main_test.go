package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := Login()
	if err != nil {
		os.Exit(1)
	}
	err = DeleteToolchainData(GetToolchainID())
	if err != nil {
		os.Exit(1)
	}
	m.Run()
}
