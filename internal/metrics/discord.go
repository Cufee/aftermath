package metrics

import (
	"fmt"
	"strings"

	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/prometheus/client_golang/prometheus"
)

type discordCollector struct {
	totalInteractions        *prometheus.CounterVec
	totalInteractionsErrored *prometheus.CounterVec
}

func (c discordCollector) Describe(ch chan<- *prometheus.Desc) {
	c.totalInteractions.Describe(ch)
	c.totalInteractionsErrored.Describe(ch)
}
func (c discordCollector) Collect(ch chan<- prometheus.Metric) {
	c.totalInteractions.Collect(ch)
	c.totalInteractionsErrored.Collect(ch)
}

func (c discordCollector) Middleware() middleware.MiddlewareFunc {
	return func(ctx common.Context, next func(common.Context) error) func(common.Context) error {
		return func(ctx common.Context) error {
			name := "unknown_interaction"
			if data, ok := ctx.CommandData(); ok {
				name = "command_" + data.Name
			}
			if data, ok := ctx.ComponentData(); ok {
				name = "component_" + strings.SplitN(data.CustomID, "#", 2)[0]
			}
			if data, ok := ctx.AutocompleteData(); ok {
				name = "autocomplete_" + data.Name
			}
			c.totalInteractions.WithLabelValues(name).Inc()
			err := next(ctx)
			if err != nil {
				c.totalInteractionsErrored.WithLabelValues(name).Inc()
			}
			return err
		}
	}
}

func NewDiscordCollector(prefix string) discordCollector {
	return discordCollector{
		totalInteractionsErrored: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: fmt.Sprintf("discord_%s_interactions_errored_total", prefix),
				Help: "Number of errored Discord interactions.",
			},
			[]string{"name"},
		),
		totalInteractions: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: fmt.Sprintf("discord_%s_interactions_handled_total", prefix),
				Help: "Number of handled Discord interactions.",
			},
			[]string{"name"},
		),
	}
}
