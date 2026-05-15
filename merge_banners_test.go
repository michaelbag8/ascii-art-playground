package main

import (
	"reflect"
	"testing"
)

func TestMergeBanners(t *testing.T) {
	base := map[rune][]string{
		'A': {"baseA"},
		'B': {"baseB"},
	}

	priority := map[rune][]string{
		'A': {"priorityA"},
		'C': {"priorityC"},
	}

	t.Run("both maps merged with priority override", func(t *testing.T) {
		got := MergeBanners(base, priority)
		if got['A'][0] != "priorityA" {
			t.Fail()
		}
	})

	t.Run("base-only keys remain", func(t *testing.T) {
		got := MergeBanners(base, priority)
		if got['B'][0] != "baseB" {
			t.Fail()
		}
	})

	t.Run("priority-only keys added", func(t *testing.T) {
		got := MergeBanners(base, priority)
		if got['C'][0] != "priorityC" {
			t.Fail()
		}
	})

	t.Run("empty base returns priority copy", func(t *testing.T) {
		got := MergeBanners(map[rune][]string{}, priority)
		if len(got) != 2 {
			t.Fail()
		}
	})

	t.Run("empty priority returns base copy", func(t *testing.T) {
		got := MergeBanners(base, map[rune][]string{})
		if len(got) != 2 {
			t.Fail()
		}
	})

	t.Run("both empty returns empty", func(t *testing.T) {
		got := MergeBanners(map[rune][]string{}, map[rune][]string{})
		if len(got) != 0 {
			t.Fail()
		}
	})

	t.Run("modifying result does not affect base", func(t *testing.T) {
		got := MergeBanners(base, priority)
		got['B'][0] = "changed"
		if base['B'][0] == "changed" {
			t.Fail()
		}
	})

	t.Run("modifying result does not affect priority", func(t *testing.T) {
		got := MergeBanners(base, priority)
		got['C'][0] = "changed"
		if priority['C'][0] == "changed" {
			t.Fail()
		}
	})

	t.Run("deep copy slices", func(t *testing.T) {
		got := MergeBanners(base, priority)
		if &got['A'][0] == &priority['A'][0] {
			t.Fail()
		}
	})

	t.Run("deterministic output", func(t *testing.T) {
		got1 := MergeBanners(base, priority)
		got2 := MergeBanners(base, priority)
		if !reflect.DeepEqual(got1, got2) {
			t.Fail()
		}
	})
}