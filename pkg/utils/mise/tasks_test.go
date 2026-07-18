package mise

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	toml "github.com/pelletier/go-toml/v2"
)

func TestMiseConfigEncodesTasksTable(t *testing.T) {
	cfg := MiseConfig{
		Tasks: map[string]MiseTask{
			"gen:crm:db:sqlc": {
				Description: "Generate SQLC code for crm",
				Run:         "sqlc generate -f sqlc.yaml",
				Dir:         "internal/crm/db",
			},
			"gen": {
				Description: "Generate all code",
				Depends:     []string{"gen:*"},
			},
		},
	}

	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(cfg); err != nil {
		t.Fatalf("encode: %v", err)
	}
	out := buf.String()

	// mise expects nested [tasks.*], not top-level task tables
	if !strings.Contains(out, "[tasks.") {
		t.Fatalf("expected [tasks.*] tables, got:\n%s", out)
	}
	if strings.Contains(out, "[gen:crm:db:sqlc]") {
		t.Fatalf("task leaked to top-level table:\n%s", out)
	}
	if !strings.Contains(out, "sqlc generate -f sqlc.yaml") {
		t.Fatalf("missing run command:\n%s", out)
	}
}

func TestWriteGalhoTomlAtomic(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "galho.toml")

	// Seed an existing file that must not be left truncated on success.
	if err := os.WriteFile(path, []byte("stale = true\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg := MiseConfig{
		Tasks: map[string]MiseTask{
			"gen": {
				Description: "Generate all code",
				Depends:     []string{"gen:*"},
			},
		},
	}
	if err := writeGalhoTomlAtomic(path, cfg); err != nil {
		t.Fatalf("write: %v", err)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	out := string(raw)
	if strings.Contains(out, "stale") {
		t.Fatalf("stale content retained:\n%s", out)
	}
	if !strings.Contains(out, "[tasks.") {
		t.Fatalf("expected [tasks.*] tables, got:\n%s", out)
	}

	// No leftover temp files in the directory.
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Name() != "galho.toml" {
		names := make([]string, 0, len(entries))
		for _, e := range entries {
			names = append(names, e.Name())
		}
		t.Fatalf("expected only galho.toml, got %v", names)
	}
}

func TestGenerateTasksTomlWritesUnderProjectRoot(t *testing.T) {
	// Minimal galho project: .galho marker, no modules.
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, ".galho"), []byte{}, 0o644); err != nil {
		t.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(cwd) })

	if err := GenerateTasksToml(""); err != nil {
		t.Fatalf("GenerateTasksToml: %v", err)
	}

	raw, err := os.ReadFile(filepath.Join(root, ".mise", "galho.toml"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "[tasks.") {
		t.Fatalf("expected tasks tables, got:\n%s", raw)
	}
	if !strings.Contains(string(raw), "Generate all code") {
		t.Fatalf("expected gen task, got:\n%s", raw)
	}
}
