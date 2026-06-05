package core

func RecursiveSplit(
	page *Page,
	sibling *Page,
) (
	InternalSplitResult,
	bool,
) {

	return InternalSplit(
		page,
		sibling,
	)
}
