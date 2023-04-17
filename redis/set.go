package redis

type set struct {
}

func init() {
	Set = &set{}
}

var Set *set

func (s *set) Set(key string, value ...any) error {
	return rdb.client.SAdd(ctx, key, value...).Err()
}
