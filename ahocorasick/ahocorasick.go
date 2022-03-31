package ahocorasick

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"stringmatch/customlog"
	"strings"
	"time"

	anknownahocorasick "github.com/anknown/ahocorasick"
	cloudflareahocorasick "github.com/cloudflare/ahocorasick"
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
	"github.com/sirupsen/logrus"
)

const CHN_DICT_FILE = "/Users/byeoungwoolee/go/src/stringmatch/ahocorasick/cn/dictionary.txt"
const CHN_TEXT_FILE = "/Users/byeoungwoolee/go/src/stringmatch/ahocorasick/cn/text.txt"
const ENG_DICT_FILE = "/Users/byeoungwoolee/go/src/stringmatch/ahocorasick/en/dictionary.txt"
const ENG_TEXT_FILE = "/Users/byeoungwoolee/go/src/stringmatch/ahocorasick/en/text.txt"

var clogrus *logrus.Logger = &logrus.Logger{}

func Dotest() {
	clogrus = customlog.DebugLogInit("ahocorasick")
	Strcompt()
	Ex1()
	TestcloudflareEnglish()
	TestanKnownEnglish()
	TestcloudflareChinese()
	TestanKnownChinese()
}

// func _ahocorasick() string {
// 	patterns := []string{
// 		"mercury", "venus", "earth", "mars",
// 		"jupiter", "saturn", "uranus", "pluto",
// 	}

// 	m := ahocorasick.NewStringMatcher(patterns)

// 	found := m.Match([]byte(`earth`))
// 	return fmt.Sprintln("found patterns", found)
// }

// func main() {

// 	strcompt()
// }

func Strcompt() {
	builder := ahocorasick.NewAhoCorasickBuilder(ahocorasick.Opts{
		AsciiCaseInsensitive: true,
		MatchOnlyWholeWords:  true,
		MatchKind:            ahocorasick.LeftMostLongestMatch,
		DFA:                  true,
	})
	ac := builder.Build([]string{"bear", "masha"})

	haystack := "The Bear and Masha"

	// r := ahocorasick.NewReplacer(ac)
	// replaced := r.ReplaceAllFunc(haystack, func(match ahocorasick.Match) (string, bool) {
	// 	return haystack[match.Start():match.End()], true
	// })
	// clogrus.Infof("replaced : %v", replaced)

	matches := ac.FindAll(haystack)

	clogrus.Info("FIND BY AHO-CORASICK")
	for _, match := range matches {
		clogrus.Info(haystack[match.Start():match.End()])
	}
}

func Ex1() {
	patterns := []string{
		"mercury", "venus", "earth", "mars",
		"jupiter", "saturn", "uranus", "pluto",
	}

	m := cloudflareahocorasick.NewStringMatcher(patterns)

	found := m.Match([]byte(`earth`))
	clogrus.Infof("found patterns : %v %v", found[len(found)-1], patterns[found[len(found)-1]])

}

func ReadBytes(filename string) ([][]byte, error) {
	dict := [][]byte{}

	f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)
	for {
		l, err := r.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		l = bytes.TrimSpace(l)
		dict = append(dict, l)
	}

	return dict, nil
}

func ReadRunes(filename string) ([][]rune, error) {
	dict := [][]rune{}

	f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)
	for {
		l, err := r.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		l = bytes.TrimSpace(l)
		dict = append(dict, bytes.Runes(l))
	}

	return dict, nil
}

func TestSEnglish() {
	content, err := ioutil.ReadFile(ENG_TEXT_FILE)
	if err != nil {
		clogrus.Error(err)
		return
	}
	scontent := string(content)
	if strings.EqualFold(scontent, "divinity") {
		clogrus.Infof("FOUND 'divinity'")
	}

}
func TestcloudflareEnglish() {
	clogrus.Info("** English Benchmark of cloudflare/ahocorasick **")
	clogrus.Info("-------------------------------------------------")
	clogrus.Info("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadBytes(ENG_DICT_FILE)
	if err != nil {
		clogrus.Error(err)
		return
	}

	content, err := ioutil.ReadFile(ENG_TEXT_FILE)
	if err != nil {
		clogrus.Error(err)
		return
	}
	end := time.Now()
	clogrus.Infof("load file cost:%d(ms)", (end.UnixNano()-start.UnixNano())/(1000*1000))

	clogrus.Info("=> Start to Search... ")
	start = time.Now()
	m := cloudflareahocorasick.NewMatcher(dict)

	//res := m.Match(content)
	// res := m.Match(content)
	m.Match(content)
	end = time.Now()

	clogrus.Infof("search cost:%d(ms)", (end.UnixNano()-start.UnixNano())/(1000*1000))

	// for _, v := range res {
	// 	clogrus.Infof("%d", v)
	// }

}

func TestcloudflareChinese() {
	clogrus.Info("** Chinese Benchmark of cloudflare/ahocorasick **")
	clogrus.Info("---------------------------------------------------")
	clogrus.Info("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadBytes(CHN_DICT_FILE)
	if err != nil {
		clogrus.Error(err)
		return
	}

	content, err := ioutil.ReadFile(CHN_TEXT_FILE)
	if err != nil {
		clogrus.Error(err)
		return
	}
	end := time.Now()
	clogrus.Infof("load file cost : %d(ms)", (end.UnixNano()-start.UnixNano())/(1000*1000))

	clogrus.Info("=> Start to Search... ")
	start = time.Now()
	m := cloudflareahocorasick.NewMatcher(dict)

	//res := m.Match(content)
	// res := m.Match(content)
	m.Match(content)
	end = time.Now()

	clogrus.Infof("search cost : %d(ms)", (end.UnixNano()-start.UnixNano())/(1000*1000))

	// for _, v := range res {
	// 	clogrus.Infof("%d", v)
	// }

}

func TestanKnownEnglish() {
	clogrus.Info("** English Benchmark of anknown/ahocorasick **")
	clogrus.Info("------------------------------------------------")
	clogrus.Info("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadRunes(ENG_DICT_FILE)
	if err != nil {
		clogrus.Error(err)
		return
	}

	content, err := ioutil.ReadFile(ENG_TEXT_FILE)
	if err != nil {
		clogrus.Error(err)
		return
	}

	contentRune := bytes.Runes([]byte(content))
	end := time.Now()
	clogrus.Info("load file cost : %d(ms)", (end.UnixNano()-start.UnixNano())/(1000*1000))

	clogrus.Info("=> Start to Search... ")
	start = time.Now()
	m := new(anknownahocorasick.Machine)
	if err := m.Build(dict); err != nil {
		clogrus.Error(err)
		return
	}
	//terms := m.Search(contentRune)
	// terms := m.MultiPatternSearch(contentRune, false)
	m.MultiPatternSearch(contentRune, false)
	end = time.Now()
	clogrus.Infof("search cost : %d(ms)", (end.UnixNano()-start.UnixNano())/(1000*1000))

	// for _, t := range terms {
	// 	clogrus.Infof("%d %s", t.Pos, string(t.Word))
	// }

}

func TestanKnownChinese() {
	clogrus.Info("** Korean Benchmark of anknown/ahocorasick **")
	clogrus.Info("------------------------------------------------")
	clogrus.Info("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadRunes(CHN_DICT_FILE)
	if err != nil {
		clogrus.Error(err)
		return
	}

	content, err := ioutil.ReadFile(CHN_TEXT_FILE)
	if err != nil {
		clogrus.Error(err)
		return
	}

	contentRune := bytes.Runes([]byte(content))
	end := time.Now()
	clogrus.Infof("load file cost : %d(ms)", (end.UnixNano()-start.UnixNano())/(1000*1000))

	clogrus.Info("=> Start to Search... ")
	start = time.Now()

	m := new(anknownahocorasick.Machine)
	if err := m.Build(dict); err != nil {
		clogrus.Error(err)
		return
	}

	// terms := m.Search(contentRune)
	// terms := m.MultiPatternSearch(contentRune, false)
	m.MultiPatternSearch(contentRune, false)
	end = time.Now()
	clogrus.Infof("search cost : %d(ms)", (end.UnixNano()-start.UnixNano())/(1000*1000))

	// for _, t := range terms {
	// 	clogrus.Infof("%d %s", t.Pos, string(t.Word))
	// }

}
