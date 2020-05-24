package aerospike

import (
	"context"
	aero "github.com/aerospike/aerospike-client-go"
	aeroTypes "github.com/aerospike/aerospike-client-go/types"
	"github.com/goibibo/intent-score/internal/infrastructure"
	"github.com/goibibo/intent-score/internal/infrastructure/internal"
	"github.com/pkg/errors"
	"sync"
)

// keyPool is the buffer used for reusing aerospike's Key.
type keyPool struct {
	buffer []*aero.Key
	sync.Mutex
}

// Get gets or create an aerospike key for the given paremeter
func (p *keyPool) Get(namespace, setName, key string) *aero.Key {
	if len(p.buffer) == 0 {
		// Create a new key if the buffer is empty
		k, _ := aero.NewKey(namespace, setName, key)
		return k
	}

	p.Lock()

	// Pop out the recent key from buffer.
	k := p.buffer[len(p.buffer)-1]
	p.buffer = p.buffer[:len(p.buffer)-1]

	if k.Namespace() == namespace && k.SetName() == setName {
		// Set value only if the namespace & setname is same.
		k.SetValue(aero.NewValue(key))
	} else {
		k, _ = aero.NewKey(namespace, setName, key)
	}

	p.Unlock()
	return k
}

// Put puts the aerospike key back to the pool.
func (p *keyPool) Put(key *aero.Key) {
	p.Lock()
	p.buffer = append(p.buffer, key)
	p.Unlock()
}

// Client is the main module's struct for interacting with the aerospike server.
type Client struct {
	Client  internal.AeroClient
	keyPool *keyPool
}

// NewAeroClient creates a client on top of aerospike's client interface.
func NewAeroClient(c internal.AeroClient) (*Client, error) {
	cli := Client{}
	cli.keyPool = &keyPool{buffer: make([]*aero.Key, 0, 1024)}
	cli.Client = c

	return &cli, nil
}

func ignoreKeyNotFoundError(err error) error {
	if aeroErr, ok := err.(aeroTypes.AerospikeError); ok {
		if aeroErr.ResultCode() == aeroTypes.KEY_NOT_FOUND_ERROR {
			err = nil
		}
	}
	return err
}

// OrderedListAppend appends the data to the specified bin for the input key in an ordered fashion.
// The input is a map hence the order of execution is not guaranteed.
func (cli *Client) OrderedListAppend(ctx context.Context, key infrastructure.Key, bin string, data []interface{}) (err error) {

	policy := aero.NewWritePolicy(0, 0)
	listPolicy := aero.NewListPolicy(aero.ListOrderOrdered, aero.ListWriteFlagsDefault)

	aeroKey := cli.keyPool.Get(key.Namespace, key.SetName, key.UserKey)

	operations := make([]*aero.Operation, len(data))
	counter := 0
	for _, value := range data {
		operations[counter] = aero.ListAppendWithPolicyOp(listPolicy, bin, value)
		counter++
	}

	if _, setErr := cli.Client.Operate(policy, aeroKey, operations...); setErr != nil {
		err = errors.Wrap(ignoreKeyNotFoundError(setErr), "Client Put error")
	}

	return
}

func (cli *Client) OrderedListGetByValueRange(ctx context.Context, key infrastructure.Key, bin string, beginValue interface{}, endValue interface{}) (out interface{}, err error) {

	policy := aero.NewWritePolicy(0, 0)
	aeroKey := cli.keyPool.Get(key.Namespace, key.SetName, key.UserKey)

	operations := make([]*aero.Operation, 1)
	operations[0] = aero.ListGetByValueRangeOp(bin, beginValue, endValue, aero.ListReturnTypeValue)

	if record, getErr := cli.Client.Operate(policy, aeroKey, operations...); getErr != nil {
		err = errors.Wrap(ignoreKeyNotFoundError(getErr), "Client OrderedListGetByValueRange error")
		return out, err
	} else {
		return record.Bins[bin], err
	}
}

func (cli *Client) OrderedListRemoveByValueList(ctx context.Context, key infrastructure.Key, bin string, value []interface{}) (out interface{}, err error) {
	policy := aero.NewWritePolicy(0, 0)
	aeroKey := cli.keyPool.Get(key.Namespace, key.SetName, key.UserKey)

	operations := make([]*aero.Operation, 1)
	operations[0] = aero.ListRemoveByValueListOp(bin, value, aero.ListReturnTypeCount)

	if record, getErr := cli.Client.Operate(policy, aeroKey, operations...); getErr != nil {
		err = errors.Wrap(ignoreKeyNotFoundError(getErr), "Client OrderedListRemoveByValueList error")
		return out, err
	} else {
		return record.Bins[bin], err
	}
}

func (cli *Client) MapGetByKey(ctx context.Context, key infrastructure.Key, bin string, mapKey interface{}) (out interface{}, err error) {

	policy := aero.NewWritePolicy(0, 0)
	aeroKey := cli.keyPool.Get(key.Namespace, key.SetName, key.UserKey)

	operations := make([]*aero.Operation, 1)
	operations[0] = aero.MapGetByKeyOp(bin, mapKey, aero.MapReturnType.VALUE)

	if record, setErr := cli.Client.Operate(policy, aeroKey, operations...); setErr != nil {
		err = errors.Wrap(ignoreKeyNotFoundError(setErr), "Client MapGetByKey Operate error")
		return out, err
	} else {
		return record.Bins[bin], nil
	}

}

func (cli *Client) MapPutByKey(ctx context.Context, key infrastructure.Key, bin string, mapKey interface{}, mapValue interface{}) (err error) {

	policy := aero.DefaultMapPolicy()
	aeroKey := cli.keyPool.Get(key.Namespace, key.SetName, key.UserKey)

	operations := make([]*aero.Operation, 1)
	operations[0] = aero.MapPutOp(policy, bin, mapKey, mapValue)

	writePolicy := aero.NewWritePolicy(0, 0)

	if _, setErr := cli.Client.Operate(writePolicy, aeroKey, operations...); setErr != nil {
		err = errors.Wrap(ignoreKeyNotFoundError(setErr), "Client MapPutByKey Operate error")
	}

	return
}
