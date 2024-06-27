package builder

type parameters struct {
	nameKey string
	descKey string
}

type Param func(*parameters)

func SetNameKey(key string) Param {
	return func(p *parameters) {
		p.nameKey = key
	}
}

func SetDescKey(key string) Param {
	return func(p *parameters) {
		p.descKey = key
	}
}
