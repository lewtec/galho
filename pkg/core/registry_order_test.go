package core

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestGetGenerateCommandsSortedByName(t *testing.T) {
	// Isolate registry for this test.
	prev := generateCommands
	generateCommands = make(map[string]*cobra.Command)
	t.Cleanup(func() { generateCommands = prev })

	// Register out of order.
	RegisterGenerateCommand("zeta", &cobra.Command{Use: "zeta"})
	RegisterGenerateCommand("alpha", &cobra.Command{Use: "alpha"})
	RegisterGenerateCommand("mu", &cobra.Command{Use: "mu"})

	got := GetGenerateCommands()
	want := []string{"alpha", "mu", "zeta"}
	if len(got) != len(want) {
		t.Fatalf("len=%d want %d", len(got), len(want))
	}
	for i, name := range want {
		if got[i].Use != name {
			t.Errorf("index %d: got %q want %q", i, got[i].Use, name)
		}
	}
}

func TestGetEntityCommandsSortedByName(t *testing.T) {
	prev := entityCommands
	entityCommands = make(map[string]*EntityCommand)
	t.Cleanup(func() { entityCommands = prev })

	RegisterEntityCommand(EntityCommand{
		Name: "zeta", EntityType: "zeta",
		Command: &cobra.Command{Use: "zeta"},
	})
	RegisterEntityCommand(EntityCommand{
		Name: "alpha", EntityType: "alpha",
		Command: &cobra.Command{Use: "alpha"},
	})
	RegisterEntityCommand(EntityCommand{
		Name: "mu", EntityType: "mu",
		Command: &cobra.Command{Use: "mu"},
	})

	got := GetEntityCommands()
	want := []string{"alpha", "mu", "zeta"}
	if len(got) != len(want) {
		t.Fatalf("len=%d want %d", len(got), len(want))
	}
	for i, name := range want {
		if got[i].Name != name {
			t.Errorf("index %d: got %q want %q", i, got[i].Name, name)
		}
	}
}
