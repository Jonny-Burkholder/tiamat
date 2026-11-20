package consume

import "github.com/jonny-burkholder/tiamat/internal/models"

type thresholdType int

const (
	TypeMaxErr     thresholdType = iota // run until a max number of errors is reached
	TypeErrPercent                      // run until a certain percent of responses are errors TODO: probably use something like RMS for this, maybe have different calcs
	TypeErrRate                         // run until a certain number of errors are reached for a sliding time window
)

func New() Dragon {
	return Dragon{
		threshold: threshold{},
	}
}

// Dragon consumes the API's resources
type Dragon struct {
	models.Config
	threshold
}

type threshold struct {
	// typ thresholdType
}
