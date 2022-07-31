package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		offset   int64
		limit    int64
		expected string
	}{
		{
			name:     "offset = 0,limit = 0",
			from:     "testdata/input.txt",
			offset:   0,
			limit:    0,
			expected: "testdata/out_offset0_limit0.txt",
		},
		{
			name:     "offset = 0, limit = 10",
			from:     "testdata/input.txt",
			offset:   0,
			limit:    10,
			expected: "testdata/out_offset0_limit10.txt",
		},
		{
			name:     "offset = 0, limit = 1000",
			from:     "testdata/input.txt",
			offset:   0,
			limit:    1000,
			expected: "testdata/out_offset0_limit1000.txt",
		},
		{
			name:     "offset = 0, limit = 10000",
			from:     "testdata/input.txt",
			offset:   0,
			limit:    10000,
			expected: "testdata/out_offset0_limit10000.txt",
		},
		{
			name:     "offset = 100, limit = 1000",
			from:     "testdata/input.txt",
			offset:   100,
			limit:    1000,
			expected: "testdata/out_offset100_limit1000.txt",
		},
		{
			name:     "offset = 6000, limit = 1000",
			from:     "testdata/input.txt",
			offset:   6000,
			limit:    1000,
			expected: "testdata/out_offset6000_limit1000.txt",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			fileTo, _ := os.CreateTemp("testdata", "test_case")
			defer os.Remove(fileTo.Name())
			defer fileTo.Close()
			tmpFile, _ := os.Open(tc.expected)
			defer tmpFile.Close()

			expectedFile, _ := ioutil.ReadAll(tmpFile)
			err := Copy(tc.from, fileTo.Name(), tc.offset, tc.limit)
			require.NoError(t, err)

			gotFile, err := ioutil.ReadAll(fileTo)
			require.NoError(t, err)
			assert.Equal(t, expectedFile, gotFile)
		})
	}
}

func TestCopyErrors(t *testing.T) {
	fromFile := "testdata/input.txt"
	toFile := "testdata/out.txt"

	t.Run("offset exceed file size", func(t *testing.T) {
		var offset int64 = 100000000
		var limit int64
		err := Copy(fromFile, toFile, offset, limit)
		require.Equal(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("offset is negative", func(t *testing.T) {
		var offset int64 = -1
		var limit int64
		err := Copy(fromFile, toFile, offset, limit)
		require.Equal(t, err, ErrOffsetIsNegative)
	})
}
