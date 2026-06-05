package core

import "testing"

func TestRepeatedSplitsDoNotBreakSearch(
	t *testing.T,
) {

	pager := NewPager()

	var root Page

	InitPage(
		&root,
		1,
	)

	pager.Add(
		&root,
	)

	for i := 0; i < 500; i++ {

		key :=
			[]byte(
				string(
					rune('a' + (i % 26)),
				),
			)

		_,
			ok :=
			InsertRecursive(
				pager,
				1,
				key,
				[]byte("value"),
			)

		if !ok {
			t.Fatal()
		}
	}

	_, ok :=
		BTreeSearch(
			pager,
			1,
			[]byte("z"),
		)

	if !ok {
		t.Fatal()
	}
}
