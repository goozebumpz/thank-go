package main

type IntMap map[int]bool

func (im IntMap) Contains(elem int) bool {
	_, has := im[elem]
	return has
}

func (im IntMap) Add(elem int) bool {
	if im.Contains(elem) {
		return false
	}
	im[elem] = true

	return true
}

func MakeIntMap() IntMap {
	return IntMap{}
}
