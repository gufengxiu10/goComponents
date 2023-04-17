package redis

type key struct{}

func init() {
	Key = &key{}
}

var Key *key

func (k *key) Exists(key string) bool {
	count, _ := rdb.client.Exists(ctx, key).Result()
	return count > 0
}
