package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Resolvers(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	start, err := time.Parse(time.RFC822, "01 Jan 21 00:00 EST")
	r.NoError(err)

	table := []struct {
		name string
		res  Resolveable
		int  time.Duration
	}{
		{name: "day", res: DayResolver(start), int: DAY},
		{name: "week", res: WeekResolver(start), int: WEEK},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			exp := start
			res := tt.res

			r.Equal(tt.int, res.Interval())
			r.Equal(exp, res.Current())
			r.Equal(exp, res.Start())

			tk, ok := res.(TickableResolver)
			r.True(ok)
			cur := tk.Tick(4 * time.Hour)
			r.Equal(start.Add(4*time.Hour), cur)
			r.Equal(cur, res.Current())
			r.Equal(start, res.Start())

			next := res.Next()

			exp = start.Add(res.Interval())
			r.Equal(exp, next)
			r.Equal(exp, res.Current())
			r.Equal(start, res.Start())
		})
	}

}
