package conf

import "flag"

var (
	Env string
)

func init() {
	var env string
	flag.StringVar(&env, "env", "prod", "set env")
	flag.Parse()

	Env = env
	return
}
