package concurrency

type SearchContext struct {
        RootPage uint64
        Key      uint64
}

func SearchPath(
        root uint64,
        key uint64,
) SearchContext {

        return SearchContext{
                RootPage: root,
                Key:      key,
        }
}
