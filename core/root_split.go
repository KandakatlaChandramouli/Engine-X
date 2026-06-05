package core

func HandleRootSplit(
	root *Page,
	split InsertRecursiveResult,
) bool {

	return PromoteRoot(
		root,
		PageID(root),
		split.RightPageID,
		split.SeparatorKey,
	)
}
