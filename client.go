package client

import (
	"github.com/lueurxax/go-admitad-client/defaults"
	"github.com/lueurxax/go-admitad-client/internal"
	"github.com/lueurxax/go-admitad-client/news"
)

type AClient struct {
	Statistics
	news.News

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
	return &AClient{Statistics: Statistics{bc}, bc: bc}
}

func (ac *AClient) SetAuth(clientID, clientSecret, scope string) error {
	return ac.bc.SetAuth(clientID, clientSecret, scope)
}
