package core

func RecursivePropagate(
	parent *Page,
	right *Page,
) (
	InternalSplitResult,
	bool,
) {

	return InternalSplit(
		parent,
		right,
	)
}
