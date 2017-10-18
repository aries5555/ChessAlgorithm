package dzpktest

import (
	"dzpkcardtype"
	"fmt"
	"poker"
	"sort"
)

func TestVoid() {

}
func TestCardType() {
	index := 0
	for {
		var CARDTYPE int = -1
		for {
			fmt.Println("please input cardtype:")
			fmt.Scanln(&CARDTYPE)
			if CARDTYPE > -1 {
				break
			}
		}
		for {
			cardRandomGroup := poker.MakeRandomCards(0, 1, 1, 1)
			cardRandomGroup.Print()
			for i := 0; i < 6; i++ {
				cardsGroup := cardRandomGroup.Deal(5)
				cardType := dzpkcardtype.GetCardType(cardsGroup)
				if cardType != CARDTYPE {
					continue
				}
				fmt.Printf("T=%02d:\n", cardType)
				sort.Sort(cardsGroup)
				cardsGroup.Print()
				index++
				if index == 5 {
					break
				}
				fmt.Printf("--------------\n")
			}
			if index == 5 {
				index = 0
				break
			}
		}

	}

}
