package testhelper

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestLoadFixture(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NotNil(t, LoadTestDB())
		})
	}
}
