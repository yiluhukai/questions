package filter

func Replace(text, replace string) (result string, hit bool) {
	result, hit = trie.Check(text, replace)
	return
}
