package goleveldb

const (
	DBName    = "goleveldb"
	MemDBName = "memory"
)

type StoreType int

const (
	StoreTypeUnkonw StoreType = iota
	StoreTypeDB
	StoreTypeMemory
)

const defaultFilterBits int = 10
