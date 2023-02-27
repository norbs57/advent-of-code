package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day5() {
	sc := bufio.NewScanner(os.Stdin)
	const offset = 4
	stackStrings := []string{}
	for sc.Scan() {
		line := sc.Text()
		if len(line) > 0 {
			lib.PushBack(&stackStrings, line)
		} else {
			break
		}
	}
	numLine := lib.PopBack(&stackStrings)
	nums := strings.Fields(numLine)
	n, _ := strconv.Atoi(nums[len(nums)-1])
	stacks1 := make([][]byte, n)
	stacks2 := make([][]byte, n)
	for len(stackStrings) > 0 {
		line := lib.PopBack(&stackStrings)
		for j := range stacks1 {
			c := line[1+j*offset]
			if c != ' ' {
				lib.PushBack(&stacks1[j], c)
				lib.PushBack(&stacks2[j], c)
			}
		}
	}
	for sc.Scan() {
		items := strings.Fields(sc.Text())
		n, _ := strconv.Atoi(items[1])
		src, _ := strconv.Atoi(items[3])
		tgt, _ := strconv.Atoi(items[5])
		src--
		tgt--
		// part 1
		for j := 0; j < n; j++ {
			c := lib.PopBack(&stacks1[src])
			lib.PushBack(&stacks1[tgt], c)
		}
		// part 2
		movingStack := []byte{}
		for j := 0; j < n; j++ {
			c := lib.PopBack(&stacks2[src])
			lib.PushBack(&movingStack, c)
		}
		for j := 0; j < n; j++ {
			c := lib.PopBack(&movingStack)
			lib.PushBack(&stacks2[tgt], c)
		}

	}
	for _, stack := range stacks1 {
		if len(stack) > 0 {
			fmt.Print(string(lib.PopBack(&stack)))
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for _, stack := range stacks2 {
		if len(stack) > 0 {
			fmt.Print(string(lib.PopBack(&stack)))
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
