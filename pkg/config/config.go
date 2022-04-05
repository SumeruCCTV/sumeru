package config

// TODO: process from file.

type Config struct {
	Database *Database
	Web      *Web
	Camera   *Camera
}

type Database struct {
	DisableGormLogger bool
	PgDSN             string
	RedisDSN          string
	RedisPassword     string
}

type Web struct {
	Port int
}

type Camera struct {
}
