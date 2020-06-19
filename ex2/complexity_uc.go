package main

type ComplexityUC struct {
	tags TagsConfig
}

func (c *ComplexityUC) calculate(tag int, iteration int) []int {
	var result []int
	step := 0
	for iteration >= step {
		val, ok := c.tags.config()[step]
		if ok && tag >= val {
			result = append(result, val)
		}
		step += 1
	}
	return result
}

func NewComplexityUC(tags TagsConfig) *ComplexityUC {
	return &ComplexityUC{tags: tags}
}
