package prepare

import (
	"fmt"
	"math"
)

type valueFormat byte

const (
	valueFormatInt valueFormat = iota
	valueFormatDecimal
	valueFormatPercent
	valueFormatInvalid
)

type statValue struct {
	format valueFormat
	data   float32
}

/*
Return a string representation of the stats value
*/
func (value statValue) String() string {
	switch value.format {
	default:
		return fmt.Sprintf("%.0f", value.data)
	case valueFormatDecimal:
		return fmt.Sprintf("%.2f", value.data)
	case valueFormatPercent:
		return fmt.Sprintf("%.2f%%", value.data)
	case valueFormatInvalid:
		return "-"
	}
}

/*
Return a float32 representation of the stats value
*/
func (value statValue) Value() float32 {
	return value.data
}

/*
Check if the value is equal to InvalidValue
*/
func (value statValue) Valid() bool {
	return value.Value() != InvalidValue.Value()
}

/*
Check if the value is equal to zero
*/
func (value statValue) IsZero() bool {
	return value.Value() == 0
}

/*
InvalidValue indicates that the value is impossible to calculate
*/
var InvalidValue = statValue{valueFormatInvalid, -1}

/*
StatsFrame is a frame to structure statistics overview
*/
type StatsFrame struct {
	Battles         statValue `json:"battles" bson:"battles"`
	BattlesWon      statValue `json:"battlesWon" bson:"battlesWon"`
	BattlesSurvived statValue `json:"battlesSurvived" bson:"battlesSurvived"`
	DamageDealt     statValue `json:"damageDealt" bson:"damageDealt"`
	DamageReceived  statValue `json:"damageReceived" bson:"damageReceived"`

	ShotsHit   statValue `json:"shotsHit" bson:"shotsHit"`
	ShotsFired statValue `json:"shotsFired" bson:"shotsFired"`

	Frags                statValue `json:"frags" bson:"frags"`
	MaxFrags             statValue `json:"maxFrags" bson:"maxFrags"`
	EnemiesSpotted       statValue `json:"enemiesSpotted" bson:"enemiesSpotted"`
	CapturePoints        statValue `json:"capturePoints" bson:"capturePoints"`
	DroppedCapturePoints statValue `json:"droppedCapturePoints" bson:"droppedCapturePoints"`

	RawRating statValue `json:"mmRating" bson:"mmRating"`

	wn8             statValue `json:"-" bson:"-"`
	winRate         statValue `json:"-" bson:"-"`
	accuracy        statValue `json:"-" bson:"-"`
	avgDamage       statValue `json:"-" bson:"-"`
	survivalPercent statValue `json:"-" bson:"-"`
	damageRatio     statValue `json:"-" bson:"-"`
	survivalRatio   statValue `json:"-" bson:"-"`
}

/*
Calculate and return average damage per battle
*/
func (r *StatsFrame) AvgDamage() statValue {
	if r.Battles.IsZero() {
		return InvalidValue
	}
	if r.avgDamage.IsZero() {
		r.avgDamage = statValue{valueFormatInt, r.DamageDealt.Value() / r.Battles.Value()}
	}
	return r.avgDamage
}

/*
Calculate and return damage dealt vs damage received ration
*/
func (r *StatsFrame) DamageRatio() statValue {
	if r.Battles.IsZero() || r.DamageReceived.IsZero() {
		return InvalidValue
	}
	if r.damageRatio.IsZero() {
		r.damageRatio = statValue{valueFormatDecimal, r.DamageDealt.Value() / r.DamageReceived.Value()}
	}
	return r.damageRatio
}

/*
Calculate and return survival battles vs total battles ratio
*/
func (r *StatsFrame) SurvivalRatio() statValue {
	if r.Battles.IsZero() {
		return InvalidValue
	}
	if r.survivalRatio.IsZero() {
		r.survivalRatio = statValue{valueFormatDecimal, r.BattlesSurvived.Value() / r.Battles.Value()}
	}
	return r.survivalRatio
}

/*
Calculate and return survival percentage
*/
func (r *StatsFrame) Survival() statValue {
	if r.Battles.IsZero() {
		return InvalidValue
	}
	if r.survivalPercent.IsZero() {
		r.survivalPercent = statValue{valueFormatPercent, r.BattlesSurvived.Value() / r.Battles.Value() * 100}
	}
	return r.survivalPercent
}

/*
Calculate and return win percentage
*/
func (r *StatsFrame) WinRate() statValue {
	if r.Battles.IsZero() {
		return InvalidValue
	}
	if r.winRate.IsZero() {
		r.winRate = statValue{valueFormatPercent, r.BattlesWon.Value() / r.Battles.Value() * 100}
	}
	return r.winRate
}

/*
Calculate and return accuracy
*/
func (r *StatsFrame) Accuracy() statValue {
	if r.Battles.IsZero() || r.ShotsFired.IsZero() {
		return InvalidValue
	}
	if r.accuracy.IsZero() {
		r.accuracy = statValue{valueFormatPercent, r.ShotsHit.Value() / r.ShotsFired.Value() * 100}
	}
	return r.accuracy
}

/*
Calculate and return user-facing wargaming rating
*/
func (r *StatsFrame) Rating() statValue {
	if r.RawRating.IsZero() {
		return InvalidValue
	}
	return statValue{valueFormatInt, r.RawRating.Value()*10 + 3000}
}

/*
Set the WN8 value manually
*/
func (r *StatsFrame) SetWN8(new int) {
	r.wn8 = statValue{valueFormatInt, float32(new)}
}

/*
	 Calculate WN8 Rating for a tank using the following formula:
		(980*rDAMAGEc + 210*rDAMAGEc*rFRAGc + 155*rFRAGc*rSPOTc + 75*rDEFc*rFRAGc + 145*MIN(1.8,rWINc))/EXPc
*/
func (r *StatsFrame) WN8(vehicleAverages ...StatsFrame) statValue {
	if r.wn8.Value() > 0 {
		return r.wn8
	}

	if r.Battles.IsZero() || len(vehicleAverages) < 1 {
		return InvalidValue
	}

	average := vehicleAverages[0]
	if average.Battles.IsZero() {
		return InvalidValue
	}

	// Expected values for WN8
	expDef := average.DroppedCapturePoints.Value() / average.Battles.Value()
	expFrag := average.Frags.Value() / average.Battles.Value()
	expSpot := average.EnemiesSpotted.Value() / average.Battles.Value()
	expDmg := average.AvgDamage().Value()
	expWr := average.WinRate().Value() / 100

	// Actual performance
	pDef := r.DroppedCapturePoints.Value() / r.Battles.Value()
	pFrag := r.Frags.Value() / r.Battles.Value()
	pSpot := r.EnemiesSpotted.Value() / r.Battles.Value()
	pDmg := r.AvgDamage().Value()
	pWr := r.WinRate().Value() / 100

	// Calculate WN8 metrics
	rDef := pDef / expDef
	rFrag := pFrag / expFrag
	rSpot := pSpot / expSpot
	rDmg := pDmg / expDmg
	rWr := pWr / expWr

	adjustedWr := math.Max(0, float64((rWr-0.71)/(1-0.71)))
	adjustedDmg := math.Max(0, float64((rDmg-0.22)/(1-0.22)))
	adjustedDef := math.Max(0, (math.Min(adjustedDmg+0.1, float64((rDef-0.10)/(1-0.10)))))
	adjustedSpot := math.Max(0, (math.Min(adjustedDmg+0.1, float64((rSpot-0.38)/(1-0.38)))))
	adjustedFrag := math.Max(0, (math.Min(adjustedDmg+0.2, float64((rFrag-0.12)/(1-0.12)))))

	r.wn8 = statValue{valueFormatInt, float32(math.Round(((980 * adjustedDmg) + (210 * adjustedDmg * adjustedFrag) + (155 * adjustedFrag * adjustedSpot) + (75 * adjustedDef * adjustedFrag) + (145 * math.Min(1.8, adjustedWr)))))}
	return r.wn8
}

/*
Add all values from a passed in frame to the current frame
*/
func (r *StatsFrame) Add(other StatsFrame) {
	r.Battles.data += other.Battles.data
	r.BattlesWon.data += other.BattlesWon.data
	r.BattlesSurvived.data += other.BattlesSurvived.data
	r.DamageDealt.data += other.DamageDealt.data
	r.DamageReceived.data += other.DamageReceived.data

	r.ShotsHit.data += other.ShotsHit.data
	r.ShotsFired.data += other.ShotsFired.data

	r.Frags.data += other.Frags.data
	if r.MaxFrags.data < other.MaxFrags.data {
		r.MaxFrags.data = other.MaxFrags.data
	}

	r.EnemiesSpotted.data += other.EnemiesSpotted.data
	r.CapturePoints.data += other.CapturePoints.data
	r.DroppedCapturePoints.data += other.DroppedCapturePoints.data
}

/*
Subtract all values of the passed in frame from the current frame
*/
func (r *StatsFrame) Subtract(other StatsFrame) {
	r.Battles.data -= other.Battles.data
	r.BattlesWon.data -= other.BattlesWon.data
	r.BattlesSurvived.data -= other.BattlesSurvived.data
	r.DamageDealt.data -= other.DamageDealt.data
	r.DamageReceived.data -= other.DamageReceived.data

	r.ShotsHit.data -= other.ShotsHit.data
	r.ShotsFired.data -= other.ShotsFired.data

	r.Frags.data -= other.Frags.data
	if r.MaxFrags.data > other.MaxFrags.data {
		r.MaxFrags.data = other.MaxFrags.data
	}

	r.EnemiesSpotted.data -= other.EnemiesSpotted.data
	r.CapturePoints.data -= other.CapturePoints.data
	r.DroppedCapturePoints.data -= other.DroppedCapturePoints.data
}

/*
VehicleStatsFrame is a frame to structure vehicle statistics overview
*/
type VehicleStatsFrame struct {
	VehicleID   int `json:"vehicleId" bson:"vehicleId"`
	*StatsFrame `bson:",inline"`

	MarkOfMastery  int `json:"markOfMastery" bson:"markOfMastery"`
	LastBattleTime int `json:"lastBattleTime" bson:"lastBattleTime"`
}

/*
Add all values from a passed in frame to the current frame
*/
func (r *VehicleStatsFrame) Add(other VehicleStatsFrame) {
	r.StatsFrame.Add(*other.StatsFrame)
	if other.MarkOfMastery > r.MarkOfMastery {
		r.MarkOfMastery = other.MarkOfMastery
	}
	if other.LastBattleTime > r.LastBattleTime {
		r.LastBattleTime = other.LastBattleTime
	}
}

/*
Subtract all values of the passed in frame from the current frame
*/
func (r *VehicleStatsFrame) Subtract(other VehicleStatsFrame) {
	r.StatsFrame.Subtract(*other.StatsFrame)
	if other.MarkOfMastery > r.MarkOfMastery {
		r.MarkOfMastery = other.MarkOfMastery
	}
	if other.LastBattleTime > r.LastBattleTime {
		r.LastBattleTime = other.LastBattleTime
	}
}
