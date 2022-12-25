package chain

type Option func(*Chain)

func OptionStart(start int) func(*Chain) {
	return func(c *Chain) { c.start = start }
}

func OptionStep(step int) func(*Chain) {
	return func(c *Chain) { c.step = step }
}

type Chain struct {
	start int
	stop  int
	step  int
	end   int
}

func New(stop int, options ...Option) *Chain {
	chain := &Chain{start: 0, stop: stop, step: 1}
	for _, option := range options {
		option(chain)
	}
	return chain
}

func (c *Chain) Bounds() (lb int, rb int) {
	c.end = c.start + c.step
	if c.end > c.stop {
		c.end = c.stop
	}
	lb, rb = c.start, c.end
	c.start += c.step
	return
}

func (c *Chain) Next() bool {
	return c.start < c.stop
}
