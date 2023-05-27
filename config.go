package goleveldb

type LevelDBConfig struct {
	Compression            bool
	BlockSize              int
	WriteBufferSize        int
	CacheSize              int
	MaxOpenFiles           int
	CompactionTableSize    int
	WriteL0SlowdownTrigger int
	WriteL0PauseTrigger    int
}

func DefaultLevelDBConfig() *LevelDBConfig {
	return &LevelDBConfig{
		Compression:            false,
		BlockSize:              32_768,
		WriteBufferSize:        67_108_864,
		CacheSize:              524_288_000,
		MaxOpenFiles:           1024,
		CompactionTableSize:    32 * 1024 * 1024,
		WriteL0SlowdownTrigger: 16,
		WriteL0PauseTrigger:    64,
	}
}
