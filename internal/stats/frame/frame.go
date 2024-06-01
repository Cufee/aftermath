package frame

import (
	"math"
	"time"

	"go.dedis.ch/protobuf"
)

/*
StatsFrame is a frame to structure statistics overview
*/
type StatsFrame struct {
	Battles         ValueInt `json:"battles" protobuf:"10"`
	BattlesWon      ValueInt `json:"battlesWon" protobuf:"11"`
	BattlesSurvived ValueInt `json:"battlesSurvived" protobuf:"12"`

	DamageDealt    ValueInt `json:"damageDealt" protobuf:"15"`
	DamageReceived ValueInt `json:"damageReceived" protobuf:"16"`

	ShotsHit   ValueInt `json:"shotsHit" protobuf:"20"`
	ShotsFired ValueInt `json:"shotsFired" protobuf:"21"`

	Frags    ValueInt `json:"frags" protobuf:"25"`
	MaxFrags ValueInt `json:"maxFrags" protobuf:"26"`

	EnemiesSpotted ValueInt `json:"enemiesSpotted" protobuf:"30"`

	CapturePoints        ValueInt `json:"capturePoints" protobuf:"35"`
	DroppedCapturePoints ValueInt `json:"droppedCapturePoints" protobuf:"36"`

	Rating ValueSpecialRating `json:"mmRating" protobuf:"40"`

	wn8             ValueInt          `json:"-" bson:"-"`
	winRate         ValueFloatPercent `json:"-" bson:"-"`
	accuracy        ValueFloatPercent `json:"-" bson:"-"`
	avgDamage       ValueFloatDecimal `json:"-" bson:"-"`
	survivalPercent ValueFloatPercent `json:"-" bson:"-"`
	damageRatio     ValueFloatDecimal `json:"-" bson:"-"`
	survivalRatio   ValueFloatDecimal `json:"-" bson:"-"`
}

func DecodeStatsFrame(encoded string) (StatsFrame, error) {
	var data StatsFrame
	err := protobuf.Decode([]byte(encoded), &data)
	if err != nil {
		return StatsFrame{}, err
	}
	return data, nil
}

/*
Encode StatsFrame to string
*/
func (r *StatsFrame) Encode() (string, error) {
	data, err := protobuf.Encode(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Calculate and return average damage per battle
*/
func (r *StatsFrame) AvgDamage() Value {
	if r.Battles == 0 {
		return InvalidValue
	}
	if r.avgDamage == 0 {
		r.avgDamage = ValueFloatDecimal(r.DamageDealt.Float() / r.Battles.Float())
	}
	return r.avgDamage
}

/*
Calculate and return damage dealt vs damage received ration
*/
func (r *StatsFrame) DamageRatio() Value {
	if r.Battles == 0 || r.DamageReceived == 0 {
		return InvalidValue
	}
	if r.damageRatio == 0 {
		r.damageRatio = ValueFloatDecimal(r.DamageDealt.Float() / r.DamageReceived.Float())
	}
	return r.damageRatio
}

/*
Calculate and return survival battles vs total battles ratio
*/
func (r *StatsFrame) SurvivalRatio() Value {
	if r.Battles == 0 {
		return InvalidValue
	}
	if r.survivalRatio == 0 {
		r.survivalRatio = ValueFloatDecimal(r.DamageDealt.Float() / r.Battles.Float())
	}
	return r.survivalRatio
}

/*
Calculate and return survival percentage
*/
func (r *StatsFrame) Survival() Value {
	if r.Battles == 0 {
		return InvalidValue
	}
	if r.survivalPercent == 0 {
		r.survivalPercent = ValueFloatPercent(r.BattlesSurvived.Float() / r.Battles.Float() * 100)
	}
	return r.survivalPercent
}

/*
Calculate and return win percentage
*/
func (r *StatsFrame) WinRate() Value {
	if r.Battles == 0 {
		return InvalidValue
	}
	if r.winRate == 0 {
		r.winRate = ValueFloatPercent(r.BattlesWon.Float() / r.Battles.Float() * 100)
	}
	return r.winRate
}

/*
Calculate and return accuracy
*/
func (r *StatsFrame) Accuracy() Value {
	if r.Battles == 0 || r.ShotsFired == 0 {
		return InvalidValue
	}
	if r.accuracy == 0 {
		r.accuracy = ValueFloatPercent(r.ShotsHit.Float() / r.ShotsFired.Float() * 100)
	}
	return r.accuracy
}

/*
Set the WN8 value manually
*/
func (r *StatsFrame) SetWN8(wn8 int) {
	r.wn8 = ValueInt(wn8)
}

/*
	 Calculate WN8 Rating for a tank using the following formula:
		(980*rDAMAGEc + 210*rDAMAGEc*rFRAGc + 155*rFRAGc*rSPOTc + 75*rDEFc*rFRAGc + 145*MIN(1.8,rWINc))/EXPc
*/
func (r *StatsFrame) WN8(vehicleAverages ...StatsFrame) Value {
	if r.wn8 > 0 {
		return r.wn8
	}

	if r.Battles == 0 || len(vehicleAverages) < 1 {
		return InvalidValue
	}

	average := vehicleAverages[0]
	if average.Battles == 0 {
		return InvalidValue
	}

	// Expected values for WN8
	expDef := average.DroppedCapturePoints.Float() / average.Battles.Float()
	expFrag := average.Frags.Float() / average.Battles.Float()
	expSpot := average.EnemiesSpotted.Float() / average.Battles.Float()
	expDmg := average.AvgDamage().Float()
	expWr := average.WinRate().Float() / 100

	// Actual performance
	pDef := r.DroppedCapturePoints.Float() / r.Battles.Float()
	pFrag := r.Frags.Float() / r.Battles.Float()
	pSpot := r.EnemiesSpotted.Float() / r.Battles.Float()
	pDmg := r.AvgDamage().Float()
	pWr := r.WinRate().Float() / 100

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

	r.wn8 = ValueInt(int(math.Round(((980 * adjustedDmg) + (210 * adjustedDmg * adjustedFrag) + (155 * adjustedFrag * adjustedSpot) + (75 * adjustedDef * adjustedFrag) + (145 * math.Min(1.8, adjustedWr))))))
	return r.wn8
}

/*
Add all values from a passed in frame to the current frame
*/
func (r *StatsFrame) Add(other StatsFrame) {
	r.Battles += other.Battles
	r.BattlesWon += other.BattlesWon
	r.BattlesSurvived += other.BattlesSurvived
	r.DamageDealt += other.DamageDealt
	r.DamageReceived += other.DamageReceived

	r.ShotsHit += other.ShotsHit
	r.ShotsFired += other.ShotsFired

	r.Frags += other.Frags
	if r.MaxFrags < other.MaxFrags {
		r.MaxFrags = other.MaxFrags
	}

	r.EnemiesSpotted += other.EnemiesSpotted
	r.CapturePoints += other.CapturePoints
	r.DroppedCapturePoints += other.DroppedCapturePoints
}

/*
Subtract all values of the passed in frame from the current frame
*/
func (r *StatsFrame) Subtract(other StatsFrame) {
	r.Battles -= other.Battles
	r.BattlesWon -= other.BattlesWon
	r.BattlesSurvived -= other.BattlesSurvived
	r.DamageDealt -= other.DamageDealt
	r.DamageReceived -= other.DamageReceived

	r.ShotsHit -= other.ShotsHit
	r.ShotsFired -= other.ShotsFired

	r.Frags -= other.Frags
	if r.MaxFrags > other.MaxFrags {
		r.MaxFrags = other.MaxFrags
	}

	r.EnemiesSpotted -= other.EnemiesSpotted
	r.CapturePoints -= other.CapturePoints
	r.DroppedCapturePoints -= other.DroppedCapturePoints
}

/*
VehicleStatsFrame is a frame to structure vehicle statistics overview
*/
type VehicleStatsFrame struct {
	VehicleID   string `json:"vehicleId" bson:"vehicleId"`
	*StatsFrame `bson:",inline"`

	MarkOfMastery  int       `json:"markOfMastery" bson:"markOfMastery"`
	LastBattleTime time.Time `json:"lastBattleTime" bson:"lastBattleTime"`
}

/*
Add all values from a passed in frame to the current frame
*/
func (r *VehicleStatsFrame) Add(other VehicleStatsFrame) {
	r.StatsFrame.Add(*other.StatsFrame)
	if other.MarkOfMastery > r.MarkOfMastery {
		r.MarkOfMastery = other.MarkOfMastery
	}
	if other.LastBattleTime.After(r.LastBattleTime) {
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
	if other.LastBattleTime.Before(r.LastBattleTime) {
		r.LastBattleTime = other.LastBattleTime
	}
}
