package service

import (
	"math/big"
)

// 分页
func Paging(page, limit int64) int64 {
	if page < 1 {
		page = 1
	}
	bigInt := big.Int{}
	bigIntPage := bigInt.Sub(big.NewInt(page), big.NewInt(1))
	offset := bigInt.Mul(bigIntPage, big.NewInt(limit))
	return  offset.Int64()
}