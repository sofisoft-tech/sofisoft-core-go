package repositories

import "context"

type IBaseRepository interface {
	Count(ctx context.Context, filter interface{}) (int64, error)
	DeleteById(ctx context.Context, id string) (int64, error)
	FilterBy(ctx context.Context, filter interface{}, receiver []interface{}) error
	FindById(ctx context.Context, id string, receiver interface{}) error
	FindOne(ctx context.Context, filter interface{}, receiver interface{}) error
	InsertMany(ctx context.Context, documents []interface{}) ([]string, error)
	InsertOne(ctx context.Context, document interface{}) (string, error)
	Paginated(ctx context.Context, filter interface{}, sort interface{}, pageSize int64, start int64, receiver interface{}) error
	UpdateOne(ctx context.Context, document interface{}) error
}
