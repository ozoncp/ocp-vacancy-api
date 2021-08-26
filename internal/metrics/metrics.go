package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createCounter prometheus.Counter
	updateCounter prometheus.Counter
	removeCounter prometheus.Counter
)

func RegisterMetrics() {
	createCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_vacancy_api_create_total",
		Help: "Total number of created vacancies",
	})

	updateCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_vacancy_api_update_total",
		Help: "Total number of updated vacancies",
	})

	removeCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_vacancy_api_remove_total",
		Help: "Total number of removed vacancies",
	})
}

func IncCreateCounter() {
	createCounter.Inc()
}

func IncUpdateCounter() {
	updateCounter.Inc()
}

func IncRemoveCounter() {
	removeCounter.Inc()
}
