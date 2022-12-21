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

type Mineral int

const (
	Ore      = Mineral(0)
	Clay     = Mineral(1)
	Obsidian = Mineral(2)
	Geode    = Mineral(3)
)

type State struct {
	Previous *State
	T        int

	Robots   [4]int
	Minerals [4]int
}

func (s State) Less(other State) bool {
	return s != other &&
		s.Robots[Ore] <= other.Robots[Ore] &&
		s.Minerals[Ore] <= other.Minerals[Ore] &&
		s.Robots[Clay] <= other.Robots[Clay] &&
		s.Minerals[Clay] <= other.Minerals[Clay] &&
		s.Robots[Obsidian] <= other.Robots[Obsidian] &&
		s.Minerals[Obsidian] <= other.Minerals[Obsidian] &&
		s.Robots[Geode] <= other.Robots[Geode] &&
		s.Minerals[Geode] <= other.Minerals[Geode]
	// if s.GeodeRobot != other.GeodeRobot {
	// 	return s.GeodeRobot < other.GeodeRobot
	// }
	// if s.ObsidianRobot != other.ObsidianRobot {
	// 	return s.ObsidianRobot < other.ObsidianRobot
	// }
	// if s.ClayRobot != other.ClayRobot {
	// 	return s.ClayRobot < other.ClayRobot
	// }
	// if s.OreRobot != other.OreRobot {
	// 	return s.OreRobot < other.OreRobot
	// }
	// return false
}

func (s State) MineMineralsInto(bp BP) func(dest State) State {
	return func(dest State) State {
		return dest.Mod(func(dest *State) {
			dest.T++
			dest.Previous = &s
			dest.Minerals[Ore] += s.Robots[Ore]
			dest.Minerals[Clay] += s.Robots[Clay]
			dest.Minerals[Obsidian] += s.Robots[Obsidian]
			dest.Minerals[Geode] += s.Robots[Geode]

			// Set dest to INFINITY (maxInt) if were never running out again
			if dest.Minerals[Ore] >= (minutes-dest.T-1)*max(bp.OreRobotOre, bp.ClayRobotOre, bp.ObsidianRobotOre, bp.GeodeRobotOre) {
				dest.Minerals[Ore] = lib.Infinity
			}
			if dest.Minerals[Clay] >= (minutes-dest.T-1)*bp.ObsidianRobotClay {
				dest.Minerals[Clay] = lib.Infinity
			}
			if dest.Minerals[Obsidian] >= (minutes-dest.T-1)*bp.GeodeRobotObsidian {
				dest.Minerals[Obsidian] = lib.Infinity
			}
		})
	}
}

func (s State) Mod(fn func(*State)) State {
	minerals := [4]int{}
	robots := [4]int{}
	copy(minerals[:], s.Minerals[:])
	copy(robots[:], s.Robots[:])
	s.Minerals = minerals
	s.Robots = robots
	fn(&s)
	return s
}

func (s State) String() (out string) {
	if s.Previous != nil {
		out = s.Previous.String() + " -> \n"
	}
	out += fmt.Sprintf("%#v (%d)", s, s.Minerals[Geode])
	return out
}

// func FeasibleGeodes(geodes int, bp BP) (bool, State) {
// 	bp.GeodeRobotObsidian
// 	bp.GeodeRobotOre
// 	State{T: T-1, Geode: }

//		mineOnly := s
//		for timeLeft := desired.T - s.T; timeLeft > 0; timeLeft-- {
//			s.MineMineralsInto(mineOnly)
//		}
//		if mineOnly.Gt(desired) {
//			return true, mineOnly
//		}
//	}
//
//	func (s State) Feasible(desired State) (bool, State) {
//		mineOnly := s
//		for timeLeft := desired.T - s.T; timeLeft > 0; timeLeft-- {
//			s.MineMineralsInto(mineOnly)
//		}
//		if mineOnly.Gt(desired) {
//			return true, mineOnly
//		}
//	}
func (s State) NeedOreRobot(bp BP) bool {
	return minutes-s.T > 2 // TODO check if we dont already have plenty!
}
func (s State) NeedClayRobot(bp BP) bool {
	return minutes-s.T > 6 // TODO check if we dont already have plenty!
}
func (s State) NeedObsidianRobot(bp BP) bool {
	return minutes-s.T > 4 // TODO check if we dont already have plenty!
}
func (s State) NeedGeodeRobot(bp BP) bool {
	return minutes-s.T > 2
}

const minutes = 24

func (s State) Build(bp BP) (out []State) {
	// defer func() {
	// 	fmt.Printf("%+v -> [%d][\n", s, len(out))
	// 	for _, o := range out {
	// 		fmt.Printf("%+v\n", o)
	// 	}
	// 	fmt.Println("]")
	// }()

	if s.Minerals[Ore] >= bp.OreRobotOre && s.Minerals[Ore] < lib.Infinity {
		out = append(out, s.Mod(func(s *State) {
			s.Minerals[Ore] -= bp.OreRobotOre
			s.Robots[Ore] += 1
		}))
	}
	if s.Minerals[Ore] >= bp.ClayRobotOre && s.Minerals[Clay] < lib.Infinity {
		out = append(out, s.Mod(func(s *State) {
			s.Minerals[Ore] -= bp.ClayRobotOre
			s.Robots[Clay] += 1
		}))
	}
	if s.Minerals[Ore] >= bp.ObsidianRobotOre && s.Minerals[Clay] >= bp.ObsidianRobotClay && s.Minerals[Obsidian] < lib.Infinity {
		out = append(out, s.Mod(func(s *State) {
			s.Minerals[Ore] -= bp.ObsidianRobotOre
			s.Minerals[Clay] -= bp.ObsidianRobotClay
			s.Robots[Obsidian] += 1
		}))
	}
	if s.Minerals[Ore] >= bp.GeodeRobotOre && s.Minerals[Obsidian] >= bp.GeodeRobotObsidian && s.T < minutes-2 {
		out = append(out, s.Mod(func(s *State) {
			s.Minerals[Ore] -= bp.GeodeRobotOre
			s.Minerals[Obsidian] -= bp.GeodeRobotObsidian
			s.Robots[Geode] += 1
		}))
	}

	out = lib.UniqueUsingKey(append(out, s))
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
	states := []State{{T: 0, Robots: [4]int{1, 0, 0, 0}}}
	max := 0
	var maxState State
	for t := 0; t < minutes; t++ {
		newstates := []State{}
		for _, oldstate := range states {
			newstates = append(newstates, lib.Map(oldstate.Build(bp), oldstate.MineMineralsInto(bp))...)
		}
		states = lib.UniqueUsingKey(newstates)
		max = 0
		for _, state := range newstates {
			if max < state.Minerals[Geode] {
				max = state.Minerals[Geode]
				maxState = state
			}
		}
		fmt.Println("t", t, "states", len(states), "max", max)
		sort.Stable(sort.Reverse(StateSlice(states)))
		if len(states) > 20 {
			states = states[0:20]
		}
	}
	fmt.Println(maxState)
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

func max(ss ...int) (max int) {
	for _, s := range ss {
		if s > max {
			max = s
		}
	}
	return max
}
