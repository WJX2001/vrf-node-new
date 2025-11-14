package common

import (
	"math/big"

	"github.com/WJX2001/vrf-node-new/database/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	_ "github.com/WJX2001/vrf-node-new/database/utils/serializers"
)

type BlockHeader1 struct {
	GUID       uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	Hash       common.Hash `gorm:"serializer:bytes"`
	ParentHash common.Hash `gorm:"serializer:bytes"`
	Number     *big.Int    `gorm:"serializer:u256"`
	Timestamp  uint64
	RLPHeader  *utils.RLPHeader `gorm:"serializer:rlp;column:rlp_bytes"`
}

func (BlockHeader) TableName1() string {
	return "block_headers"
}
