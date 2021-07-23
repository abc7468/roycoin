package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

//singleton pattern 단 하나의 instance만을 공유하는 구조

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

// block의 주소를 저장하는 배열
type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

//가진 데이터와 이전의 hash를 합한 새로운 hash 생성
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

// blockchain의 마지막 hash값 return
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

// block생성 후 생성된 block의 주솟값 return
func createBlock(data string) *block {
	newBlock := block{Data: data, Hash: "", PrevHash: getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// createBlock함수로 생성된 block을 blockchain에 추가
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))

}

// blockchain instance를 가지고 오는 함수
func GetBlockchain() *blockchain {
	if b == nil {
		//병렬적으로 처리한다고 해도 딱 한번만 실행하도록 도와주는 함수
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

//blockchain이 가지고있는 block의 주솟값을 가진 list를 return
func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}
