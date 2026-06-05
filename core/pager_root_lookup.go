package core

func PagerRootLookup(
	pager *Pager,
	rootID uint64,
) (
	*Page,
	bool,
) {

	return pager.Get(
		rootID,
	)
}
