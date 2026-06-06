package lockmgr

type Mode uint8

const (
        Shared Mode = iota
        Exclusive
)

func Compatible(
        existing Mode,
        requested Mode,
) bool {

        if existing == Shared &&
                requested == Shared {
                return true
        }

        return false
}
