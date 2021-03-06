package blockchain

import (
	"reflect"
	"sync"
	"testing"

	"github.com/abc7468/roycoin/utils"
)

type fakeDB struct {
	fakeLoadCahin func() []byte
	fakeFindBlock func() []byte
}

func (f fakeDB) FindBlock(hash string) []byte {
	return f.fakeFindBlock()

}
func (f fakeDB) LoadChain() []byte {
	return f.fakeLoadCahin()
}
func (fakeDB) SaveChain(data []byte)              {}
func (fakeDB) SaveBlock(hash string, data []byte) {}
func (fakeDB) DeleteAllBlocks()                   {}

var fakeDBstorage storage = fakeDB{}

func TestBlockchain(t *testing.T) {
	t.Run("Should create blockchain", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeLoadCahin: func() []byte {
				return nil
			},
		}
		bc := Blockchain()
		if bc.Height != 1 {
			t.Error("Blockchain() should create a blockchain")
		}
	})
	t.Run("Should restore blockchain", func(t *testing.T) {
		once = *new(sync.Once)
		dbStorage = fakeDB{
			fakeLoadCahin: func() []byte {
				bc := &blockchain{Height: 2, NewestHash: "xxx", CurrentDifficulty: 1}
				return utils.ToBytes(bc)
			},
		}
		bc := Blockchain()
		if bc.Height != 2 {
			t.Errorf("Blockchain() should restore a blockchain with a height of %d, got %d", 2, bc.Height)
		}
	})
}

func TestBlocks(t *testing.T) {
	fakeBlocks := 0
	dbStorage = fakeDB{
		fakeFindBlock: func() []byte {
			var b *Block
			if fakeBlocks == 0 {
				b = &Block{
					Height:   2,
					PrevHash: "x",
				}
			}
			if fakeBlocks == 1 {
				b = &Block{
					Height: 1,
				}
			}
			fakeBlocks++
			return utils.ToBytes(b)
		},
	}
	bc := &blockchain{}
	blocks := Blocks(bc)
	if reflect.TypeOf(blocks) != reflect.TypeOf([]*Block{}) {
		t.Error("Blocks() should return a slice of blocks")
	}
}

func TestFindTx(t *testing.T) {
	t.Run("Tx not found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height:       2,
					Transactions: []*Tx{},
				}

				return utils.ToBytes(b)
			},
		}
		tx := FindTx(&blockchain{NewestHash: "x"}, "test")
		if tx != nil {
			t.Error("Tx should be not found")
		}
	})
	t.Run("Tx should be found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height: 2,
					Transactions: []*Tx{
						{Id: "test"},
					},
				}

				return utils.ToBytes(b)
			},
		}
		tx := FindTx(&blockchain{NewestHash: "x"}, "test")
		if tx == nil {
			t.Error("Tx should be found")
		}
	})
}

func TestGetDifficulty(t *testing.T) {
	blocks := []Block{
		{PrevHash: "x"},
		{PrevHash: "x"},
		{PrevHash: "x"},
		{PrevHash: "x"},
		{PrevHash: ""},
	}
	fakeBlock := 0
	dbStorage = fakeDB{
		fakeFindBlock: func() []byte {
			defer func() {
				fakeBlock++
			}()
			return utils.ToBytes(blocks[fakeBlock])
		},
	}
	type test struct {
		height int
		want   int
	}
	tests := []test{
		{height: 0, want: defaultDefficulty},
		{height: 2, want: defaultDefficulty},
		{height: 5, want: 3},
	}
	for _, tc := range tests {
		bc := &blockchain{Height: tc.height, CurrentDifficulty: defaultDefficulty}
		got := getDifficulty(bc)
		if got != tc.want {
			t.Errorf("getDifficulty() should return %d got %d", tc.want, got)
		}
	}
}
