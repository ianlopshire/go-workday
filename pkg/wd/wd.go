package wd

// Bool returns a pointer to b.
func Bool(b bool) *bool {
	return &b
}

// String returns a pointer to s.
func String(s string) *string {
	return &s
}

// Int returns a pointer to i.
func Int(i int) *int {
	return &i
}

// Float returns a pointer to f.
func Float(f float64) *float64 {
	return &f
}
