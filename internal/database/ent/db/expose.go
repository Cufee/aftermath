// Code generated by ent, DO NOT EDIT.

package db

import "fmt"

func (a *Account) AssignValues(columns []string, values []any) error {
	if a == nil {
		return fmt.Errorf("Account(nil)")
	}
	return a.assignValues(columns, values)
}
func (a *Account) ScanValues(columns []string) ([]any, error) {
	if a == nil {
		return nil, fmt.Errorf("Account(nil)")
	}
	return a.scanValues(columns)
}

func (as *AccountSnapshot) AssignValues(columns []string, values []any) error {
	if as == nil {
		return fmt.Errorf("AccountSnapshot(nil)")
	}
	return as.assignValues(columns, values)
}
func (as *AccountSnapshot) ScanValues(columns []string) ([]any, error) {
	if as == nil {
		return nil, fmt.Errorf("AccountSnapshot(nil)")
	}
	return as.scanValues(columns)
}

func (ac *AppConfiguration) AssignValues(columns []string, values []any) error {
	if ac == nil {
		return fmt.Errorf("AppConfiguration(nil)")
	}
	return ac.assignValues(columns, values)
}
func (ac *AppConfiguration) ScanValues(columns []string) ([]any, error) {
	if ac == nil {
		return nil, fmt.Errorf("AppConfiguration(nil)")
	}
	return ac.scanValues(columns)
}

func (ac *ApplicationCommand) AssignValues(columns []string, values []any) error {
	if ac == nil {
		return fmt.Errorf("ApplicationCommand(nil)")
	}
	return ac.assignValues(columns, values)
}
func (ac *ApplicationCommand) ScanValues(columns []string) ([]any, error) {
	if ac == nil {
		return nil, fmt.Errorf("ApplicationCommand(nil)")
	}
	return ac.scanValues(columns)
}

func (an *AuthNonce) AssignValues(columns []string, values []any) error {
	if an == nil {
		return fmt.Errorf("AuthNonce(nil)")
	}
	return an.assignValues(columns, values)
}
func (an *AuthNonce) ScanValues(columns []string) ([]any, error) {
	if an == nil {
		return nil, fmt.Errorf("AuthNonce(nil)")
	}
	return an.scanValues(columns)
}

func (c *Clan) AssignValues(columns []string, values []any) error {
	if c == nil {
		return fmt.Errorf("Clan(nil)")
	}
	return c.assignValues(columns, values)
}
func (c *Clan) ScanValues(columns []string) ([]any, error) {
	if c == nil {
		return nil, fmt.Errorf("Clan(nil)")
	}
	return c.scanValues(columns)
}

func (ct *CronTask) AssignValues(columns []string, values []any) error {
	if ct == nil {
		return fmt.Errorf("CronTask(nil)")
	}
	return ct.assignValues(columns, values)
}
func (ct *CronTask) ScanValues(columns []string) ([]any, error) {
	if ct == nil {
		return nil, fmt.Errorf("CronTask(nil)")
	}
	return ct.scanValues(columns)
}

func (di *DiscordInteraction) AssignValues(columns []string, values []any) error {
	if di == nil {
		return fmt.Errorf("DiscordInteraction(nil)")
	}
	return di.assignValues(columns, values)
}
func (di *DiscordInteraction) ScanValues(columns []string) ([]any, error) {
	if di == nil {
		return nil, fmt.Errorf("DiscordInteraction(nil)")
	}
	return di.scanValues(columns)
}

func (gm *GameMap) AssignValues(columns []string, values []any) error {
	if gm == nil {
		return fmt.Errorf("GameMap(nil)")
	}
	return gm.assignValues(columns, values)
}
func (gm *GameMap) ScanValues(columns []string) ([]any, error) {
	if gm == nil {
		return nil, fmt.Errorf("GameMap(nil)")
	}
	return gm.scanValues(columns)
}

func (gm *GameMode) AssignValues(columns []string, values []any) error {
	if gm == nil {
		return fmt.Errorf("GameMode(nil)")
	}
	return gm.assignValues(columns, values)
}
func (gm *GameMode) ScanValues(columns []string) ([]any, error) {
	if gm == nil {
		return nil, fmt.Errorf("GameMode(nil)")
	}
	return gm.scanValues(columns)
}

func (ls *LeaderboardScore) AssignValues(columns []string, values []any) error {
	if ls == nil {
		return fmt.Errorf("LeaderboardScore(nil)")
	}
	return ls.assignValues(columns, values)
}
func (ls *LeaderboardScore) ScanValues(columns []string) ([]any, error) {
	if ls == nil {
		return nil, fmt.Errorf("LeaderboardScore(nil)")
	}
	return ls.scanValues(columns)
}

func (s *Session) AssignValues(columns []string, values []any) error {
	if s == nil {
		return fmt.Errorf("Session(nil)")
	}
	return s.assignValues(columns, values)
}
func (s *Session) ScanValues(columns []string) ([]any, error) {
	if s == nil {
		return nil, fmt.Errorf("Session(nil)")
	}
	return s.scanValues(columns)
}

func (u *User) AssignValues(columns []string, values []any) error {
	if u == nil {
		return fmt.Errorf("User(nil)")
	}
	return u.assignValues(columns, values)
}
func (u *User) ScanValues(columns []string) ([]any, error) {
	if u == nil {
		return nil, fmt.Errorf("User(nil)")
	}
	return u.scanValues(columns)
}

func (uc *UserConnection) AssignValues(columns []string, values []any) error {
	if uc == nil {
		return fmt.Errorf("UserConnection(nil)")
	}
	return uc.assignValues(columns, values)
}
func (uc *UserConnection) ScanValues(columns []string) ([]any, error) {
	if uc == nil {
		return nil, fmt.Errorf("UserConnection(nil)")
	}
	return uc.scanValues(columns)
}

func (uc *UserContent) AssignValues(columns []string, values []any) error {
	if uc == nil {
		return fmt.Errorf("UserContent(nil)")
	}
	return uc.assignValues(columns, values)
}
func (uc *UserContent) ScanValues(columns []string) ([]any, error) {
	if uc == nil {
		return nil, fmt.Errorf("UserContent(nil)")
	}
	return uc.scanValues(columns)
}

func (us *UserSubscription) AssignValues(columns []string, values []any) error {
	if us == nil {
		return fmt.Errorf("UserSubscription(nil)")
	}
	return us.assignValues(columns, values)
}
func (us *UserSubscription) ScanValues(columns []string) ([]any, error) {
	if us == nil {
		return nil, fmt.Errorf("UserSubscription(nil)")
	}
	return us.scanValues(columns)
}

func (v *Vehicle) AssignValues(columns []string, values []any) error {
	if v == nil {
		return fmt.Errorf("Vehicle(nil)")
	}
	return v.assignValues(columns, values)
}
func (v *Vehicle) ScanValues(columns []string) ([]any, error) {
	if v == nil {
		return nil, fmt.Errorf("Vehicle(nil)")
	}
	return v.scanValues(columns)
}

func (va *VehicleAverage) AssignValues(columns []string, values []any) error {
	if va == nil {
		return fmt.Errorf("VehicleAverage(nil)")
	}
	return va.assignValues(columns, values)
}
func (va *VehicleAverage) ScanValues(columns []string) ([]any, error) {
	if va == nil {
		return nil, fmt.Errorf("VehicleAverage(nil)")
	}
	return va.scanValues(columns)
}

func (vs *VehicleSnapshot) AssignValues(columns []string, values []any) error {
	if vs == nil {
		return fmt.Errorf("VehicleSnapshot(nil)")
	}
	return vs.assignValues(columns, values)
}
func (vs *VehicleSnapshot) ScanValues(columns []string) ([]any, error) {
	if vs == nil {
		return nil, fmt.Errorf("VehicleSnapshot(nil)")
	}
	return vs.scanValues(columns)
}

func (ws *WidgetSettings) AssignValues(columns []string, values []any) error {
	if ws == nil {
		return fmt.Errorf("WidgetSettings(nil)")
	}
	return ws.assignValues(columns, values)
}
func (ws *WidgetSettings) ScanValues(columns []string) ([]any, error) {
	if ws == nil {
		return nil, fmt.Errorf("WidgetSettings(nil)")
	}
	return ws.scanValues(columns)
}
