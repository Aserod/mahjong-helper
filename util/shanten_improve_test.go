package util

import (
	"testing"
	"github.com/EndlessCheng/mahjong-helper/util/model"
	"github.com/stretchr/testify/assert"
)

var exampleMelds = []model.Meld{{MeldType: model.MeldTypePon, Tiles: MustStrToTiles("666z")}}

func TestCalculateShantenAndWaits13(t *testing.T) {
	toString := func(shanten int, waits Waits) string {
		return NumberToChineseShanten(shanten) + " " + waits.String()
	}

	// closed
	assert.Equal(t, "三向听 32 进张 [46m 2468p 24s]", toString(CalculateShantenAndWaits13(MustStrToTiles34("11357m 13579p 135s"), nil)))
	assert.Equal(t, "听牌 4 进张 [4s]", toString(CalculateShantenAndWaits13(MustStrToTiles34("123456789m 1135s"), nil)))
	assert.Equal(t, "听牌 8 进张 [25s]", toString(CalculateShantenAndWaits13(MustStrToTiles34("123456789m 1134s"), nil)))
	assert.Equal(t, "两向听 12 进张 [1234z]", toString(CalculateShantenAndWaits13(MustStrToTiles34("123456789m 1234z"), nil)))
	assert.Equal(t, "一向听 61 进张 [12345678m 47p 12345678s]", toString(CalculateShantenAndWaits13(MustStrToTiles34("3456m 3456s 44456p"), nil)))

	// open
	assert.Equal(t, "听牌 6 进张 [14p]", toString(CalculateShantenAndWaits13(MustStrToTiles34("1234p"), nil)))
	assert.Equal(t, "两向听 12 进张 [1234z]", toString(CalculateShantenAndWaits13(MustStrToTiles34("1234z"), nil)))
	assert.Equal(t, "听牌 3 进张 [5p]", toString(CalculateShantenAndWaits13(MustStrToTiles34("5p"), nil)))
}

func TestCalculateShantenWithImproves13Closed(t *testing.T) {
	for _, tiles := range []string{
		//"11357m 13579p 135s",
		//"123456789m 1135s",
		//"123456789m 1134s",
		//"123456789m 1234z",
		//"3m 12668p 5678s 222z",
		//"6m 12668p 5678s 222z",
		//"557m 34789p 26s 111z",
		//"111333555m 23p 23s",
		//"23m 234p 234888s 44z",
		//"23467m 234p 23488s",
		//"34567m 22334455p",
		//"1199m 112235566z",
		//"123456789m 23p 88s",
		//"12399m 123p 12999s",
		//"3577m 345p 345678s",
		//"23467m 234p 23488s",
		//"13789m 11p 345s 555z",
		//"12346789m 123p 88s",
		//"3456m 111s 999p 777z",
		//"123m 44p 34888s 777z",
		//"13789m 111789p 77z",
		//"23467m 222p 23488s",
		//"13789m 111789p 11s",
		//"12346789m 123p 88s",
		"56778p 112345s 77z",
		"56778p 122345s 77z",
	} {
		tiles34 := MustStrToTiles34(tiles)
		if CountOfTiles34(tiles34) != 13 {
			t.Error(tiles, "不是13张牌")
			continue
		}
		playerInfo := model.NewSimplePlayerInfo(tiles34, nil)
		//playerInfo.DiscardTiles = []int{MustStrToTile34("4s")}
		//playerInfo.IsRiichi = true
		//playerInfo.DoraCount = 2
		result := CalculateShantenWithImproves13(playerInfo)
		t.Log(tiles, "=\n"+result.String())
	}
}

func TestCalculateShantenWithImproves13Open(t *testing.T) {
	for _, tiles := range []string{
		"1234m",
		"1135m",
		"5p",
	} {
		tiles34 := MustStrToTiles34(tiles)
		result := CalculateShantenWithImproves13(model.NewSimplePlayerInfo(tiles34, exampleMelds))
		t.Log(tiles, "=\n"+result.String())
	}
}

func TestCalculateShantenWithImproves14Closed(t *testing.T) {
	tiles := "124679m 3678p 2366s"
	tiles = "11379m 347p 277s 111z"
	tiles = "334578m 11468p 235s"
	tiles = "478m 33588p 457899s"
	tiles = "2233688m 1234p 379s"
	tiles = "1233347m 23699p 88s"
	tiles = "334m 22457p 23456s 1z"
	tiles = "334m 122478p 23456s"
	tiles = "1m 258p 258s 1234567z"
	tiles = "4567m 4579p 344588s"
	tiles = "2479999m 45667p 13s" // 切任何一张都不会向听倒退
	tiles = "25667m 27789p 37s 44z"
	tiles = "111444777m 11177s"
	tiles = "2468m 33578p 22356s"
	tiles = "57m 4455p 12345699s"
	tiles = "57m 3445667p 12399s"
	tiles = "2335578899m 5677p"
	tiles = "123p 3445668m 6799s"
	tiles = "455678m 11566p 234s" // TODO 振听 9m 的场合，切 6p 振听听牌的概率比切 5m 低
	tiles = "1245m 12789p 34588s"
	tiles = "4456778p 2245s 111z"
	tiles = "388m 113668p 56s 456z"
	tiles = "56778p 1122345s 77z"
	tiles = "66778p 1122345s 77z"
	tiles = "67778p 1122345s 77z"
	tiles = "3336888m 678p 5678s"
	tiles = "145679p 1s 1234567z"
	playerInfo := model.NewSimplePlayerInfo(MustStrToTiles34(tiles), nil)
	//playerInfo.NumRedFives[2] = 1
	//playerInfo.IsTsumo = true
	//playerInfo.LeftTiles34 = InitLeftTiles34WithTiles34(MustStrToTiles34("388m 113668p 566s 45556z")) // 注意手牌也算上
	//playerInfo.DiscardTiles = []int{MustStrToTile34("9p")}
	shanten, results, incShantenResults := CalculateShantenWithImproves14(playerInfo)
	t.Log(NumberToChineseShanten(shanten))
	for _, result := range results {
		t.Log(result)
	}
	if len(incShantenResults) > 0 {
		t.Log(NumberToChineseShanten(shanten + 1))
		for _, result := range incShantenResults {
			t.Log(result)
		}
	} else {
		t.Log("无向听倒退的切牌")
	}
}

func BenchmarkCalculateShantenWithImproves14Closed(b *testing.B) {
	tiles34 := MustStrToTiles34("124679m 3678p 2366s")
	playerInfo := model.NewSimplePlayerInfo(tiles34, nil)
	for i := 0; i < b.N; i++ {
		// 剪枝前：0.28s
		// 剪枝后：0.22s
		CalculateShantenWithImproves14(playerInfo)
	}
}

func TestCalculateShantenWithImproves14Open(t *testing.T) {
	tiles := "35m"
	tiles = "13m 456s 778p"
	tiles = "6888m 678p 5678s"
	//leftTiles34 := InitLeftTiles34WithTiles34(MustStrToTiles34(tiles))
	//leftTiles34[1] = 0
	playerInfo := model.NewSimplePlayerInfo(MustStrToTiles34(tiles), exampleMelds)
	//playerInfo.LeftTiles34 = leftTiles34
	shanten, results, incShantenResults := CalculateShantenWithImproves14(playerInfo)
	t.Log(NumberToChineseShanten(shanten))
	for _, result := range results {
		t.Log(result)
	}
	t.Log(NumberToChineseShanten(shanten + 1))
	for _, result := range incShantenResults {
		t.Log(result)
	}
}

func TestCalculateMeld(t *testing.T) {
	tiles := "1234m 112z"
	tiles = "23445667m 11z"
	tiles = "112356799m 1233z"
	tiles = "78m 12355p 789s" // ***
	tiles = "245689s 1z"
	tiles34 := MustStrToTiles34(tiles)
	result := CalculateShantenWithImproves13(model.NewSimplePlayerInfo(tiles34, exampleMelds))
	t.Log("原手牌" + NumberToChineseShanten(result.Shanten))
	t.Log(result)

	tile := "1m"
	tile = "3m" // "1z"
	tile = "4m"
	tile = "4p"
	tile = "3s"
	shanten, results, incShantenResults := CalculateMeld(model.NewSimplePlayerInfo(tiles34, exampleMelds), MustStrToTile34(tile), false, true)
	t.Log("鸣牌后" + NumberToChineseShanten(shanten))
	for _, result := range results {
		t.Log(result)
	}
	t.Log("鸣牌后" + NumberToChineseShanten(shanten+1))
	for _, result := range incShantenResults {
		t.Log(result)
	}
}

//

// 何切 300
func TestQ300(t *testing.T) {
	tiles := "56778p 1122345s 77z" // Q001
	playerInfo := model.NewSimplePlayerInfo(MustStrToTiles34(tiles), nil)
	//playerInfo.IsTsumo = true
	//playerInfo.LeftTiles34 = InitLeftTiles34WithTiles34(MustStrToTiles34("388m 113668p 566s 45556z")) // 注意手牌也算上
	shanten, results, incShantenResults := CalculateShantenWithImproves14(playerInfo)
	t.Log(NumberToChineseShanten(shanten))
	for _, result := range results {
		t.Log(result)
	}
	if len(incShantenResults) > 0 {
		t.Log(NumberToChineseShanten(shanten + 1))
		for _, result := range incShantenResults {
			t.Log(result)
		}
	} else {
		t.Log("无向听倒退的切牌")
	}
}

func Test_calculateIsolatedTileValue(t *testing.T) {
	newPI := func(selfWindTile int, roundWindTile int, discardedHumanTiles string) *model.PlayerInfo {
		return &model.PlayerInfo{
			SelfWindTile:  selfWindTile,
			RoundWindTile: roundWindTile,
			LeftTiles34:   InitLeftTiles34WithTiles34(MustStrToTiles34(discardedHumanTiles)),
		}
	}

	const eps = 1e-3

	assert.InDelta(t, 100, float64(calculateIsolatedTileValue(MustStrToTile34("9m"), newPI(27, 27, "2s"))), eps)
	assert.InDelta(t, 130, float64(calculateIsolatedTileValue(MustStrToTile34("1z"), newPI(27, 27, "2s"))), eps)
	assert.InDelta(t, 117, float64(calculateIsolatedTileValue(MustStrToTile34("1z"), newPI(27, 27, "2s11z"))), eps)
	assert.InDelta(t, 97, float64(calculateIsolatedTileValue(MustStrToTile34("2z"), newPI(27, 27, "2s"))), eps)
	assert.InDelta(t, 98, float64(calculateIsolatedTileValue(MustStrToTile34("3z"), newPI(27, 27, "2s"))), eps)
	assert.InDelta(t, 99, float64(calculateIsolatedTileValue(MustStrToTile34("4z"), newPI(27, 27, "2s"))), eps)
	assert.InDelta(t, 115, float64(calculateIsolatedTileValue(MustStrToTile34("5z"), newPI(27, 27, "2s"))), eps)
	assert.InDelta(t, 115, float64(calculateIsolatedTileValue(MustStrToTile34("5z"), newPI(27, 27, "2s"))), eps)
	assert.InDelta(t, 114, float64(calculateIsolatedTileValue(MustStrToTile34("6z"), newPI(27, 27, "2s"))), eps)
	assert.InDelta(t, 115, float64(calculateIsolatedTileValue(MustStrToTile34("7z"), newPI(27, 27, "2s"))), eps)
	assert.InDelta(t, 103.5, float64(calculateIsolatedTileValue(MustStrToTile34("7z"), newPI(27, 27, "2s77z"))), eps)
	assert.InDelta(t, 23, float64(calculateIsolatedTileValue(MustStrToTile34("7z"), newPI(27, 27, "2s777z"))), eps)

	assert.InDelta(t, 114, float64(calculateIsolatedTileValue(MustStrToTile34("1z"), newPI(29, 27, "2s"))), eps)
	assert.InDelta(t, 102.6, float64(calculateIsolatedTileValue(MustStrToTile34("1z"), newPI(29, 27, "2s11z"))), eps)
	assert.InDelta(t, 99, float64(calculateIsolatedTileValue(MustStrToTile34("2z"), newPI(29, 27, "2s"))), eps)
	assert.InDelta(t, 116, float64(calculateIsolatedTileValue(MustStrToTile34("3z"), newPI(29, 27, "2s"))), eps)
	assert.InDelta(t, 97, float64(calculateIsolatedTileValue(MustStrToTile34("4z"), newPI(29, 27, "2s"))), eps)
}
