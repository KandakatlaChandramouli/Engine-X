package core

import (
	"fmt"
	"testing"
)

func TestDeepTreeGrowth(
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

	for i := 0; i < 5000; i++ {

		key :=
			[]byte(
				fmt.Sprintf(
					"key-%05d",
					i,
				),
			)

		if !LeafInsert(
			&root,
			key,
			[]byte("value"),
		) {
			break
		}
	}
}
