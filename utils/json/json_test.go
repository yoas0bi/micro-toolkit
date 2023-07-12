package json

import (
	"github.com/yoas0bi/micro-toolkit/utils/helper"
	"testing"
)

func TestParseJson(t *testing.T) {
	_, err := helper.TJson.ParseJson(helper.jsonExample)
	if err != nil {
		t.Errorf("parse json errors of %v \n", err.Error())
	}
	_, err = helper.TJson.ParseJson(helper.exampleErrJson)
	if err == nil {
		t.Errorf("parse json errors")
	}
}

func TestMapToJson(t *testing.T) {
	jsonStr := helper.TJson.MapToJson(helper.exampleJsonMap)
	_, err := helper.TJson.ParseJson(jsonStr)
	if err != nil {
		t.Errorf("parse json errors of %v \n", err.Error())
	}
}

func TestJsonToMap(t *testing.T) {
	m := helper.TJson.JsonToMap(helper.jsonExample)
	if !helper.isMap(m) {
		t.Errorf("reflect valueof does not map \n")
	}
	if _, ok := m["k1"]; !ok {
		t.Errorf("map conv unit test fail \n")
	}
}

func TestJsonToMapArray(t *testing.T) {
	m := helper.TJson.JsonToMapArr(helper.exampleJsonArr)
	if !helper.isMap(m) || len(m) != 2 {
		t.Errorf("JsonToMapArr unit test fail \n")
	}
}

func TestStructToMap(t *testing.T) {
	example := helper.Example{
		Examples: "test",
	}
	m := helper.TJson.StructToMap(example)
	if !helper.isMap(m) {
		t.Errorf("StructToMap does not map")
	}

	if _, ok := m["Examples"]; !ok {
		t.Errorf("StructToMap key does not exists\n")
	}
}

func TestMapToStruct(t *testing.T) {

	var ex helper.Example
	ex1, err := helper.TJson.MapToStruct(helper.exampleJsonStruct, ex)
	if err != nil {
		t.Errorf("MapToStruct unit test fail \n")
	}
	if ex1.(helper.Example).Examples != "test" {
		t.Errorf("MapToStruct values of %v, not test \n", ex1.(helper.Example).Examples)
	}
}

func TestJsonEncode(t *testing.T) {
	bJson, err := helper.TJson.JsonEncode(helper.jsonExample)
	if err != nil {
		t.Errorf("JsonEncode errors: %v\n", err)
	}

	_, err = helper.TJson.ParseJson(string(bJson))
	if err != nil {
		t.Errorf("JsonEncode ParseJson errors of %v \n", err.Error())
	}
}

func TestJsonDecode(t *testing.T) {
	var i interface{}
	err := helper.TJson.JsonDecode([]byte(helper.jsonExample), &i)
	if err != nil {
		t.Errorf("JsonDecode unit test fail, errors: %v\n", err)
	}
}
