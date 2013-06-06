package chewing

import (
    "io/ioutil"
    "os"
    "path"
    "testing"
)

func TestNew(t *testing.T) {
    tmpFile, err := ioutil.TempFile("", "")
    if err != nil {
        t.Fatalf("Cannot create temp file: %s", err.Error())
    }
    tmpFileName := tmpFile.Name()
    defer os.Remove(tmpFileName)

    tmpFile.WriteString(
        "# This is comment\n" +
        "\x09\x0b\x20測試 5 ㄘㄜˋ ㄕˋ\n" +
        "側室 4 ㄘㄜˋ ㄕˋ\x09\x0b\x20\n" +
        "側視 3 ㄘㄜˋ ㄕ # This is commentˋ\n" +
        "策士 2 ㄘㄜˋ ㄕˋ\n" +
        "策試 1 ㄘㄜˋ ㄕˋ\n")
    tmpFile.Close()

    params := ChewingParameters {
        PhraseFile: tmpFileName,
    }

    chewing, err := New(&params)
    if chewing == nil {
        t.Error("Shall not return nil")
    }
    if err != nil {
        t.Errorf("Shall not return error %s", err.Error())
    }
}

func TestNewNoPhraseFile(t *testing.T) {
    params := ChewingParameters{
        PhraseFile: "NoSuchFile",
    }

    chewing, err := New(&params)
    if chewing != nil {
        t.Error("Shall return nil")
    }
    if err == nil {
        t.Error("Shall return error")
    }
}

func TestNewBadFrequency(t *testing.T) {
    tmpFile, err := ioutil.TempFile("", "")
    if err != nil {
        t.Fatalf("Cannot create temp file: %s", err.Error())
    }
    tmpFileName := tmpFile.Name()
    defer os.Remove(tmpFileName)

    tmpFile.WriteString("測試 a ㄘㄜˋ ㄕˋ\n")
    tmpFile.Close()

    params := ChewingParameters {
        PhraseFile: tmpFileName,
    }

    chewing, err := New(&params)
    if chewing != nil {
        t.Error("Shall return nil")
    }
    if err == nil {
        t.Error("Shall return error")
    }
}

func TestNewBadBopomofo(t *testing.T) {
    tmpFile, err := ioutil.TempFile("", "")
    if err != nil {
        t.Fatalf("Cannot create temp file: %s", err.Error())
    }
    tmpFileName := tmpFile.Name()
    defer os.Remove(tmpFileName)

    tmpFile.WriteString("測試 a ㄘㄜ1 ㄕˋ\n")
    tmpFile.Close()

    params := ChewingParameters {
        PhraseFile: tmpFileName,
    }

    chewing, err := New(&params)
    if chewing != nil {
        t.Error("Shall return nil")
    }
    if err == nil {
        t.Error("Shall return error")
    }
}

func BenchmarkNew(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _, err := New(&ChewingParameters{
            PhraseFile: path.Join(os.Getenv("GOPATH"), "data", "tsi.src"),
        })
        if err != nil {
            panic("New shall not return error")
        }
    }
}

func BenchmarkBKForestQuery(b *testing.B) {
    ctx, err := New(&ChewingParameters{
        PhraseFile: path.Join(os.Getenv("GOPATH"), "data", "tsi.src"),
    })
    if err != nil {
        panic("New shall not return error")
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ctx.phraseBKForest.query([]uint16{ 10268, 8708 }, 2)
    }
}
