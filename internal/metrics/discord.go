package metrics

import (
	"fmt"
	"strings"

	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/prometheus/client_golang/prometheus"
)

type discordInteractionCollector struct {
	totalInteractions        *prometheus.CounterVec
	totalInteractionsErrored *prometheus.CounterVec
}

func (c discordInteractionCollector) Describe(ch chan<- *prometheus.Desc) {
	c.totalInteractions.Describe(ch)
	c.totalInteractionsErrored.Describe(ch)
}
func (c discordInteractionCollector) Collect(ch chan<- prometheus.Metric) {
	c.totalInteractions.Collect(ch)
	c.totalInteractionsErrored.Collect(ch)
}

func (c discordInteractionCollector) Middleware() middleware.MiddlewareFunc {
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

func NewDiscordInteractionCollector(prefix string) discordInteractionCollector {
	return discordInteractionCollector{
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

type discordGatewayCollector struct {
	totalEvents        *prometheus.CounterVec
	totalEventsErrored *prometheus.CounterVec
}

func (c discordGatewayCollector) Describe(ch chan<- *prometheus.Desc) {
	c.totalEvents.Describe(ch)
	c.totalEventsErrored.Describe(ch)
}
func (c discordGatewayCollector) Collect(ch chan<- prometheus.Metric) {
	c.totalEvents.Collect(ch)
	c.totalEventsErrored.Collect(ch)
}

func (c discordGatewayCollector) Middleware() middleware.MiddlewareFunc {
	return func(ctx common.Context, next func(common.Context) error) func(common.Context) error {
		return func(ctx common.Context) error {
			name := "gateway_event_" + ctx.ID()
			c.totalEvents.WithLabelValues(name).Inc()
			err := next(ctx)
			if err != nil {
				c.totalEventsErrored.WithLabelValues(name).Inc()
			}
			return err
		}
	}
}

func NewDiscordGatewayCollector(prefix string) discordGatewayCollector {
	return discordGatewayCollector{
		totalEventsErrored: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: fmt.Sprintf("discord_%s_events_errored_total", prefix),
				Help: "Number of errored Discord gateway events.",
			},
			[]string{"name"},
		),
		totalEvents: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: fmt.Sprintf("discord_%s_events_handled_total", prefix),
				Help: "Number of handled Discord gateway events.",
			},
			[]string{"name"},
		),
	}
}
