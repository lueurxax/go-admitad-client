package responses

type Statistics struct {
    Results []ResultsStats `json:"results"`
    Meta    Meta           `json:"_meta"`
}

type ResultsStats struct {
    LeadsSum           int     `json:"leads_sum"`
    PaymentSumApproved float64 `json:"payment_sum_approved"`
    Currency           string  `json:"currency"`
    PaymentSumDeclined float64 `json:"payment_sum_declined"`
    PaymentSumOpen     float64 `json:"payment_sum_open"`
    Subid              string  `json:"subid"`
    SalesSum           int     `json:"sales_sum"`
    Cr                 float64 `json:"cr"`
    Ecpc               float64 `json:"ecpc"`
    Clicks             int     `json:"clicks"`
}

