package AppliancesInfo

import "fmt"

type StructAppliances struct {
	Name string
}

func InitAppliances() []StructAppliances {
	tv := StructAppliances{"телевизор"}
	computer := StructAppliances{"компьютер"}
	washingMachine := StructAppliances{"стиральная машина"}
	fridge := StructAppliances{"холодильник"}
	kettle := StructAppliances{"чайник"}
	return []StructAppliances{tv, computer, washingMachine, fridge, kettle}
}

func ShowAppliances(appliances []StructAppliances) {
	for _, ObjAppliances := range appliances {
		fmt.Println(ObjAppliances.Name, " ")
	}
}
