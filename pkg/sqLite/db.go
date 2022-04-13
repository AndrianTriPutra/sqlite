package sqLite

import "context"

type IDatabase interface {
	Connect(ctx context.Context) error
	Db(ctx context.Context) interface{}
}
