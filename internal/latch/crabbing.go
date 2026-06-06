package latch

type CrabbingPath struct {
        Pages []uint64
}

func (c *CrabbingPath) Push(
        pageID uint64,
) {
        c.Pages = append(
                c.Pages,
                pageID,
        )
}

func (c *CrabbingPath) Pop() uint64 {

        n := len(c.Pages)

        if n == 0 {
                return 0
        }

        v := c.Pages[n-1]

        c.Pages =
                c.Pages[:n-1]

        return v
}
