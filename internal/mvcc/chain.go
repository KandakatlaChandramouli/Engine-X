package mvcc

func ReadVersion(
        head *Version,
        snapshot Snapshot,
) *Version {

        current := head

        for current != nil {

                if Visible(
                        current,
                        snapshot.ReadTS,
                ) {
                        return current
                }

                current =
                        current.Prev
        }

        return nil
}

func InsertVersion(
        head *Version,
        next *Version,
) *Version {

        next.Prev = head

        return next
}
