package chewing

import (
    "os"
)

type Chewing struct {
}

func New(phrase_file string) (chewing *Chewing, err error) {
    chewing = new(Chewing)

    file, err := os.Open(phrase_file)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    return chewing, nil
}
