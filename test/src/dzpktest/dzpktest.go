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
package dzpktest

import (
	"dzpkcardtype"
	"fmt"
	"poker"
	"sort"
)

func TestVoid() {
	cards := new(poker.CardGroup)

	card2 := new(poker.PokerCard)
	card2.CardNumber = 2
	card2.Color = 3
	cards.Card = append(cards.Card, card2)

	card3 := new(poker.PokerCard)
	card3.CardNumber = 2
	card3.Color = 2
	cards.Card = append(cards.Card, card3)

	card4 := new(poker.PokerCard)
	card4.CardNumber = 5
	card4.Color = 4
	cards.Card = append(cards.Card, card4)

	card5 := new(poker.PokerCard)
	card5.CardNumber = 9
	card5.Color = 4
	cards.Card = append(cards.Card, card5)

	card1 := new(poker.PokerCard)
	card1.CardNumber = 2
	card1.Color = 1
	cards.Card = append(cards.Card, card1)

	cards.Print()
	cardType := dzpkcardtype.GetCardType(cards)
	fmt.Printf("cardType=%02d", cardType)
}
func TestCardType() {
	count := 5

	for {
		var CARDTYPE int = -1
		for {
			fmt.Println("please input cardtype:")
			fmt.Scanln(&CARDTYPE)
			if CARDTYPE > -1 {
				break
			}
		}
		index := 0
		for {
			// paizi：皮子个数；king：0无，1有，是否有大小王；begin：A开始值;index：几副牌
			cardRandomGroup := poker.MakeRandomCards(0, 1, 2, 1)
			for i := 0; i < 6; i++ {
				cardsGroup := cardRandomGroup.Deal(5)
				cardType := dzpkcardtype.GetCardType(cardsGroup)
				if cardType != CARDTYPE {
					continue
				}
				index++
				sort.Sort(cardsGroup)
				cardsGroup.Print()
				if index > count {
					break
				}
				fmt.Printf("--------------\n")
			}
			if index > count {
				break
			}
		}

	}

}
func TestCardTypeRate() {

	count := 50000
	for a := 0; a < 10; a++ {
		index, sum := 0, 0
		for {
			// paizi：皮子个数；king：0无，1有，是否有大小王；begin：A开始值;index：几副牌
			cardRandomGroup := poker.MakeRandomCards(0, 0, 2, 1)
			for i := 0; i < 7; i++ {
				if sum > count {
					break
				}
				cardsGroup := cardRandomGroup.Deal(7)
				sum++
				if sum == count {
					fmt.Printf("Type=%d,T=%02d,e=%.6f\n", a, sum, float32(index)/float32(sum))
				}
				cardType := dzpkcardtype.GetCardType(cardsGroup)
				if cardType != a {
					continue
				}
				index++

			}
			if sum > count {
				break
			}
		}
	}
}
