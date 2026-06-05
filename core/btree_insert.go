package core

func BTreeInsert(
	p *Page,
	key []byte,
	value []byte,
) (
	BTreeInsertResult,
	bool,
) {

	ok :=
		LeafInsert(
			p,
			key,
			value,
		)

	if !ok {

		return BTreeInsertResult{
			Split: true,
		}, true
	}

	return BTreeInsertResult{
		Split: false,
	}, true
}
