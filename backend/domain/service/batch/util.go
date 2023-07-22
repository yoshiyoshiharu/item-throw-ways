package service

import "github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"

func itemExists(name string, items []entity.Item) bool {
	for _, item := range items {
		if name == item.Name {
			return true
		}
	}
	return false
}

func findKind(kindName string, allKinds []*entity.Kind) *entity.Kind {
	for _, kind := range allKinds {
		if kind.Name == kindName {
			return kind
		}
	}

	return nil
}
