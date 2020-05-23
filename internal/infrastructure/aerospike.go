package infrastructure

import "context"

type Key struct {
	Namespace string
	SetName   string
	UserKey   string
}

type Aerospike interface {
	OrderedListAppend(ctx context.Context, collectionName string, key Key, bin string, data []interface{}) (err error)
	MapGetByKey(ctx context.Context, collectionName string, key Key, bin string, mapKey interface{}) (out interface{}, err error)
	MapPutByKey(ctx context.Context, collectionName string, key Key, bin string, mapKey interface{}, mapValue interface{}) (err error)
}
