package ingester

import (
	"time"

	"github.com/grafana/tempo/modules/generator/registry"
	"github.com/grafana/tempo/modules/overrides"
	"github.com/grafana/tempo/tempodb/backend"
)

type ingesterOverrides interface {
	registry.Overrides

	DedicatedColumns(userID string) backend.DedicatedColumns
	IngesterMaxBlockDuration(userID string) time.Duration
}

var _ ingesterOverrides = (overrides.Interface)(nil)
