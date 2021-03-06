package util

import (
	"testing"
	"github.com/EndlessCheng/mahjong-helper/util/model"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_search13(t *testing.T) {
	humanTiles := "5555m"
	humanTiles = "55678m 3467p 2466s"
	tiles34 := MustStrToTiles34(humanTiles)
	pi := model.NewSimplePlayerInfo(tiles34, nil)
	shanten := CalculateShanten(tiles34)
	fmt.Println(NumberToChineseShanten(shanten))
	fmt.Print(_search13(shanten, pi, shanten-1))
}

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

func Test_searchShanten14(t *testing.T) {
	humanTiles := "466m 234467p 77s"
	tiles34 := MustStrToTiles34(humanTiles)
	pi := model.NewSimplePlayerInfo(tiles34, nil)
	shanten := CalculateShanten(tiles34)
	fmt.Println(NumberToChineseShanten(shanten))
	fmt.Print(searchShanten14(shanten, pi, -1))
	fmt.Println("倒退回" + NumberToChineseShanten(shanten+1))
	fmt.Print(searchShanten14(shanten+1, pi, -1))
}

func BenchmarkSearchShanten0(b *testing.B) {
	tiles34 := MustStrToTiles34("234788m 234567s 33z")
	shanten := CalculateShanten(tiles34)
	pi := model.NewSimplePlayerInfo(tiles34, nil)
	for i := 0; i < b.N; i++ {
		// 11,536 ns/op
		searchShanten14(shanten, pi, -1)
	}
}

func BenchmarkSearchShanten1(b *testing.B) {
	tiles34 := MustStrToTiles34("33455m 668p 345667s")
	shanten := CalculateShanten(tiles34)
	pi := model.NewSimplePlayerInfo(tiles34, nil)
	for i := 0; i < b.N; i++ {
		// 361,680 ns/op
		searchShanten14(shanten, pi, -1)
	}
}

func BenchmarkSearchShanten2(b *testing.B) {
	tiles34 := MustStrToTiles34("55678m 3467p 24668s")
	shanten := CalculateShanten(tiles34)
	pi := model.NewSimplePlayerInfo(tiles34, nil)
	for i := 0; i < b.N; i++ {
		// 19,343,607 ns/op
		searchShanten14(shanten, pi, -1)
	}
}

func BenchmarkSearchShanten3(b *testing.B) {
	tiles34 := MustStrToTiles34("12688m 33579p 24s 56z")
	shanten := CalculateShanten(tiles34)
	pi := model.NewSimplePlayerInfo(tiles34, nil)
	for i := 0; i < b.N; i++ {
		// 92,369,360 ns/op
		searchShanten14(shanten, pi, -1)
	}
}
