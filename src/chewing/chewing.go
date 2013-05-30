package chewing

import (
    "bufio"
    "os"
)

type Chewing struct {
    logger ChewingLogger
}


type ChewingParameters struct {
    phraseFile string
    logger ChewingLogger
}

func New(params *ChewingParameters) (chewing *Chewing, err error) {
    chewing = new(Chewing)
    if params.logger == nil {
        params.logger = new(ChewingDefaultLogger)
    }
    chewing.logger = params.logger

    file, err := os.Open(params.phraseFile)
    if err != nil {
        chewing.logger.Printf("%s", err.Error())
        return nil, err
    }
    defer file.Close()

    for scanner := bufio.NewScanner(file); scanner.Scan(); {
        text := scanner.Text()
        chewing.logger.Printf("%s", text)
    }

    return chewing, nil
}
