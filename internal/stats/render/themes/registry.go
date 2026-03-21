package themes

import (
	"github.com/cufee/aftermath/internal/render/common"
	"github.com/cufee/aftermath/internal/stats/render/themes/spring2026"
)

var registry = map[string]func() common.Theme{
	"default":    common.DefaultTheme,
	"spring2026": spring2026.Theme,
}

func GetTheme(id string) (common.Theme, bool) {
	fn, ok := registry[id]
	if !ok {
		return common.Theme{}, false
	}
	return fn(), true
}

func AvailableThemes() []string {
	var ids []string
	for id := range registry {
		ids = append(ids, id)
	}
	return ids
}
