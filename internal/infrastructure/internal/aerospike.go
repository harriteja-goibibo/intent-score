package internal

import (
	aero "github.com/aerospike/aerospike-client-go"
)

// AeroClient is the interface replacement for aerospike's client.
// All the functions should have exact same name and signature as aerospike's.
type AeroClient interface {
	Operate(policy *aero.WritePolicy, key *aero.Key, operations ...*aero.Operation) (*aero.Record, error)
	Get(policy *aero.BasePolicy, key *aero.Key, binNames ...string) (*aero.Record, error)
}
