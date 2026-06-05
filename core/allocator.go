package core

func AllocatePage(
	meta *Page,
	freelist *Page,
) uint64 {

	pageID,
		ok :=
		FreelistPop(
			freelist,
		)

	if ok {
		return pageID
	}

	last :=
		LastPageID(
			meta,
		)

	last++

	SetLastPageID(
		meta,
		last,
	)

	return last
}

func FreePage(
	freelist *Page,
	pageID uint64,
) bool {

	return FreelistPush(
		freelist,
		pageID,
	)
}
