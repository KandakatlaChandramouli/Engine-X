package core

type Entry struct {
	Key []byte
	Val []byte
}

func ReadEntry(
	p *Page,
	slot uint16,
) (
	Entry,
	bool,
) {

	k,
		v,
		ok :=
		Get(
			p,
			slot,
		)

	if !ok {
		return Entry{}, false
	}

	return Entry{
		Key: k,
		Val: v,
	}, true
}
