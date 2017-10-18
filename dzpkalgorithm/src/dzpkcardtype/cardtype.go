/*
Copyright 2017 The TensorFlow Authors. All Rights Reserved.

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

/**
*Copyright (C), 2017.
*@file: cardtype.go (UTF8)
*@brief:
*@author: Billikeu
*@version: 1.0.0
*@date: 2017/10/17
*@warning
*@History:
 */
package dzpkcardtype

import (
	"poker"
	"sort"
)

// card type 牌型
const (
	HIGHCARD     = 0
	ONEPAIR      = 1
	TWOPAIR      = 2
	THREEOFAKING = 3
	STRIGHT      = 4
	FLUSH        = 5
	FULLHOURSE   = 6
	FOUROFAKING  = 7
	FLUSHSTRIGHT = 8
	YORALFLUSH   = 9
)

// is Yoralflush 皇家同花顺
func isYoralFlush(cardGroup *poker.CardGroup) bool {
	cardType, max := isFlushStright(cardGroup)
	if cardType && max == 14 {
		return true
	}
	return false
}

// isFlushStright 同花顺
func isFlushStright(cardGroup *poker.CardGroup) (cardType bool, max int) {
	if cardGroup.Len() < 5 {
		return false, 0
	}
	cardType, color := isFlush(cardGroup)
	cardGroupSameColor := new(poker.CardGroup)
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

// isFourOfAKing 四条
func isFourOfAKing(cardGroup *poker.CardGroup) (cardType bool, max int) {
	if cardGroup.Len() < 4 {
		return false, 0
	}
	mapCard := make(map[int]int)
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

// isFullHourse 葫芦
func isFullHourse(cardGroup *poker.CardGroup) (cardType bool, max int) {
	if cardGroup.Len() < 5 {
		return false, 0
	}
	mapCard := make(map[int]int)
	for _, cd := range cardGroup.Card {
		mapCard[cd.CardNumber]++
	}
	var maxThree int = 0
	for cardNumber, sum := range mapCard {
		if sum == 3 {
			if maxThree < cardNumber {
				maxThree = cardNumber
			}
		}
	}
	if maxThree == 0 {
		return false, 0
	}
	for cardNumber, sum := range mapCard {
		if cardNumber == maxThree {
			continue
		}
		if sum == 2 || sum == 3 {
			return true, maxThree
		}
	}
	return false, 0
}

// isFlush 同花
func isFlush(cardGroup *poker.CardGroup) (cardType bool, color int) {
	if cardGroup.Len() < 5 {
		return false, 0
	}
	mapCard := make(map[int]int)
	for _, cd := range cardGroup.Card {
		mapCard[cd.Color]++
	}
	for cardColor, sum := range mapCard {
		if sum >= 5 {
			return true, cardColor
		}
	}
	return false, 0
}

// isStright 顺子
func isStright(cardGroup *poker.CardGroup) (cardType bool, max int) {
	if cardGroup.Len() < 5 {
		return false, 0
	}
	mapCard := make(map[int]int)
	for _, cd := range cardGroup.Card {
		mapCard[cd.CardNumber]++
	}
	var tmp int = 15
	for i := 14; i > 1; i-- {
		if mapCard[i] == 0 {
			tmp = i
		}
		if tmp-i >= 5 {
			max = tmp - 1
			return true, max
		}
	}
	return false, 0
}

// isThreeOfAKing 三条
func isThreeOfAKing(cardGroup *poker.CardGroup) (cardType bool, max int) {
	if cardGroup.Len() < 3 {
		return false, 0
	}
	mapCard := make(map[int]int)
	for _, cd := range cardGroup.Card {
		mapCard[cd.CardNumber]++
	}
	for cardNumber, sum := range mapCard {
		if sum == 3 {
			return true, cardNumber
		}
	}
	return false, 0
}

// isTwoPair 两队
func isTwoPair(cardGroup *poker.CardGroup) (cardType bool, max int, min int) {
	if cardGroup.Len() < 4 {
		return false, 0, 0
	}
	mapCard := make(map[int]int)
	for _, cd := range cardGroup.Card {
		mapCard[cd.CardNumber]++
	}
	var maxCard, minCard, count int = 0, 0, 0
	for cardNumber, sum := range mapCard {
		if sum == 2 {
			count++
			if maxCard < cardNumber {
				maxCard = cardNumber
			}
		}
	}
	if count < 2 {
		return false, 0, 0
	}
	for cardNumber, sum := range mapCard {
		if cardNumber == maxCard {
			continue
		}
		if sum == 2 {
			if minCard < cardNumber {
				minCard = cardNumber
			}
			return true, maxCard, minCard
		}
	}
	return false, 0, 0
}

// isOnePair一对
func isOnePair(cardGroup *poker.CardGroup) (cardType bool, max int) {
	if cardGroup.Len() < 2 {
		return false, 0
	}
	mapCard := make(map[int]int)
	for _, cd := range cardGroup.Card {
		mapCard[cd.CardNumber]++
	}
	for cardNumber, sum := range mapCard {
		if sum == 2 {
			return true, cardNumber
		}
	}
	return false, 0
}

// GetCardType 获取牌型
func GetCardType(cardGroup *poker.CardGroup) int {

	if isYoralFlush(cardGroup) {
		return YORALFLUSH
	}
	bFlag, _ := isFlushStright(cardGroup)
	if bFlag {
		return FLUSHSTRIGHT
	}
	bFlag, _ = isFourOfAKing(cardGroup)
	if bFlag {
		return FOUROFAKING
	}
	bFlag, _ = isFullHourse(cardGroup)
	if bFlag {
		return FULLHOURSE
	}
	bFlag, _ = isFlush(cardGroup)
	if bFlag {
		return FLUSH
	}
	bFlag, _ = isStright(cardGroup)
	if bFlag {
		return STRIGHT
	}
	bFlag, _ = isThreeOfAKing(cardGroup)
	if bFlag {
		return THREEOFAKING
	}
	bFlag, _, _ = isTwoPair(cardGroup)
	if bFlag {
		return TWOPAIR
	}
	bFlag, _ = isOnePair(cardGroup)
	if bFlag {
		return ONEPAIR
	}
	return HIGHCARD
}

// CompareCardGroups 比较两副手牌大小，1为大于，0为等于，-1小于
func CompareCardGroups(left *poker.CardGroup, right *poker.CardGroup) int {
	leftType := GetCardType(left)
	rightType := GetCardType(right)
	if leftType > rightType {
		return 1
	}
	if leftType < rightType {
		return -1
	}
	return compareTheSameType(left, right, leftType)
}

// compareTheSameType 比价相同牌型大小，1为大于，0为等于，-1小于
func compareTheSameType(left *poker.CardGroup, right *poker.CardGroup, cardType int) int {
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
		_, cardColor := isFlush(left)
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
		return compareCardsNumber(left, right, 0, 5)
	default:
		return 0 //
	}
}

// deleteCardsFrom 从card中剔除deleteNumber
func deleteCardsFrom(deleteNumber int, card *poker.CardGroup) *poker.CardGroup {
	returnCard := new(poker.CardGroup)
	for _, cd := range card.Card {
		if cd.CardNumber == deleteNumber {
			continue
		}
		returnCard.Card = append(returnCard.Card, cd)
	}
	return returnCard
}

// compareCardsNumber 比较点数大小，1为大于，0为等于，-1小于
func compareCardsNumber(leftCard *poker.CardGroup, rightCard *poker.CardGroup, color int, number int) int {
	if leftCard.Len() < number || rightCard.Len() < number {
		return -2
	}
	sort.Sort(leftCard)
	sort.Sort(rightCard)

	if color == 0 {
		for i := 0; i < number; i++ {
			if leftCard.Card[i].CardNumber > rightCard.Card[i].CardNumber {
				return 1
			}
			if leftCard.Card[i].CardNumber < rightCard.Card[i].CardNumber {
				return -1
			}
		}
		return 0
	} else {
		for i := 0; i < number; i++ {
			if leftCard.Card[i].Color != color {
				continue
			}
			if leftCard.Card[i].CardNumber > rightCard.Card[i].CardNumber {
				return 1
			}
			if leftCard.Card[i].CardNumber < rightCard.Card[i].CardNumber {
				return -1
			}
		}
		return 0
	}
}
