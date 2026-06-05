package core

import "testing"

func TestLeafSplitCanBeInsertedIntoParent(
	t *testing.T,
) {

	pager := NewPager()

	var parent Page
	var leaf Page
	var sibling Page

	InitInternalPage(
		&parent,
		1,
	)

	InitPage(
		&leaf,
		2,
	)

	InitPage(
		&sibling,
		3,
	)

	SetLeftChild(
		&parent,
		2,
	)

	pager.Add(&parent)
	pager.Add(&leaf)
	pager.Add(&sibling)

	inserted := 0

	for {

		ok :=
			LeafInsert(
				&leaf,
				[]byte(string(rune('a'+(inserted%26)))),
				[]byte("value"),
			)

		if !ok {
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

	ok =
		InsertIntoParent(
			&parent,
			split.SeparatorKey,
			PageID(&sibling),
		)

	if !ok {
		t.Fatal()
	}

	entry, ok := ReadInternalEntry(&parent, 0)

	if !ok {
		t.Fatal()
	}

	if len(entry.Key) == 0 {
		t.Fatal()
	}
}
