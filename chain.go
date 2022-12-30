package chain

// Chain is a lower and upper bounds generator
type Chain struct {
	start    int
	stop     int
	step     int
	end      int
	backward bool
}

// SetStart sets the start interval value
func (c *Chain) SetStart(start int) {
	c.start = start
}

// SetStop sets the maximum interval value
func (c *Chain) SetStop(stop int) {
	c.stop = stop
}

// SetStep sets the step value for every iteration
func (c *Chain) SetStep(step int) {
	c.step = step
}

// Next returns true or false whether there are any subintervals left
func (c *Chain) Next() bool {
	return !c.isOnLastPosition()
}

// Bounds returns the lower and upper bounds of current subinterval
func (c *Chain) Bounds() (lb int, rb int) {
	if !c.backward {
		c.end = c.start + c.step
		if c.end > c.stop {
			c.end = c.stop
		}
		lb, rb = c.start, c.end
		c.start += c.step
		if c.start > c.stop {
			c.start = c.stop
		}
	} else {
		c.end = c.start
		c.start -= c.step
		if c.start < 0 {
			c.start = 0
		}
		lb, rb = c.start, c.end
	}
	return
}

// Reset resets current start and end positions according to the chosen direction
func (c *Chain) Reset() {
	if !c.backward {
		c.start = 0
		c.end = 0
	} else {
		c.end = c.stop
		c.start = c.stop
	}
}

// Reverse changes the direction to iterate
func (c *Chain) Reverse() {
	c.backward = !c.backward
	if c.isOnLastPosition() {
		c.Reset()
	}
}

func (c *Chain) isOnLastPosition() bool {
	if !c.backward {
		return c.end == c.stop
	}
	return c.start == 0
}
