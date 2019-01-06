package looker

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kr/pretty"
)

func TestReflect(t *testing.T) {
	pkg, err := Reflect("github.com/go-gad/dekit/looker/testdata", []string{"CreateOrderReq"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Package %# v", pretty.Formatter(pkg))
}

func TestEncodeGob(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	filename := f.Name()
	//t.Log("filename ", filename)
	defer os.Remove(filename)
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	pkg, err := Reflect("github.com/go-gad/dekit/looker/testdata", []string{"CreateOrderReq"})
	if err != nil {
		t.Fatal(err)
	}

	if err := EncodeGob(filename, pkg); err != nil {
		t.Fatal(err)
	}

	fb, _ := ioutil.ReadFile(filename)
	t.Logf("File content:\n%s", string(fb))

	gb := bytes.NewBuffer(fb)
	var pkgD Package
	if err := gob.NewDecoder(gb).Decode(&pkgD); err != nil {
		t.Fatal(err)
	}
	t.Logf("Package %# v", pretty.Formatter(pkg))
}
