package core

func (p *Pager) AllocatePageID() uint64 {

	var max uint64

	for id := range p.pages {

		if id > max {
			max = id
		}
	}

	return max + 1
}
