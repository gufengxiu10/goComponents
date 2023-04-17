package redis

type strings struct {
}

func init() {
	String = &strings{}
}

var String *strings

func (s *strings) Set(key string, value any) error {
	return rdb.client.Set(ctx, key, value, 0).Err()
}

func (s *strings) Get(key string) (string, error) {
	return rdb.client.Get(ctx, key).Result()
}
