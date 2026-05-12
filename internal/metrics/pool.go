package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
    PoolActiveConns = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "core_api_db_pool_active",
        Help: "Active connections in the DB pool",
    })
    PoolIdleConns = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "core_api_db_pool_idle",
        Help: "Idle connections in the DB pool",
    })
)

func init() {
    prometheus.MustRegister(PoolActiveConns, PoolIdleConns)
}
