package core

import (
	"fmt"
	"testing"
)

func TestBTreeInsertAPI(
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

		key :=
			[]byte(
				fmt.Sprintf(
					"key-%06d",
					i,
				),
			)

		ok :=
			BTreeInsert(
				pager,
				&root,
				key,
				[]byte("value"),
			)

		if !ok {
			break
		}

		inserted++
	}

	if inserted == 0 {
		t.Fatal()
	}

	for i := 0; i < inserted; i++ {

		key :=
			[]byte(
				fmt.Sprintf(
					"key-%06d",
					i,
				),
			)

		_,
			ok :=
			BTreeGet(
				pager,
				1,
				key,
			)

		if !ok {
			t.Fatalf(
				"missing key %d",
				i,
			)
		}
	}
}
