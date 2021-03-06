package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	APPLE = "ð"
	STAR  = "ð"
	BELL  = "ð"
	TREE  = "ð"
	DOOR  = "ðª"
	GIFT  = "ð"
)

func main() {
	//åæ¥ä¸ª5å±é«ç
	ct := newChristmasTree(4)
	fmt.Print(ct)
}

type ChristmasTree struct {
	floor       int
	treeBuilder strings.Builder
}

func (ct *ChristmasTree) String() string {
	return ct.treeBuilder.String()
}

func newChristmasTree(floor int) *ChristmasTree {
	ct := &ChristmasTree{floor: floor}
	ct.gen()
	return ct
}

//è·åç¬¬floorå±çç¬¬lineè¡çå£è¯æ æ°é
func getLineAmount(floor, line int) int {
	return 1 + line*2 + floor*4 + int(floor/2*2)*int((floor+1)/2)
}

//éæºææ¯ä¾åéææãééãè¹æåå£è¯æ 
func randAppleTree() string {
	r := rand.Intn(100)
	if r < 1 {
		return STAR
	} else if r < 2 {
		return BELL
	} else if r < 10 {
		return APPLE
	} else {
		return TREE
	}
}

//çæä¸æ´é¢å£è¯æ 
func (ct *ChristmasTree) gen() {
	bottomAmount := getLineAmount(ct.floor, ct.floor+4)

	//ä¸å±ï¼ä¸è¡ççæ
	for floor := 0; floor < ct.floor; floor++ {
		for line := 0; line < floor+5; line++ {
			lineAmount := getLineAmount(floor, line)

			for i := (bottomAmount-lineAmount)/2 - 1; i > 0; i-- {
				ct.treeBuilder.WriteString("  ")
			}

			for i := 0; i < lineAmount; i++ {
				ct.treeBuilder.WriteString(randAppleTree())
			}

			ct.treeBuilder.WriteString("\n")
		}
	}

	//å±ä¸­ãçæå£è¯æ æ ¹
	for i := 0; i < ct.floor; i++ {
		lineAmount := ct.floor + (ct.floor+1)%2 //ä¸ä¸ªæ´æ¥è¿å±æ°çè¿ä¼¼å¼

		for i := (bottomAmount-lineAmount)/2 - 1; i > 0; i-- {
			ct.treeBuilder.WriteString("  ")
		}

		for i := 0; i < lineAmount; i++ {
			ct.treeBuilder.WriteString(DOOR)
		}
		ct.treeBuilder.WriteString("\n")
	}

	//å¨å£è¯æ ä¸æ¾ç¹ç¤¼ç©
	if ct.floor > 3 {
		//ä¸»è¦æ¯ä»åæ°ç¬¬2è¡å¼å§æ¾
		lines := strings.Split(ct.treeBuilder.String(), "\n")
		lines[len(lines)-4] += "     " + GIFT
		lines[len(lines)-3] += "    " + GIFT + GIFT + GIFT
		lines[len(lines)-2] += "   " + GIFT + GIFT + GIFT + GIFT + GIFT
		ct.treeBuilder.Reset()
		ct.treeBuilder.WriteString(strings.Join(lines, "\n"))
	}
}
