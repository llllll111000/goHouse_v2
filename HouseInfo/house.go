package HouseInfo

import (
	"goHouse/AppliancesInfo"
	"goHouse/FurnitureInfo"
	"goHouse/PeopleInfo"
	"goHouse/RoomsInfo"
)

func ShowHouse() {
	appliances := AppliancesInfo.InitAppliances()
	furniture := FurnitureInfo.InitFurniture()
	people := PeopleInfo.InitPeople()
	rooms := RoomsInfo.InitRooms()

	AppliancesInfo.ShowAppliances(appliances)
	FurnitureInfo.ShowFurniture(furniture)
	PeopleInfo.ShowPeople(people)
	RoomsInfo.ShowRooms(rooms)
}
