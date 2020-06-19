package main

import "github.com/stretchr/testify/mock"

type ComplexityUCMock struct {
	mock.Mock
}

func (c *ComplexityUCMock) calculate(tag int, iteration int) []int {
	args := c.Called(tag, iteration)
	return args.Get(0).([]int)
}
