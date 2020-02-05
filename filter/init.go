package filter

import (
	"bufio"
	"io"
	"os"
	"questions/util"
	"strings"
)

var (
	trie *util.Trie
)

func Init(filepath string) (err error) {
	trie = util.NewTrie()
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, errREet := reader.ReadString('\n')
		//去除换行符和空格
		word := strings.TrimSpace(strings.Replace(line, "\r\n", "", -1))
		if errREet == io.EOF {
			break
		}
		if errREet != nil {
			err = errREet
			return
		}
		if len(word) == 0 {
			continue
		}
		err = trie.Add(word, nil)
	}

	return
}
