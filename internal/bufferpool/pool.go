package bufferpool

type Pool struct {
        Frames []Frame
}

func New(
        size int,
) *Pool {

        return &Pool{
                Frames: make(
                        []Frame,
                        size,
                ),
        }
}
