//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of redis-kurs.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	err := rdb.Set("client:kwiedor", "Jonas Kwiedor", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	client, err := rdb.Get("client:kwiedor").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client)

}
