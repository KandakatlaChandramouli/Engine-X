package core

import "testing"

func TestHandleRootSplit(
	t *testing.T,
) {

	var root Page

	InitPage(
		&root,
		1,
	)

	split :=
		InsertRecursiveResult{
			Split:        true,
			SeparatorKey: []byte("m"),
			RightPageID:  2,
		}

	if !HandleRootSplit(
		&root,
		split,
	) {
		t.Fatal()
	}

	if PageType(&root) != PageTypeInternal {
		t.Fatal()
	}
}
