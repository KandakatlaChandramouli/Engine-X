package core

import "testing"

func TestAllocatePageIDStress(
	t *testing.T,
) {

	pager := NewPager()

	lastID := uint64(0)

	for i := 0; i < 1000; i++ {

		id :=
			pager.AllocatePageID()

		if id <= lastID {
			t.Fatalf(
				"allocation regression %d <= %d",
				id,
				lastID,
			)
		}

		var page Page

		InitPage(
			&page,
			id,
		)

		pager.Add(
			&page,
		)

		lastID = id
	}
}
