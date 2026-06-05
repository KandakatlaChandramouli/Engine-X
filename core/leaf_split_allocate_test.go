package core

import "testing"

func TestAllocateLeafSibling(
	t *testing.T,
) {

	pager := NewPager()

	sibling,
		ok :=
		AllocateLeafSibling(
			pager,
			99,
		)

	if !ok {
		t.Fatal()
	}

	if PageID(sibling) != 99 {
		t.Fatal()
	}

	found,
		ok :=
		pager.Get(
			99,
		)

	if !ok {
		t.Fatal()
	}

	if PageID(found) != 99 {
		t.Fatal()
	}
}
