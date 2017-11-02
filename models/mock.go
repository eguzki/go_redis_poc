package models

import (
	"log"
	"time"

	"github.com/mediocregopher/radix.v2/pool"
	cache "github.com/patrickmn/go-cache"
)

var db *pool.Pool
var c *cache.Cache

func init() {
	var err error
	db, err = pool.New("tcp", "redis-storage.benchmark.3sca.net:6379", 10)
	if err != nil {
		log.Panic(err)
	}
	c = cache.New(1*time.Minute, 1*time.Minute)
}

type Auth struct {
	ID string
}

func Mock(id string) (*Auth, error) {
	a := new(Auth)

	value, found := c.Get("auth:" + id)

	if found {
		a.ID = value.(string)
		conn, err := db.Get()
		conn.Cmd("set", "auth:"+id, "yeah")
		return a, nil

	} else {

		conn, err := db.Get()
		if err != nil {
			return nil, err
		}
		defer db.Put(conn)
		conn.Cmd("set", "auth:"+id, "yeah")
		reply := conn.Cmd("get", "auth:"+id).String()
		c.Set("auth:"+id, reply, cache.DefaultExpiration)
		a.ID = reply
		return a, nil
	}
}
