package daterange

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type key int

const DateRangeKey key = 0

type DateRange struct {
	Begin, End time.Time
}

func NewContext(ctx context.Context, date DateRange) context.Context {
	return context.WithValue(ctx, DateRangeKey, date)
}

func FromContext(ctx context.Context) (DateRange, bool) {
	// type assertion, ok=false if assertion fails
	t, ok := ctx.Value(DateRangeKey).(DateRange)
	return t, ok
}

func FromRequest(r *http.Request) (date DateRange, err error) {
	begin, err := time.Parse("20060102", r.FormValue("begin"))
	if err != nil {
		err = fmt.Errorf("Error parsing begin time. %s\n", err.Error())
		return
	}
	end, err := time.Parse("20060102", r.FormValue("end"))
	if err != nil {
		err = fmt.Errorf("Error parsing end time. %s\n", err.Error())
		return
	}
	date = DateRange{begin, end}
	return
}

func NewContextFromRequest(ctx context.Context, r *http.Request) (res context.Context, err error) {
	date, err := FromRequest(r)
	if err != nil {
		return
	}
	res = NewContext(ctx, date)
	return
}
