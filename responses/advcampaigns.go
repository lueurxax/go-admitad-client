package responses

type AdvCampaigns struct {
    Results []Results `json:"results"`
    Meta    Meta      `json:"_meta"`
}
type Traffics struct {
    Enabled bool   `json:"enabled"`
    Name    string `json:"name"`
    Type    string `json:"type"`
    ID      int    `json:"id"`
}

type ActionsAdv struct {
    PaymentSize string `json:"payment_size"`
    Type        string `json:"type"`
    Name        string `json:"name"`
    ID          int    `json:"id"`
}

type Regions struct {
    Region string `json:"region"`
}
type Categories struct {
    Language string      `json:"language"`
    Name     string      `json:"name"`
    Parent   interface{} `json:"parent"`
    ID       int         `json:"id"`
}
type Results struct {
    Status                   string       `json:"status"`
    Rating                   string       `json:"rating"`
    Image                    string       `json:"image"`
    Description              string       `json:"description"`
    Traffics                 []Traffics   `json:"traffics"`
    Actions                  []ActionsAdv `json:"actions"`
    SiteURL                  string       `json:"site_url"`
    Regions                  []Regions    `json:"regions"`
    Currency                 string       `json:"currency"`
    Geotargeting             bool         `json:"geotargeting"`
    CouponIframeDenied       bool         `json:"coupon_iframe_denied"`
    Connected                bool         `json:"connected"`
    ID                       int          `json:"id"`
    CR                       float64      `json:"cr"`
    ECPC                     float64      `json:"ecpc"`
    EPC                      float64      `json:"epc"`
    CRTrend                  float64      `json:"cr_trend"`
    ECPCTrend                float64      `json:"ecpc_trend"`
    EPCTrend                 float64      `json:"epc_trend"`
    Categories               []Categories `json:"categories"`
    Name                     string       `json:"name"`
    ActionType               string       `json:"action_type"`
    IndividualTerms    bool `json:"individual_terms"`
    AllowDeepLink      bool `json:"allow_deeplink"`
    ActionTestingLimit int  `json:"action_testing_limit"`
    MobileDeviceType         string       `json:"mobile_device_type"`
    MobileOsType             string       `json:"mobile_os_type"`
    MobileOs                 string       `json:"mobile_os"`
    ActionCountries          []string     `json:"action_countries"`
    AllowActionsAllCountries bool         `json:"allow_actions_all_countries"`
}
