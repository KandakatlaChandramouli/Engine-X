package mvcc

func BeginTransaction(
        id uint64,
) *Transaction {

        return &Transaction{
                ID: id,
                StartTS: NextTS(),
        }
}

func CommitTransaction(
        tx *Transaction,
) {

        tx.CommitTS =
                NextTS()
}
