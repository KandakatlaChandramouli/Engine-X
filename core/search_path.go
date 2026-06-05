package core

type SearchPath struct {
	Pages []*Page
}

func FindInsertionPath(
	pager *Pager,
	rootID uint64,
	key []byte,
) (
	SearchPath,
	bool,
) {

	var path SearchPath

	page,
		ok :=
		pager.Get(
			rootID,
		)

	if !ok {
		return SearchPath{},
			false
	}

	for {

		path.Pages =
			append(
				path.Pages,
				page,
			)

		if PageType(page) != PageTypeInternal {
			break
		}

		childID,
			ok :=
			InternalSearch(
				page,
				key,
			)

		if !ok {
			return SearchPath{},
				false
		}

		page,
			ok =
			pager.Get(
				childID,
			)

		if !ok {
			return SearchPath{},
				false
		}
	}

	return path,
		true
}
