package models

type HistoryStack struct {
	items []string
}

func HistoryStackNew() *HistoryStack {
	hs := HistoryStack{}
	hs.items = []string{}
	return &hs
}

func (hs *HistoryStack) Pop() string {
	if len(hs.items) == 0 {
		return ""
	}

	item := hs.items[len(hs.items)-1]

	hs.items = hs.items[0 : len(hs.items)-1]

	return item
}

func (hs *HistoryStack) Push(item string) {
	hs.items = append(hs.items, item)
}

func (hs *HistoryStack) Peek() string {
	return hs.items[len(hs.items)-1]
}

func (hs *HistoryStack) Length() int {
	return len(hs.items)
}
