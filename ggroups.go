// ggroups is a Go-equivalent for groups(1), used to validate go-->libc behavior.
package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	u, err := user.Lookup(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("lookup failed: %v", err))
	}
	gids, err := u.GroupIds()
	if err != nil {
		panic(fmt.Sprintf("groupids failed: %v", err))
	}
	fmt.Println(strings.Join(gids, " "))
}
