package simulator

import "mdp_algo/common"

const height = 20
const width = 20

var startPosition = common.Position{
	Cell: common.Cell{
		Xcoord: 2,
		Ycoord: 17,
	},
	Direction: common.North,
}
