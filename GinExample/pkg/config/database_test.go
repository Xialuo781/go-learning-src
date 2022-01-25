package config

import (
	"fmt"
	"testing"
)

func TestMysql(t *testing.T) {
	Mysql.Load("../../conf/database.ini").Init()
	fmt.Println(Mysql)
	if Mysql == nil {
		t.Fail()
	}
}
