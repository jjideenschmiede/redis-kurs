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

	// Zur Datenbank verbinden
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	// Einen string erstellen
	err := rdb.Set("client:kwiedor", "Jonas Kwiedor", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// String auslesen
	client, err := rdb.Get("client:kwiedor").Result()
	if err != nil {
		fmt.Println(err)
	}

	// Client ausgeben
	fmt.Println(client)

	// Hash erstellen
	fields := make(map[string]interface{})

	fields["name"] = "Jonas Kwiedor"
	fields["age"] = 23
	fields["job"] = "Developer"

	err = rdb.HMSet("client:kwiedor2", fields).Err()
	if err != nil {
		fmt.Println(err)
	}

	// Feld aus hash entfernen
	err = rdb.HDel("client:kwiedor", "age").Err()
	if err != nil {
		fmt.Println(err)
	}

}
