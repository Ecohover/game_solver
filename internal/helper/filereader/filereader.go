package filereader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"game_solver/internal/model"
)

func ReadFile(path string) string {
	file, err := ioutil.ReadFile("./configs/config.json")
	check(err)

	data := model.Config{}

	_ = json.Unmarshal([]byte(file), &data)

	return string(data.Item.Dice[1].Value[0])
}

func check(e error) {
	if e != nil {
		fmt.Println(string(e.Error()))
		panic(e)
	}
}
