package govidat

type County struct {
	id             string
	Name           string
	Residents      int64
	Cases          int64
	Deaths         int64
	Incidence7d    int64
	Municipalities []Municipality
}
