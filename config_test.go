package opencc

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestConfig(t *testing.T) {
	cases := []string{
		"config/s2t.json",
		"config/t2s.json",
	}
	for _, c := range cases {
		body, err := ioutil.ReadFile(c)
		if err != nil {
			t.Error(err)
		}
		var config interface{}
		err = json.Unmarshal(body, &config)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", config)
		m := config.(map[string]interface{})
		for k, v := range m {
			t.Logf("%+v", k)
			switch vv := v.(type) {
			case map[string]interface{}:
				for i, u := range vv {
					t.Logf("%+v, %+v", i, u)
				}
			default:
				t.Logf("%+v", vv)
			}
		}
	}
}
