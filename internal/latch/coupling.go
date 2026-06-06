package latch

type CoupledPath struct {
        Parent uint64
        Child  uint64
}

func SafeToRelease(
        isLeaf bool,
        freeSpace int,
) bool {

        if isLeaf {
                return freeSpace > 1
        }

        return freeSpace > 2
}

func Couple(
        parent uint64,
        child uint64,
) CoupledPath {

        return CoupledPath{
                Parent: parent,
                Child:  child,
        }
}
