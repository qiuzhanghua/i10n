package i10n

import "golang.org/x/text/language"

var _data = map[language.Tag]map[string]string{language.English: {}}

var _tag = language.English

// T Get default tag resource/translated
func T(key string) string {
	return TT(key, _tag)
}

// SetDefaultLang set default lang
func SetDefaultLang(lang string) error {
	tag, err := language.Parse(lang)
	if err != nil {
		return err
	}
	_tag = tag
	return nil
}

func SetDefaultTag(tag language.Tag) {
	_tag = tag
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
