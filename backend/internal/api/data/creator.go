package data // calculate the creator with the most active products

import (
	"sort"
	"time"
)

func CalcActiveCreators(p *Payload, limit int) []string {
	// ac = active creators
	acMap := make(map[string]int, len(p.Products))

	// count the number of products created by each creator
	for _, creator := range p.Products {
		acMap[creator.ID]++
	}

	// check if all values in acMap are equal
	entriesEqual := allValuesEqual(acMap)
	// create a slice to store the top active creators initially empty
	topActiveCreators := make([]string, 0)
	if entriesEqual {
		// Sort products by creation time
		sort.Slice(p.Products, func(i, j int) bool {
			// parse the creation time to time.Time and compare it with other products creation time
			return parseTime(p.Products[i].CreateTime).After(parseTime(p.Products[j].CreateTime))
		})

		// Find the most recent creator's email
		for i := len(p.Creators) - 1; i >= 0; i-- {
			if p.Creators[i].ID == p.Products[0].CreatorID {
				topActiveCreators = append(topActiveCreators, p.Creators[i].Email)
				break
			}
		}
	} else {
		// Sort creators by activity (descending order)
		sort.Slice(p.Creators, func(i, j int) bool {
			return acMap[p.Creators[i].ID] > acMap[p.Creators[j].ID]
		})

		// Find the top active creators
		topActiveCreators = make([]string, limit)
		for i := 0; i < limit; i++ {
			topActiveCreators[i] = p.Creators[i].Email
		}

	}
	return topActiveCreators
}

func parseTime(s string) time.Time {
	// time formatter
	layout := time.RFC3339
	parsedTime, _ := time.Parse(layout, s)
	return parsedTime
}

func allValuesEqual(m map[string]int) bool {
	last := -1
	for _, v := range m {
		if last != -1 && last != v {
			return false
		}
		last = v
	}
	return true
}
