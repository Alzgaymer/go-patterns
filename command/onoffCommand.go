package main

type onCommand struct {
	device IDevice
}

func (c *onCommand) do() {
	c.device.on()
}

type offCommand struct {
	device IDevice
}

func (c *offCommand) do() {
	c.device.off()
}
