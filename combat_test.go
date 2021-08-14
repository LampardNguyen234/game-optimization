package main

import (
	"github.com/LampardNguyen234/game-optimization/character"
	"github.com/incognitochain/go-incognito-sdk-v2/common"
	"log"
	"testing"
)

func TestCombat(t *testing.T) {
	countW := 0
	countA := 0

	for i := 0; i < 100; i+= 1 {
		isFlipped := (common.RandInt() % 2) == 1
		if i % 100 == 0 {
			log.Printf("TEST COMBAT %v\n", i)
		}
		//log.Printf("COMBAT %v\n", i)
		w := character.NewWizard(115, 4, 1, 1, 0)
		a := character.NewAssassin(130,1, 2,1, 0)

		//log.Println(w.Invariant(), a.Invariant())

		var res int
		if isFlipped {
			res = Combat(a, w)
			if res == 1 {
				countW++
			} else {
				countA++
			}
		} else {
			res = Combat(w, a)
			if res == 0 {
				countW++
			} else {
				countA++
			}
		}

		//log.Printf("CONBAT %v RESULT: %v, [%v : %v]\n\n", i, res, countW, countA)
	}

	log.Printf("COMBAT FINISHED: [%v : %v]\n", countW, countA)

}
