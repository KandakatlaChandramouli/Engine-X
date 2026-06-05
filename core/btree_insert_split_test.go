package core

import (
	"fmt"
	"testing"
)

func TestBTreeInsertSplitSignal(
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

	for i := 0; i < 1000; i++ {

		ok :=
			BTreeInsert(
				pager,
				&root,
				[]byte(
					fmt.Sprintf(
						"key-%04d",
						i,
					),
				),
				[]byte("x"),
			)

		if !ok {
			break
		}

		inserted++
	}

	if inserted == 0 {
		t.Fatal()
	}
}
