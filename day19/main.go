package main

import (
	"aoc/lib"
	"fmt"
	"regexp"
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

func (s State) Build(bp BP) (out []State) {
	out = append(out, s)
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
	if s.Ore >= bp.GeodeRobotOre && s.Obsidian >= bp.GeodeRobotObsidian {
		out = append(out, s.Mod(func(s *State) {
			s.Ore -= bp.GeodeRobotOre
			s.Obsidian -= bp.GeodeRobotObsidian
			s.GeodeRobot += 1
		}).Build(bp)...)
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
	for t := 0; t <= 24; t++ {
		newstates := map[State]State{}
		for _, oldstate := range states {
			for _, newstate := range lib.Map(oldstate.Build(bp), func(after State) State { return oldstate.MineMineralsInto(after) }) {
				newstates[newstate] = newstate
			}
		}
		states = nil
		for _, state := range newstates {
			states = append(states, state)
		}
		fmt.Println("t", t, "states", len(states))
	}
	return 0
}
