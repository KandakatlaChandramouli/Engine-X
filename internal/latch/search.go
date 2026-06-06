package latch

type SearchFrame struct {
        PageID uint64
}

type SearchPath struct {
        Frames []SearchFrame
}

func (s *SearchPath) Add(
        pageID uint64,
) {
        s.Frames =
                append(
                        s.Frames,
                        SearchFrame{
                                PageID: pageID,
                        },
                )
}

func (s *SearchPath) Depth() int {
        return len(s.Frames)
}
