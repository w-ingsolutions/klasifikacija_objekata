package app

import (
	"fmt"
)

func Ispis(klasa Str) {
	fmt.Println(klasa.Text)
	fmt.Println(klasa.Opis)
	if len(klasa.Pod) > 0 {
		for i := range klasa.Pod {
			Ispis(klasa.Pod[i])
		}
	}
}
