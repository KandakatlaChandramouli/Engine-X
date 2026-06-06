package concurrency

type RootLatch struct {
        RootPage uint64
        Held     bool
}

func AcquireRoot(
        pageID uint64,
) RootLatch {

        return RootLatch{
                RootPage: pageID,
                Held:     true,
        }
}
