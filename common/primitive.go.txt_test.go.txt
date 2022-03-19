package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransition(t *testing.T) {
	FrontRadius = 1
	SideRadius = 4
	start := Position{
		Cell: Cell{
			Xcoord: 5,
			Ycoord: 5,
		},
		Direction: East,
	}

	testCases := []struct {
		inputs   Transition
		expected Position
	}{
		{
			Transition{Forward, Forward},
			Position{
				Cell:      Cell{Xcoord: 7, Ycoord: 5},
				Direction: East,
			},
		},
		{
			Transition{ForwardRight},
			Position{
				Cell:      Cell{Xcoord: 6, Ycoord: 9},
				Direction: South,
			},
		},
		{
			Transition{Backward, Forward},
			start,
		},
		{
			Transition{BackwardLeft},
			Position{
				Cell:      Cell{Xcoord: 4, Ycoord: 1},
				Direction: South,
			},
		},
		{
			Transition{ForwardLeft, BackwardLeft},
			Position{
				Cell:      Cell{Xcoord: 2, Ycoord: 2},
				Direction: East,
			},
		},
		{
			Transition{ForwardLeft},
			Position{
				Cell:      Cell{Xcoord: 6, Ycoord: 1},
				Direction: North,
			},
		},
		{
			Transition{ForwardLeftRotation},
			Position{
				Cell:      Cell{Xcoord: 5, Ycoord: 5},
				Direction: North,
			},
		},
		{
			Transition{ForwardRightRotation},
			Position{
				Cell:      Cell{Xcoord: 5, Ycoord: 5},
				Direction: South,
			},
		},
	}

	for _, testCase := range testCases {
		actual := start.Transition(testCase.inputs...)
		assert.Equal(t, testCase.expected, actual)
	}
}

func TestPosition_Footprint(t *testing.T) {
	FrontRadius = 2
	SideRadius = 2
	start := Position{
		Cell: Cell{
			Xcoord: 13,
			Ycoord: 16,
		},
		Direction: North,
	}
	footprint := start.Footprint(ForwardLeft)
	expected := []Cell{
		{Xcoord: 11, Ycoord: 14}, {Xcoord: 12, Ycoord: 14}, {Xcoord: 13, Ycoord: 14},
		{Xcoord: 12, Ycoord: 15}, {Xcoord: 13, Ycoord: 15},
		{Xcoord: 13, Ycoord: 16},
	}
	assert.ElementsMatch(t, expected, footprint)

	start = Position{
		Cell: Cell{
			Xcoord: 13,
			Ycoord: 16,
		},
		Direction: West,
	}
	footprint = start.Footprint(ForwardRight)
	expected = []Cell{
		{Xcoord: 11, Ycoord: 16}, {Xcoord: 12, Ycoord: 16}, {Xcoord: 13, Ycoord: 16},
		{Xcoord: 11, Ycoord: 15}, {Xcoord: 12, Ycoord: 15},
		{Xcoord: 11, Ycoord: 14},
	}
	assert.ElementsMatch(t, expected, footprint)

	start = Position{
		Cell: Cell{
			Xcoord: 5,
			Ycoord: 5,
		},
		Direction: East,
	}

	footprint = start.Footprint(BackwardLeft)
	expected = []Cell{
		{Xcoord: 3, Ycoord: 3},
		{Xcoord: 3, Ycoord: 4}, {Xcoord: 4, Ycoord: 4},
		{Xcoord: 3, Ycoord: 5}, {Xcoord: 4, Ycoord: 5}, {Xcoord: 5, Ycoord: 5},
	}
	assert.ElementsMatch(t, expected, footprint)

	footprint = start.Footprint(Forward)
	expected = []Cell{
		{Xcoord: 5, Ycoord: 5}, {Xcoord: 6, Ycoord: 5},
	}
	assert.ElementsMatch(t, expected, footprint)

	FrontRadius = 3
	SideRadius = 3

	start = Position{
		Cell: Cell{
			Xcoord: 5,
			Ycoord: 5,
		},
		Direction: East,
	}

	footprint = start.Footprint(BackwardRight)
	expected = []Cell{
		{Xcoord: 2, Ycoord: 5}, {Xcoord: 3, Ycoord: 5}, {Xcoord: 4, Ycoord: 5}, {Xcoord: 5, Ycoord: 5},
		{Xcoord: 2, Ycoord: 6}, {Xcoord: 3, Ycoord: 6}, {Xcoord: 4, Ycoord: 6},
		{Xcoord: 2, Ycoord: 7}, {Xcoord: 3, Ycoord: 7},
		{Xcoord: 2, Ycoord: 8},
	}
	assert.ElementsMatch(t, expected, footprint)

	FrontRadius = 1
	SideRadius = 4

	start = Position{
		Cell: Cell{
			Xcoord: 10,
			Ycoord: 10,
		},
		Direction: West,
	}

	footprint = start.Footprint(ForwardLeft)
	expected = []Cell{
		{Xcoord: 10, Ycoord: 10},
		{Xcoord: 9, Ycoord: 10}, {Xcoord: 9, Ycoord: 11}, {Xcoord: 9, Ycoord: 12}, {Xcoord: 9, Ycoord: 13}, {Xcoord: 9, Ycoord: 14},
	}
	assert.ElementsMatch(t, expected, footprint)
}

func TestTurnOnTheSpotFootprint(t *testing.T) {
	start := Position{
		Cell: Cell{
			Xcoord: 13,
			Ycoord: 16,
		},
		Direction: North,
	}

	footprint := start.Footprint(ForwardLeftRotation)
	expected := []Cell{
		{Xcoord: 12, Ycoord: 14},
		{Xcoord: 12, Ycoord: 15}, {Xcoord: 13, Ycoord: 15},
		{Xcoord: 12, Ycoord: 16}, {Xcoord: 13, Ycoord: 16},
	}
	assert.ElementsMatch(t, expected, footprint)

	footprint = start.Footprint(ForwardRightRotation)
	expected = []Cell{
		{Xcoord: 14, Ycoord: 14},
		{Xcoord: 13, Ycoord: 15}, {Xcoord: 14, Ycoord: 15},
		{Xcoord: 13, Ycoord: 16}, {Xcoord: 14, Ycoord: 16},
	}
	assert.ElementsMatch(t, expected, footprint)
}
