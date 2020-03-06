package responses

type OfflineReceipts struct {
    Results []ResultsReceipts `json:"results"`
    Meta    Meta              `json:"_meta"`
}
type Items struct {
    Sum             string `json:"sum"`
    Quantity        string `json:"quantity"`
    OfflineSalesIds []int  `json:"offline_sales_ids"`
    Price           string `json:"price"`
    Name            string `json:"name"`
}

type ResultsReceipts struct {
    ID                          string  `json:"id"`
    WebsiteID                   int     `json:"website_id"`
    PurchaserID                 int     `json:"purchaser_id"`
    Status                      string  `json:"status"`
    PaymentIds                  []int   `json:"payment_ids"`
    RetailPlaceAddress          string  `json:"retail_place_address"`
    User                        string  `json:"user"`
    TotalSum                    string  `json:"total_sum"`
    DatetimeCreated             string  `json:"datetime_created"`
    DatetimeUpdated             string  `json:"datetime_updated"`
    OfflineSalesIds             []int   `json:"offline_sales_ids"`
    NotActivatedOfflineSalesIds []int   `json:"not_activated_offline_sales_ids"`
    Items                       []Items `json:"items"`
    FNSData                     string  `json:"FNS_data"`
}

