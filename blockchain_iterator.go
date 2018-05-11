package main

import (
	"log"

	"github.com/boltdb/bolt"
)

// BlockchainIterator is used to iterate over blockchian blocks
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Next return next block starting from the tip
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodeBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodeBlock)

		return nil

	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
