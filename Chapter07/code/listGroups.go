package main

import (
	"fmt"
	"os"
	"os/user"
	"sort"
	"strconv"
)

func main() {
	arguments := os.Args
	var u *user.User
	var err error
	if len(arguments) == 1 {
		u, err = user.Current()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		username := arguments[1]
		u, err = user.Lookup(username)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	gids, _ := u.GroupIds()
	gids_int := make([]int, len(gids))
	for i, s := range gids {
		gids_int[i], _ = strconv.Atoi(s)
	}

	sort.Ints(gids_int)

	for _, gid := range gids_int {
		group, err := user.LookupGroupId(strconv.Itoa(gid))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s(%s)\n", group.Gid, group.Name)
	}
	fmt.Println()
}
