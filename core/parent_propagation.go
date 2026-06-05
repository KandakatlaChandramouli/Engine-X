package core

func ParentAwarePropagation(
	child *Page,
) (
	uint64,
	bool,
) {

	parent :=
		ParentPageID(
			child,
		)

	if parent == 0 {
		return 0,
			false
	}

	return parent,
		true
}
