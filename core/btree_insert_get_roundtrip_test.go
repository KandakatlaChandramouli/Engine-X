package core

import (
	"fmt"
	"testing"
)

func TestInsertGetRoundTrip(
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
					"key-%04d",
					inserted,
				),
			)

		value :=
			[]byte(
				fmt.Sprintf(
					"value-%04d",
					inserted,
				),
			)

		ok :=
			LeafInsert(
				&root,
				key,
				value,
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
					"key-%04d",
					i,
				),
			)

		value,
			ok :=
			BTreeGet(
				pager,
				1,
				key,
			)

		if !ok {
			t.Fatal()
		}

		expected :=
			fmt.Sprintf(
				"value-%04d",
				i,
			)

		if string(value) != expected {
			t.Fatalf(
				"expected %s got %s",
				expected,
				string(value),
			)
		}
	}
}
