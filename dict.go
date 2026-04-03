package sprig

import (
	"fmt"
	"reflect"

	"dario.cat/mergo"
	"github.com/mitchellh/copystructure"
)

func get(d map[string]interface{}, key string) interface{} {
	if val, ok := d[key]; ok {
		return val
	}
	return ""
}

func set(d map[string]interface{}, key string, value interface{}) map[string]interface{} {
	d[key] = value
	return d
}

func unset(d map[string]interface{}, key string) map[string]interface{} {
	delete(d, key)
	return d
}

func hasKey(d interface{}, key string) bool {
	v := reflect.ValueOf(d)
	if v.Kind() == reflect.Map {
		if v.Type().Key().Kind() != reflect.String {
			return false
		}
		return v.MapIndex(reflect.ValueOf(key)).IsValid()
	}
	if m, ok := d.(map[string]interface{}); ok {
		_, found := m[key]
		return found
	}
	return false
}

func pluck(key string, d ...map[string]interface{}) []interface{} {
	res := []interface{}{}
	for _, dict := range d {
		if val, ok := dict[key]; ok {
			res = append(res, val)
		}
	}
	return res
}

func keys(dicts ...interface{}) []string {
	k := []string{}
	for _, dictIface := range dicts {
		v := reflect.ValueOf(dictIface)
		if v.Kind() == reflect.Map {
			if v.Type().Key().Kind() != reflect.String {
				panic(fmt.Sprintf("keys: map key type must be string, got %s", v.Type().Key().Kind()))
			}
			for _, key := range v.MapKeys() {
				k = append(k, key.String())
			}
		}
	}
	return k
}

func pick(dictIface interface{}, keys ...string) map[string]interface{} {
	res := map[string]interface{}{}
	v := reflect.ValueOf(dictIface)
	if v.Kind() == reflect.Map {
		if v.Type().Key().Kind() != reflect.String {
			panic(fmt.Sprintf("pick: map key type must be string, got %s", v.Type().Key().Kind()))
		}
		for _, k := range keys {
			val := v.MapIndex(reflect.ValueOf(k))
			if val.IsValid() {
				res[k] = val.Interface()
			}
		}
	}
	return res
}

func omit(dictIface interface{}, keys ...string) map[string]interface{} {
	res := map[string]interface{}{}

	omitSet := make(map[string]bool, len(keys))
	for _, k := range keys {
		omitSet[k] = true
	}

	v := reflect.ValueOf(dictIface)
	if v.Kind() == reflect.Map {
		if v.Type().Key().Kind() != reflect.String {
			panic(fmt.Sprintf("omit: map key type must be string, got %s", v.Type().Key().Kind()))
		}
		for _, key := range v.MapKeys() {
			k := key.String()
			if !omitSet[k] {
				res[k] = v.MapIndex(key).Interface()
			}
		}
	}
	return res
}

func dict(v ...interface{}) map[string]interface{} {
	dict := map[string]interface{}{}
	lenv := len(v)
	for i := 0; i < lenv; i += 2 {
		key := strval(v[i])
		if i+1 >= lenv {
			dict[key] = ""
			continue
		}
		dict[key] = v[i+1]
	}
	return dict
}

func merge(dst map[string]interface{}, srcs ...map[string]interface{}) interface{} {
	for _, src := range srcs {
		if err := mergo.Merge(&dst, src); err != nil {
			// Swallow errors inside of a template.
			return ""
		}
	}
	return dst
}

func mustMerge(dst map[string]interface{}, srcs ...map[string]interface{}) (interface{}, error) {
	for _, src := range srcs {
		if err := mergo.Merge(&dst, src); err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func mergeOverwrite(dst map[string]interface{}, srcs ...map[string]interface{}) interface{} {
	for _, src := range srcs {
		if err := mergo.MergeWithOverwrite(&dst, src); err != nil {
			// Swallow errors inside of a template.
			return ""
		}
	}
	return dst
}

func mustMergeOverwrite(dst map[string]interface{}, srcs ...map[string]interface{}) (interface{}, error) {
	for _, src := range srcs {
		if err := mergo.MergeWithOverwrite(&dst, src); err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func values(dictIface interface{}) []interface{} {
	vals := []interface{}{}
	v := reflect.ValueOf(dictIface)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			vals = append(vals, v.MapIndex(key).Interface())
		}
	}
	return vals
}

func deepCopy(i interface{}) interface{} {
	c, err := mustDeepCopy(i)
	if err != nil {
		panic("deepCopy error: " + err.Error())
	}

	return c
}

func mustDeepCopy(i interface{}) (interface{}, error) {
	return copystructure.Copy(i)
}

func dig(ps ...interface{}) (interface{}, error) {
	if len(ps) < 3 {
		panic("dig needs at least three arguments")
	}
	dict := ps[len(ps)-1].(map[string]interface{})
	def := ps[len(ps)-2]
	ks := make([]string, len(ps)-2)
	for i := 0; i < len(ks); i++ {
		ks[i] = ps[i].(string)
	}

	return digFromDict(dict, def, ks)
}

func digFromDict(dict map[string]interface{}, d interface{}, ks []string) (interface{}, error) {
	k, ns := ks[0], ks[1:len(ks)]
	step, has := dict[k]
	if !has {
		return d, nil
	}
	if len(ns) == 0 {
		return step, nil
	}
	return digFromDict(step.(map[string]interface{}), d, ns)
}
