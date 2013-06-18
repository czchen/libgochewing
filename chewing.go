package libgochewing

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Chewing struct {
	phraseArray    *PhraseArray
	phraseTree     *PhraseTree
	phraseBKForest *PhraseBKForest
	logger         ChewingLogger
	keyboardType   int
}

type ChewingLogger interface {
	Printf(format string, v ...interface{})
}

type ChewingParameters struct {
	PhraseFile string
	Logger     ChewingLogger
}

type ChewingDefaultLogger struct{}

func New(params *ChewingParameters) (chewing *Chewing, err error) {
	chewing = new(Chewing)

	chewing.setupLogger(params)
	err = chewing.setupPhraseArray(params)
	if err != nil {
		return nil, err
	}

	chewing.setupPhraseTree(params)
	chewing.setupPhraseBKForest(params)

	return chewing, nil
}

func (this *ChewingDefaultLogger) Printf(format string, v ...interface{}) {}

func (this *Chewing) setupLogger(params *ChewingParameters) {
	if params.Logger == nil {
		params.Logger = new(ChewingDefaultLogger)
	}
	this.logger = params.Logger
}

func (this *Chewing) setupPhraseArray(params *ChewingParameters) (err error) {
	file, err := os.Open(params.PhraseFile)
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
			return errors.New(fmt.Sprintf("`%s' is invalid in PhraseFile", text))
		}

		var frequency uint32
		count, _ := fmt.Sscanf(token[1], "%d", &frequency)
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

func (this *Chewing) setupPhraseTree(params *ChewingParameters) {
	this.phraseTree = newPhraseTree()
	for _, item := range this.phraseArray.array {
		this.phraseTree.insert(item)
	}
}

func (this *Chewing) setupPhraseBKForest(params *ChewingParameters) {
	this.phraseBKForest = newPhraseBKForest()
	for _, item := range this.phraseArray.array {
		this.phraseBKForest.insert(item)
	}
}

func (this *Chewing) SetKeyboardType(keyboard int) error {
	if keyboard < KEYBOARD_MIN || keyboard > KEYBOARD_MAX {
		return errors.New(fmt.Sprintf("illegal keyboard type %d", keyboard))
	}
	this.keyboardType = keyboard
	return nil
}
