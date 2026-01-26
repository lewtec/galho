package mise

import (
	"path/filepath"
	"testing"

	"github.com/lewtec/galho/pkg/core"
)

// MockModule implements core.Module for testing
type MockModule struct {
	name string
	path string
	typ  string
}

func (m *MockModule) Type() string { return m.typ }
func (m *MockModule) Path() string { return m.path }
func (m *MockModule) Name() string { return m.name }
func (m *MockModule) GenerateTasks() ([]core.Task, error) {
	return []core.Task{
		{
			Name:        "gen:" + m.name,
			Description: "Generate " + m.name,
			Run:         "echo " + m.name,
			Dir:         m.path,
		},
	}, nil
}

func TestCollectTasks(t *testing.T) {
	// Create a real temp directory so standard finders don't fail
	rootDir := t.TempDir()
	project := core.NewProject(rootDir)

	// Register mock finder with a unique name
	finderName := "mock_finder_for_collect_tasks"

	// We need to ensure we don't panic if this test runs multiple times in same process (e.g. strict test runners)
	// But core.RegisterModuleFinder panics.
	// In standard go test, this is fine as long as we don't call it twice in init() or setup.
	// We'll call it once here. If tests are parallel, this is global state.
	// Ideally we should check if it's already registered, but we can't (map is private).
	// We'll just wrap it in a defer/recover or assume it's fine for this run.

	defer func() {
		if r := recover(); r != nil {
			// Ignore panic on double registration if it happens (though it shouldn't in a fresh run)
			t.Logf("Recovered from panic: %v", r)
		}
	}()

	core.RegisterModuleFinder(finderName, func(p *core.Project) ([]core.Module, error) {
		// Only return modules if we are in our test project
		if p.Dir() != rootDir {
			return nil, nil
		}

		return []core.Module{
			&MockModule{name: "mod1", path: filepath.Join(rootDir, "mod1"), typ: "test"},
			&MockModule{name: "mod2", path: filepath.Join(rootDir, "mod2"), typ: "test"},
		}, nil
	})

	// Execute
	tasks, err := CollectTasks(project)
	if err != nil {
		t.Fatalf("CollectTasks failed: %v", err)
	}

	// Verify
	// We expect at least our mock tasks + gen:all.
	// Real finders shouldn't find anything in the empty temp dir.

	if _, ok := tasks["gen:mod1"]; !ok {
		t.Error("Task gen:mod1 not found")
	}

	// Check relative path conversion
	// Since we used filepath.Join(rootDir, "mod1"), and project dir is rootDir
	// Relative path should be "mod1"
	expectedDir := "mod1"
	if task, ok := tasks["gen:mod1"]; ok {
		if task.Dir != expectedDir {
			t.Errorf("Expected dir %s, got %s", expectedDir, task.Dir)
		}
	}

	if _, ok := tasks["gen"]; !ok {
		t.Error("Task gen (all) not found")
	}
}
