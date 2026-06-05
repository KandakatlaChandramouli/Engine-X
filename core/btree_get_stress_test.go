package core

import (
	"fmt"
	"testing"
)

func TestBTreeGetStress(
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

		key :=
			[]byte(
				fmt.Sprintf(
					"k-%04d",
					inserted,
				),
			)

		ok :=
			LeafInsert(
				&root,
				key,
				[]byte("value"),
			)

		if !ok {
			break
		}

		inserted++
	}

	for i := 0; i < inserted; i++ {

		key :=
			[]byte(
				fmt.Sprintf(
					"k-%04d",
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
				"missing key %s",
				string(key),
			)
		}
	}
}
