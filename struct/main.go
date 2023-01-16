package main

import (
	"errors"
	"fmt"
	"reflect"
)

type S2F105HsmsData struct {
	CeId     string
	ProdInfo *ProdInfo
}

var b = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18"}
var _err error

type ProdInfo1 interface {
	ProdInfo()
}
type ProdInfo struct {
	FactoryName string
	FlowId      string
	EqpId       string
	UnitId      string
	LotId       string
	CstId       string
	GlsId       string
	ChId        string
	OperId      string
	ProdId      string
	PpId        string
	DvInfo      []DVInfo
}

type DVInfo struct {
	DvName  string
	DvValue string
}

func main() {

	s := &S2F105HsmsData{ // 初始化结构体
		CeId: "",
		ProdInfo: &ProdInfo{
			FactoryName: "",
			FlowId:      "",
			EqpId:       "",
			UnitId:      "",
			LotId:       "",
			CstId:       "",
			GlsId:       "",
			ChId:        "",
			OperId:      "",
			ProdId:      "",
			PpId:        "", //12
			DvInfo:      []DVInfo{{DvName: "", DvValue: ""}, {DvName: "", DvValue: ""}},
		},
	}
	DecodeElement(s, b)
	fmt.Println(s.CeId)
	fmt.Println(s.ProdInfo.FactoryName)
	fmt.Println(s.ProdInfo.FlowId)
	fmt.Println(s.ProdInfo.EqpId)
	fmt.Println(s.ProdInfo.UnitId)
	fmt.Println(s.ProdInfo.LotId)
	fmt.Println(s.ProdInfo.CstId)
	fmt.Println(s.ProdInfo.ChId)
	fmt.Println(s.ProdInfo.OperId)
	fmt.Println(s.ProdInfo.ProdId)
	fmt.Println(s.ProdInfo.PpId)
	fmt.Println(s.ProdInfo.DvInfo[0].DvName)
	fmt.Println(s.ProdInfo.DvInfo[0].DvValue)
	fmt.Println(s.ProdInfo.DvInfo[1].DvName)
	fmt.Println(s.ProdInfo.DvInfo[1].DvValue)
}

func DecodeElement(v any, b []string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer { //因为要修改v必须要传指针
		return errors.New("non-pointer passed to Unmarshal")
	}
	return display(val, b)
}

var str string

func display(_v reflect.Value, b []string) error {
	value := _v.Elem()
	typ := value.Type()

	for k := 0; k < typ.NumField(); k++ {
		if value.Field(k).Kind() == reflect.String {
			b, str, _err = extractItemData(b)
			if _err != nil {
				return _err
			}
			field := _v.Elem().FieldByName(typ.Field(k).Name)
			field.SetString(str)
		}
		if value.Field(k).Kind() == reflect.Ptr {
			fmt.Println("1111")
			vv := value.Field(k).Interface()
			vvv := reflect.ValueOf(vv)
			v := vvv.Elem()
			t := v.Type()
			if t.Kind() == reflect.Struct {
				fmt.Println("2222")
				for n := 0; n < t.NumField(); n++ {
					if v.Field(n).Kind() == reflect.String {
						b, str, _err = extractItemData(b)
						if _err != nil {
							return _err
						}
						field := v.FieldByName(t.Field(n).Name)
						field.SetString(str)
					}
					if v.Field(n).Kind() == reflect.Slice {
						v2 := v.Field(n)
						for i := 0; i < v2.Len(); i++ {
							v3 := v2.Index(i)
							for j := 0; j < v3.NumField(); j++ {
								b, str, _err = extractItemData(b)
								if _err != nil {
									return _err
								}
								v4 := v3.Addr()
								field := v4.Elem().FieldByName(v3.Type().Field(j).Name)
								field.SetString(str)
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func extractItemData(b []string) ([]string, string, error) {
	str := b[0]
	bb := b[1:]
	return bb, str, nil

}

// func display(t reflect.Type, v reflect.Value, _v reflect.Value) {

// 	for k := 0; k < t.NumField(); k++ {

// 		if v.Field(k).Kind() == reflect.String {
// 			field := _v.Elem().FieldByName(t.Field(k).Name)
// 			field.SetString("gerrylon")
// 		}

// 		if v.Field(k).Kind() == reflect.Slice {
// 			t2 := v.Field(k).Type()
// 			v2 := v.Field(k)
// 			fmt.Println(t2.String() + " --")
// 			for i := 0; i < v2.Len(); i++ {
// 				v3 := v2.Index(i)
// 				for j := 0; j < v3.NumField(); j++ {
// 					v4 := v3.Addr()
// 					field := v4.Elem().FieldByName(v3.Type().Field(j).Name)
// 					field.SetString("gerrylon")
// 					fmt.Printf("%s -- %v \n", v3.Type().Field(j).Name, v3.Field(j).Interface())
// 				}
// 			}
// 		}
// 	}
// }

//func FieldByIndex(index []int) StructField //这个方法使得访问结构的内嵌字段成为可能。将访问各个层次的字段的索引排列起来，就形成了一个[]int，参数index不可越界，否则panic

// func display(v reflect.Value) {
// 	switch v.Kind() {
// 	case reflect.Invalid:
// 		fmt.Printf("%s = invalid\n", path)
// 	case reflect.Slice, reflect.Array:
// 		for i := 0; i < v.Len(); i++ {
// 			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
// 		}
// 	case reflect.Struct:
// 		for i := 0; i < v.NumField(); i++ {
// 			fieldPath := fmt.Sprintf(("%S.%S"), path, v.Type().Field(i).Name)
// 			display(fieldPath, v.Field(i))
// 		}
// 	case reflect.Map:
// 		for _, key := range v.MapKeys() {
// 			display(fmt.Sprintf("#{path}{#{formatAtoml(key)}}"), v.MapIndex(key))
// 		}
// 	case reflect.Interface:
// 		if v.IsNil() {
// 			fmt.Print("#{path}.type = nil\n")

// 		} else {
// 			fmt.Printf("#{path}.type = #{v.Elem().Type()}\n")
// 			display(path+".value", v.Elem())
// 		}
// 	default:
// 		formatAtoml(&v)
// 	}
// }
