package compare

import "errors"

var (
	ErrNotEnoughResults = errors.New("command 'compare' requires at least 2 benchmark result files")
)
