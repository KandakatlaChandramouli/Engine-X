package core

import "testing"

func TestMetaInit(
	t *testing.T,
) {

	var p Page

	InitMeta(&p)

	if !ValidateMeta(&p) {
		t.Fatal("meta validation failed")
	}

	if RootPageID(&p) != 3 {
		t.Fatal("bad root page")
	}

	if LastPageID(&p) != 3 {
		t.Fatal("bad last page")
	}
}

func TestMetaUpdate(
	t *testing.T,
) {

	var p Page

	InitMeta(&p)

	SetLastPageID(
		&p,
		100,
	)

	if LastPageID(&p) != 100 {
		t.Fatal("update failed")
	}

	if !ValidateMeta(&p) {
		t.Fatal("crc invalid")
	}
}
