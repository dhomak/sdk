

package cache

const (
	KVStorePrefix                               = "store"
	LRUSize                                     = 1000000
	LevelSubscriptionNotificationsBufferMaxSize = 30000 // ~16Mb: elemenets := 16 * 1024 * 1024 / (64 + 512), where 512 - avg value size, 64 - avg key size
	PrometricsEnabled                           = false
)

type Config struct {
	id                                          string
	kvStorePrefix                               string
	lruSize                                     int
	levelSubscriptionNotificationsBufferMaxSize int
	prometricsEnabled                           bool
}

func NewCacheConfig(id string) *Config {
	return &Config{
		id:            id,
		kvStorePrefix: KVStorePrefix,
		lruSize:       LRUSize,
		levelSubscriptionNotificationsBufferMaxSize: LevelSubscriptionNotificationsBufferMaxSize,
		prometricsEnabled: PrometricsEnabled,
	}
}

func (ro *Config) SetKVStorePrefix(kvStorePrefix string) *Config {
	ro.kvStorePrefix = kvStorePrefix
	return ro
}

func (ro *Config) SetLRUSize(lruSize int) *Config {
	ro.lruSize = lruSize
	return ro
}

func (ro *Config) SetLevelSubscriptionNotificationsBufferMaxSize(levelSubscriptionNotificationsBufferMaxSize int) *Config {
	ro.levelSubscriptionNotificationsBufferMaxSize = levelSubscriptionNotificationsBufferMaxSize
	return ro
}

func (ro *Config) SetPrometricsEnabled(enabled bool) *Config {
	ro.prometricsEnabled = enabled
	return ro
}
