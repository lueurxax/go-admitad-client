package tests

import (
	admitad "github.com/lueurxax/go-admitad-client"
	"github.com/lueurxax/go-admitad-client/defaults"
	"github.com/lueurxax/go-admitad-client/requests"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestStatistics(t *testing.T) {
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	logger := logrus.NewEntry(l)

	url := "https://api.admitad.com"
	clientID := os.Getenv("clientID")
	clientSecret := os.Getenv("clientSecret")
	client := admitad.New(url, &defaults.Logrus{Entry: logger}, nil)
	require.NoError(t, client.SetAuth(clientID, clientSecret, "statistics public_data"))
	dateStart := time.Now().Add(-time.Hour * time.Duration(48))
	t.Log(dateStart)
	n := 0
	for {
		offset := n * 500
		id := 492430099
		data, err := client.Statistics.Actions(&requests.Actions{ActionID: &id, Limit: 500, Offset: &offset, DateStart: &dateStart})
		require.NoError(t, err)
		t.Log(len(data.Results))
		for _, el := range data.Results {
			t.Logf("ActionDate: %+v", el.ActionDate)
			t.Logf("ClickDate: %+v", el.ClickDate)
			t.Logf("ClosingDate: %+v", el.ClosingDate)
			t.Logf("StatusUpdated: %+v", el.StatusUpdated)
		}
		if offset+500 > data.Meta.Count {
			break
		}
		n++
	}
}
