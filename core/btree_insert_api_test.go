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

	for i := 0; i < 1000; i++ {

		key :=
			[]byte(
				fmt.Sprintf(
					"key-%04d",
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
			t.Fatalf(
				"insert failed %d",
				i,
			)
		}
	}
}
