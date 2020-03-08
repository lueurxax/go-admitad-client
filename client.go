package client

import (
	"github.com/lueurxax/go-admitad-client/currencies"
	"github.com/lueurxax/go-admitad-client/defaults"
	"github.com/lueurxax/go-admitad-client/internal"
	"github.com/lueurxax/go-admitad-client/me"
	"github.com/lueurxax/go-admitad-client/news"
	offlineReceipts "github.com/lueurxax/go-admitad-client/offline_receipts"
	"github.com/lueurxax/go-admitad-client/tickets"
)

type AClient struct {
	Statistics
	news.News
	currencies.Currencies
	me.Me
	offlineReceipts.OfflineReceipts
	tickets.Ticket

	bc *internal.BaseClient
}

func New(url string, logger internal.Logger, mc internal.MetricsCollector) *AClient {
	if logger == nil {
		logger = defaults.NewLogrus()
	}
	if mc == nil {
		mc = &defaults.EmptyMetrics{}
	}
	bc := internal.NewBaseClient(url, logger, mc)
	return &AClient{
		Statistics:      Statistics{bc},
		News:            news.News{BaseClient: bc},
		Currencies:      currencies.Currencies{BaseClient: bc},
		Me:              me.Me{BaseClient: bc},
		OfflineReceipts: offlineReceipts.OfflineReceipts{BaseClient: bc},
		Ticket:          tickets.NewTicket(bc),
		bc:              bc}
}

func (ac *AClient) SetAuth(clientID, clientSecret, scope string) error {
	return ac.bc.SetAuth(clientID, clientSecret, scope)
}
