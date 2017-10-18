package poker

import (
	"fmt"
)

// 牌结构
type PokerCard struct {
	Color      int // 花色
	CardNumber int // 点数
}

// 牌组
type CardGroup struct {
	Card []*PokerCard // 牌组
}

// Cards are divided into four kinds: spades, diamonds, clubs and hearts.
const (
	HEARTS   = 1 // 红桃
	SPADES   = 2 // 黑桃
	CLUBS    = 3 // 梅花
	DIAMONDS = 4 // 方块

	KING  = 21 //王
	QUEEN = 22 //后
	PIZI  = 30 //皮子
)

func (cardGroup *CardGroup) Len() int {
	return len(cardGroup.Card)
}

// 大到小的顺序排序
func (cardGroup *CardGroup) Swap(i, j int) {
	cardGroup.Card[i], cardGroup.Card[j] = cardGroup.Card[j], cardGroup.Card[i]
}
func (cardGroup *CardGroup) Less(i, j int) bool {
	return cardGroup.Card[i].CardNumber > cardGroup.Card[j].CardNumber
}
func (cardGroup *CardGroup) Deal(number int) *CardGroup {
	dealCardGroup := new(CardGroup)
	dealCardGroup.Card = cardGroup.Card[:number]
	cardGroup.Card = cardGroup.Card[number:]
	return dealCardGroup
}
func (cardGroup *CardGroup) Print() {
	for index, card := range cardGroup.Card {
		fmt.Printf("index=%02d: N=%02d, C=%02d\n", index, card.CardNumber, card.Color)
	}
}

// paizi：皮子个数；king：0无，1有，是否有大小王；begin：A开始值;index：几副牌
func MakeRandomCards(paizi, king, begin, index int) *CardGroup {

	cards := new(CardGroup)
	for k := 0; k < index; k++ {

		for j := 1; j < 5; j++ {
			for i := begin; i < begin+13; i++ {
				card := new(PokerCard)
				card.CardNumber = i
				card.Color = j
				cards.Card = append(cards.Card, card)
			}
		}
		for i := 0; i < paizi; i++ {
			card := new(PokerCard)
			card.CardNumber = PIZI + i
			card.Color = PIZI
			cards.Card = append(cards.Card, card)
		}
		if king > 0 {
			card := new(PokerCard)
			card.CardNumber = KING
			card.Color = KING
			cards.Card = append(cards.Card, card)
			card.CardNumber = QUEEN
			card.Color = QUEEN
			cards.Card = append(cards.Card, card)
		}
	}
	ShuffleTimes(cards, 3)
	return cards
}

// Shuffle
func Shuffle(cards *CardGroup) {
	mcard := make(map[int]*PokerCard)
	for index, card := range cards.Card {
		mcard[index] = card
	}
	cards.Card = make([]*PokerCard, 0)
	for _, card := range mcard {
		cards.Card = append(cards.Card, card)
	}
}

func ShuffleTimes(cards *CardGroup, times int) {
	for i := 0; i < times; i++ {
		Shuffle(cards)
	}
}

func TestMakeCards() {
	cards := MakeRandomCards(0, 1, 1, 1)
	for index, card := range cards.Card {
		fmt.Printf("index=%02d:N=%02d,C=%02d\n", index, card.CardNumber, card.Color)
	}
}
func TestVoid() {
}
