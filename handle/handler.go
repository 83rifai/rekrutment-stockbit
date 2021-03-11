package handle

import (
	"stockbit/cnf"
	"stockbit/cnf/databases"
	"stockbit/cnf/env"
	HandleAnagram "stockbit/module/anagram/presenter"
	HandleImdb "stockbit/module/imdb/presenter"
)

type Service struct {
	AnagramHandle *HandleAnagram.HTTPAnagramHandler
	ImdbHandle    *HandleImdb.HTTPImdbHandler
}

func MakeHandler() *Service {

	env.LoadEnv()

	DbConn, err := databases.MysqlDB()
	if err != nil {
		panic("error database")
	}

	configuration := cnf.Config{
		DB: DbConn,
	}

	anagramHandle := HandleAnagram.NewHTTPHandler(configuration)
	imdbHandle := HandleImdb.NewHTTPHandler(configuration)

	return &Service{
		AnagramHandle: anagramHandle,
		ImdbHandle:    imdbHandle,
	}
}
