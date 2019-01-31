package entity

import (
	"github.com/seehuhn/mt19937"
	"math/rand"
	"reflect"
	"time"
)

type PickPattern int

const (
	AllRandom	PickPattern = iota
	BuildRandom
)

type EquipmentPicker struct{
	resourceDir string
}

func (pattern PickPattern) Name() string {
	switch pattern {
		case BuildRandom:
			return "BuildRandom"
		default:
			return "AllRandom"
	}
}

func (pattern *PickPattern)Value(name string) PickPattern{
	switch name {
		case "build":
			return BuildRandom
		default:
			return AllRandom
	}
}

func (pattern *PickPattern)Patterns() []string{
	return []string{"all","build"}
}

func (picker *EquipmentPicker) PickAllRandom(player Player) Player {
	var role Role
	park := picker.ReadPark(picker.resourceDir + "./survivor_park.json")
	if player.Role == role.Value("killer"){
		park = picker.ReadPark(picker.resourceDir + "./killer_park.json")
	}
	shufflePark := shuffle(park,func(list []Park, a int, b int) []Park {
		list[a], list[b] = list[b], list[a]
		return list
	},)
	player.Park = shufflePark.([]Park)
	if player.Role == role.Value("survivor") {
		items := picker.ReadItems(picker.resourceDir + "./items.json")
		player.Item = pickRandom(items).(Item)
	}

	offering := picker.ReadOffering(picker.resourceDir + "./common_offering.json")
	roleOffering := picker.ReadOffering(picker.resourceDir + "./survivor_offering.json")
	if player.Role == role.Value("killer"){
		roleOffering = picker.ReadOffering(picker.resourceDir + "./killer_offering.json")
	}

	offering = append(offering, roleOffering...)
	player.Offering = pickRandom(offering).(Offering)
	return player
}


func (picker *EquipmentPicker) PickSurvivorBuildRandom(player Player) Player {
	survivorBuild := picker.ReadBuild(picker.resourceDir + "./test_build.json")
	build := pickRandom(survivorBuild).(Build)
	parks := build.Park
	if len(parks) > 4 {
		parks= shuffle(parks,func(list []Park, a int, b int) []Park {
			list[a], list[b] = list[b], list[a]
			return list
		},).([]Park)
	}

	if len(parks) < 4 {
		num := 4 - len(parks)
		for i := 0; i < num ; i++{
			parks = append(parks, Park{"フリーパーク","FreePark","Unknown"})
		}
	}
	player.Park = parks

	if !(build.Offering.JapaneseName == nil && build.Offering.EnglishName == nil){
		player.Offering = build.Offering
	}else{
		survivorOffering := picker.ReadOffering(picker.resourceDir + "./survivor_offering.json")
		pickOffering := pickType(build.Offering.Type,survivorOffering, func(o []Offering,t string) []Offering {
			var offerings []Offering
			for _, v := range o {
				if t == v.Type {
					offerings = append(offerings,v)
				}
			}
			return  offerings
		})
		player.Offering = pickOffering.(Offering)
	}

	if !(build.Item.JapaneseName == nil && build.Item.EnglishName == nil){
		player.Item = build.Item
	}else{
		items := picker.ReadItems(picker.resourceDir + "./items.json")
		pickItem := pickType(build.Item.Type,items, func(i []Item,t string) []Item{
			var tmp []Item
			for _, v := range i {
				if t == v.Type {
					tmp = append(tmp,v)
				}
			}
			return tmp
		})
		player.Item = pickItem.(Item)
	}
	return player
}

func pickType(buildType string ,slice interface{}, function interface{}) interface{}{
	rv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(function)
	if buildType == "random" {
		return pickRandom(slice)
	}else {
		v := fv.Call([]reflect.Value{rv, reflect.ValueOf(buildType)})[0]
		return pickRandom(v.Interface())
	}
}

func pickRandom(l interface{}) interface{} {
	lv := reflect.ValueOf(l)
	size := lv.Len()
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	vi := lv.Index(rng.Intn(size))
	return vi.Interface()
}

func shuffle(l interface{}, f interface{}) interface{} {
	lv := reflect.ValueOf(l)
	fv := reflect.ValueOf(f)
	size := lv.Len()
	v := lv.Index(0)
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	for i := 1; i < size; i++ {
		j := rng.Intn(size)
		v = fv.Call([]reflect.Value{lv, reflect.ValueOf(i),reflect.ValueOf(j)})[0]
	}
	return v.Interface()
}
