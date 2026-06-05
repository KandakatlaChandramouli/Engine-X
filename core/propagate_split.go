package core

func PropagateSplit(
	parent *Page,
	separator []byte,
	rightPageID uint64,
) (
	PropagationResult,
	bool,
) {

	ok :=
		InsertIntoParent(
			parent,
			separator,
			rightPageID,
		)

	if !ok {
		return PropagationResult{},
			false
	}

	return PropagationResult{
		SeparatorKey: separator,
		RightPageID:  rightPageID,
	}, true
}
