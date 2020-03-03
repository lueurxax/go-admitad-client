package defaults

import (
	"time"
)

type EmptyMetrics struct {
}

func (p *EmptyMetrics) Collect(_ string, _ int, _ int, _ time.Duration) {}
