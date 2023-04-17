package redis

type options func(*redisClient)

func WithPassword(password string) func(*redisClient) {
	return func(rc *redisClient) {
		rc.password = password
	}
}

func WithPort(port string) func(*redisClient) {
	return func(rc *redisClient) {
		rc.port = port
	}
}
