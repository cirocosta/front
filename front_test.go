package front

import (
	"os"
	"testing"
)

type frontMatter struct {
	Title string `json:"title" yaml:"title"`
}

var expectedBody = `# Body
Over my dead body`

func TestMatterJson(t *testing.T) {
	m := NewMatter()
	m.Handle("+++", JSONHandler)

	file, err := os.Open("testdata/front/json.md")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	front := &frontMatter{}

	body, err := m.Parse(file, front)
	if err != nil {
		t.Error(err)
	}

	if body != expectedBody {
		t.Errorf("expected '%s' got '%s'", expectedBody, body)
	}

	if front.Title != "bongo" {
		t.Error("expected front matter to contain title")
	}
}


func TestMatterYaml(t *testing.T) {
	m := NewMatter()
	m.Handle("---", YAMLHandler)

	file, err := os.Open("testdata/front/yaml.md")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	front := &frontMatter{}

	body, err := m.Parse(file, front)
	if err != nil {
		t.Errorf("failed to parse - %+v", err)
	}

	if body != expectedBody {
		t.Errorf("expected '%s' got '%s'", expectedBody, body)
	}

	if front.Title != "bongo" {
		t.Error("expected front matter to contain title")
	}
}


