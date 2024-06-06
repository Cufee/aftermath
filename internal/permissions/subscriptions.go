package permissions

var (
	SubscriptionAftermathPlus = UseRealTimeCommands.
					Add(CreatePersonalContent).
					Add(UpdatePersonalContent).
					Add(RemovePersonalContent)

	SubscriptionAftermathPro = SubscriptionAftermathPlus
)
