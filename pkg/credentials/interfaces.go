package credentials

type Credentials struct {
	X_API_TOKEN string `json:"X-API-TOKEN"`
	From        string `json:"from"`
	DB          string
	DB_USER     string
	DB_PASS     string
	DB_PORT     string
	DB_HOST     string
	DB_DIALECT  string
}
