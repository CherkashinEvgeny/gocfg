package cfg

import "github.com/knadh/koanf/v2"

var inst = koanf.New(".")

func SetUp(i *koanf.Koanf) {
	inst = i
}
