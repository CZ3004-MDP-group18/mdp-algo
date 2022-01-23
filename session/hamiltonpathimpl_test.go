package session

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"mdp_algo/common"
	"mdp_algo/utils"
	"strconv"
	"testing"
)

type HamiltonPathTestSuite struct {
	suite.Suite
	Session Session
	arena   [][]string
}

func (s *HamiltonPathTestSuite) SetupTest() {
	startPos := common.Position{
		Cell: common.Cell{
			Xcoord: 1,
			Ycoord: 18,
		},
		Direction: common.East,
	}
	s.Session = NewSession(20, 20, startPos)

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

func TestSuite(t *testing.T) {
	suite.Run(t, new(HamiltonPathTestSuite))
}
