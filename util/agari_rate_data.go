package util

const (
	// 参考:「統計学」のマージャン戦術
	FuritenAgariMulti = 0.4706727
)

var (
	// TODO: 基于巡目的和了率数据
	// 例如 566778m 122345s 77z

	// TODO: 考虑读山的和了率？

	// 6~10巡目 [牌0-4][剩余数1-4]
	// 参考: 勝つための現代麻雀技術論
	// TODO: 仅为无筋数据，未考虑筋牌、早外、NC、是否为宝牌、其他场况等，仅供参考
	// https://github.com/EndlessCheng/mahjong-helper/issues/46

	// 传入 tileType，返回剩余枚数的和率
	agariMap = map[tileType][...]float64{
		1: {0, 26.3, 41.6, 50.1, 55.0},
		{19.2, 31.7, 38.2, 42.0},
		{14.8, 25.5, 32.0, 36.8},
		{11.8, 20.3, 26.7, 31.0},
		{11.8, 20.3, 26.7, 31.0},
	}

	// 8巡目 [剩余数1-3]
	// 参考:「統計学」のマージャン戦術
	honorTileAgariTable = [3]float64{47.5, 58.0, 49.5}


)
