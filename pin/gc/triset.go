package gc

import (
	"context"

	dag "github.com/ipfs/go-ipfs/merkledag"

	cid "gx/ipfs/QmTprEaAA2A9bst5XH7exuyi5KzNMK3SEDNN8rBDnKWcUS/go-cid"
)

type color uint8

const (
	colorNull color = iota
	color1
	color2
	color3
)

// try to keep trielement as small as possible
// 8bit color is overkill, 2 bits would be enough, maybe pack in future
// if ever it blows over 64bits total size change the colmap in triset to use
// pointers onto this structure
type trielement struct {
	c color
}

type triset struct {
	// colors are per triset allowing fast color swap after sweep,
	// the update operation in map of structs is about 3x as fast
	// as insert and requres 0 allocations (keysize allocation in case of insert)
	white, gray, black color

	freshColor color

	// grays is used as stack to enumerate elements that are still gray
	grays []cid.Cid

	// if item doesn't exist in the colmap it is treated as white
	colmap map[string]trielement
}

func newTriset() *triset {
	tr := &triset{
		white: color1,
		gray:  color2,
		black: color3,

		grays:  make([]cid.Cid, 0, 1<<10),
		colmap: make(map[string]trielement),
	}

	tr.freshColor = tr.white
	return tr
}

// InsertFresh inserts fresh item into a set
// it marks it with freshColor if it is currently white
func (tr *triset) InsertFresh(c cid.Cid) {
	e := tr.colmap[c.KeyString()]

	if e.c == colorNull || (e.c == tr.white && tr.freshColor != tr.white) {
		tr.colmap[c.KeyString()] = trielement{tr.freshColor}
	}
}

// InsertWhite inserts white item into a set if it doesn't exist
func (tr *triset) InsertWhite(c cid.Cid) {
	_, ok := tr.colmap[c.KeyString()]
	if !ok {
		tr.colmap[c.KeyString()] = trielement{tr.white}
	}
}

// InsertGray inserts new item into set as gray or turns white item into gray
func (tr *triset) InsertGray(c cid.Cid) {
	e := tr.colmap[c.KeyString()]
	if e.c == colorNull || e.c == tr.white {
		tr.colmap[c.KeyString()] = trielement{tr.gray}
		tr.grays = append(tr.grays, c)
	}
}

func (tr *triset) blacken(key string) {
	tr.colmap[key] = trielement{tr.black}
}

// EnumerateStep performs one Links lookup in search for elements to gray out
// it returns error is the getLinks function errors
// if the gray set is empty after this step it returns (true, nil)
func (tr *triset) EnumerateStep(ctx context.Context, getLinks dag.GetLinks) (bool, error) {
	var c cid.Cid
	for next := true; next; next = tr.colmap[c.KeyString()].c != tr.gray {
		if len(tr.grays) == 0 {
			return true, nil
		}
		// get element from top of queue
		c := tr.grays(len(tr.grays) - 1)
		tr.grays = tr.grays[:len(tr.grays)-1]
	}
}
