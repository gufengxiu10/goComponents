package redis

type hash struct{}

func init() {
	Hash = &hash{}
}

var Hash *hash

// 设置一个值
func (h *hash) Set(key, field string, value interface{}) error {
	return rdb.client.HSet(ctx, key, field, value).Err()
}

func (h *hash) MSet(key string, value ...interface{}) error {
	return rdb.client.HMSet(ctx, key, value...).Err()
}

func (h *hash) Exists(key string, field string) bool {
	bl, _ := rdb.client.HExists(ctx, key, field).Result()
	return bl
}

func (h *hash) Get(key, field string) (string, error) {
	return rdb.client.HGet(ctx, key, field).Result()
}

func (h *hash) GetAll(key string) map[string]string {
	data, _ := rdb.client.HGetAll(ctx, key).Result()
	return data
}
