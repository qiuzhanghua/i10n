package i10n

import (
	"golang.org/x/text/language"
	"testing"
)

func TestGetDefaultTag(t *testing.T) {
	expected := "en"
	actual := GetDefaultTag().String()
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestSetDefaultLang(t *testing.T) {
	_ = SetDefaultLang("zh_CN")
	expected := "zh"
	actual := GetDefaultTag().Parent().String()
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestAddResource(t *testing.T) {
	AddResource(language.English, "Hello", "world")
	expected := "world"
	actual := GetExact("Hello", language.English)
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestTT(t *testing.T) {
	SetDefaultTag(language.SimplifiedChinese)
	zh, _ := language.Parse("zh")
	AddResource(language.English, "Hello", "world")
	AddResource(zh, "Hello", "世界")
	expected := "世界"
	actual := TT("Hello", language.SimplifiedChinese)
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
	expected = "world"
	actual = TT("Hello", language.English)
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}

	jp, _ := language.Parse("jp")
	expected = "world"
	actual = TT("Hello", jp)
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
}

func TestT(t *testing.T) {
	_ = SetDefaultLang("zh-CN")
	zh, _ := language.Parse("zh")
	AddResource(language.English, "Hello", "world")
	AddResource(zh, "Hello", "世界")
	expected := "世界"
	actual := T("Hello")
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}
	expected = ""
	actual = T("Hello2")
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got: '%v'", expected, actual)
	}

}
