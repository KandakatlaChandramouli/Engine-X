package core

import "testing"

func TestBTreeInsertSplitSignal(
	t *testing.T,
) {

	var p Page

	InitPage(
		&p,
		1,
	)

	inserted := 0

	for {

		result,
			ok :=
			BTreeInsert(
				&p,
				[]byte(string(rune('a'+inserted))),
				[]byte("x"),
			)

		if !ok {
			break
		}

		if result.Split {
			return
		}

		inserted++

		if inserted > 1000 {
			t.Fatal()
		}
	}
}
