package mvcc

type Transaction struct {
        ID       uint64
        StartTS  uint64
        CommitTS uint64
}
