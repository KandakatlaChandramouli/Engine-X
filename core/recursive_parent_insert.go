package core

func RecursiveParentInsert(
	pager *Pager,
	path SearchPath,
	split InsertRecursiveResult,
) bool {

	if len(path.Pages) < 2 {
		return true
	}

	parent := path.Pages[len(path.Pages)-2]

	return InsertIntoParent(
		parent,
		split.SeparatorKey,
		split.RightPageID,
	)
}
