package core

import "testing"

func TestAllocatorGrowth(
	t *testing.T,
) {

	var meta Page
	var freelist Page

	InitMeta(
		&meta,
	)

	InitFreelist(
		&freelist,
	)

	id :=
		AllocatePage(
			&meta,
			&freelist,
		)

	if id != 4 {
		t.Fatalf(
			"expected 4 got %d",
			id,
		)
	}

	id =
		AllocatePage(
			&meta,
			&freelist,
		)

	if id != 5 {
		t.Fatalf(
			"expected 5 got %d",
			id,
		)
	}
}

func TestAllocatorReuse(
	t *testing.T,
) {

	var meta Page
	var freelist Page

	InitMeta(
		&meta,
	)

	InitFreelist(
		&freelist,
	)

	ok :=
		FreePage(
			&freelist,
			100,
		)

	if !ok {
		t.Fatal(
			"free failed",
		)
	}

	id :=
		AllocatePage(
			&meta,
			&freelist,
		)

	if id != 100 {
		t.Fatalf(
			"expected 100 got %d",
			id,
		)
	}
}
