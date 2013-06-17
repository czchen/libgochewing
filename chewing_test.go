package libgochewing

import (
	"io/ioutil"
	"launchpad.net/gocheck"
	"os"
	"path"
	"runtime"
	"testing"
)

type ChewingSuite struct{
	phraseFile string
}

var _ = gocheck.Suite(&ChewingSuite{})

func (this *ChewingSuite) SetUpSuite(c *gocheck.C) {
	phraseFile, err := ioutil.TempFile("", "")
	c.Assert(phraseFile, gocheck.FitsTypeOf, &os.File{})
	c.Assert(err, gocheck.IsNil)

	this.phraseFile = phraseFile.Name()

	phraseFile.WriteString(
		"# This is comment\n" +
			"\x09\x0b\x20測試 5 ㄘㄜˋ ㄕˋ\n" +
			"側室 4 ㄘㄜˋ ㄕˋ\x09\x0b\x20\n" +
			"側視 3 ㄘㄜˋ ㄕ # This is commentˋ\n" +
			"策士 2 ㄘㄜˋ ㄕˋ\n" +
			"策試 1 ㄘㄜˋ ㄕˋ\n")
	phraseFile.Close()
}

func (this *ChewingSuite) TearDownSuite(c *gocheck.C) {
	os.Remove(this.phraseFile)
}

func (this *ChewingSuite) TestNew(c *gocheck.C) {
	chewing, err := New(&ChewingParameters{
		PhraseFile: this.phraseFile,
	})

	c.Check(chewing, gocheck.FitsTypeOf, &Chewing{})
	c.Check(err, gocheck.IsNil)
}

func (this *ChewingSuite) TestNewNoPhraseFile(c *gocheck.C) {
	chewing, err := New(&ChewingParameters{
		PhraseFile: "NoSuchFile",
	})

	c.Check(chewing, gocheck.IsNil)
	c.Check(err, gocheck.NotNil)
}

func (this *ChewingSuite) TestNewBadFrequency(c *gocheck.C) {
	tmpFile, err := ioutil.TempFile("", "")
	c.Assert(tmpFile, gocheck.FitsTypeOf, &os.File{})
	c.Assert(err, gocheck.IsNil)

	tmpFileName := tmpFile.Name()
	defer os.Remove(tmpFileName)

	tmpFile.WriteString("測試 a ㄘㄜˋ ㄕˋ\n")
	tmpFile.Close()

	chewing, err := New(&ChewingParameters{
		PhraseFile: tmpFileName,
	})
	c.Check(chewing, gocheck.IsNil)
	c.Check(err, gocheck.NotNil)
}

func (this *ChewingSuite) TestNewBadBopomofo(c *gocheck.C) {
	tmpFile, err := ioutil.TempFile("", "")
	c.Assert(tmpFile, gocheck.FitsTypeOf, &os.File{})
	c.Assert(err, gocheck.IsNil)

	tmpFileName := tmpFile.Name()
	defer os.Remove(tmpFileName)

	tmpFile.WriteString("測試 a ㄘㄜ1 ㄕˋ\n")
	tmpFile.Close()

	chewing, err := New(&ChewingParameters{
		PhraseFile: tmpFileName,
	})
	c.Check(chewing, gocheck.IsNil)
	c.Check(err, gocheck.NotNil)
}

func (this *ChewingSuite) TestSetKeyboardType(c *gocheck.C) {
	chewing, err := New(&ChewingParameters{
		PhraseFile: this.phraseFile,
	})

	c.Assert(chewing, gocheck.FitsTypeOf, &Chewing{})
	c.Assert(err, gocheck.IsNil)

	c.Check(chewing.SetKeyboardType(KEYBOARD_DEFAULT), gocheck.IsNil)
	c.Check(chewing.SetKeyboardType(KEYBOARD_MIN - 1), gocheck.NotNil)
	c.Check(chewing.SetKeyboardType(KEYBOARD_MAX + 1), gocheck.NotNil)
}

func BenchmarkNew(b *testing.B) {
	_, phraseFile, _, _ := runtime.Caller(0)
	phraseFile = path.Join(path.Dir(phraseFile), "data", "tsi.src")

	for i := 0; i < b.N; i++ {
		_, err := New(&ChewingParameters{
			PhraseFile: phraseFile,
		})
		if err != nil {
			panic("New shall not return error")
		}
	}
}

func BenchmarkBKForestQuery(b *testing.B) {
	_, phraseFile, _, _ := runtime.Caller(0)
	phraseFile = path.Join(path.Dir(phraseFile), "data", "tsi.src")

	ctx, err := New(&ChewingParameters{
		PhraseFile: phraseFile,
	})
	if err != nil {
		panic("New shall not return error")
	}

	pivot := newFakePhoneSeq([]uint16{10268, 8708})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.phraseBKForest.query(pivot, 2)
	}
}
