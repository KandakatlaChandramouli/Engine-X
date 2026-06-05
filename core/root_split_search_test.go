package core

import "testing"

func TestSearchAfterRootPromotion(
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

			if !HandleRootSplit(
				&root,
				result,
			) {
				t.Fatal()
			}

			_, ok =
				BTreeGet(
					pager,
					1,
					[]byte("z"),
				)

			if !ok {
				t.Fatal(
					"search failed after root promotion",
				)
			}

			return
		}

		inserted++
	}
}
