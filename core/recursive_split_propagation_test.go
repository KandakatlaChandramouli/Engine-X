package core

import "testing"

func TestRecursiveSplitPropagation(
	t *testing.T,
) {

	pager := NewPager()

	var root Page
	var parent Page
	var leaf Page
	var sibling Page
	var parentSibling Page

	InitInternalPage(
		&root,
		1,
	)

	InitInternalPage(
		&parent,
		2,
	)

	InitPage(
		&leaf,
		3,
	)

	InitPage(
		&sibling,
		4,
	)

	InitInternalPage(
		&parentSibling,
		5,
	)

	pager.Add(&root)
	pager.Add(&parent)
	pager.Add(&leaf)
	pager.Add(&sibling)
	pager.Add(&parentSibling)

	SetLeftChild(
		&root,
		2,
	)

	SetLeftChild(
		&parent,
		3,
	)

	inserted := 0

	for {

		if !LeafInsert(
			&leaf,
			[]byte{
				byte(inserted),
				byte(inserted + 1),
			},
			[]byte("value"),
		) {
			break
		}

		inserted++
	}

	split,
		ok :=
		LeafSplit(
			&leaf,
			&sibling,
		)

	if !ok {
		t.Fatal()
	}

	if !InsertIntoParent(
		&parent,
		split.SeparatorKey,
		PageID(&sibling),
	) {
		t.Fatal()
	}

	for i := 0; i < 512; i++ {

		if !InternalInsert(
			&parent,
			[]byte{
				byte(i),
				byte(i + 1),
			},
			uint64(i+1000),
		) {
			break
		}
	}

	result,
		ok :=
		HandleParentOverflow(
			&parent,
			&parentSibling,
		)

	if !ok {
		t.Fatal()
	}

	if len(
		result.SeparatorKey,
	) == 0 {
		t.Fatal()
	}

	if !InsertIntoParent(
		&root,
		result.SeparatorKey,
		PageID(&parentSibling),
	) {
		t.Fatal()
	}

	if InternalEntryCount(
		&root,
	) == 0 {
		t.Fatal()
	}
}
