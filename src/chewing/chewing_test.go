package chewing

import (
    "os"
    "path"
    "testing"
)

func TestNew(t *testing.T) {
    phrase_file := path.Join(os.Getenv("GOPATH"), "data", "tsi.src")

    _, err := New(phrase_file)
    if err != nil {
        t.Fatalf("New shall success, but it fails. %s\n", err.Error())
    }
}

func Test_New_NoPhraseFile(t *testing.T) {
    _, err := New("NoSuchFile")
    if err == nil {
        t.Fatal("New shall fail when there is no phrase file")
    }
}
