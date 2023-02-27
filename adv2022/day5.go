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
	stackStrings := &lib.Stack[string]{}
	for sc.Scan() {
		line := sc.Text()
		if len(line) > 0 {
			stackStrings.Push(line)
		} else {
			break
		}
	}
	numLine := stackStrings.Pop()
	nums := strings.Fields(numLine)
	n, _ := strconv.Atoi(nums[len(nums)-1])
	stacks1 := make([]lib.Stack[byte], n)
	stacks2 := make([]lib.Stack[byte], n)
	for stackStrings.Size() > 0 {
		line := stackStrings.Pop()
		for j := range stacks1 {
			c := line[1+j*offset]
			if c != ' ' {
				stacks1[j].Push(c)
				stacks2[j].Push(c)
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
			c := stacks1[src].Pop()
			stacks1[tgt].Push(c)
		}
		// part 2
		movingStack := &lib.Stack[byte]{}
		for j := 0; j < n; j++ {
			c := stacks2[src].Pop()
			movingStack.Push(c)
		}
		for j := 0; j < n; j++ {
			c := movingStack.Pop()
			stacks2[tgt].Push(c)
		}

	}
	for _, stack := range stacks1 {
		if stack.Size() > 0 {
			fmt.Print(string(stack.Pop()))
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for _, stack := range stacks2 {
		if stack.Size() > 0 {
			fmt.Print(string(stack.Pop()))
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
