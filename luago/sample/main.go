package main

import (
	"fmt"
	"luago/api"
	"luago/state"
	"os"
)

func print(ls api.LuaState) int {
	nArgs := ls.GetTop()
	for i := 1; i <= nArgs; i++ {
		if ls.IsBoolean(i) {
			fmt.Printf("%t", ls.ToBoolean(i))
		} else if ls.IsString(i) {
			fmt.Print(ls.ToString(i))
		} else {
			fmt.Print(ls.TypeName(ls.Type(i)))
		}
		if i < nArgs {
			fmt.Print("\t")
		}
	}
	fmt.Println()
	return 0
}
func main() {
	filename := "../lua/h.luac"
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	ls := state.New()
	ls.Register("print", print)
	ls.Load(data, filename, "b")
	ls.Call(0, 0)
}
