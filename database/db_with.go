package database

type options func(*database)

var definedConfig = map[string]string{
	"port": "3306",
}

func WithPort(port string) func(*database) {
	return func(d *database) {
		d.port = port
	}
}

func WithPrefix(prefix string) func(*database) {
	return func(d *database) {
		d.prefix = prefix
	}
}

// func WithDeBug()
