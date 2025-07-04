// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

package debug

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/spacemonkeygo/monkit/v3"
)

// ApplyNewTransformers adds the default set of monkit.CallbackTransformers.
// This needs to happen individually for each output type and endpoint.
func ApplyNewTransformers(r *monkit.Registry) *monkit.Registry {
	return r.WithTransformers(monkit.NewDeltaTransformer())
}

// PrometheusEndpointConfig is a configuration for PrometheusEndpoint.
type PrometheusEndpointConfig struct {
	SanitizeValues bool `help:"Sanitize all values" default:"true"`
}

// PrometheusEndpoint includes all the information to server Prometheus compatible HTTP pages.
type PrometheusEndpoint struct {
	registryMu     sync.Mutex
	registries     map[string]*monkit.Registry
	baseRegistry   *monkit.Registry
	sanitizeValues bool
}

// NewPrometheusEndpointFromConfig creates an initialized PrometheusEndpoint based on configuration..
func NewPrometheusEndpointFromConfig(registry *monkit.Registry, cfg PrometheusEndpointConfig) *PrometheusEndpoint {
	return &PrometheusEndpoint{
		baseRegistry:   registry,
		registries:     map[string]*monkit.Registry{},
		sanitizeValues: cfg.SanitizeValues,
	}
}

// NewPrometheusEndpoint creates an initialized PrometheusEndpoint.
func NewPrometheusEndpoint(registry *monkit.Registry) *PrometheusEndpoint {
	return &PrometheusEndpoint{
		baseRegistry:   registry,
		registries:     map[string]*monkit.Registry{},
		sanitizeValues: true,
	}
}

// supportsGzip checks if the client supports gzip compression.
func supportsGzip(r *http.Request) bool {
	acceptEncoding := r.Header.Get("Accept-Encoding")
	return strings.Contains(acceptEncoding, "gzip")
}

// PrometheusMetrics writes monkit data in  https://prometheus.io/docs/instrumenting/exposition_formats/.
func (server *PrometheusEndpoint) PrometheusMetrics(w http.ResponseWriter, r *http.Request) {
	// We have to collect all of the metrics before we write. This is because we do not
	// get all of the metrics from the registry sorted by measurement, and from the docs:
	//
	//     All lines for a given metric must be provided as one single group, with the
	//     optional HELP and TYPE lines first (in no particular order). Beyond that,
	//     reproducible sorting in repeated expositions is preferred but not required,
	//     i.e. do not sort if the computational cost is prohibitive.

	var writer io.Writer
	writer = w
	if supportsGzip(r) {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		gzipWriter := gzip.NewWriter(w)
		writer = gzipWriter
		defer func() {
			if err := gzipWriter.Close(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
	}

	data := make(map[string][]string)
	var components []string

	server.registryForRequest(r).Stats(func(key monkit.SeriesKey, field string, val float64) {
		components = components[:0]

		measurement := sanitize(key.Measurement)
		for tag, tagVal := range key.Tags.All() {
			if server.sanitizeValues {
				tagVal = sanitize(tagVal)
			}
			components = append(components, fmt.Sprintf("%s=%q", sanitize(tag), tagVal))
		}
		components = append(components,
			fmt.Sprintf("field=%q", sanitize(field)))

		data[measurement] = append(data[measurement],
			fmt.Sprintf("{%s} %g", strings.Join(components, ","), val))
	})

	for measurement, samples := range data {
		_, _ = fmt.Fprintln(writer, "# TYPE", measurement, "gauge")
		for _, sample := range samples {
			_, _ = fmt.Fprintf(writer, "%s%s\n", measurement, sample)
		}
	}
}

func (server *PrometheusEndpoint) registryForRequest(r *http.Request) *monkit.Registry {
	outputID := r.URL.Query().Get("output-id")
	server.registryMu.Lock()
	defer server.registryMu.Unlock()
	// okay if outputID is ""
	reg, found := server.registries[outputID]
	if !found {
		reg = ApplyNewTransformers(server.baseRegistry)
		server.registries[outputID] = reg
	}
	return reg
}
