package RoomsInfo

import "fmt"

type RoomsStruct struct {
	Name string
	h    int
	w    int
	l    int
}

func InitRooms() []RoomsStruct {
	livingRoom := RoomsStruct{"гостинная", 400, 250, 600}
	kitchen := RoomsStruct{"кухня", 400, 250, 500}
	bedroom := RoomsStruct{"спальня", 400, 300, 550}
	return []RoomsStruct{livingRoom, kitchen, bedroom}
}

func ShowRooms(cm []RoomsStruct) {
	for _, size := range cm {
		fmt.Print("Комната: %s\nВысота: %d\nШирина: %d\nДлина: %d", cm.Name, cm.h, cm.w, cm.l)
	}
}
