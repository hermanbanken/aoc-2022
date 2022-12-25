package main

import (
	"aoc/lib"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type BP struct {
	idx       int
	maxGeodes int

	OreRobotOre  int
	ClayRobotOre int

	ObsidianRobotOre  int
	ObsidianRobotClay int

	GeodeRobotOre      int
	GeodeRobotObsidian int
}

type State struct {
	Previous *State
	T        int

	OreRobot      int
	ClayRobot     int
	ObsidianRobot int
	GeodeRobot    int

	Ore      int
	Clay     int
	Obsidian int
	Geode    int
}

func (s State) MineMineralsInto(bp BP) func(dest State) State {
	return func(dest State) State {
		dest.T++
		dest.Previous = &s
		dest.Ore += s.OreRobot
		dest.Clay += s.ClayRobot
		dest.Obsidian += s.ObsidianRobot
		dest.Geode += s.GeodeRobot

		// Set dest to INFINITY (maxInt) if were never running out again
		if dest.OreRobot >= max(bp.OreRobotOre, bp.ClayRobotOre, bp.ObsidianRobotOre, bp.GeodeRobotOre) {
			dest.Ore = lib.Infinity
		}
		if dest.ClayRobot >= bp.ObsidianRobotClay {
			dest.Clay = lib.Infinity
		}
		if dest.ObsidianRobot >= bp.GeodeRobotObsidian {
			dest.Obsidian = lib.Infinity
		}
		return dest
	}
}

func (s State) Mod(fn func(*State)) State {
	fn(&s)
	return s
}

func (s State) String() (out string) {
	if s.Previous != nil {
		out = s.Previous.String() + " -> \n"
	}
	out += fmt.Sprintf("State{T:%d, OreRobot:%s, ClayRobot:%s, ObsidianRobot:%s, GeodeRobot:%s, Ore:%s, Clay:%s, Obsidian:%s, Geode:%s}",
		s.T,
		Int(s.OreRobot),
		Int(s.ClayRobot),
		Int(s.ObsidianRobot),
		Int(s.GeodeRobot),
		Int(s.Ore),
		Int(s.Clay),
		Int(s.Obsidian),
		Int(s.Geode),
	)
	return out
}

type Int int

func (i Int) String() string {
	if i == Int(lib.Infinity) {
		return "Inf"
	}
	return fmt.Sprint(int(i))
}

func (s State) Build(bp BP) (out []State) {

	if s.Ore >= bp.OreRobotOre &&
		s.Ore < lib.Infinity {
		out = append(out, s.Mod(func(s *State) {
			if s.Ore < lib.Infinity {
				s.Ore -= bp.OreRobotOre
			}
			s.OreRobot += 1
		}))
	}

	if s.Ore >= bp.ClayRobotOre &&
		s.Clay < lib.Infinity {
		out = append(out, s.Mod(func(s *State) {
			if s.Ore < lib.Infinity {
				s.Ore -= bp.ClayRobotOre
			}
			s.ClayRobot += 1
		}))
	}

	if s.Ore >= bp.ObsidianRobotOre &&
		s.Clay >= bp.ObsidianRobotClay &&
		s.Obsidian < lib.Infinity {
		out = append(out, s.Mod(func(s *State) {
			if s.Ore < lib.Infinity {
				s.Ore -= bp.ObsidianRobotOre
			}
			if s.Clay < lib.Infinity {
				s.Clay -= bp.ObsidianRobotClay
			}
			s.ObsidianRobot += 1
		}))
	}

	if s.Ore >= bp.GeodeRobotOre &&
		s.Obsidian >= bp.GeodeRobotObsidian {
		out = append(out, s.Mod(func(s *State) {
			if s.Ore < lib.Infinity {
				s.Ore -= bp.GeodeRobotOre
			}
			if s.Obsidian < lib.Infinity {
				s.Obsidian -= bp.GeodeRobotObsidian
			}
			s.GeodeRobot += 1
		}))
	}

	// only add idle state, if we could not already build everything
	if s.Ore < max(bp.OreRobotOre, bp.ClayRobotOre, bp.ObsidianRobotOre, bp.GeodeRobotOre) ||
		s.Clay < bp.ObsidianRobotClay ||
		s.Obsidian < bp.GeodeRobotObsidian {
		out = append(out, s)
	}

	return out
}

var re = regexp.MustCompile(`Blueprint \d+: Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

/*`Blueprint (\d+): ` + strings.Join([]string{
	`Each ore robot costs (\d+) ore.`,
	`Each clay robot costs (\d+) ore.`,
	`Each obsidian robot costs (\d+) ore and (\d+) clay.`,
	`Each geode robot costs (\d+) ore and (\d+) obsidian.`,
}, "\r\n ")*/

func main() {
	idx := 0
	plans := lib.Map(lib.Lines(), func(bp string) BP {
		matches := re.FindStringSubmatch(bp)
		if len(matches) == 0 {
			panic(bp + "; not readable")
		}
		idx++
		plan := BP{idx: idx}
		plan.OreRobotOre = lib.Int(matches[1])
		plan.ClayRobotOre = lib.Int(matches[2])
		plan.ObsidianRobotOre = lib.Int(matches[3])
		plan.ObsidianRobotClay = lib.Int(matches[4])
		plan.GeodeRobotOre = lib.Int(matches[5])
		plan.GeodeRobotObsidian = lib.Int(matches[6])

		fmt.Printf("\nBlueprint %d\n%+v", idx, plan)
		plan.maxGeodes = maxGeodes(plan)
		return plan
	})

	qlSum := 0
	for _, plan := range plans {
		qlSum += plan.idx * plan.maxGeodes
		fmt.Println("plan", plan.idx, "geodes", plan.maxGeodes)
	}
	// part1: 1785 is too low
	fmt.Println(qlSum)
}

func maxGeodes(bp BP) int {
	states := []State{{T: 0, OreRobot: 1}}
	max := 0
	var maxState State
	for t := 0; t < 24; t++ {
		newstates := []State{}
		for _, oldstate := range states {
			newstates = append(newstates, lib.Map(oldstate.Build(bp), oldstate.MineMineralsInto(bp))...)
		}
		if bp.idx == 1 && strings.Contains(os.Args[1], "input0.txt") {
			tests1(t, newstates)
		}
		if bp.idx == 2 && strings.Contains(os.Args[1], "input0.txt") {
			tests2(t, newstates)
		}
		states = lib.UniqueUsingKeyPrepared(newstates, erasePrev)
		max = 0
		for _, state := range newstates {
			if max < state.Geode {
				max = state.Geode
				maxState = state
			}
		}
		fmt.Println("t", t, "states", len(states), "max", max)
	}
	fmt.Println(maxState)
	return max
}

func max(ss ...int) (max int) {
	for _, s := range ss {
		if s > max {
			max = s
		}
	}
	return max
}

func tests1(t int, newstates []State) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("test failed; t==%d\nStates:\n", t)
			for _, s := range newstates {
				fmt.Printf("o) %#v\n", s)
			}
			os.Exit(1)
		}
	}()
	if t == 2 {
		assert("t==2", some(newstates, State{T: 3, OreRobot: 1, ClayRobot: 1, Ore: 1}))
	}
	if t == 3 {
		assert("t==3", some(newstates, State{T: 4, OreRobot: 1, ClayRobot: 1, Ore: 2, Clay: 1}))
	}
}

func tests2(t int, newstates []State) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("test failed; t==%d\nStates:\n", t)
			for _, s := range newstates {
				fmt.Printf("o) %#v\n", s)
			}
			os.Exit(1)
		}
	}()
	if t == 23 {
		with12 := []State{}
		for _, s := range newstates {
			if s.Geode == 12 {
				with12 = append(with12, s)
			}
		}
		newstates = with12
		assert("t==23", len(with12) > 0)
		fmt.Printf("With 12 geodes (%d):\n%s\n\n", len(with12), with12[0])
	}
}

func some(hay []State, needle State) bool {
	needle.Previous = nil
	for _, h := range hay {
		h.Previous = nil
		if h == needle {
			return true
		}
	}
	return false
}

func assert(name string, b bool) {
	if !b {
		panic("assertion failed: " + name)
	}
}

func erasePrev(s State) State { s.Previous = nil; return s }
