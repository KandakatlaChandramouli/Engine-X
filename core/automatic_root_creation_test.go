package core

import "testing"

func TestAutomaticRootCreation(
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

		result,
			ok :=
			InsertRecursive(
				pager,
				1,
				[]byte{
					byte(inserted),
					byte(inserted + 1),
				},
				[]byte("value"),
			)

		if !ok {
			t.Fatal()
		}

		if result.Split {

			if !HandleRootSplit(
				&root,
				result,
			) {
				t.Fatal()
			}

			break
		}

		inserted++
	}

	if PageType(
		&root,
	) != PageTypeInternal {
		t.Fatal()
	}

	if InternalEntryCount(
		&root,
	) == 0 {
		t.Fatal()
	}
}
