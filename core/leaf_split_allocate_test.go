package core

import "testing"

func TestAllocateLeafSiblingRejectsDuplicateID(
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

	if sibling == nil {
		t.Fatal()
	}

	_,
		ok =
		AllocateLeafSibling(
			pager,
			99,
		)

	if ok {
		t.Fatal()
	}
}
