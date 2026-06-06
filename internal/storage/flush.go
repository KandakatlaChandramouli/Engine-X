package storage

func FlushCandidates(dpt *DirtyPageTable) []uint64 {
	var pages []uint64

	for pageID := range dpt.pages {
		pages = append(pages, pageID)
	}

	return pages
}
