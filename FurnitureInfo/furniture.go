package FurnitureInfo

import "fmt"

type StructFurniture struct {
	Name string
	h    int
	w    int
	l    int
}

func InitFurniture() []StructFurniture {
	couch := StructFurniture{"диван", 100, 120, 300}
	chair := StructFurniture{"кресло", 100, 120, 80}
	table := StructFurniture{"стол", 90, 90, 90}
	wardrobe := StructFurniture{"шкаф", 200, 70, 100}
	shelf := StructFurniture{"полка", 50, 30, 100}
	return []StructFurniture{couch, chair, table, wardrobe, shelf}
}

func ShowFurniture(furniture []StructFurniture) {
	for _, ObjFurniture := range furniture {
		fmt.Println(ObjFurniture.Name, " ")
	}
}
