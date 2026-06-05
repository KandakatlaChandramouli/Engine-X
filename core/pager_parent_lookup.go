package core

func PagerParentLookup(
	pager *Pager,
	child *Page,
) (
	*Page,
	bool,
) {

	parentID :=
		ParentPageID(
			child,
		)

	if parentID == 0 {
		return nil,
			false
	}

	return pager.Get(
		parentID,
	)
}
