// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package debug

import (
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"storj.io/common/testcontext"
)

func TestServer_PrometheusMetrics(t *testing.T) {
	ctx := testcontext.New(t)
	registry := monkit.NewRegistry()
	srv := NewServer(zaptest.NewLogger(t), nil, registry, Config{})

	registry.ScopeNamed("test").Chain(
		monkit.StatSourceFunc(func(cb func(key monkit.SeriesKey, field string, val float64)) {
			cb(monkit.NewSeriesKey("m1"), "f1", 1)
			cb(monkit.NewSeriesKey("m2"), "f2", 2)
			cb(monkit.NewSeriesKey("m1"), "f3", 3)
			cb(monkit.NewSeriesKey("m3"), "", 4)
		}))

	rec := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	require.NoError(t, err)
	srv.PrometheusEndpoint.PrometheusMetrics(rec, req)

	const (
		m1 = `# TYPE m1 gauge
m1{scope="test",field="f1"} 1
m1{scope="test",field="f3"} 3
`
		m2 = `# TYPE m2 gauge
m2{scope="test",field="f2"} 2
`
		m3 = `# TYPE m3 gauge
m3{scope="test",field=""} 4
`
	)

	body := rec.Body.String()

	if body != m1+m2+m3 && body != m3+m2+m1 && body != m3+m1+m2 && body != m2+m1+m3 && body != m2+m3+m1 {
		t.Fatalf("string not a combination of m1,m2,m3:\nbody:%q\nm1:%q\nm2:%q\nm3:%q", body, m1, m2, m3)
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodGet, "/top", nil)
	require.NoError(t, err)
	rec = httptest.NewRecorder()

	uriCounter := Top.NewTagCounter("http_requests", "uri")
	uriCounter("/one")
	uriCounter("/one")
	uriCounter("/two")

	ServeTop(rec, req)

	s := rec.Body.String()
	require.Contains(t, s, "http_requests_count uri=/one 2")

}

func TestServer_Close(t *testing.T) {
	ctx := testcontext.New(t)

	for i := 0; i < 1000; i++ {
		registry := monkit.NewRegistry()
		lis := newFakeListener()
		srv := NewServer(zaptest.NewLogger(t), lis, registry, Config{
			Crawlspace: true,
		})
		go func() { _ = srv.Close() }()
		require.NoError(t, srv.Run(ctx))
	}
}

type fakeListener struct {
	once sync.Once
	ch   chan struct{}
}

func newFakeListener() net.Listener {
	return &fakeListener{
		ch: make(chan struct{}),
	}
}

func (f *fakeListener) Addr() net.Addr { return nil }

func (f *fakeListener) Accept() (net.Conn, error) {
	<-f.ch
	return nil, errors.New("closed")
}

func (f *fakeListener) Close() error {
	f.once.Do(func() { close(f.ch) })
	return nil
}
