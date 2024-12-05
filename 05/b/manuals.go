package main

import "fmt"

type Update []int

type Instructions map[int]map[int]bool // We are going to make a " this page map[page] has value{map[pages as keys]} before"

func (i Instructions) hasKey(key int) bool {
	_, ok := i[key]
	return ok
}

func (i Instructions) hasValue(key, value int) bool {
	_, ok := i[key][value]
	return ok
}

func (i Instructions) init(key int) {
	i[key] = make(map[int]bool)
}

func (i Instructions) printRules() {
	for page, pages := range i {
		fmt.Printf("%d == %v\n", page, pages)
	}
}

type ManualUpdate struct {
	after   Instructions // Pages that should come after key
	updates []Update
}

func (m ManualUpdate) getCorrectUpdates() (updates []Update) {
	for _, update := range m.updates {
		if m.validateUpdate(update) {
			updates = append(updates, update)
		}
	}
	return updates
}

func (m ManualUpdate) getWrongUpdates() (updates []Update) {
	for _, update := range m.updates {
		if !m.validateUpdate(update) {
			updates = append(updates, update)
		}
	}
	return updates
}

func (m ManualUpdate) validateUpdate(update Update) bool {
	// Let's validate page by page
	for pos, page := range update {
		// First we check the pages that came before
		if pos != 0 {
			for i := pos - 1; i >= 0; i-- {
				// Is it in the after list?
				// We don't have to check the before list, there can be pages without a rule anyway
				// We are just looking for rulebreakers!
				if m.after.hasValue(page, update[i]) {
					// fmt.Printf("%d was before page %d, but it's in the after list!\n", update[i], page)
					return false
				}
			}
		}
	}
	return true
}

// returns false if wrong, with the offenders: return 1 is before, but should be after return 2
func (m ManualUpdate) validateUpdateVerbose(update Update) (bool, int, int) {
	// Let's validate page by page
	for pos, page := range update {
		// First we check the pages that came before
		if pos != 0 {
			for i := pos - 1; i >= 0; i-- {
				// Is it in the after list?
				// We don't have to check the before list, there can be pages without a rule anyway
				// We are just looking for rulebreakers!
				if m.after.hasValue(page, update[i]) {
					// fmt.Printf("%d was before page %d, but it's in the after list!\n", update[i], page)
					return false, i, pos
				}
			}
		}
	}
	return true, 0, 0
}
