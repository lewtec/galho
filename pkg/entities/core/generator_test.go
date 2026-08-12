package core_test

import (
	"testing"

	galhocore "github.com/lewtec/galho/pkg/core"
	// Register generate app via init.
	_ "github.com/lewtec/galho/pkg/entities/core"
)

func TestGenerateAppCommandRegistered(t *testing.T) {
	cmds := galhocore.GetGenerateCommands()
	cmd, ok := cmds["app"]
	if !ok {
		t.Fatal("expected generate command \"app\" to be registered")
	}
	if cmd.Use != "app [path]" {
		t.Fatalf("unexpected Use: %q", cmd.Use)
	}
}
