package task

import (
	
)

var (
	ColorUrl string = "@{!wB}"
	ColorCheckedTrue string = ColorFGGreenBGWhite
	ColorCheckedFalse string = ColorFGRedBGWhite
	ColorFGRedBGWhite string = "@{!wR}"
	ColorFGMagentaBGWhite string = "@{!wM}"
	ColorFGYellowBGWhite string = "@{!wY}"
	ColorFGGreenBGWhite string = "@{!wG}"
	ColorFGWhiteBgBlack string = "@{!wK}"
	ColorNone string = ""

	PriorityColors = map[string]string{
		PriorityUrgent: ColorFGMagentaBGWhite,
		PriorityHigh: ColorFGRedBGWhite,
		PriorityMedium: ColorFGYellowBGWhite,
		PriorityLow: ColorFGWhiteBgBlack,
	}

	StatusColors = map[string]string{
		StatusOpen: ColorFGMagentaBGWhite,
		StatusNew: ColorFGWhiteBgBlack,
		StatusPause: ColorFGYellowBGWhite,
		StatusClosed: ColorFGGreenBGWhite,
	}
)