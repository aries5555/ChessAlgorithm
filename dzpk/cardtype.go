/*
Copyright 2016 The TensorFlow Authors. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package dzpk

import (
	"sort"
)

// 牌结构
type PokerCard struct {
	Color      int32 //花色
	CardNumber int32 //点数
}

// 牌组
type CardGroup struct {
	Card []*PokerCard //牌组
}

const (
	HIGHCARD = iota
	ONEPAIR
	TWOPAIR
	THREEOFAKING
	STRIGHT
	FLUSH
	FULLHOURSE
	FOUROFAKING
	FLUSHSTRIGHT
	YORALFLUSH
)

func isYoralFlush(cardGroup *CardGroup) bool {
	cardType, max := isFlushStright(cardGroup)
	if cardType && max == 14 {
		return true
	}
	return false
}
func isFlushStright(cardGroup *CardGroup) (cardType bool, max int32) {
	if cardGroup.Len() < 5 {
		return false
	}
	cardType, color := isFlush(cardGroup)
	cardGroupSameColor := new(CardGroup)
	if cardType {
		for _, cd := range cardGroup.Card {
			if cd.Color != color {
				continue
			}
			cardGroupSameColor.Card = append(cardGroupSameColor.Card, cd)
		}
	}
	return isStright(cardGroupSameColor)
}

func isFourOfAKing(cardGroup *CardGroup) (cardType bool, max int32) {
	if cardGroup.Len() < 4 {
		return false, 0
	}
	mapCard := new(map[int32]int32)
	for _, cd := range cardGroup.Card {
		mapCard[cd.CardNumber]++
	}
	for cardNumber, sum := range mapCard {
		if sum == 4 {
			return true, cardNumber
		}
	}
	return false, 0
}
func isFullHourse(cardGroup *CardGroup) bool {
	if cardGroup.Len() < 5 {
		return false
	}

	return false
}
func isFlush(cardGroup *CardGroup) (cardType bool, color int32) {
	return false
}
func isStright(cardGroup *CardGroup) bool {
	return false
}
func isThreeOfAKing(cardGroup *CardGroup) bool {
	return false
}
func isTwoPair(cardGroup *CardGroup) bool {
	return false
}
func isOnePair(cardGroup *CardGroup) bool {
	return false
}
func isHighCard(cardGroup *CardGroup) bool {
	return false
}

func getCardType(cardGroup *CardGroup) int32 {
	if isYoralFlush(cardGroup) {
		return YORALFLUSH
	}
	if isFlushStright(cardGroup) {
		return FLUSHSTRIGHT
	}
	if isFourOfAKing(cardGroup) {
		return FOUROFAKING
	}
	if isFullHourse(cardGroup) {
		return FULLHOURSE
	}
	if isFlush(cardGroup) {
		return FLUSH
	}
	if isStright(cardGroup) {
		return STRIGHT
	}
	if isThreeOfAKing(cardGroup) {
		return THREEOFAKING
	}
	if isTwoPair(cardGroup) {
		return TWOPAIR
	}
	if isOnePair(cardGroup) {
		return ONEPAIR
	}
	return HIGHCARD
}

func compareCardGroups(left *CardGroup, right *CardGroup) int32 {
	leftType := getCardType(left)
	rightType := getCardType(right)
	if leftType > rightType {
		return 1
	}
	if leftType < rightType {
		return -1
	}
	return compareTheSameType(left, right, leftType)
}
func compareTheSameType(left *CardGroup, right *CardGroup, cardType int32) int32 {
	switch cardType {
	case YORALFLUSH:
		return 0
	case FLUSHSTRIGHT:
		_, leftMax := isFlushStright(left)
		_, rightMax := isFlushStright(right)
		if leftMax > rightMax {
			return 1
		}
		if leftMax < rightMax {
			return -1
		}
		return 0
	case FOUROFAKING:
		_, leftMax := isFourOfAKing(left)
		_, rightMax := isFourOfAKing(right)
		if leftMax > rightMax {
			return 1
		} else {
			return -1
		}
	case FULLHOURSE:
		_, leftMax := isFullHourse(left)
		_, rightMax := isFullHourse(right)
		if leftMax > rightMax {
			return 1
		} else {
			return -1
		}
	case FLUSH:
		return compareCardsNumber(left, right, cardColor, 5)
	case STRIGHT:
		_, leftMax := isStright(left)
		_, rightMax := isStright(right)
		if leftMax > rightMax {
			return 1
		}
		if leftMax < rightMax {
			return -1
		}
		return 0
	case THREEOFAKING:
		_, leftMax := isThreeOfAKing(left)
		_, rightMax := isThreeOfAKing(right)
		if leftMax > rightMax {
			return 1
		} else {
			return -1
		}
	case TWOPAIR:
		_, leftMax, leftMin := isTwoPair(left)
		_, rightMax, rightMin := isTwoPair(right)
		if leftMax > rightMax {
			return 1
		}
		if leftMax < rightMax {
			return -1
		}
		if leftMin > rightMin {
			return -1
		}
		if leftMin < rightMin {
			return -1
		}
		leftTmpCards := deleteCardsFrom(leftMax, left)
		leftTmpCards = deleteCardsFrom(leftMin, leftTmpCards)
		rightTmpCards := deleteCardsFrom(rightMax, right)
		rightTmpCards = deleteCardsFrom(rightMin, rightTmpCards)
		return compareCardsNumber(leftTmpCards, rightTmpCards, 0, 1)
	case ONEPAIR:
		_, leftMax := isOnePair(left)
		_, rightMax := isOnePair(right)
		if leftMax > rightMax {
			return 1
		}
		if leftMax < rightMax {
			return -1
		}
		leftTmpCards := deleteCardsFrom(leftMax, left)
		rightTmpCards := deleteCardsFrom(rightMax, right)
		return compareCardsNumber(leftTmpCards, rightTmpCards, 0, 3)
	case HIGHCARD:
		return compareCardsNumber(leftTmpCards, rightTmpCards, 0, 5)
	default:
		return 0 //
	}
}
func deleteCardsFrom(deleteNumber int32, card *CardGroup) *CardGroup {
	returnCard := new(CardGroup)
	for _, cd := range card.Card {
		if cd.CardNumber == deleteNumber {
			continue
		}
		returnCard.Card = append(returnCard.Card, cd)
	}
	return returnCard
}
func compareCardsNumber(leftCard *CardGroup, rightCard *CardGroup, color int32, number int32) int32 {
	if leftCard.Card.len() < number || rightCard.Card.len() < number {
		return -2
	}
	sort.Sort(leftCard)
	sort.Sort(rightCard)

	if color == 0 {
		for i := 0; i < number; i++ {
			if leftCard.Card[i] > rightCard.Card[i] {
				return 1
			}
			if leftCard.Card[i] < rightCard.Card[i] {
				return -1
			}
		}
		return 0
	} else {
		for i := 0; i < number; i++ {
			if leftCard.Card[i].Color != color {
				continue
			}
			if leftCard.Card[i] > rightCard.Card[i] {
				return 1
			}
			if leftCard.Card[i] < rightCard.Card[i] {
				return -1
			}
		}
		return 0
	}
}
