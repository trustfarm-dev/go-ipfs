package gc

import (
	"testing"

	"github.com/ipfs/go-ipfs/blocks/blocksutil"

	cid "gx/ipfs/QmTprEaAA2A9bst5XH7exuyi5KzNMK3SEDNN8rBDnKWcUS/go-cid"
)

func TestInsertWhite(t *testing.T) {
	tri := newTriset()
	blkgen := blocksutil.NewBlockGenerator()

	whites := make([]*cid.Cid, 1000)
	for i := range whites {
		blk := blkgen.Next()
		whites[i] = blk.Cid()

		tri.InsertWhite(blk.Cid())
	}

	for _, v := range whites {
		if tri.colmap[v.KeyString()].getColor() != tri.white {
			t.Errorf("cid %s should be white and is not %s", v, tri.colmap[v.KeyString()])
		}
	}

}
