package timex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LookupDay(t *testing.T) {
	t.Parallel()

	table := []struct {
		in  string
		exp Day
		err bool
	}{
		{in: "MONDAY", exp: MONDAY},
		{in: "TUESDAY", exp: TUESDAY},
		{in: "WEDNESDAY", exp: WEDNESDAY},
		{in: "THURSDAY", exp: THURSDAY},
		{in: "FRIDAY", exp: FRIDAY},
		{in: "SATURDAY", exp: SATURDAY},
		{in: "SUNDAY", exp: SUNDAY},
		{in: "thrusday", err: true},
		{in: "0", exp: MONDAY},
		{in: "1", exp: TUESDAY},
		{in: "2", exp: WEDNESDAY},
		{in: "3", exp: THURSDAY},
		{in: "4", exp: FRIDAY},
		{in: "5", exp: SATURDAY},
		{in: "6", exp: SUNDAY},
		{in: "42", err: true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			r := require.New(t)

			act, err := LookupDay(tt.in)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)

			r.Equal(tt.exp, act)

		})
	}

}
