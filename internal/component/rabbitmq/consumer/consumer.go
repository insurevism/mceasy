package consumer

import (
	"context"
)

type Consumer[Data any] interface {
	Handle(ctx context.Context, data Data) error
	ParseData(data []byte) (Data, error)
}
