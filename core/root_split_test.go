package core

import "testing"

func TestPromoteRootAfterLeafSplit(
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

			ok =
				PromoteRoot(
					&root,
					1,
					result.RightPageID,
					result.SeparatorKey,
				)

			if !ok {
				t.Fatal()
			}

			if PageType(&root) != PageTypeInternal {
				t.Fatal()
			}

			return
		}

		inserted++
	}
}
