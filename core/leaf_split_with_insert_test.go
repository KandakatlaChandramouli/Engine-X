package core

import "testing"

func TestLeafSplitPreservesOverflowRecord(
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

			found,
				ok :=
				BTreeGet(
					pager,
					1,
					[]byte(string(rune('a'+(inserted%26)))),
				)

			if !ok {
				t.Fatal(
					"overflow key lost during split",
				)
			}

			if string(found) != "value" {
				t.Fatal()
			}

			return
		}

		inserted++
	}
}
