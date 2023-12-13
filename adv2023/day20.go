package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day20() {
	type Mdl struct {
		ty         byte
		inputs     []string
		dests      []string
		inputState map[string]bool
		state      int // 0/1 for ff, count of true inputs for conj
	}
	mMap := make(map[string]*Mdl)
	sc := bufio.NewScanner(os.Stdin)
	// parse input data
	for sc.Scan() {
		items := strings.Fields(sc.Text())
		ty := items[0][0]
		dests := []string{}
		for _, d := range items[2:] {
			dests = append(dests, strings.TrimSuffix(d, ","))
		}
		m := &Mdl{ty, []string{}, dests, nil, 0}
		name := items[0][1:]
		switch ty {
		case 'b':
			name = items[0]
		case '&':
			m.inputState = make(map[string]bool)
		}
		mMap[name] = m
	}
	// set module inputs
	for name, m := range mMap {
		for _, dest := range m.dests {
			mDest, ok := mMap[dest]
			if ok {
				lib.PushBack(&mDest.inputs, name)
				if mDest.ty == '&' {
					mDest.inputState[name] = false
				}
			}
		}
	}
	const brd = "broadcaster"
	type Signal struct {
		src   string
		dest  string
		pulse bool
	}
	qs := make([]Signal, 0)
	pulses := [2]int{}
	rxSent := false
	bPressed := 0
	// Module names obtained from manual inspection of input!
	obsModules := []string{"cd", "qx", "rk", "zf"}
	obs := make([]int, 4)
	send := func(src string, dests []string, pulse bool) {
		if pulse {
			pulses[1] += len(dests)
		} else {
			pulses[0] += len(dests)
		}
		for _, dest := range dests {
			qs = append(qs, Signal{src, dest, pulse})
			if dest == "rx" && !pulse {
				rxSent = true
			}
			for i, ob := range obsModules {
				if dest == ob && !pulse && obs[i] == 0 {
					obs[i] = bPressed
				}
			}
		}
	}
	buttonPress := func() {
		bPressed++
		send("button", []string{brd}, false)
		for len(qs) > 0 {
			sg := lib.PopFront(&qs)
			m, ok := mMap[sg.dest]
			if !ok {
				continue
			}
			switch m.ty {
			case 'b':
				send(brd, m.dests, sg.pulse)
			case '%':
				if !sg.pulse {
					m.state = 1 - m.state
					pulse := m.state == 1
					send(sg.dest, m.dests, pulse)
				}
			case '&':
				current := m.inputState[sg.src]
				if current != sg.pulse {
					if sg.pulse {
						m.state++
					} else {
						m.state--
					}
				}
				m.inputState[sg.src] = sg.pulse
				pulse := m.state != len(m.inputState)
				send(sg.dest, m.dests, pulse)
			}
		}
	}
	res := [2]int{}
	// Part 1
	for i := 1; i <= 1000; i++ {
		buttonPress()
	}
	res[0] = pulses[0] * pulses[1]
	// Part 2
	// reset state
	bPressed = 0
	for _, m := range mMap {
		m.state = 0
		for input := range m.inputState {
			m.inputState[input] = false
		}
	}
	// loop until rxSent or all periods have been determined
	for !rxSent && min(obs[0], obs[1], obs[2], obs[3]) == 0 {
		buttonPress()
	}
	if rxSent {
		res[1] = bPressed
	} else {
		res[1] = lib.Lcm(obs...)
	}
	fmt.Println(res[0], res[1])
}
