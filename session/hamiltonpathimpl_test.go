package session

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"math"
	"mdp_algo/common"
	"mdp_algo/utils"
	"strconv"
	"testing"
)

type HamiltonPathTestSuite struct {
	suite.Suite
	Session  Session
	arena    [][]string
	startPos common.Position
}

func (s *HamiltonPathTestSuite) SetupTest() {
	s.startPos = common.Position{
		Cell: common.Cell{
			Xcoord: 2,
			Ycoord: 17,
		},
		Direction: common.East,
	}
	s.Session = NewSession(20, 20, s.startPos)

	s.arena = [][]string{
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "S", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "S", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "W", "O", "O"},
		{"E", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "N", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "W", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "S", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "N", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
	}
	s.Session.LoadArena(s.arena)
}

func (s *HamiltonPathTestSuite) TestHamiltonPath() {
	plan := s.Session.HamiltonPath()
	fmt.Println("cost: ", plan.Cost)
	for i, pos := range plan.Path {
		if s.arena[pos.Cell.Ycoord][pos.Cell.Xcoord] == "O" {
			s.arena[pos.Cell.Ycoord][pos.Cell.Xcoord] = strconv.Itoa(i)
		} else {
			s.arena[pos.Cell.Ycoord][pos.Cell.Xcoord] += ";" + strconv.Itoa(i)
		}
	}
	utils.WriteCSV("tmp_hamilton.csv", s.arena)
}

func (s *HamiltonPathTestSuite) TestShortestPath() {
	arenaStr, err := ioutil.ReadFile("./assets/map1.txt")
	if err != nil {
		panic(err)
	}
	arena := parseArenaStr(string(arenaStr), 20)
	s.Session.LoadArena(arena)

	goalTest := func(pos common.Position) bool {
		goalPos := common.Position{
			Cell: common.Cell{
				Xcoord: 10,
				Ycoord: 14,
			},
			Direction: common.North,
		}
		return pos == goalPos
	}
	goalNum := 1
	moveableTest := func(cell common.Cell) bool {
		cs := s.Session.(*sessionImpl).getAugmentCellState(cell)
		return cs.state == empty
	}
	plan := s.Session.(*sessionImpl).shortestPath(s.startPos, goalTest, goalNum, moveableTest, hamiltonCostModel)
	assert.NotEqual(s.T(), math.MaxInt64, plan.Cost)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(HamiltonPathTestSuite))
}
