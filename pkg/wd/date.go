package wd

import "time"

const DateFormatString = "2006-01-02-07:00"

// Date is a time.Time that understands the format workday uses for dates.
//
// The date format used by workday is "2006-01-02-07:00".
type Date struct {
	time.Time
}

func (t *Date) UnmarshalText(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t.Time, err = time.Parse(DateFormatString, string(data))
	return err
}

func (t Date) MarshalText() (text []byte, err error) {
	return []byte(t.Format(DateFormatString)), nil
}

// NewDate creates a new data given a Workday formatted date string.
//
// The date format used by workday is "2006-01-02-07:00".
func NewDate(s string) (Date, error) {
	t, err := time.Parse(DateFormatString, s)
	if err != nil {
		return Date{}, err
	}
	return Date{Time: t}, nil
}

// MustDate is similar to NewDate but panics in the event of an error.
//
// This is useful for testing but should not be used in place of NewDate.
func MustDate(s string) Date {
	t, err := time.Parse(DateFormatString, s)
	if err != nil {
		panic(err)
	}
	return Date{Time: t}
}
