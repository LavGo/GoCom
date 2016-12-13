package test

import (
	"testing"
	"fmt"
	"github.com/LavGo/GoCom/mongo"
)

func TestConnect(t *testing.T){
	sess,coll:=mongo.Connect("lavblog","lavblog")
	fmt.Println(sess,coll)

}