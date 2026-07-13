package mise

import (
	"bytes"
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
