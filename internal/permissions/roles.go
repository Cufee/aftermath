package permissions

var (
	User Permissions = UseTextCommands.
		Add(UseImageCommands).
		Add(UsePromotionalPersonalContent).
		Add(CreatePersonalContent).
		Add(UpdatePersonalContent).
		Add(RemovePersonalContent).
		Add(CreatePersonalConnection).
		Add(UpdatePersonalConnection).
		Add(RemovePersonalConnection).
		Add(CreatePersonalSubscription).
		Add(ExtendPersonalSubscription)

	ContentModerator = User.
				Add(RemoveUserPersonalContent).
				Add(RemoveGlobalBackgroundPreset).
				Add(CreateSoftUserRestriction)

	GlobalModerator = ContentModerator.
			Add(UseDebugFeatures).
			Add(ViewUserSubscriptions).
			Add(CreateUserSubscription).
			Add(ExtendUserSubscription).
			Add(TerminateUserSubscription).
			Add(ViewUserConnections).
			Add(CreateUserConnection).
			Add(UpdateUserConnection).
			Add(RemoveUserConnection).
			Add(ViewUserRestrictions).
			Add(UpdateSoftRestriction).
			Add(RemoveSoftRestriction).
			Add(CreateHardUserRestriction)

	GlobalAdmin = GlobalModerator.
			Add(DebugAccess).
			Add(CreateGlobalBackgroundPreset).
			Add(UpdateGlobalBackgroundPreset).
			Add(RemoveGlobalBackgroundPreset).
			Add(CreateUserPersonalContent).
			Add(UpdateUserPersonalContent).
			Add(RemoveUserPersonalContent).
			Add(ViewUserSubscriptions).
			Add(CreateUserSubscription).
			Add(ExtendUserSubscription).
			Add(TerminateUserSubscription).
			Add(ViewUserConnections).
			Add(CreateUserConnection).
			Add(UpdateUserConnection).
			Add(RemoveUserConnection).
			Add(ViewUserRestrictions).
			Add(CreateSoftUserRestriction).
			Add(CreateHardUserRestriction).
			Add(UpdateSoftRestriction).
			Add(UpdateHardRestriction).
			Add(RemoveSoftRestriction).
			Add(RemoveHardRestriction)

	DebugAccess = UseDebugFeatures.
			Add(ViewTaskLogs)
)
