package mined

import (
	"time"

	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/libs/hash"
)

type block struct {
	hash      hash.Hash
	block     blocks.Block
	results   string
	createdOn time.Time
}

func createBlock(
	hash hash.Hash,
	blk blocks.Block,
	results string,
	createdOn time.Time,
) Block {
	out := block{
		hash:      hash,
		block:     blk,
		results:   results,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
}

// Block returns the block
func (obj *block) Block() blocks.Block {
	return obj.block
}

// Results returns the results
func (obj *block) Results() string {
	return obj.results
}

// CreatedOn returns the creation time
func (obj *block) CreatedOn() time.Time {
	return obj.createdOn
}
