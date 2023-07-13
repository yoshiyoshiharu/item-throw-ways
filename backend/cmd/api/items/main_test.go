package main

import (
	"testing"

	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

func TestHandler(t *testing.T) {
  fmt.Println("test")

  var kinds []entity.Kind

  repository.Db.Find(&kinds)
  fmt.Println(kinds)
}
