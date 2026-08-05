package scaffold

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestMustSubReturnsSubdir(t *testing.T) {
	root := fstest.MapFS{
		"_template/hello.txt": {Data: []byte("hi")},
	}
	sub := MustSub(root, "_template")
	b, err := fs.ReadFile(sub, "hello.txt")
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "hi" {
		t.Fatalf("got %q", string(b))
	}
}

func TestMustSubPanicsOnInvalidPath(t *testing.T) {
	// fs.Sub only fails eagerly for invalid paths; missing dirs are deferred to Open
	// (except embed.FS SubFS). Invalid ".." must panic via MustSub.
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for invalid path")
		}
	}()
	_ = MustSub(fstest.MapFS{}, "../nope")
}
