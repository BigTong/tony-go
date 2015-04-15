// Copyright (c) 2015, The Tony Authors.
// All rights reserved.
//
// Author: Rentong Zhang <rentongzh@gmail.com>

package base

import (
	"code.google.com/p/mahonia"
	"github.com/saintfish/chardet"
	"strings"
	"unicode/utf8"
)

type Utf8Converter struct {
	detector *chardet.Detector
	gbk      mahonia.Decoder
	big5     mahonia.Decoder
}

func NewUtf8Converter() *Utf8Converter {
	ret := Utf8Converter{}
	ret.detector = chardet.NewHtmlDetector()
	ret.gbk = mahonia.NewDecoder("gb18030")
	ret.big5 = mahonia.NewDecoder("big5")
	return &ret
}

func (self *Utf8Converter) DetectCharset(src []byte) string {
	rets, err := self.detector.DetectAll(src)
	if err != nil {
		return ""
	}
	maxret := ""
	w := 0
	for _, ret := range rets {
		cs := strings.ToLower(ret.Charset)
		if strings.HasPrefix(cs, "gb") || strings.HasPrefix(cs, "utf") {
			if w < ret.Confidence {
				w = ret.Confidence
				maxret = cs
			}
		} else {
			continue
		}
	}
	return maxret
}

func (self *Utf8Converter) ToUTF8(src []byte) []byte {
	charset := self.DetectCharset(src)

	if !strings.Contains(charset, "gb") && !strings.Contains(charset, "big") {
		charset = "utf-8"
	}
	switch charset {
	case "utf-8", "utf8":
		return src
	case "gb2312", "gb-2312", "gbk", "gb18030", "gb-18030":
		ret, ok := self.gbk.ConvertStringOK(string(html))
		if ok {
			return []byte(ret)
		}
	case "big5":
		ret, ok := self.big5.ConvertStringOK(string(html))
		if ok {
			return []byte(ret)
		}
	}
	return nil
}
