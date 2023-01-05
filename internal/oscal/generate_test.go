package oscal

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var examplesDir string = "../../examples"
var pkgName string = "oscal"

// TestSimpleJson tests that a simple JSON string with a single key and a single (string) value returns no error
// It does not (yet) test for correctness of the end result
func TestSimpleJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar"}`)
	if _, err := Generate(i, ParseJson, "TestStruct", pkgName, []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestNullableJson tests that a null JSON value is handled properly
func TestNullableJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar", "baz" : null}`)
	if _, err := Generate(i, ParseJson, "TestStruct", pkgName, []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestSimpleArray tests that an array without conflicting types is handled correctly
func TestSimpleArray(t *testing.T) {
	i := strings.NewReader(`{"foo" : [{"bar": 24}, {"bar" : 42}]}`)
	if _, err := Generate(i, ParseJson, "TestStruct", pkgName, []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestInvalidFieldChars tests that a document with invalid field chars is handled correctly
func TestInvalidFieldChars(t *testing.T) {
	i := strings.NewReader(`{"f.o-o" : 42}`)
	if _, err := Generate(i, ParseJson, "TestStruct", pkgName, []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestInferFloatInt tests that we can correctly infer a float or an int from a
// JSON number when no command-line flag is provided.
func TestInferFloatInt(t *testing.T) {
	f, err := os.Open(filepath.Join(examplesDir, "floats.json"))
	if err != nil {
		t.Fatalf("error opening"+examplesDir+"/floats.json: %s", err)
	}
	defer f.Close()

	expected, err := os.ReadFile(filepath.Join(examplesDir, "expected_floats.go.out"))
	if err != nil {
		t.Fatalf("error reading expected_floats.go.out: %s", err)
	}

	actual, err := Generate(f, ParseJson, "Stats", pkgName, []string{"json"}, false, true)
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}

}

// TestYamlNumbers tests that we handle Yaml's number system that has both floats and ints correctly
func TestYamlNumbers(t *testing.T) {
	f, err := os.Open(filepath.Join(examplesDir, "numbers.yaml"))
	if err != nil {
		t.Fatalf("error opening"+examplesDir+"/numbers.yaml: %s", err)
	}
	defer f.Close()

	expected, err := os.ReadFile(filepath.Join(examplesDir, "expected_numbers.go.out"))
	if err != nil {
		t.Fatalf("error reading expected_numbers.go.out: %s", err)
	}

	actual, err := Generate(f, ParseYaml, "Stats", pkgName, []string{"yaml"}, false, false)
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}

// Test example document
func TestExample(t *testing.T) {
	i, err := os.Open(filepath.Join(examplesDir, "example.json"))
	if err != nil {
		t.Error("error opening example.json", err)
	}

	expected, err := os.ReadFile(filepath.Join(examplesDir, "expected_output_test.go.out"))
	if err != nil {
		t.Error("error reading expected_output_test.go", err)
	}

	actual, err := Generate(i, ParseJson, "User", pkgName, []string{"json"}, false, true)
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}

func TestFmtFieldName(t *testing.T) {
	type TestCase struct {
		in  string
		out string
	}

	testCases := []TestCase{
		{in: "foo_id", out: "FooID"},
		{in: "fooId", out: "FooID"},
		{in: "foo_url", out: "FooURL"},
		{in: "foobar", out: "Foobar"},
		{in: "url_sample", out: "URLSample"},
		{in: "_id", out: "ID"},
		{in: "__id", out: "ID"},
	}

	for _, testCase := range testCases {
		lintField := FmtFieldName(testCase.in)
		if lintField != testCase.out {
			t.Errorf("error fmtFiledName %s != %s (%s)", testCase.in, testCase.out, lintField)
		}
	}
}

// Test example document
func TestExampleArray(t *testing.T) {
	i, err := os.Open(filepath.Join(examplesDir, "example_array.json"))
	if err != nil {
		t.Fatalf("error opening example.json: %s", err)
	}
	defer i.Close()

	expectedf, err := os.Open(filepath.Join(examplesDir, "example_array.go.out"))
	if err != nil {
		t.Fatalf("error opening example_array.go: %s", err)
	}
	defer expectedf.Close()

	expectedBts, err := io.ReadAll(expectedf)
	if err != nil {
		t.Fatalf("error reading example_array.go: %s", err)
	}

	actual, err := Generate(i, ParseJson, "Users", pkgName, []string{"json"}, false, true)
	if err != nil {
		t.Fatal(err)
	}
	sactual, sexpected := string(actual), string(expectedBts)
	if sactual != sexpected {
		t.Fatalf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}
