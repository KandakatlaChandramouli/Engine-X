package core

import "testing"

func TestMultiLevelTreeStructure(
	t *testing.T,
) {

	var root Page
	var child Page
	var grandchild Page

	InitInternalPage(
		&root,
		1,
	)

	InitInternalPage(
		&child,
		2,
	)

	InitPage(
		&grandchild,
		3,
	)

	SetLeftChild(
		&root,
		2,
	)

	if !InternalInsert(
		&root,
		[]byte("m"),
		2,
	) {
		t.Fatal()
	}

	SetLeftChild(
		&child,
		3,
	)

	if !InternalInsert(
		&child,
		[]byte("m"),
		3,
	) {
		t.Fatal()
	}

	if PageType(
		&root,
	) != PageTypeInternal {
		t.Fatal()
	}

	if PageType(
		&child,
	) != PageTypeInternal {
		t.Fatal()
	}

	if PageType(
		&grandchild,
	) != PageTypeLeaf {
		t.Fatal()
	}
}
