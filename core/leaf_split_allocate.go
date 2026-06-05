package core

func AllocateLeafSibling(
	pager *Pager,
	pageID uint64,
) (
	*Page,
	bool,
) {

	if _, ok := pager.Get(pageID); ok {
		return nil,
			false
	}

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
