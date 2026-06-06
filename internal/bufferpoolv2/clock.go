package bufferpoolv2

type Clock struct {
        pages []int
        hand  int
}

func NewClock() *Clock {
        return &Clock{}
}

func (c *Clock) Add(id int) {
        c.pages = append(c.pages, id)
}

func (c *Clock) Victim() (int, bool) {
        if len(c.pages) == 0 {
                return 0, false
        }

        v := c.pages[0]

        c.pages = c.pages[1:]

        if c.hand > len(c.pages) {
                c.hand = 0
        }

        return v, true
}

func (c *Clock) Len() int {
        return len(c.pages)
}
