package core

import "testing"

func TestInsertRecursivePerformsLeafSplit(
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

	inserted := 0

	for {

		result,
			ok :=
			InsertRecursive(
				pager,
				1,
				[]byte(string(rune('a'+(inserted%26)))),
				[]byte("value"),
			)

		if !ok {
			t.Fatal()
		}

		if result.Split {

			if len(result.SeparatorKey) == 0 {
				t.Fatal()
			}

			if result.RightPageID == 0 {
				t.Fatal()
			}

			return
		}

		inserted++
	}
}
