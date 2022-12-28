package chain

const (
	DirectionForward = iota
	DirectionBackward
)

const (
	defaultStart = 0
	defaultStop  = 0
	defaultStep  = 0
	defaultEnd   = 0
)

type Chain struct {
	start     int
	stop      int
	step      int
	end       int
	direction int
}

func New() *Chain {
	return &Chain{
		start:     defaultStart,
		stop:      defaultStop,
		step:      defaultStep,
		end:       defaultEnd,
		direction: DirectionForward,
	}
}

func (c *Chain) SetDirection(direction int) {
	c.direction = direction
	if c.isOnLastPosition() {
		c.Reset()
	}
}

func (c *Chain) SetStart(start int) {
	c.start = start
}

func (c *Chain) SetStop(stop int) {
	c.stop = stop
}

func (c *Chain) SetStep(step int) {
	c.step = step
}

func (c *Chain) Next() bool {
	return !c.isOnLastPosition()
}

func (c *Chain) Bounds() (lb int, rb int) {
	switch c.direction {
	case DirectionForward:
		c.end = c.start + c.step
		if c.end > c.stop {
			c.end = c.stop
		}
		lb, rb = c.start, c.end
		c.start += c.step
		if c.start > c.stop {
			c.start = c.stop
		}
	case DirectionBackward:
		c.end = c.start
		c.start -= c.step
		if c.start < 0 {
			c.start = 0
		}
		lb, rb = c.start, c.end
	}
	return
}

func (c *Chain) Reset() {
	switch c.direction {
	case DirectionForward:
		c.start = defaultStart
		c.end = defaultEnd
	case DirectionBackward:
		c.end = c.stop
		c.start = c.stop
	}
}

func (c *Chain) isOnLastPosition() bool {
	switch c.direction {
	case DirectionForward:
		return c.end == c.stop
	case DirectionBackward:
		return c.start == defaultStart
	}
	return false
}
