// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package util

const (
	Escape    = "\u001B["
	Bold      = "1"
	FGBlack   = "30m"
	FGRed     = "31m"
	FGGreen   = "32m"
	FGYellow  = "33m"
	FGBlue    = "34m"
	FGMagenta = "35m"
	FGCyan    = "36m"
	FGWhite   = "37m"
	BGBlack   = "40m"
	BGRed     = "41m"
	BGGreen   = "42m"
	BGYellow  = "43m"
	BGBlue    = "44m"
	BGMagenta = "45m"
	BGCyan    = "46m"
	BGWhite   = "47m"
)

func Colors(enable bool) ANSIColors {
	if enable {
		return ANSIColors{
			KerriaFlower:   Escape + Bold + ";" + FGYellow,
			Action:         Escape + Bold + ";" + FGBlue,
			Success:        Escape + FGGreen,
			VersionControl: Escape + FGRed,
			Timestamp:      Escape + FGMagenta,
			ExplainSection: Escape + Bold + ";" + FGBlue,
			ExplainField:   Escape + FGCyan,
			ExplainType:    Escape + FGGreen,
			Reset:          Escape + "0m",
		}
	}
	return ANSIColors{}
}

type ANSIColors struct {
	KerriaFlower   string
	Action         string
	Success        string
	VersionControl string
	Timestamp      string
	ExplainSection string
	ExplainField   string
	ExplainType    string
	Reset          string
}
