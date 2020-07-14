package infrastructure

import "context"

type Key struct {
	Namespace string
	SetName   string
	UserKey   string
}

type Aerospike interface {
	OrderedListAppend(ctx context.Context, key Key, bin string, data []interface{}) (err error)
	OrderedListGetByValueRange(ctx context.Context, key Key, bin string, beginValue interface{}, endValue interface{}) (out interface{}, err error)
	OrderedListRemoveByValueList(ctx context.Context, key Key, bin string, value []interface{}) (out interface{}, err error)
	MapGetByKey(ctx context.Context, key Key, bin string, mapKey interface{}) (out interface{}, err error)
	MapPutByKey(ctx context.Context, key Key, bin string, mapKey interface{}, mapValue interface{}) (err error)
}
