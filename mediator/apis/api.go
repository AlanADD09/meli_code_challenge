package api

type API interface {
	FetchMultiData(input []SiteID) ([]ResponseData, error)
	FetchData(input string) (string, error)
	FetchNumericData(input int64) (string, error)
}

type SiteID struct {
	Site string `json:"site"`
	ID   string `json:"id"`
}

type ResponseData struct {
	Site        string  `json:"site_id"`
	ID          string  `json:"id"`
	Price       float64 `json:"price"`
	StartTime   string  `json:"date_created"`
	CategoryID  string  `json:"category_id"`
	CurrencyID  string  `json:"currency_id"`
	SellerID    int64   `json:"seller_id"`
	Name        string
	Description string
	Nickname    string
	Error       string `json:"error"`
}
