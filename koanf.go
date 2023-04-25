package cfg

import (
	"github.com/knadh/koanf/v2"
)

func Load(p koanf.Provider, pa koanf.Parser, opts ...koanf.Option) error {
	return inst.Load(p, pa, opts...)
}

func Keys() []string {
	return inst.Keys()
}

func KKeyMap() koanf.KeyMap {
	return inst.KeyMap()
}

func All() map[string]interface{} {
	return inst.All()
}

func Raw() map[string]interface{} {
	return inst.All()
}

func Sprint() string {
	return inst.Sprint()
}

func Print() {
	inst.Print()
}

func Cut(path string) *koanf.Koanf {
	return inst.Cut(path)
}

func Copy() *koanf.Koanf {
	return inst.Copy()
}

func Merge(in *koanf.Koanf) error {
	return inst.Merge(in)
}

func MergeAt(in *koanf.Koanf, path string) error {
	return inst.MergeAt(in, path)
}

func Set(key string, val interface{}) error {
	return inst.Set(key, val)
}

func Marshal(p koanf.Parser) ([]byte, error) {
	return inst.Marshal(p)
}

func Unmarshal(path string, o interface{}) error {
	return inst.Unmarshal(path, o)
}

func UnmarshalWithConf(path string, o interface{}, c koanf.UnmarshalConf) error {
	return inst.UnmarshalWithConf(path, o, c)
}

func Delete(path string) {
	inst.Delete(path)
}

func Get(path string) interface{} {
	return inst.Get(path)
}

func Slices(path string) []*koanf.Koanf {
	return inst.Slices(path)
}

func Exists(path string) bool {
	return inst.Exists(path)
}

func MapKeys(path string) []string {
	return inst.MapKeys(path)
}

func Delim() string {
	return inst.Delim()
}
