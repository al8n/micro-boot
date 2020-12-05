package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
	"time"
)



func TestConfig(t *testing.T) {
	var yamlCases = map[string]Config {
		`
path: /metrics
counter-options:
  opt1: 
    namespace: ns1
    name: n1
    subsystem: s1
    help: h1
    labels:
      l1: v1
      l2: v2
gauge-options:
  opt1: 
    namespace: ns1
    name: n1
    subsystem: s1
    help: h1
    labels:
      l1: v1
      l2: v2
histogram-options:
  opt1: 
    namespace: ns1
    name: n1
    subsystem: s1
    help: h1
    labels:
      l1: v1
      l2: v2
    buckets: 
      - 1
      - 5
      - 6
summary-options:
  opt1: 
    namespace: ns1
    name: n1
    subsystem: s1
    help: h1
    labels:
      l1: v1
      l2: v2
    objectives:
      3.5: 6.8
      1.4: 8.9
    max-age: 1h
    age-buckets: 20
    buf-cap: 10
    
handler-options: 
  compression: true
  max-requests-in-flight: 20
  timeout: 15s
  error-handling: 10
`: {
			Path: "/metrics",
			CounterOptions: map[string]CounterOptions{
				"opt1": {
					Namespace:   "ns1",
					Subsystem:   "s1",
					Name:        "n1",
					Help:        "h1",
					ConstLabels: map[string]string{
						"l1": "v1",
						"l2": "v2",
					},
				},
			},
			GaugeOptions: map[string]GaugeOptions{
				"opt1": {
					Namespace:   "ns1",
					Subsystem:   "s1",
					Name:        "n1",
					Help:        "h1",
					ConstLabels: map[string]string{
						"l1": "v1",
						"l2": "v2",
					},
				},
			},
			SummaryOptions: map[string]SummaryOptions{
				"opt1": {
					Namespace:   "ns1",
					Subsystem:   "s1",
					Name:        "n1",
					Help:        "h1",
					ConstLabels: map[string]string{
						"l1": "v1",
						"l2": "v2",
					},
					Objectives: map[float64]float64{
						3.5: 6.8,
						1.4: 8.9,
					},
					MaxAge: time.Hour,
					BufCap: 10,
					AgeBuckets: 20,
				},
			},
			HistogramOptions: map[string]HistogramOptions{
				"opt1": {
					Namespace:   "ns1",
					Subsystem:   "s1",
					Name:        "n1",
					Help:        "h1",
					ConstLabels: map[string]string{
						"l1": "v1",
						"l2": "v2",
					},
					Buckets: []float64{
						1,
						5,
						6,
					},
				},
			},
		},
	}

	for mock, exp := range yamlCases {
		var settings Config
		err := yaml.Unmarshal([]byte(mock), &settings)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, settings.Path, exp.Path)
		for key, options := range settings.CounterOptions {
			expOpt := exp.CounterOptions[key]
			assert.Equal(t, options, expOpt)
		}
		for key, options := range settings.GaugeOptions {
			expOpt := exp.GaugeOptions[key]
			assert.Equal(t, options, expOpt)
		}
		for key, options := range settings.HistogramOptions {
			expOpt := exp.HistogramOptions[key]
			assert.Equal(t, options, expOpt)
		}
		for key, options := range settings.SummaryOptions {
			expOpt := exp.SummaryOptions[key]
			assert.Equal(t, options, expOpt)
		}

		assert.Equal(t, settings.HandlerOptions.Standardize(), promhttp.HandlerOpts{
			ErrorHandling:       10,
			DisableCompression:  true,
			MaxRequestsInFlight: 20,
			Timeout:             15 * time.Second,
		})
	}
}

func TestOpts_Standardize(t *testing.T) {
	var opts =  CounterOptions{
		Name: "opt1",
	}

	assert.Equal(t, opts.Standardize(), prometheus.CounterOpts{Name: "opt1"})
}
