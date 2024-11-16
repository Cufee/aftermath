package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Instrument struct {
	reg *prometheus.Registry
}

func (i Instrument) Handler() http.HandlerFunc {
	return promhttp.Handler().ServeHTTP
}

func (i Instrument) Register(c prometheus.Collector) error {
	return prometheus.DefaultRegisterer.Register(c)
}

func NewInstrument(collectors ...prometheus.Collector) (Instrument, error) {
	instrument := Instrument{reg: prometheus.NewRegistry()}
	for _, c := range collectors {
		err := instrument.Register(c)
		if err != nil {
			return Instrument{}, err
		}
	}
	return instrument, nil
}
