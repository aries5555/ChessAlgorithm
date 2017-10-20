德州扑克牌型判断
==

德克萨斯扑克全称Texas Hold’em poker，中文简称德州扑克。它是一种玩家对玩家的公共牌类游戏。一张台面至少2人，最多22人，一般是由2-10人参加。在各种系列的赌神电影中看到的比赛扑克就是这种德州扑克，霸气的AllIn（全押）。

下面将要介绍的就是在德州扑克游戏开发中经常遇到的扑克牌型判断的Go语言算法，算法可能可能不是最优的，还请多多指教。

## 正文

### 定义

注释中是各种牌型的说明。
```
// card type 牌型
const (
    HIGHCARD     = 0 //高牌：不是下面9种牌型中的任何一种
    ONEPAIR      = 1 //一对：手牌+公牌有两张点数一样的牌，并且不是下面8种牌型中的任何一种
    TWOPAIR      = 2 //两队：手牌+公牌可以组成两队，并且不是下面7种牌型中的任何一种
    THREEOFAKING = 3 //三条：手牌+公牌有三张点数一样的牌，并且不是下面6种牌型中的任何一种
    STRIGHT      = 4 //顺子：手牌+公牌至少有5张点数连续，并且不是下面5种牌型中的任何一种
    FLUSH        = 5 //同花：手牌+公牌至少5张花色连续，并且不是下面4种牌型中的任何一种
    FULLHOURSE   = 6 //葫芦：手牌+公牌至少有一个三条和一对，并且不是下面3种牌型中的任何一种
    FOUROFAKING  = 7 //四条：手牌+公牌有4张花色一样，并且不是下面2种牌型中的任何一种
    FLUSHSTRIGHT = 8 //同花顺：手牌+公牌至少有5张点数连续并且花色一样，并且不是下面1种牌型
    YORALFLUSH   = 9 //皇家同花顺：手牌+公牌可以同花顺，并且最大牌为A
)
```

### 获取牌型

通过下面的牌型判断算法获取牌型

```
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

```
### 皇家同花顺

皇家同花顺：手牌+公牌可以同花顺，并且最大牌为A

1. 先检测该牌型是不是同花顺。
2. 再判断最大点数是否为A

```
// is Yoralflush 皇家同花顺
func isYoralFlush(cardGroup *poker.CardGroup) bool {
    cardType, max := isFlushStright(cardGroup)
    if cardType && max == 14 {
        return true
    }
    return false
}
```
### 同花顺
同花顺：手牌+公牌至少有5张点数连续并且花色一样，并且不是下面1种牌型
1. 判断是否为同花，如果是返回可组成的花色
2. 判断该花色的牌是否可以组成顺子
```
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
```
### 四条

四条：手牌+公牌有4张花色一样，并且不是下面2种牌型中的任何一种
1. 按照牌点数放入map中，累计点数出现的次数
2. map中有点数出现四次，说明为四条

```
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
```
### 葫芦

葫芦：手牌+公牌至少有一个三条和一对，并且不是下面3种牌型中的任何一种

1. 按照牌点数放入map中，累计点数出现的次数
2. 统计map中点数出现的次数，出现一次3，一次2或者两次3

```
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
```
### 同花

同花：手牌+公牌至少5张花色连续，并且不是下面4种牌型中的任何一种

1. 按照花色存入map中
2. 花色数量大于等于5为真

```
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
```
### 顺子

顺子：手牌+公牌至少有5张点数连续，并且不是下面5种牌型中的任何一种

1. 按照牌点数放入map中，累计点数出现的次数
2. 判断点数是否有连续5个或以上

```
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
```
### 三条

三条：手牌+公牌有三张点数一样的牌，并且不是下面6种牌型中的任何一种

1. 按照牌点数放入map中，累计点数出现的次数
2. map中有点数出现3次，说明为3条

```
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
```
### 两队

两队：手牌+公牌可以组成两队，并且不是下面7种牌型中的任何一种

1. 按照牌点数放入map中，累计点数出现的次数
2. map中有点数出现2次的个数有两个

```
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
```
### 一对

一对：手牌+公牌有两张点数一样的牌，并且不是下面8种牌型中的任何一种

1. 按照牌点数放入map中，累计点数出现的次数
2. map中有点数出现2次的个数有1个 

```
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
```
### 高牌

高牌：不是上面9种牌型中的任何一种

## 结尾
完整的代码请参考我的[GitHub](https://github.com/billikeu/ChessAlgorithm)！