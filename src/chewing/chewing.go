package chewing

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

type Chewing struct {
    dict *PhraseDictionary
    phraseArray *PhraseArray
    logger ChewingLogger
}


type ChewingParameters struct {
    phraseFile string
    logger ChewingLogger
}

func New(params *ChewingParameters) (chewing *Chewing, err error) {
    chewing = new(Chewing)

    chewing.setupLogger(params)
    err = chewing.setupPhraseArray(params)
    if err != nil {
        return nil, err
    }

    err = chewing.setupDictionary(params)
    if err != nil {
        return nil, err
    }

    return chewing, nil
}

func (this *Chewing) setupLogger(params *ChewingParameters) {
    if params.logger == nil {
        params.logger = new(ChewingDefaultLogger)
    }
    this.logger = params.logger
}

func (this *Chewing) setupPhraseArray(params *ChewingParameters) (err error) {
    file, err := os.Open(params.phraseFile)
    if err != nil {
        return err
    }
    defer file.Close()

    this.phraseArray = newPhraseArray()

    for scanner := bufio.NewScanner(file); scanner.Scan(); {
        text := scanner.Text()

        comment := strings.Index(text, "#")
        if comment != -1 {
            text = text[:comment]
        }

        text = strings.TrimSpace(text)
        if text == "" {
            continue
        }

        token := strings.Split(text, " ")
        if len(token) < 3 {
            return errors.New(fmt.Sprintf("`%s' is invalid in phraseFile", text))
        }

        var frequency uint32
        count, _:= fmt.Sscanf(token[1], "%d", &frequency)
        if count != 1 {
            return errors.New(fmt.Sprintf("`%s' is not a valid frequency", token[1]))
        }

        bopomofoSeq := token[2:]
        phoneSeq := make([]uint16, len(bopomofoSeq))

        for i := 0; i < len(bopomofoSeq); i++ {
            phoneSeq[i], err = convertBopomofoToPhone(bopomofoSeq[i])
            if err != nil {
                return err
            }
        }

        phrase, err := newPhrase(token[0], frequency)
        if err != nil {
            return err
        }

        this.phraseArray.insert(phrase, phoneSeq)
    }

    return nil
}


func (this *Chewing) setupDictionary(params *ChewingParameters) (err error) {
    file, err := os.Open(params.phraseFile)
    if err != nil {
        return err
    }
    defer file.Close()

    this.dict = newPhraseDictionary()

    for scanner := bufio.NewScanner(file); scanner.Scan(); {
        text := scanner.Text()

        comment := strings.Index(text, "#")
        if comment != -1 {
            text = text[:comment]
        }

        text = strings.TrimSpace(text)
        if text == "" {
            continue
        }

        token := strings.Split(text, " ")
        if len(token) < 3 {
            return errors.New(fmt.Sprintf("`%s' is invalid in phraseFile", text))
        }

        var frequency uint32
        count, _:= fmt.Sscanf(token[1], "%d", &frequency)
        if count != 1 {
            return errors.New(fmt.Sprintf("`%s' is not a valid frequency", token[1]))
        }

        bopomofoSeq := token[2:]
        phoneSeq := make([]uint16, len(bopomofoSeq))

        for i := 0; i < len(bopomofoSeq); i++ {
            phoneSeq[i], err = convertBopomofoToPhone(bopomofoSeq[i])
            if err != nil {
                return err
            }
        }

        phrase, err := newPhrase(token[0], frequency)
        if err != nil {
            return err
        }

        this.dict.insertPhrase(phrase, phoneSeq)
    }

    return nil
}
