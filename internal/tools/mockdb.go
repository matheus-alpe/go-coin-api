package tools

import "time"

var mockLoginDetails = map[string]LoginDetails{
	"matt": {
		AuthToken: "123ZXC",
		Username:  "Matheus",
	},
	"nesouro": {
		AuthToken: "456DFG",
		Username:  "Aline",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"matt": {
		Coins:    69,
		Username: "Matheus",
	},
	"nesouro": {
		Coins:    100,
		Username: "Aline",
	},
}

type mockDB struct{}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second)

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second)

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
