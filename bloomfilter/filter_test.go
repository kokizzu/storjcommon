// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package bloomfilter_test

import (
	"flag"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"storj.io/common/bloomfilter"
	"storj.io/common/memory"
	"storj.io/common/storj"
	"storj.io/common/testrand"
)

func TestNoFalsePositive(t *testing.T) {
	const numberOfPieces = 10000
	pieceIDs := generateTestIDs(numberOfPieces)

	for _, ratio := range []float32{0.5, 1, 2} {
		size := int64(numberOfPieces * ratio)
		filter := bloomfilter.NewOptimal(size, 0.1)
		for _, pieceID := range pieceIDs {
			filter.Add(pieceID)
		}
		for _, pieceID := range pieceIDs {
			require.True(t, filter.Contains(pieceID))
		}
	}
}

func TestBytes(t *testing.T) {
	for _, count := range []int64{0, 100, 1000, 10000} {
		filter := bloomfilter.NewOptimal(count, 0.1)
		for range count {
			id := testrand.PieceID()
			filter.Add(id)
		}

		bytes := filter.Bytes()
		unmarshaled, err := bloomfilter.NewFromBytes(bytes)
		require.NoError(t, err)

		require.Equal(t, filter, unmarshaled)
	}
}

func TestBytes_Failing(t *testing.T) {
	failing := [][]byte{
		{},
		{0},
		{1},
		{1, 0},
		{255, 10, 10, 10},
	}
	for _, bytes := range failing {
		_, err := bloomfilter.NewFromBytes(bytes)
		require.Error(t, err)
	}
}

// generateTestIDs generates n piece ids.
func generateTestIDs(n int) []storj.PieceID {
	ids := make([]storj.PieceID, n)
	for i := range ids {
		ids[i] = testrand.PieceID()
	}
	return ids
}

func BenchmarkFilterAdd(b *testing.B) {
	ids := generateTestIDs(100000)
	filter := bloomfilter.NewOptimal(int64(len(ids)), 0.1)

	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filter.Add(ids[i%len(ids)])
		}
	})

	b.Run("Contains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filter.Contains(ids[i%len(ids)])
		}
	})
}

func TestGolden(t *testing.T) {
	ids := []storj.PieceID{
		{0x52, 0xfd, 0xfc, 0x7, 0x21, 0x82, 0x65, 0x4f, 0x16, 0x3f, 0x5f, 0xf, 0x9a, 0x62, 0x1d, 0x72, 0x95, 0x66, 0xc7, 0x4d, 0x10, 0x3, 0x7c, 0x4d, 0x7b, 0xbb, 0x4, 0x7, 0xd1, 0xe2, 0xc6, 0x49},
		{0x81, 0x85, 0x5a, 0xd8, 0x68, 0x1d, 0xd, 0x86, 0xd1, 0xe9, 0x1e, 0x0, 0x16, 0x79, 0x39, 0xcb, 0x66, 0x94, 0xd2, 0xc4, 0x22, 0xac, 0xd2, 0x8, 0xa0, 0x7, 0x29, 0x39, 0x48, 0x7f, 0x69, 0x99},
		{0xeb, 0x9d, 0x18, 0xa4, 0x47, 0x84, 0x4, 0x5d, 0x87, 0xf3, 0xc6, 0x7c, 0xf2, 0x27, 0x46, 0xe9, 0x95, 0xaf, 0x5a, 0x25, 0x36, 0x79, 0x51, 0xba, 0xa2, 0xff, 0x6c, 0xd4, 0x71, 0xc4, 0x83, 0xf1},
		{0x5f, 0xb9, 0xb, 0xad, 0xb3, 0x7c, 0x58, 0x21, 0xb6, 0xd9, 0x55, 0x26, 0xa4, 0x1a, 0x95, 0x4, 0x68, 0xb, 0x4e, 0x7c, 0x8b, 0x76, 0x3a, 0x1b, 0x1d, 0x49, 0xd4, 0x95, 0x5c, 0x84, 0x86, 0x21},
		{0x63, 0x25, 0x25, 0x3f, 0xec, 0x73, 0x8d, 0xd7, 0xa9, 0xe2, 0x8b, 0xf9, 0x21, 0x11, 0x9c, 0x16, 0xf, 0x7, 0x2, 0x44, 0x86, 0x15, 0xbb, 0xda, 0x8, 0x31, 0x3f, 0x6a, 0x8e, 0xb6, 0x68, 0xd2},
	}
	expected := []byte{0x1, 0x99, 0x3, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x80, 0x0, 0x0, 0x20}

	filter := bloomfilter.NewExplicit(153, 3, 256)
	for _, id := range ids {
		filter.Add(id)
	}

	for _, id := range ids {
		require.True(t, filter.Contains(id))
	}

	require.Equal(t, expected, filter.Bytes())
}

var approximateFalsePositives = flag.Bool("approximate-false-positive-rate", false, "")

func TestApproximateFalsePositives(t *testing.T) {
	if !*approximateFalsePositives {
		t.Skip("Use --approximate-false-positive-rate to enable diagnostic test.")
	}

	const measurements = 100
	const validation = 1000

	for _, p := range []float64{0.01, 0.05, 0.1, 0.2, 0.3} {
		for _, n := range []int64{1000, 10000, 100000, 1000000} {
			fpp := []float64{}

			for range measurements {
				filter := bloomfilter.NewOptimal(n, p)
				for range n {
					filter.Add(testrand.PieceID())
				}

				positive := 0
				for range validation {
					if filter.Contains(testrand.PieceID()) {
						positive++
					}
				}
				fpp = append(fpp, float64(positive)/validation)
			}

			hashCount, size := bloomfilter.NewOptimal(n, p).Parameters()
			summary := summarize(p, fpp)
			t.Logf("n=%8d p=%.2f avg=%.2f min=%.2f mean=%.2f max=%.2f mse=%.3f hc=%d sz=%s", n, p, summary.avg, summary.min, summary.mean, summary.max, summary.mse, hashCount, memory.Size(size))
		}
	}
}

func TestAddFilter(t *testing.T) {
	doesNotContainAtLeastOne := func(filter *bloomfilter.Filter, ids []storj.PieceID) bool {
		for _, id := range ids {
			if !filter.Contains(id) {
				return true
			}
		}
		return false
	}

	ids1 := generateTestIDs(50000)
	ids2 := generateTestIDs(50000)

	filter1 := bloomfilter.NewOptimal(25000, 0.1)
	for _, id := range ids1 {
		filter1.Add(id)
	}

	filter2 := bloomfilter.NewExplicit(filter1.SeedAndParameters())
	for _, id := range ids2 {
		filter2.Add(id)
	}

	require.True(t, doesNotContainAtLeastOne(filter1, ids2), "at least one ID from the 2nd set should not be contained before merge")

	err := filter1.AddFilter(filter2)
	require.NoError(t, err)

	require.False(t, doesNotContainAtLeastOne(filter1, ids2), "all IDs from the 2nd set should be contained after merge")
}

func TestAddFilter_Bad(t *testing.T) {
	t.Run("mismatched seed", func(t *testing.T) {
		filter1 := bloomfilter.NewExplicit(100, 4, 300)
		filter2 := bloomfilter.NewExplicit(101, 4, 300)
		err := filter1.AddFilter(filter2)
		require.EqualError(t, err, "cannot merge: mismatched seed: expected 100 but got 101")
	})
	t.Run("mismatched heap count", func(t *testing.T) {
		filter1 := bloomfilter.NewExplicit(100, 4, 300)
		filter2 := bloomfilter.NewExplicit(100, 5, 300)
		err := filter1.AddFilter(filter2)
		require.EqualError(t, err, "cannot merge: mismatched hash count: expected 4 but got 5")
	})
	t.Run("mismatched table size", func(t *testing.T) {
		filter1 := bloomfilter.NewExplicit(100, 4, 300)
		filter2 := bloomfilter.NewExplicit(100, 4, 400)
		err := filter1.AddFilter(filter2)
		require.EqualError(t, err, "cannot merge: mismatched table size: expected 300 but got 400")
	})
}

type stats struct {
	avg, min, mean, max, mse float64
}

// summarize calculates average, minimum, maximum and mean squared error.
func summarize(expected float64, values []float64) (r stats) {
	sort.Float64s(values)

	for _, v := range values {
		r.avg += v
		r.mse += (v - expected) * (v - expected)
	}
	r.avg /= float64(len(values))
	r.mse /= float64(len(values))

	r.min, r.mean, r.max = values[0], values[len(values)/2], values[len(values)-1]

	return r
}

func TestFillRate(t *testing.T) {
	const half = 0b_00001111
	const quarter = 0b_00000011

	tests := []struct {
		data   []byte
		expect float64
	}{
		{[]byte{1, 0x00, 0x04, 0x00, 0x00, 0x00}, 0},
		{[]byte{1, 0x00, 0x04, 0xFF, 0xFF, 0xFF}, 1},
		{[]byte{1, 0x00, 0x04, half, half, half}, 0.5},
		{[]byte{1, 0x00, 0x04, quarter, quarter, quarter}, 0.25},
	}

	for _, test := range tests {
		filter, err := bloomfilter.NewFromBytes(test.data)
		require.NoError(t, err)
		require.InDelta(t, test.expect, filter.FillRate(), 0.001)
	}
}
