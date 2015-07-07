package elasticsearch

import (
	"github.com/influxdb/telegraf/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestElasticsearchGenerateMetrics(t *testing.T) {
	e := &Elasticsearch{}
	var acc testutil.Accumulator

	err := e.Gather(&acc)
	require.NoError(t, err)
}
