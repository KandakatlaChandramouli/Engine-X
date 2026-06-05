package core

import "testing"

func TestFindInsertionPath(
	t *testing.T,
) {

	pager := NewPager()

	var root Page
	var internal Page
	var leaf Page

	InitInternalPage(&root, 1)
	InitInternalPage(&internal, 2)
	InitPage(&leaf, 3)

	SetLeftChild(&root, 2)
	SetLeftChild(&internal, 3)

	pager.Add(&root)
	pager.Add(&internal)
	pager.Add(&leaf)

	path,
		ok :=
		FindInsertionPath(
			pager,
			1,
			[]byte("abc"),
		)

	if !ok {
		t.Fatal()
	}

	if len(path.Pages) != 3 {
		t.Fatal()
	}
}
