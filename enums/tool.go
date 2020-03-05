package enums

type Tool string

// https://www.admitad.com/ru/developers/doc/api_ru/methods/advcampaigns/advcampaigns-list/
const (
	DeepLink             Tool = "deeplink"
	Products             Tool = "products"
	ReTag                Tool = "retag"
	PostView             Tool = "postview"
	LostOrdersTool           Tool = "lost_orders"
	BrokerTraffic        Tool = "broker_traffic"
	CouponsTool          Tool = "coupons"
	BasketTracking       Tool = "basket_tracking"
	TrackingInMobileSite Tool = "tracking_in_mobile_site"
	TrackingInMobileApp  Tool = "tracking_in_mobile_app"
)
