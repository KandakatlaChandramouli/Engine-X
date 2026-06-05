package core

import "testing"

func TestPropagateSplitInsert(
	t *testing.T,
) {

	var parent Page

	InitInternalPage(
		&parent,
		1,
	)

	SetLeftChild(
		&parent,
		100,
	)

	if !InternalInsert(
		&parent,
		[]byte("m"),
		200,
	) {
		t.Fatal()
	}

	result,
		ok :=
		PropagateSplit(
			&parent,
			[]byte("g"),
			300,
		)

	if !ok {
		t.Fatal()
	}

	if string(result.SeparatorKey) != "g" {
		t.Fatal()
	}

	if PageSlotCount(&parent) != 2 {
		t.Fatal()
	}

	e0, _ := ReadInternalEntry(&parent, 0)
	e1, _ := ReadInternalEntry(&parent, 1)

	if string(e0.Key) != "g" {
		t.Fatal(string(e0.Key))
	}

	if string(e1.Key) != "m" {
		t.Fatal(string(e1.Key))
	}
}
