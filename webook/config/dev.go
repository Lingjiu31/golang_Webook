//go:build !k8s

package config

var Config = config{
	// 本地
	DB: DBConfig{
		DSN: "root:root@tcp(localhost:13316)/webook",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}
