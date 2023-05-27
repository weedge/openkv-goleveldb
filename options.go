package goleveldb

import (
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type options struct {
	config LevelDBConfig

	// optional, more details see pebble Options
	// if use goleveldb options, config options can't use
	options *opt.Options
}

type Option interface {
	apply(*options)
}

type pebbleKVStoreOption struct {
	f func(*options)
}

func (fdo *pebbleKVStoreOption) apply(do *options) {
	fdo.f(do)
}

func newOption(f func(*options)) *pebbleKVStoreOption {
	return &pebbleKVStoreOption{
		f: f,
	}
}

func WithConfig(config LevelDBConfig) Option {
	return newOption(func(o *options) {
		o.config = config
	})
}

func WithLeveldbOptions(opts *opt.Options) Option {
	return newOption(func(o *options) {
		o.options = opts
	})
}

func getOptions(opts ...Option) *options {
	options := &options{
		config: *DefaultLevelDBConfig(),
	}

	for _, o := range opts {
		o.apply(options)
	}

	return options
}
