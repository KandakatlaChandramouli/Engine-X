package concurrency

type InsertContext struct {
        RootPage uint64
        Key      uint64
        Value    []byte
}

func InsertPath(
        root uint64,
        key uint64,
        value []byte,
) InsertContext {

        return InsertContext{
                RootPage: root,
                Key:      key,
                Value:    value,
        }
}
