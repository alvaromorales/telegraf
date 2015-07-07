package elasticsearch

import (
	"fmt"
	"github.com/influxdb/telegraf/plugins"
	"github.com/olivere/elastic"
)

type Elasticsearch struct {
	Client elastic.Client
}

var sampleConfig = ""

func (e *Elasticsearch) SampleConfig() string {
	return sampleConfig
}

func (e *Elasticsearch) Description() string {
	return "Read metrics from one or many Elasticsearch nodes"
}

// Reads stats from all configured servers accumulates stats.
// Returns one of the errors encountered while gather stats (if any).
func (e *Elasticsearch) Gather(acc plugins.Accumulator) error {
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
	}

	stats, err := client.ClusterStats().Do()
	if err != nil {
		// Handle error
	}

	fmt.Printf("%s", stats.Indices)
	return nil

}

func init() {
	plugins.Add("elasticsearch", func() plugins.Plugin {
		return &Elasticsearch{}
	})
}
