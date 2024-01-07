package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "abc123",
		Username:  "alex",
	},
	"geddy": {
		AuthToken: "qwe456",
		Username:  "geddy",
	},
	"neil": {
		AuthToken: "asd789",
		Username:  "neil",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    250,
		Username: "alex",
	},
	"geddy": {
		Coins:    320,
		Username: "geddy",
	},
	"neil": {
		Coins:    280,
		Username: "neil",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Mock DB call
	time.Sleep(time.Second * 2)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Mock DB call
	time.Sleep(time.Second * 2)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
