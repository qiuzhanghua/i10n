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

func ParseTagWithDefault(name string) (tag language.Tag) {
	tag, _ = ParseTag(name)
	return tag
}

func ParseTag(name string) (tag language.Tag, err error) {
	tag = language.English
	dotIndex := strings.LastIndex(name, ".")
	if dotIndex <= 0 {
		return tag, errors.New("name error, should be xxx_zh-CN.properties or XXX.en-US.yaml etc")
	}
	name = name[:dotIndex]
	dotIndex = strings.LastIndex(name, "_")
	if dotIndex <= 0 {
		dotIndex = strings.LastIndex(name, ".")
	}
	if dotIndex > 0 {
		tag, err = language.Parse(name[dotIndex+1:])
		return tag, err
	}
	return tag, nil
}

func AddTagMap(tag language.Tag, m map[string]string) {
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

//Nearest
// suppose locale is longest, languages should be less
func Nearest(locale string, languages []string) string {
	for i := range languages {
		if languages[i] == locale {
			return locale
		}
	}
	for i := range languages {
		if strings.HasPrefix(locale, languages[i]) {
			return languages[i]
		}
	}
	return "en"
}
