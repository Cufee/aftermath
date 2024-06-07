package permissions

import (
	"fmt"
	"math/big"
	"strings"
)

const version = "v4"

type Permissions struct {
	value *big.Int
}

func (p Permissions) Encode() string {
	text, _ := p.value.MarshalText()
	return fmt.Sprintf("%s/%s", version, string(text))
}

func Parse(input string, fallback Permissions) Permissions {
	split := strings.Split(input, "/")
	if !strings.HasPrefix(input, version+"/") || len(split) != 2 {
		return fallback
	}

	value := big.NewInt(0)
	err := value.UnmarshalText([]byte(split[1]))
	if err != nil {
		return fallback
	}
	return Permissions{value}
}

var (
	Blank Permissions = Permissions{big.NewInt(0)}
)

func (p Permissions) Has(permission Permissions) bool {
	result := big.NewInt(0)
	return result.And(p.value, permission.value).Cmp(permission.value) == 0
}

func (p Permissions) Add(permission Permissions) Permissions {
	result := big.NewInt(0)
	return Permissions{result.Or(p.value, permission.value)}
}

func (p Permissions) Remove(permission Permissions) Permissions {
	result := big.NewInt(0)
	return Permissions{result.AndNot(p.value, permission.value)}
}

func (p Permissions) Toggle(permission Permissions) Permissions {
	result := big.NewInt(0)
	return Permissions{result.Xor(p.value, permission.value)}
}

func (p Permissions) Set(permission Permissions, enabled bool) Permissions {
	if enabled {
		return p.Add(permission)
	}
	return p.Remove(permission)
}

func (p Permissions) String() string {
	return p.Encode()
}

func fromLsh(value uint) Permissions {
	return Permissions{big.NewInt(0).Lsh(big.NewInt(1), value)}
}
