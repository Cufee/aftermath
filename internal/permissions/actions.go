package permissions

var (
	// Basic user actions
	UseTextCommands     Permissions = fromLsh(1)
	UseImageCommands                = fromLsh(2)
	UseRealTimeCommands             = fromLsh(3)

	// Content
	UsePromotionalPersonalContent Permissions = fromLsh(10)
	CreatePersonalContent                     = fromLsh(11)
	UpdatePersonalContent                     = fromLsh(12)
	RemovePersonalContent                     = fromLsh(13)

	// Connections
	CreatePersonalConnection Permissions = fromLsh(15)
	UpdatePersonalConnection             = fromLsh(16)
	RemovePersonalConnection             = fromLsh(17)

	// Subscriptions
	CreatePersonalSubscription Permissions = fromLsh(20)
	ExtendPersonalSubscription             = fromLsh(21)
)

// Moderation actions
var (
	// Background Presets
	CreateGlobalBackgroundPreset Permissions = fromLsh(25)
	UpdateGlobalBackgroundPreset             = fromLsh(26)
	RemoveGlobalBackgroundPreset             = fromLsh(27)

	// Manage User Content
	CreateUserPersonalContent Permissions = fromLsh(30)
	UpdateUserPersonalContent             = fromLsh(31)
	RemoveUserPersonalContent             = fromLsh(32)

	// Subscriptions
	ViewUserSubscriptions     Permissions = fromLsh(33)
	CreateUserSubscription                = fromLsh(34)
	ExtendUserSubscription                = fromLsh(35)
	TerminateUserSubscription             = fromLsh(36)

	// Connections
	ViewUserConnections  Permissions = fromLsh(37)
	CreateUserConnection             = fromLsh(38)
	UpdateUserConnection             = fromLsh(39)
	RemoveUserConnection             = fromLsh(30)

	// Restrictions
	ViewUserRestrictions      Permissions = fromLsh(45)
	CreateSoftUserRestriction             = fromLsh(46)
	CreateHardUserRestriction             = fromLsh(47)
	UpdateSoftRestriction                 = fromLsh(48)
	UpdateHardRestriction                 = fromLsh(49)
	RemoveSoftRestriction                 = fromLsh(50)
	RemoveHardRestriction                 = fromLsh(51)
)

const (
// Admin
// _ Permissions = = fromLsh(60)
)
