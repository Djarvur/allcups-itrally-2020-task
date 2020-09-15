package game

// Config contains game configuration.
type Config struct {
	Seed              int64
	MaxActiveLicenses int
	Density           int // One treasure per Density cells.
	SizeX             int
	SizeY             int
	Depth             uint8
}

func (cfg Config) treasures() int {
	return cfg.area() * int(cfg.Depth) / cfg.Density
}

func (cfg Config) area() int {
	return cfg.SizeX * cfg.SizeY
}

func (cfg Config) volume() int {
	return cfg.area() * int(cfg.Depth)
}

// TotalCash returns amount of coins required to cash all treasures in
// worst case (all treasures cost as much as possible).
func (cfg Config) totalCash() (total int) {
	treasures := cfg.treasures()
	for depth := cfg.Depth; treasures > 0 && depth > 0; depth-- {
		var max int
		if treasures <= cfg.area() {
			max = treasures
		} else {
			max = cfg.area()
		}
		_, cost := treasureCostAt(depth)
		total += max * cost
		treasures -= max
	}
	if treasures > 0 {
		panic("not all treasures were counted")
	}
	return total
}
