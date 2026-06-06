package mvcc

func VisibleToTransaction(
        version *Version,
        tx *Transaction,
) bool {

        if version == nil {
                return false
        }

        if version.BeginTS > tx.StartTS {
                return false
        }

        if version.EndTS != 0 &&
                tx.StartTS >= version.EndTS {
                return false
        }

        if version.Deleted {
                return false
        }

        return true
}

func HasWriteConflict(
        head *Version,
        tx *Transaction,
) bool {

        if head == nil {
                return false
        }

        return head.BeginTS > tx.StartTS
}
