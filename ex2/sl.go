package main

type TagsConfig map[int]int

func (t *TagsConfig) config() TagsConfig {
	return map[int]int{
		0:  4,
		5:  3,
		10: 2,
		15: 1,
	}
}

type Config struct {
	tags TagsConfig
}

func NewConfig() *Config {
	return &Config{
		tags: TagsConfig{},
	}
}

//-------------

type Sl struct {
	cfg *Config

	complexityUC ComplexityUCInterface
}

func NewSl(cfg *Config) *Sl {
	return &Sl{
		cfg:          cfg,
		complexityUC: NewComplexityUC(cfg.tags),
	}
}
