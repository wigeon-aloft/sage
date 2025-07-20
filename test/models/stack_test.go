package models

import (
	"testing"

	"gitlab.wige.one/wigeon/sage/internal/models"
)

func TestStackPush(t *testing.T) {

	item1 := "Lister"
	s := models.HistoryStack{}

	s.Push(item1)

	if s.Peek() != item1 {
		t.Errorf("s.Peek() == %s, expected pushed item %s", s.Peek(), item1)
	}
}

func TestStackLength(t *testing.T) {

	item1 := "Lister"
	item2 := "Rimmer"

	s := models.HistoryStack{}

	s.Push(item1)
	s.Push(item2)

	if s.Length() != 2 {
		t.Errorf("s.Length() == %d, expected %d", s.Length(), 2)
	}

}

func TestStackPop(t *testing.T) {

	item1 := "Lister"
	item2 := "Rimmer"

	s := models.HistoryStack{}

	s.Push(item1)
	s.Push(item2)

	poppedItem := s.Pop()

	if poppedItem != item2 {
		t.Errorf("poppedItem == %s, expected %s", poppedItem, item2)
	}

	if s.Length() != 1 {
		t.Errorf("s.Length() == %d, expected %d", s.Length(), 1)
	}

}

func TestStackPeek(t *testing.T) {

	item1 := "Lister"
	item2 := "Rimmer"

	s := models.HistoryStack{}

	s.Push(item1)
	s.Push(item2)

	peekedItem := s.Peek()

	if peekedItem != item2 {
		t.Errorf("peekedItem == %s, expected %s", peekedItem, item2)
	}

	if s.Length() != 2 {
		t.Errorf("s.Length() == %d, expected %d", s.Length(), 2)
	}

	s.Pop()

	peekedItem = s.Peek()

	if peekedItem != item1 {
		t.Errorf("peekedItem == %s, expected %s", peekedItem, item1)
	}

	if s.Length() != 1 {
		t.Errorf("s.Length() == %d, expected %d", s.Length(), 1)
	}

}
