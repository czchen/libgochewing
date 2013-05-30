package chewing

import (
    "os"
    "runtime"
    "path"
    "testing"
)

type TestLogger struct {
    t *testing.T
}

func (this *TestLogger) PrintRuntimeInformation() {
    _, file, line, ok := runtime.Caller(2)
    if ok {
        this.t.Logf("%s:%d", file, line)
    }
}

func (this *TestLogger) Printf(format string, v ...interface{}) {
    this.PrintRuntimeInformation();
    this.t.Logf(format, v)
}

func TestNew(t *testing.T) {
    logger := TestLogger{t: t}
    params := ChewingParameters{
        phraseFile: path.Join(os.Getenv("GOPATH"), "data", "tsi.src"),
        logger: &logger,
    }

    _, err := New(&params)
    if err != nil {
        t.Errorf("New shall success, but it fails. %s\n", err.Error())
    }
}

func Test_New_NoPhraseFile(t *testing.T) {
    logger := TestLogger{t: t}
    params := ChewingParameters{
        phraseFile: "NoSuchFile",
        logger: &logger,
    }

    _, err := New(&params)
    if err == nil {
        t.Errorf("New shall fail when there is no phrase file")
    }
}
