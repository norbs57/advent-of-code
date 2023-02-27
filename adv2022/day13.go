package adv2022

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func Day13() {
	sc := bufio.NewScanner(os.Stdin)
	items := make([]interface{}, 2)
	allItems := make([]interface{}, 0)
	sum := 0
	for idx := 1; true; idx++ {
		lines := make([]string, 2)
		for i := range lines {
			sc.Scan()
			lines[i] = sc.Text()
			json.Unmarshal([]byte(lines[i]), &items[i])
		}
		if CmpItem(items[0], items[1]) == -1 {
			sum += idx
		}
		allItems = append(allItems, items[0], items[1])
		if !sc.Scan() {
			break
		}
	}
	fmt.Println(sum)
	delimiters := []string{"[[2]]", "[[6]]"}
	for i := range delimiters {
		json.Unmarshal([]byte(delimiters[i]), &items[i])
	}
	allItems = append(allItems, items[0], items[1])
	sort.Slice(allItems, func(i, j int) bool {
		return CmpItem(allItems[i], allItems[j]) == -1
	})
	prod := 1
	for i, item := range allItems {
		if CmpItem(items[0], item) == 0 || CmpItem(items[1], item) == 0 {
			prod *= i + 1
		}
	}
	fmt.Println(prod)
}

func CmpItem(s, t interface{}) int {
	// return -1 if s < t, 0 if s== t and +1 if s > t
	switch v := s.(type) {
	case float64:
		switch w := t.(type) {
		case float64:
			if v < w {
				return -1
			}
			if v > w {
				return +1
			}
			return 0
		case []interface{}:
			return CmpItem([]interface{}{v}, t)
		}
	case []interface{}:
		switch w := t.(type) {
		case float64:
			return CmpItem(s, []interface{}{w})
		case []interface{}:
			if len(v) == 0 {
				if len(w) == 0 {
					return 0
				} else {
					return -1
				}
			}
			if len(w) == 0 {
				return 1
			}
			s0, t0 := v[0], w[0]
			cmp0 := CmpItem(s0, t0)
			if cmp0 != 0 {
				return cmp0
			} else {
				return CmpItem(v[1:], w[1:])
			}
		}
	}
	panic("Compare")
}
