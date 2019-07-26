package billing

import (
	"fmt"
	"time"

	"github.com/orangesys/janus/pkg/promql"
	"github.com/orangesys/janus/pkg/util"
)

func CountNodesFromQuerier(addr string) float64 {
	currentTime := time.Now()
	start, end := util.OneDaysAgoTimestamp(currentTime)

	// count prometheus server with usage record
	// count(kube_pod_start_time{pod=~"prometheus-k8s.*"})
	query := "count(node_boot_time_seconds)"

	prom, err := promql.NewClient(addr)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := prom.QueryRange(query, start, end)
	if err != nil {
		fmt.Println(err)
	}
	type valueEntry struct {
		Metric map[string]string `json:"metric"`
		Value  float64           `json:"value`
	}

	type timeEntry struct {
		Time   int64         `jsno:"time"`
		Values []*valueEntry `json:"values"`
	}

	entryByTime := map[int64]*timeEntry{}
	var nodes float64
	// Save count node number to firestore
	// Show billing with app.orangesys.io dashboard
	for _, r := range resp.Data.Result {
		for _, v := range r.Values {
			t := v.Time()
			u := t.Unix()
			e, ok := entryByTime[u]

			if !ok {
				e = &timeEntry{
					Time:   u,
					Values: []*valueEntry{},
				}
				entryByTime[u] = e
			}
			val, err := v.Value()
			nodes = nodes + val
			if err != nil {
				fmt.Println(err)
			}
			e.Values = append(e.Values, &valueEntry{
				Metric: r.Metric,
				Value:  val,
			})
		}
	}
	return nodes
}
