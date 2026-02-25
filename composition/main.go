package main

import (
	"time"
)

func testUniqRevert() {
	cancel := make(chan struct{})
	generated := generate(cancel)
	unique := takeUniq(cancel, generated)
	reversed1 := reverse(cancel, unique)
	reversed2 := reverse(cancel, unique)
	merged := mergeUniqRevert(cancel, reversed1, reversed2)
	printUniqRevert(cancel, merged, 5)

	close(cancel)

	time.Sleep(2000 * time.Millisecond)
}

func main() {
	testUniqRevert()
}
