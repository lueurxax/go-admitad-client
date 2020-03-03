package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

type Prom struct {
	*prometheus.HistogramVec
}

func (p Prom) Collect(url string, httpCode int, appCode int, duration time.Duration) {
	labels := map[string]string{}
	if url != "" {
		labels["url"] = url
	}
	if httpCode != 0 {
		labels["http_code"] = strconv.Itoa(httpCode)
	}
	labels["app_code"] = strconv.Itoa(appCode)
	p.HistogramVec.With(labels).Observe(duration.Seconds())
}
