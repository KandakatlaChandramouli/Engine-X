package core

type Pager struct {
	pages map[uint64]*Page
}

func NewPager() *Pager {

	return &Pager{
		pages: make(map[uint64]*Page),
	}
}

func (p *Pager) Add(
	page *Page,
) {

	p.pages[PageID(page)] = page
}

func (p *Pager) Get(
	id uint64,
) (
	*Page,
	bool,
) {

	page,
		ok :=
		p.pages[id]

	return page,
		ok
}
