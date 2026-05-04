// Package config
package config

type config struct {
	AppConfig *appConfig
	DBConfig  *dbConfig
	JwtConfig *jwtConfig
}

type appConfig struct {
	Port      string
	RateLimit int
}

type dbConfig struct {
	Name string
	URI  string
}

type jwtConfig struct {
	Secret    string
	ExpiresIn string
}
