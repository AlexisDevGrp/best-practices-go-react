package examples

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func TestFoo(t *testing.T) {
	// optional: t.Parallel()

	for k, tc := range []struct {
		assert bool
	}{
		{assert: true},
		{assert: false},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			assert.True(t, tc.assert)
		})
	}
}




















