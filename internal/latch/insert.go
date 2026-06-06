package latch

type InsertContext struct {
        Path SearchPath
        Split bool
}

func NeedSplit(
        freeSpace int,
) bool {
        return freeSpace <= 1
}
