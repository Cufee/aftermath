package permissions

import (
	"testing"

	"github.com/matryer/is"
)

func TestPermissions(t *testing.T) {
	is := is.New(t)

	perms := Blank
	is.Equal(perms.Has(UseTextCommands), false)

	perms = perms.Add(UseTextCommands)
	is.Equal(perms.Has(UseTextCommands), true)

	perms = perms.Remove(UseTextCommands)
	is.Equal(perms.Has(UseTextCommands), false)

	perms = perms.Add(UseDebugFeatures)
	is.Equal(perms.Has(UseDebugFeatures), true)
}
