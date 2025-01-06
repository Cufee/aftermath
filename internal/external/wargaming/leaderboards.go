package wargaming

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type RatingLeaderboardClient struct {
	http *http.Client
}

func NewRatingLeaderboardClient() (*RatingLeaderboardClient, error) {
	return &RatingLeaderboardClient{http: http.DefaultClient}, nil
}

func (c *RatingLeaderboardClient) CurrentSeason(ctx context.Context, realm types.Realm) (RatingSeason, error) {
	req, err := http.NewRequest("GET", seasonURL(realm), nil)
	if err != nil {
		return RatingSeason{}, err
	}
	req = req.WithContext(ctx)

	res, err := c.http.Do(req)
	if err != nil {
		return RatingSeason{}, err
	}
	defer res.Body.Close()

	var data RatingSeason
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return RatingSeason{}, err
	}

	return data, nil
}

func (c *RatingLeaderboardClient) LeagueTop(ctx context.Context, realm types.Realm, leagueID int) ([]LeaderboardPosition, error) {
	req, err := http.NewRequest("GET", leagueTopURL(realm, leagueID), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data struct {
		Result []LeaderboardPosition `json:"result"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Result, nil
}

func (c *RatingLeaderboardClient) PlayerPosition(ctx context.Context, realm types.Realm, accountID int, neighbors int) (LeaderboardPosition, error) {
	req, err := http.NewRequest("GET", neighborsURL(realm, accountID, neighbors), nil)
	if err != nil {
		return LeaderboardPosition{}, err
	}
	req = req.WithContext(ctx)

	res, err := c.http.Do(req)
	if err != nil {
		return LeaderboardPosition{}, err
	}
	defer res.Body.Close()

	var data LeaderboardPosition
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return LeaderboardPosition{}, err
	}

	return data, nil
}

// Season
type RatingSeason struct {
	Title string `json:"title"`
	// Icon          any       `json:"icon"`
	StartAt       string    `json:"start_at"`
	FinishAt      string    `json:"finish_at"`
	CurrentSeason int       `json:"current_season"`
	Rewards       []Rewards `json:"rewards"`
	UpdatedAt     string    `json:"updated_at"`
	Leagues       []Leagues `json:"leagues"`
	Count         int       `json:"count"`
}

type RewardVehicle struct {
	ID           int    `json:"id"`
	NameKey      string `json:"name"`
	NameString   string `json:"user_string"`
	Nation       string `json:"nation"`
	Class        string `json:"type_slug"`
	Tier         int    `json:"level"`
	TierRoman    string `json:"roman_level"`
	Image        string `json:"image_url"`
	PreviewImage string `json:"preview_image_url"`
	Premium      bool   `json:"is_premium"`
	Collectible  bool   `json:"is_collectible"`
}

type RewardStuff struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Title string `json:"title"`
	Image string `json:"image_url"`
	Type  string `json:"type"`
}

type Rewards struct {
	Type         string        `json:"type"`
	Vehicle      RewardVehicle `json:"vehicle"`
	Stuff        RewardStuff   `json:"stuff"`
	FromPosition int           `json:"from_position"`
	ToPosition   int           `json:"to_position"`
	Count        int           `json:"count"`
}

type Leagues struct {
	ID         int     `json:"index"`
	Title      string  `json:"title"`
	SmallIcon  string  `json:"small_icon"`
	BigIcon    string  `json:"big_icon"`
	Background string  `json:"background"`
	Percentile float64 `json:"percentile"`
}

type LeaderboardPosition struct {
	SeasonID  int                   `json:"season_number"`
	Neighbors []LeaderboardPosition `json:"neighbors"`

	AccountID int    `json:"spa_id"`
	Nickname  string `json:"nickname"`
	ClanTag   string `json:"clan_tag"`

	Rating                 int     `json:"score"`
	RawRating              float64 `json:"mmr"`
	Position               int     `json:"number"`
	Percentile             float64 `json:"percentile"`
	CalibrationBattlesLeft int     `json:"calibrationBattlesLeft"`

	Skip      bool   `json:"skip"`
	UpdatedAt string `json:"updated_at"`
}

func seasonURL(realm types.Realm) string {
	domain, _ := realm.DomainBlitz()
	return fmt.Sprintf("https://%s/en/api/rating-leaderboards/season/", domain)
}

func neighborsURL(realm types.Realm, accountID int, neighbors int) string {
	domain, _ := realm.DomainBlitz()
	return fmt.Sprintf("https://%s/en/api/rating-leaderboards/user/%d/?neighbors=%d", domain, accountID, neighbors)
}

func leagueTopURL(realm types.Realm, leagueID int) string {
	domain, _ := realm.DomainBlitz()
	return fmt.Sprintf("https://%s/en/api/rating-leaderboards/league/%d/top/", domain, leagueID)
}
