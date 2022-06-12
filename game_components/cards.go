package game_components

type ResourceCards struct {
	CardsCount uint32
}

type DevCards struct {
	KnightsCount      uint32
	ProgressCount     uint32
	VictoryPointCount uint32
}

type BuildingCostCards struct {
	Count uint32
}

const (
	LongestRoad uint32 = iota
	LargestArmy
)
