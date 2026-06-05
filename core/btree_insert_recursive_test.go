package core

import "testing"

func TestRootBecomesSearchableAfterPromotion(
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

			page,
				ok :=
				BTreeSearch(
					pager,
					1,
					[]byte("z"),
				)

			if !ok {
				t.Fatal()
			}

			if page == nil {
				t.Fatal()
			}

			return
		}

		inserted++
	}
}
