package gocc

import (
	"testing"
)

func TestConvert(t *testing.T) {
	cases := []string{
		`我们是工农子弟兵`,
		`从正数第x行到倒数第y行，截取多行输出文本的部分内容`,
		`2017年中国住房租赁市场租金规模约为1.3万亿元`,
		`香煙（英語：Cigarette），為煙草製品的一種。滑鼠是一種很常見及常用的電腦輸入設備。`,
		`香菸（英語：Cigarette），為菸草製品的一種。記憶體是一種很常見及常用的電腦輸入裝置。`,
		`乾隆爷是谁的干爷爷？乾爷爷吗？`,
	}

	for k := range conversions {
		s2t, err := New(k)
		if err != nil {
			t.Errorf("New %s error:%+v", k, err)
		}
		t.Logf("%+v", s2t.DictChains)

		for _, c := range cases {
			str, err := s2t.Convert(c)
			if err != nil {
				t.Error(err)
			}
			t.Logf("\n%s:original\n%s:%s", c, str, k)
		}
	}
}
