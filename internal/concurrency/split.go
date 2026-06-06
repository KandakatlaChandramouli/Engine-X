package concurrency

type SplitContext struct {
        PageID uint64
        Safe   bool
}

func CanSplit(
        pageID uint64,
        safe bool,
) SplitContext {

        return SplitContext{
                PageID: pageID,
                Safe:   safe,
        }
}
