package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eklipsan/forumeather/pkg/config"
	"github.com/eklipsan/forumeather/pkg/storage"
)

const (
	addForumNumber = iota
	deleteForumNumber
	showAllForumsNumber
)

func main() {
	var (
		operation int
	)
	forumDB, err := storage.NewForumDB(getPathDB())
	if err != nil {
		panic(err)
	}
	fmt.Println("Добавить форум: 0")
	fmt.Println("Удалить форум: 1")
	fmt.Println("Показать все форумы: 2")
	fmt.Println("Введите код операции:")
	fmt.Scanln(&operation)
	switch operation {
	case addForumNumber:
		newForum := scanNewForum()
		err = forumDB.AddForum(newForum)
		if err != nil {
			panic(err)
		}
		fmt.Println("Форум добавлен")
		fmt.Println(newForum)
	case deleteForumNumber:
		var deleteId int
		fmt.Println("Введите ID форума для удаления: ")
		fmt.Scan(&deleteId)

		err := forumDB.DeleteForum(deleteId)
		if err != nil {
			panic(err)
		}
		fmt.Println("Запись о форуме была удалена")
	case showAllForumsNumber:
		var forums []storage.Forum
		forums, err := forumDB.GetAllForums()
		if err != nil {
			panic(err)
		}
		for _, forum := range forums {
			fmt.Println(forum)
		}
	default:
		fmt.Println("Неверный код операции")
	}
	main()
}

func getPathDB() string {
	config := config.WebConfig
	return config.DatabaseConfig.PathDB
}


func scanNewForum() storage.Forum {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите название: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите место: ")
	scanner.Scan()
	place := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите темы: ")
	scanner.Scan()
	topics := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите широту (latitude): ")
	scanner.Scan()
	lat, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

	fmt.Print("Введите долготу (longitude): ")
	scanner.Scan()
	lon, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

	return storage.Forum{
		Name:      name,
		Place:     place,
		Topics:    topics,
		Latitude:  lat,
		Longitude: lon,
	}
}