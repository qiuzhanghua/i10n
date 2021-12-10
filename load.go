package i10n

import (
	"errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
)

var (
	_data map[language.Tag]map[string]string
	_tag  language.Tag
	_p    *message.Printer
)

func init() {
	Reset()
}

// T Get default tag resource/translated
func T(key string, arg ...interface{}) string {
	return _p.Sprintf(TT(key, _tag), arg...)
}

// E Expend resource/translated of tag
func E(key string, tag language.Tag, arg ...interface{}) string {
	p := message.NewPrinter(tag)
	return p.Sprintf(TT(key, tag), arg...)
}

// SetDefaultLang set default lang
func SetDefaultLang(lang string) error {
	tag, err := language.Parse(lang)
	if err != nil {
		return err
	}
	SetDefaultTag(tag)
	return nil
}

func SetDefaultTag(tag language.Tag) {
	_tag = tag
	_p = message.NewPrinter(_tag)
}

func GetDefaultTag() language.Tag {
	return _tag
}

// AddResource add resource by tag and key
func AddResource(tag language.Tag, key string, value string) {
	if _, ok := _data[tag]; !ok {
		_data[tag] = map[string]string{}
	}
	_data[tag][key] = value
}

func GetExact(key string, tag language.Tag) string {
	if m, ok := _data[tag]; ok {
		if val, ok := m[key]; ok {
			return val
		}
	}
	return ""
}

// TT Get resource/translated according to key and tag
func TT(key string, tag language.Tag) string {
	for {
		if m, ok := _data[tag]; ok {
			if val, ok := m[key]; ok {
				return val
			}
		}
		tag = tag.Parent()
		if tag.IsRoot() {
			break
		}
	}
	if val, ok := _data[language.English][key]; ok {
		return val
	}
	return ""
}

/*
	Properties file handle
*/
func Parse(name string) (tag language.Tag, fileType string, err error) {
	dotIndex := strings.LastIndex(name, ".")
	if dotIndex <= 0 {
		return tag, fileType, errors.New("name error, should be xxx_zh-CN.properties or XXX.en-US.yaml etc")
	}
	fileType = name[dotIndex+1:]
	name = name[:dotIndex]
	dotIndex = strings.LastIndex(name, "_")
	if dotIndex <= 0 {
		dotIndex = strings.LastIndex(name, ".")
	}
	if dotIndex <= 0 {
		tag = language.English
	} else {
		tag, err = language.Parse(name[dotIndex+1:])
		return tag, fileType, err
	}
	return tag, fileType, nil
}

func AddMap(tag language.Tag, m map[string]string) {
	if _, ok := _data[tag]; !ok {
		_data[tag] = map[string]string{}
	}
	for key, val := range m {
		_data[tag][key] = val
	}
}

func Reset() {
	_data = map[language.Tag]map[string]string{language.English: {}}
	_tag = language.English
	_p = message.NewPrinter(_tag)
}
