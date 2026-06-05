package core

import "testing"

func TestRecursiveParentInsert(
	t *testing.T,
) {

	pager := NewPager()

	var parent Page
	var leaf Page

	InitInternalPage(
		&parent,
		1,
	)

	InitPage(
		&leaf,
		2,
	)

	SetLeftChild(
		&parent,
		2,
	)

	pager.Add(&parent)
	pager.Add(&leaf)

	path :=
		SearchPath{
			Pages: []*Page{
				&parent,
				&leaf,
			},
		}

	ok :=
		RecursiveParentInsert(
			pager,
			path,
			InsertRecursiveResult{
				Split:        true,
				SeparatorKey: []byte("m"),
				RightPageID:  3,
			},
		)

	if !ok {
		t.Fatal()
	}

	if InternalEntryCount(
		&parent,
	) != 1 {
		t.Fatal()
	}
}
