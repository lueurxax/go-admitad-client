package client

import (
	"github.com/lueurxax/go-admitad-client/defaults"
<<<<<<< HEAD
	"github.com/lueurxax/go-admitad-client/news"
	"github.com/lueurxax/go-admitad-client/referrals"
	"github.com/lueurxax/go-admitad-client/responses"
	"github.com/lueurxax/go-admitad-client/tickets"
	"time"
=======
	"github.com/lueurxax/go-admitad-client/internal"
>>>>>>> master
)

type AClient struct {
	Statistics
<<<<<<< HEAD
	news.News
	referrals.Referrals
	tickets.Ticket

	bc *baseClient
=======
	bc *internal.BaseClient
>>>>>>> master
}

func New(url string, logger internal.Logger, mc internal.MetricsCollector) *AClient {
	if logger == nil {
		logger = defaults.NewLogrus()
	}
	if mc == nil {
		mc = &defaults.EmptyMetrics{}
	}
	bc := internal.NewBaseClient(url, logger, mc)
	return &AClient{Statistics: Statistics{bc}, bc: bc}
}

func (ac *AClient) SetAuth(clientID, clientSecret, scope string) error {
	return ac.bc.SetAuth(clientID, clientSecret, scope)
}
