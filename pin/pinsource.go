package pin

import (
	cid "gx/ipfs/QmTprEaAA2A9bst5XH7exuyi5KzNMK3SEDNN8rBDnKWcUS/go-cid"
)

// Pinsource is structure describing source of the pin
type PinSource struct {
	// Get if a function that will be called to get the pins for GCing
	Get func() ([]*cid.Cid, error)
	// Strict makes the GC fail if some objects can't be fetched during
	// recursive traversal of the graph
	Strict bool
	// Direct marks the pinned object as the final object in the traversal
	Direct bool
	// Inernal marks the pin source which recursive enumeration should be
	// terminated by a direct pin
	Inernal bool
}
