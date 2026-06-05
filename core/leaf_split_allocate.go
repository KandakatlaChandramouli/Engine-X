package core

func AllocateLeafSibling(
	pager *Pager,
	pageID uint64,
) (
	*Page,
	bool,
) {

	var sibling Page

	InitPage(
		&sibling,
		pageID,
	)

	pager.Add(
		&sibling,
	)

	return &sibling,
		true
}
