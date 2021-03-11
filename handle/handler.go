package handle

import (
	"stockbit/cnf"
	"stockbit/cnf/databases"
	"stockbit/cnf/env"
	HandleAnagram "stockbit/module/anagram/presenter"
)

type Service struct {
	AnagramHandle *HandleAnagram.HTTPAnagramHandler
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

	return &Service{
		AnagramHandle: anagramHandle,
	}
}
