package timex

import (
	"sync"
	"time"
)

type Resolveable interface {
	Interval() time.Duration
	Current() time.Time
	Next() time.Time
	Start() time.Time
}

type TickableResolver interface {
	Resolveable
	Tick(d time.Duration) time.Time
}

type Resolver struct {
	current  time.Time
	interval time.Duration
	next     time.Time
	start    time.Time
	wg       sync.RWMutex
}

func (r *Resolver) Interval() time.Duration {
	r.wg.RLock()
	defer r.wg.RUnlock()
	return r.interval
}

func (r *Resolver) Current() time.Time {
	r.wg.RLock()
	defer r.wg.RUnlock()
	return r.current
}

func (r *Resolver) Start() time.Time {
	r.wg.RLock()
	defer r.wg.RUnlock()
	return r.start
}

func (r *Resolver) Next() time.Time {
	r.wg.Lock()
	defer r.wg.Unlock()
	next := r.next
	r.current = next
	r.next = next.Add(r.interval)

	return r.current
}

func (r *Resolver) Tick(d time.Duration) time.Time {
	r.wg.Lock()
	defer r.wg.Unlock()
	r.current = r.current.Add(d)
	return r.current
}

func NewResolver(interval time.Duration, start time.Time) *Resolver {
	return &Resolver{
		current:  start,
		interval: interval,
		next:     start.Add(interval),
		start:    start,
	}
}

func DayResolver(start time.Time) *Resolver {
	return NewResolver(DAY, start)
}

func WeekResolver(start time.Time) *Resolver {
	return NewResolver(WEEK, start)
}
