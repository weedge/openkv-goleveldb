package goleveldb

type LevelDBConfig struct {
	Compression            bool `mapstructure:"compression"`
	BlockSize              int  `mapstructure:"blockSize"`
	WriteBufferSize        int  `mapstructure:"writeBufferSize"`
	CacheSize              int  `mapstructure:"cacheSize"`
	MaxOpenFiles           int  `mapstructure:"maxOpenFiles"`
	CompactionTableSize    int  `mapstructure:"compactionTableSize"`
	WriteL0SlowdownTrigger int  `mapstructure:"writeL0SlowdownTrigger"`
	WriteL0PauseTrigger    int  `mapstructure:"writeL0PauseTrigger"`
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
