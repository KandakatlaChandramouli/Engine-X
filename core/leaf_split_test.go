package core

import "testing"

func TestLeafSplit(
	t *testing.T,
) {

	var left Page
	var right Page

	InitPage(
		&left,
		10,
	)

	InitPage(
		&right,
		11,
	)

	SetPageType(
		&left,
		PageTypeLeaf,
	)

	SetPageType(
		&right,
		PageTypeLeaf,
	)

	keys := []string{
		"apple",
		"banana",
		"cat",
		"dog",
		"elephant",
		"fox",
	}

	for _, k := range keys {

		if !LeafInsert(
			&left,
			[]byte(k),
			[]byte(k),
		) {
			t.Fatal()
		}
	}

	result,
		ok :=
		LeafSplit(
			&left,
			&right,
		)

	if !ok {
		t.Fatal()
	}

	if len(result.SeparatorKey) == 0 {
		t.Fatal()
	}

	if LeafEntryCount(&left) != 3 {
		t.Fatal()
	}

	if LeafEntryCount(&right) != 3 {
		t.Fatal()
	}
}
