package global

type Range[Type any] struct {
	Min Type `json:"min,omitempty"`
	Max Type `json:"max,omitempty"`
}
