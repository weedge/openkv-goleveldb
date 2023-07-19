package goleveldb

import (
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
	driver "github.com/weedge/pkg/driver/openkv"
)

type Store struct {
	opts      *options
	storeType StoreType
}

// New new db/memory store
func New(storeType StoreType, opts ...Option) (s *Store) {
	s = &Store{storeType: storeType}
	s.opts = getOptions(opts...)

	return
}

func (s *Store) Name() string {
	switch s.storeType {
	case StoreTypeDB:
		return DBName + "." + s.opts.config.Tag
	case StoreTypeMemory:
		return MemDBName + "." + s.opts.config.Tag
	default:
		return DBName + "." + s.opts.config.Tag
	}
}

// Open open db/memory leveldb
func (s *Store) Open(path string) (driver.IDB, error) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, err
	}

	db := &DB{path: path, cfg: &s.opts.config}
	db.initOpts()

	var ldb *leveldb.DB
	if s.storeType == StoreTypeMemory {
		ldb, err = leveldb.Open(storage.NewMemStorage(), db.opts)
	} else {
		ldb, err = leveldb.OpenFile(db.path, db.opts)
	}
	if err != nil {
		return nil, err
	}
	db.db = ldb

	return db, nil
}

func (s *Store) Repair(path string) error {
	if s.storeType == StoreTypeMemory {
		return nil
	}

	db, err := leveldb.RecoverFile(path, nOptions(DefaultLevelDBConfig()))
	if err != nil {
		return err
	}

	db.Close()
	return nil
}
