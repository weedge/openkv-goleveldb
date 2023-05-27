package goleveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"

	driver "github.com/weedge/pkg/driver/openkv"
)

type DB struct {
	cfg  *LevelDBConfig
	path string
	db   *leveldb.DB

	opts         *opt.Options
	iteratorOpts *opt.ReadOptions
	syncOpts     *opt.WriteOptions
}

func (db *DB) initOpts() {
	db.opts = nOptions(db.cfg)

	db.iteratorOpts = &opt.ReadOptions{}
	db.iteratorOpts.DontFillCache = true

	db.syncOpts = &opt.WriteOptions{}
	db.syncOpts.Sync = true
}

func nOptions(cfg *LevelDBConfig) *opt.Options {
	opts := &opt.Options{}
	opts.ErrorIfMissing = false
	opts.BlockCacheCapacity = cfg.CacheSize
	//we must use bloomfilter
	opts.Filter = filter.NewBloomFilter(defaultFilterBits)

	if !cfg.Compression {
		opts.Compression = opt.NoCompression
	} else {
		opts.Compression = opt.SnappyCompression
	}

	opts.BlockSize = cfg.BlockSize
	opts.WriteBuffer = cfg.WriteBufferSize
	opts.OpenFilesCacheCapacity = cfg.MaxOpenFiles

	opts.CompactionTableSize = cfg.CompactionTableSize
	opts.WriteL0SlowdownTrigger = cfg.WriteL0SlowdownTrigger
	opts.WriteL0PauseTrigger = cfg.WriteL0PauseTrigger

	return opts
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Put(key, value []byte) error {
	return db.db.Put(key, value, nil)
}

func (db *DB) Get(key []byte) ([]byte, error) {
	v, err := db.db.Get(key, nil)
	if err == leveldb.ErrNotFound {
		return nil, nil
	}
	return v, nil
}

func (db *DB) Delete(key []byte) error {
	return db.db.Delete(key, nil)
}

func (db *DB) SyncPut(key []byte, value []byte) error {
	return db.db.Put(key, value, db.syncOpts)
}

func (db *DB) SyncDelete(key []byte) error {
	return db.db.Delete(key, db.syncOpts)
}

func (db *DB) NewWriteBatch() driver.IWriteBatch {
	wb := &WriteBatch{
		db:     db,
		wbatch: new(leveldb.Batch),
	}
	return wb
}

func (db *DB) NewIterator() driver.IIterator {
	it := &Iterator{
		db.db.NewIterator(nil, db.iteratorOpts),
	}

	return it
}

func (db *DB) NewSnapshot() (driver.ISnapshot, error) {
	snapshot, err := db.db.GetSnapshot()
	if err != nil {
		return nil, err
	}

	s := &Snapshot{
		db:  db,
		snp: snapshot,
	}

	return s, nil
}

func (db *DB) Compact() error {
	return db.db.CompactRange(util.Range{
		Start: nil,
		Limit: nil,
	})
}
