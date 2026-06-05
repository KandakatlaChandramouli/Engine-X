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

	return BTreeInsertResult{
		Split: false,
	}, ok
}
