package model

// AccountTypes type of accounts eg. "twitter"
var AccountTypes map[string]AccountType

// Account holds a social media account
type Account struct {
	ID         string
	Type       AccountType
	Title      string
	OauthToken string
}

// AccountType holds the type of accounts, eg. twitter, facebook etc. Don't save to db, it's just a constant like variable
type AccountType struct {
	Title string
	ID    string
	URL   string
}

func init() {
	defineTypes()
}

func defineTypes() {
	// Define new types here
	AccountTypes = make(map[string]AccountType)
	AccountTypes["twitter"] = AccountType{ID: "01", Title: "twitter"}
	AccountTypes["facebook"] = AccountType{ID: "02", Title: "facebook"}
}
