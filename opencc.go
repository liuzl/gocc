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

type Group struct {
	Files []string
	Dicts []*da.Dict
}

func (self *Group) String() string {
	return fmt.Sprintf("%+v", self.Files)
}

type OpenCC struct {
	Conversion  string
	Description string
	DictChains  []*Group
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
	for _, group := range self.DictChains {
		r := []rune(in)
		var tokens []string
		for i := 0; i < len(r); {
			s := r[i:]
			var token string
			max := 0
			for _, dict := range group.Dicts {
				ret, err := dict.PrefixMatch(string(s))
				if err != nil {
					return "", err
				}
				if len(ret) > 0 {
					o := ""
					for k, v := range ret {
						if len(k) > max {
							max = len(k)
							token = v[0]
							o = k
						}
					}
					i += len([]rune(o))
					break
				}
			}
			if max == 0 { //no match
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
	self.Description = name.(string)
	chain, has := config["conversion_chain"]
	if !has {
		return fmt.Errorf("conversion_chain not found in %s", configFile)
	}
	if dictChain, ok := chain.([]interface{}); ok {
		for _, v := range dictChain {
			if d, ok := v.(map[string]interface{}); ok {
				if gdict, has := d["dict"]; has {
					if dict, is := gdict.(map[string]interface{}); is {
						group, err := self.addDictChain(dict)
						if err != nil {
							return err
						}
						self.DictChains = append(self.DictChains, group)
					}
				} else {
					return fmt.Errorf("should have dict inside conversion_chain")
				}
			} else {
				return fmt.Errorf("should be map inside conversion_chain")
			}
		}
	} else {
		return fmt.Errorf("format %+v not correct in %s",
			reflect.TypeOf(dictChain), configFile)
	}
	return nil
}

func (self *OpenCC) addDictChain(d map[string]interface{}) (*Group, error) {
	t, has := d["type"]
	if !has {
		return nil, fmt.Errorf("type not found in %+v", d)
	}
	if typ, ok := t.(string); ok {
		ret := &Group{}
		switch typ {
		case "group":
			ds, has := d["dicts"]
			if !has {
				return nil, fmt.Errorf("no dicts field found")
			}
			dicts, is := ds.([]interface{})
			if !is {
				return nil, fmt.Errorf("dicts field invalid")
			}

			for _, dict := range dicts {
				if d, is := dict.(map[string]interface{}); is {
					group, err := self.addDictChain(d)
					if err != nil {
						return nil, err
					}
					ret.Files = append(ret.Files, group.Files...)
					ret.Dicts = append(ret.Dicts, group.Dicts...)
				} else {
					return nil, fmt.Errorf("dicts items invalid")
				}
			}
		case "txt":
			file, has := d["file"]
			if !has {
				return nil, fmt.Errorf("no file field found")
			}
			daDict, err := da.BuildFromFile(filepath.Join(DICT_DIR, file.(string)))
			if err != nil {
				return nil, err
			}
			ret.Files = append(ret.Files, file.(string))
			ret.Dicts = append(ret.Dicts, daDict)
		default:
			return nil, fmt.Errorf("type should be txt or group")
		}
		return ret, nil
	} else {
		return nil, fmt.Errorf("type should be string")
	}
}
