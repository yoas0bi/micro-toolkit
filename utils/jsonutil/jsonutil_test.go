package jsonutil_test

import (
	"bytes"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/jsonutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var testUser = user{"inhere", 200}

func TestPretty(t *testing.T) {
	tests := []any{
		map[string]int{"a": 1},
		struct {
			A int `json:"a"`
		}{1},
	}
	want := `{
    "a": 1
}`
	for _, sample := range tests {
		got, err := jsonutil.Pretty(sample)
		assert.NoErr(t, err)
		assert.Eq(t, want, got)
	}

	bts, err := jsonutil.EncodePretty(map[string]int{"a": 1})
	assert.NoErr(t, err)
	assert.Eq(t, want, string(bts))
}

func TestEncode(t *testing.T) {
	bts, err := jsonutil.Encode(testUser)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}`, string(bts))

	bts, err = jsonutil.Encode(&testUser)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}`, string(bts))

	bts, err = jsonutil.EncodeUnescapeHTML(&testUser)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}
`, string(bts))
}

func TestEncodeUnescapeHTML(t *testing.T) {
	bts, err := jsonutil.Encode(&testUser)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}`, string(bts))
}

func TestEncodeToWriter(t *testing.T) {
	buf := &bytes.Buffer{}

	err := jsonutil.EncodeToWriter(testUser, buf)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}
`, buf.String())
}

func TestDecode(t *testing.T) {
	str := `{"name":"inhere","age":200}`
	usr := &user{}
	err := jsonutil.Decode([]byte(str), usr)

	assert.NoErr(t, err)
	assert.Eq(t, "inhere", usr.Name)
	assert.Eq(t, 200, usr.Age)
}

func TestDecodeString(t *testing.T) {
	str := `{"name":"inhere","age":200}`
	usr := &user{}
	err := jsonutil.DecodeString(str, usr)

	assert.NoErr(t, err)
	assert.Eq(t, "inhere", usr.Name)
	assert.Eq(t, 200, usr.Age)
}

func TestWriteReadFile(t *testing.T) {
	user := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{"inhere", 200}

	err := jsonutil.WriteFile("testdata/test.json", &user)
	assert.NoErr(t, err)

	err = jsonutil.WritePretty("testdata/test2.json", &user)
	assert.NoErr(t, err)

	// err = jsonutil.WritePretty("/path/to/not-exist.json", &user)
	// assert.Err(t, err)

	err = jsonutil.ReadFile("testdata/test.json", &user)
	assert.NoErr(t, err)

	assert.Eq(t, "inhere", user.Name)
	assert.Eq(t, 200, user.Age)
}

func TestIsJsonFast(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"single character", "a", false},
		{"two characters object", "{}", true},
		{"two characters slice", "[]", true},
		{"invalid json", "{a}", false},
		{"valid json object", `{"a": 1}`, true},
		{"valid json array", `[1]`, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Eq(t, tc.expected, jsonutil.IsJSON(tc.input))

			if jsonutil.IsJSONFast(tc.input) != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, !tc.expected)
			}
		})
	}
}

func TestStripComments(t *testing.T) {
	is := assert.New(t)

	str := jsonutil.StripComments(`{"name":"app"}`)
	is.Eq(`{"name":"app"}`, str)

	givens := []string{
		// single line comments
		`{
"name":"app" // comments
}`,
		`{
// comments
"name":"app" 
}`,
		`{"name":"app"} // comments
`,
		// multi line comments
		`{"name":"app"} /* comments */
`,
		`/* comments */
{"name":"app"}`,
		`/* 
comments 
*/
{"name":"app"}`,
		`/** 
comments 
*/
{"name":"app"}`,
		`/** 
comments 
**/
{"name":"app"}`,
		`/** 
* comments 
**/
{"name":"app"}`,
		`/** 
/* comments 
**/
{"name":"app"}`,
		`/** 
/* comments *
**/
{"name":"app"}`,
		`{"name": /*comments*/"app"}`,
		`{/*comments*/"name": "app"}`,
	}
	wants := []string{
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		// multi line comments
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name": "app"}`,
		`{"name": "app"}`,
	}

	for i, s := range givens {
		is.Eq(wants[i], jsonutil.StripComments(s))
	}

	str = jsonutil.StripComments(`{"name":"app"} // comments`)
	is.Eq(`{"name":"app"}`, str)

	// fix https://github.com/gookit/config/issues/2
	str = jsonutil.StripComments(`{"name":"http://abc.com"} // comments`)
	is.Eq(`{"name":"http://abc.com"}`, str)

	str = jsonutil.StripComments(`{
"address": [
	"http://192.168.1.XXX:2379"
]
} // comments`)
	is.Eq(`{"address":["http://192.168.1.XXX:2379"]}`, str)

	s := `{"name":"http://abc.com"} // comments`
	s = jsonutil.StripComments(s)
	assert.Eq(t, `{"name":"http://abc.com"}`, s)

	s = `
{// comments
    "name": "app", // comments
/*comments*/
    "debug": false,
    "baseKey": "value", // comments
	/* comments */
    "age": 123,
    "envKey1": "${NotExist|defValue}",
    "map1": { // comments
        "key": "val",
        "key1": "val1",
        "key2": "val2"
    },
    "arr1": [ // comments
        "val",
        "val1", // comments
		/* comments */
        "val2",
		"http://a.com"
    ],
	/* 
		comments 
*/
    "lang": {
		/** 
 		 * comments 
 		 */
        "dir": "res/lang",
        "allowed": {
            "en": "val",
            "zh-CN": "val2"
        }
    }
}`
	ep := `{"name":"app","debug":false,"baseKey":"value","age":123,"envKey1":"${NotExist|defValue}","map1":{"key":"val","key1":"val1","key2":"val2"},"arr1":["val","val1","val2","http://a.com"],"lang":{"dir":"res/lang","allowed":{"en":"val","zh-CN":"val2"}}}`
	s = jsonutil.StripComments(s)
	assert.Eq(t, ep, s)
}
