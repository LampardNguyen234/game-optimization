package main

import (
	"github.com/LampardNguyen234/game-optimization/character"
	"math"
	"math/rand"
	"time"
)

func RandInt() int {
	// Use this function for testing-purposes only.
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

func Hit(A, B character.Character) {
	randCritical := RandInt() % 100
	isCritical := randCritical < A.CriticalRate()
	criticalRate := 2
	if !isCritical {
		criticalRate = 1
	}

	B.Hit(criticalRate * A.Damage())
}

func Combat(A, B character.Character) int {
	t := 0
	stepA := int(math.Ceil(10000 / float64(A.AttackSpeed())))
	stepB := int(math.Ceil(10000 / float64(B.AttackSpeed())))

	//log.Printf("[BEFORE] A.HP %v, B.HP %v\n", A.HitPoints(), B.HitPoints())
	for A.HitPoints() > 0 && B.HitPoints() > 0 {
		t++
		if t%stepA == 0 {
			Hit(A, B)
			//log.Printf("[%v] Hit(A, B): %v %v\n", t, A.HitPoints(), B.HitPoints())
		}
		if B.HitPoints() == 0 {
			break
		}
		if t%stepB == 0 {
			Hit(B, A)
			//log.Printf("[%v] Hit(B, A): %v %v\n", t, A.HitPoints(), B.HitPoints())
		}

	}

	//log.Printf("[AFTER] A.HP %v, B.HP %v\n", A.HitPoints(), B.HitPoints())

	if A.HitPoints() == 0 {
		return 1
	} else {
		return 0
	}
}
