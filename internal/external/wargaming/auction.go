package wargaming

import (
	"fmt"
	"net/http"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

var auctionHttpClient = http.DefaultClient

// https://asia.wotblitz.com/en/api/events/items/auction/?page_size=10&type[]=vehicle&saleable=true

func CurrentAuction(realm types.Realm) (any, error) {
	domain, ok := realm.DomainBlitz()
	if !ok {
		return nil, ErrRealmNotSupported
	}

	url := fmt.Sprintf("https://%s/en/api/events/items/auction/?page_size=100&type[]=vehicle&saleable=true", domain)
	res, err := auctionHttpClient.Get(url)
	if err != nil {
		return nil, err
	}

	_ = res
	return nil, nil
}
