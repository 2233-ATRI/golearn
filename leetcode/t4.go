package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s string
	fmt.Fscanln(reader, &s)

	var zt, sg []string
	for i := 0; i < len(s); i++ {
		if i+2 <= len(s) && s[i:i+2] == "10" {
			if len(zt) <= len(sg) {
				zt = append(zt, "10")
			} else {
				sg = append(sg, "10")
			}
			i++
		} else {
			if len(zt) <= len(sg) {
				zt = append(zt, string(s[i]))
			} else {
				sg = append(sg, string(s[i]))
			}
		}
	}

	table := []string{}

	cur := 0
	for len(zt) > 0 && len(sg) > 0 {
		var card string
		if cur == 0 {
			card = zt[0]
			zt = zt[1:]
		} else {
			card = sg[0]
			sg = sg[1:]

		}

		table = append(table, card)

		if card == "J" {
			if len(table) > 1 {
				if cur == 0 {
					zt = append(zt, table...)

				} else {
					sg = append(sg, table...)
				}
				table = []string{}
			} else {

			}
		}
		if len(table) >= 2 && table[len(table)-2] == table[len(table)-1] {
			if cur == 0 {
				zt = append(zt, table...)
			} else {
				sg = append(sg, table...)
			}
			table = []string{}

		}
		cur = 1 - cur
	}
	if len(zt) == 0 {
		fmt.Println(writer, "sangiu")
		fmt.Println(writer, strings.Join(sg, " "))
	} else {
		fmt.Println(writer, "ztran")
		fmt.Println(writer, strings.Join(zt, " "))
	}
}
