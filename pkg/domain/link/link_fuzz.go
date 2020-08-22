// +build gofuzz

package link

func FuzzNewURL(link []byte) int { // nolint unparam
	newLink, err := NewURL(string(link))
	if err != nil {
		return 1
	}
	if len(newLink.Hash) != 7 {
		return 1
	}
	return 0
}
