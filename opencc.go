package opencc

import (
	"encoding/json"
	"fmt"
	"github.com/liuzl/da"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	CONFIG_DIR = "config"
	DICT_DIR   = "dictionary"
)

type OpenCC struct {
	Conversion     string
	ConversionName string
	DictFiles      []string
	Dicts          []*da.Dict
}

var conversions map[string]struct{} = map[string]struct{}{
	"hk2s": {}, "s2hk": {}, "s2t": {}, "s2tw": {}, "s2twp": {},
	"t2hk": {}, "t2s": {}, "t2tw": {}, "tw2s": {}, "tw2sp": {},
}

func New(conversion string) (*OpenCC, error) {
	if _, has := conversions[conversion]; !has {
		return nil, fmt.Errorf("%s not valid", conversion)
	}
	cc := &OpenCC{Conversion: conversion}
	err := cc.initDict()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

// Convert string from Simplified Chinese to Traditional Chinese or vice versa
func (self *OpenCC) Convert(in string) (string, error) {
	for _, dict := range self.Dicts {
		r := []rune(in)
		var tokens []string
		for i := 0; i < len(r); {
			s := r[i:]
			ret, err := dict.PrefixMatch(string(s))
			if err != nil {
				return "", err
			}
			var token string
			if len(ret) > 0 {
				max := 0
				o := ""
				for k, v := range ret {
					if len(k) > max {
						token = v[0]
						o = k
					}
				}
				i += len([]rune(o))
			} else {
				token = string(r[i])
				i++
			}
			tokens = append(tokens, token)
		}
		in = strings.Join(tokens, "")
	}
	return in, nil
}

func (self *OpenCC) initDict() error {
	if self.Conversion == "" {
		return fmt.Errorf("conversion is not set")
	}
	configFile := filepath.Join(CONFIG_DIR, self.Conversion+".json")
	body, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	var m interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	config := m.(map[string]interface{})
	name, has := config["name"]
	if !has {
		return fmt.Errorf("name not found in %s", configFile)
	}
	self.ConversionName = name.(string)
	chain, has := config["conversion_chain"]
	if !has {
		return fmt.Errorf("conversion_chain not found in %s", configFile)
	}
	switch dictChain := chain.(type) {
	case []interface{}:
		for _, v := range dictChain {
			if d, ok := v.(map[string]interface{}); ok {
				if gdict, has := d["dict"]; has {
					if dict, is := gdict.(map[string]interface{}); is {
						err = self.addDictChain(dict)
						if err != nil {
							return err
						}
					}
				} else {
					return fmt.Errorf("should have dict inside conversion_chain")
				}
			} else {
				return fmt.Errorf("should be map inside conversion_chain")
			}
		}
	default:
		return fmt.Errorf("format %+v not correct in %s",
			reflect.TypeOf(dictChain), configFile)
	}
	return self.addDicts()
}

func (self *OpenCC) addDictChain(d map[string]interface{}) error {
	t, has := d["type"]
	if !has {
		return fmt.Errorf("type not found in %+v", d)
	}
	switch typ := t.(type) {
	case string:
		switch typ {
		case "group":
			ds, has := d["dicts"]
			if !has {
				return fmt.Errorf("no dicts field found")
			}
			dicts, is := ds.([]interface{})
			if !is {
				return fmt.Errorf("dicts field invalid")
			}

			for _, dict := range dicts {
				if d, is := dict.(map[string]interface{}); is {
					err := self.addDictChain(d)
					if err != nil {
						return err
					}
				} else {
					return fmt.Errorf("dicts items invalid")
				}
			}
		case "txt":
			file, has := d["file"]
			if !has {
				return fmt.Errorf("no file field found")
			}
			dictFile := filepath.Join(DICT_DIR, file.(string))
			self.DictFiles = append(self.DictFiles, dictFile)
		default:
			return fmt.Errorf("type should be txt or group")
		}
	default:
		return fmt.Errorf("type should be string")
	}
	return nil
}

func (self *OpenCC) addDicts() error {
	for _, file := range self.DictFiles {
		d, err := da.BuildFromFile(file)
		if err != nil {
			return err
		}
		self.Dicts = append(self.Dicts, d)
	}
	return nil
}
