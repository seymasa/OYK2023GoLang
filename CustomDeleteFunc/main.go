/*
1. key boş string ise boş key gönderme diyecek.
2. elimde valid key var value okay syntax ı ile bakıcam bu gelen value var mı bu map te
3. varsa bize gelen key den kaldır o kaydı.
4. eger yoksa error döncez.
5. kontrol et bide nil map var mı varsa ona da error bas.
map nil
*/
package main

import (
	"fmt"
	"log"
)

func Mydelete(m map[string]string, key string) error {
	if m == nil {
		log.Println("Error: map is nil")
	}
	if key == "" {
		log.Println("Error: key is an empty string")
	}
	if _, ok := m[key]; !ok {
		log.Println("Error: not found! ", key)
	}
	delete(m, key)

	return nil
}

func main() {

	map1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(map1)
	Mydelete(map1, "key1")
	fmt.Println(map1)
}
