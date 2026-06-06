package lockmgr

type WaitQueue struct {
        requests []Request
}

func (q *WaitQueue) Enqueue(
        r Request,
) {
        q.requests =
                append(
                        q.requests,
                        r,
                )
}

func (q *WaitQueue) Dequeue() (
        Request,
        bool,
) {

        if len(q.requests) == 0 {
                return Request{},
                        false
        }

        r := q.requests[0]

        q.requests =
                q.requests[1:]

        return r,
                true
}

func (q *WaitQueue) Len() int {
        return len(q.requests)
}
