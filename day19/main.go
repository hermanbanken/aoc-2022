package main

import (
	"aoc/lib"
	"fmt"
	"regexp"
	"sort"
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

	// Each obsidian robot costs 3 ore and 20 clay.
	// Each ore robot costs 4 ore.
	// Each clay robot costs 2 ore.
	// Each clay robot costs 4 ore.
	// Each geode robot costs 2 ore and 17 obsidian.
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

func (s State) Less(other State) bool {
	if s.GeodeRobot != other.GeodeRobot {
		return s.GeodeRobot < other.GeodeRobot
	}
	if s.ObsidianRobot != other.ObsidianRobot {
		return s.ObsidianRobot < other.ObsidianRobot
	}
	if s.ClayRobot != other.ClayRobot {
		return s.ClayRobot < other.ClayRobot
	}
	if s.OreRobot != other.OreRobot {
		return s.OreRobot < other.OreRobot
	}
	return false
}

func (s State) MineMineralsInto(dest State) State {
	dest.T++
	dest.Ore += s.OreRobot
	dest.Clay += s.ClayRobot
	dest.Obsidian += s.ObsidianRobot
	dest.Geode += s.GeodeRobot
	return dest
}

func (s State) Mod(fn func(*State)) State {
	fn(&s)
	return s
}

// func FeasibleGeodes(geodes int, bp BP) (bool, State) {
// 	bp.GeodeRobotObsidian
// 	bp.GeodeRobotOre
// 	State{T: T-1, Geode: }

// 	mineOnly := s
// 	for timeLeft := desired.T - s.T; timeLeft > 0; timeLeft-- {
// 		s.MineMineralsInto(mineOnly)
// 	}
// 	if mineOnly.Gt(desired) {
// 		return true, mineOnly
// 	}
// }
// func (s State) Feasible(desired State) (bool, State) {
// 	mineOnly := s
// 	for timeLeft := desired.T - s.T; timeLeft > 0; timeLeft-- {
// 		s.MineMineralsInto(mineOnly)
// 	}
// 	if mineOnly.Gt(desired) {
// 		return true, mineOnly
// 	}
// }

func (s State) Build(bp BP) (out []State) {
	if s.Ore >= bp.OreRobotOre {
		out = append(out, s.Mod(func(s *State) {
			s.Ore -= bp.OreRobotOre
			s.OreRobot += 1
		}).Build(bp)...)
	}
	if s.Ore >= bp.ClayRobotOre {
		out = append(out, s.Mod(func(s *State) {
			s.Ore -= bp.ClayRobotOre
			s.ClayRobot += 1
		}).Build(bp)...)
	}
	if s.Ore >= bp.ObsidianRobotOre && s.Clay >= bp.ObsidianRobotClay {
		out = append(out, s.Mod(func(s *State) {
			s.Ore -= bp.ObsidianRobotOre
			s.Clay -= bp.ObsidianRobotClay
			s.ObsidianRobot += 1
		}).Build(bp)...)
	}
	//TODO why not adding geodes bots?
	if s.Ore >= bp.GeodeRobotOre && s.Obsidian >= bp.GeodeRobotObsidian {
		out = append(out, s.Mod(func(s *State) {
			s.Ore -= bp.GeodeRobotOre
			s.Obsidian -= bp.GeodeRobotObsidian
			s.GeodeRobot += 1
		}).Build(bp)...)
	}
	if len(out) == 0 {
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
		plan.maxGeodes = maxGeodes(plan)
		return plan
	})

	qlSum := 0
	for _, plan := range plans {
		qlSum += plan.idx * plan.maxGeodes
		fmt.Println("plan", plan.idx, "geodes", plan.maxGeodes)
	}
	fmt.Println(qlSum)

	maxTime := 24
	_ = maxTime
}

func maxGeodes(bp BP) int {

	states := []State{{T: 0, OreRobot: 1}}
	max := 0
	for t := 0; t <= 24; t++ {
		newstates := map[State]State{}
		for _, oldstate := range states {
			for _, newstate := range lib.Map(oldstate.Build(bp), func(after State) State { return oldstate.MineMineralsInto(after) }) {
				newstates[newstate] = newstate
			}
		}
		states = nil
		max = 0
		for _, state := range newstates {
			states = append(states, state)
			if max < state.Geode {
				max = state.Geode
			}
		}
		fmt.Println("t", t, "states", len(states), "max", max)
		sort.Sort(StateSlice(states))
		if len(states) > 20 {
			states = states[0:20]
		}
	}
	return max
}

type StateSlice []State

func (ss StateSlice) Len() int { return len(ss) }
func (ss StateSlice) Less(i, j int) bool {
	return ss[i].Less(ss[j])
}
func (ss StateSlice) Swap(i, j int) {
	tmp := ss[i]
	ss[i] = ss[j]
	ss[j] = tmp
}
