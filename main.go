package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var b = []byte{
	0x01, 0x02, 0x41, 0x03, 0x31, 0x32, 0x33, 0x01, 0x0B, 0x41, 0x0A, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
	0x30, 0x41, 0x14, 0x31, 0x31, 0x32, 0x32, 0x33, 0x33, 0x34, 0x34, 0x35, 0x35, 0x36, 0x36, 0x37, 0x37, 0x38, 0x38, 0x39,
	0x39, 0x30, 0x30, 0x41, 0x14, 0x31, 0x31, 0x32, 0x32, 0x33, 0x33, 0x34, 0x34, 0x35, 0x35, 0x36, 0x36, 0x37, 0x37, 0x38,
	0x38, 0x39, 0x39, 0x30, 0x30, 0x41, 0x07, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x14, 0x39, 0x39, 0x38, 0x38,
	0x37, 0x37, 0x36, 0x36, 0x35, 0x35, 0x34, 0x34, 0x33, 0x33, 0x32, 0x32, 0x31, 0x31, 0x30, 0x30, 0x41, 0x14, 0x31, 0x31,
	0x32, 0x32, 0x33, 0x33, 0x34, 0x34, 0x35, 0x35, 0x30, 0x30, 0x36, 0x36, 0x37, 0x37, 0x38, 0x38, 0x39, 0x39, 0x41, 0x14,
	0x31, 0x31, 0x32, 0x32, 0x30, 0x30, 0x33, 0x33, 0x34, 0x34, 0x30, 0x30, 0x35, 0x35, 0x36, 0x36, 0x30, 0x30, 0x37, 0x37,
	0x41, 0x03, 0x31, 0x32, 0x33, 0x41, 0x0A, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x41, 0x14, 0x31,
	0x31, 0x32, 0x32, 0x30, 0x30, 0x33, 0x33, 0x34, 0x34, 0x35, 0x35, 0x36, 0x36, 0x30, 0x30, 0x37, 0x37, 0x38, 0x38, 0x41,
	0x14, 0x41, 0x42, 0x43, 0x44, 0x46, 0x52, 0x45, 0x48, 0x31, 0x31, 0x31, 0x31, 0x31, 0x32, 0x32, 0x32, 0x33, 0x33, 0x34,
	0x34, 0x01, 0x02, 0x01, 0x02, 0x41, 0x0F, 0x57, 0x45, 0x52, 0x55, 0x48, 0x44, 0x44, 0x48, 0x48, 0x48, 0x48, 0x68, 0x68,
	0x68, 0x68, 0x41, 0x14, 0x32, 0x31, 0x32, 0x32, 0x33, 0x33, 0x31, 0x31, 0x33, 0x31, 0x33, 0x32, 0x31, 0x32, 0x33, 0x32,
	0x33, 0x32, 0x33, 0x33, 0x01, 0x02, 0x41, 0x0F, 0x61, 0x73, 0x73, 0x44, 0x44, 0x46, 0x46, 0x65, 0x64, 0x63, 0x65, 0x31,
	0x31, 0x31, 0x32, 0x41, 0x14, 0x31, 0x31, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32,
	0x33, 0x33, 0x33, 0x33, 0x33,
}

type S2F105HsmsData struct {
	CeId     string
	ProdInfo *ProdInfo
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

var _err error
var str string

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
	Unmarshal(b, s)
	//DecodeElement(s, b)
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
	//fmt.Println(s.ProdInfo.DvInfo[1].DvValue)
}

func DecodeElement(v any, b []byte) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer { //因为要修改v必须要传指针
		return errors.New("non-pointer passed to Unmarshal")
	}
	return display(val, b)
}
func Unmarshal(data []byte, v any) error {

	typ := reflect.TypeOf(v)
	value := reflect.ValueOf(v)

	if typ.Kind() != reflect.Pointer { //因为要修改v必须要传指针
		return errors.New("non-pointer passed to Unmarshal")
	}
	typ = typ.Elem()
	value = value.Elem()

	switch value.Kind() {
	case reflect.String:
		fmt.Println("1111")
		data, str, _err = extractItemData(data)
		if _err != nil {
			return _err
		}
		value.SetString(str)
	case reflect.Bool:
		value.SetBool(true)
	case reflect.Float32,
		reflect.Float64:
		if f, err := strconv.ParseFloat(string(data), 64); err != nil {
			return err
		} else {
			value.SetFloat(f) //通过reflect.Value修改原始数据的值
		}
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		if i, err := strconv.ParseInt(string(data), 10, 64); err != nil {
			return err
		} else {
			value.SetInt(i) //有符号整型通过SetInt
		}
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		if i, err := strconv.ParseUint(string(data), 10, 64); err != nil {
			return err
		} else {
			value.SetUint(i) //无符号整型需要通过SetUint
		}
	case reflect.Slice:
		fmt.Printf("slice")
		slice := reflect.ValueOf(v) //别忘了，v是指针
		slice.Set(reflect.MakeSlice(typ, 11, 11))
		for i := 0; i < value.Len(); i++ {
			eleValue := slice.Index(i)
			eleType := eleValue.Type()
			if eleType.Kind() != reflect.Ptr {
				eleValue = eleValue.Addr()
			}
			if err := Unmarshal(data, eleValue.Interface()); err != nil {
				return err
			}
		}
	case reflect.Struct:
		//fmt.Printf("struct")
		fieldCount := typ.NumField()
		for i := 0; i < fieldCount; i++ {
			fieldName := value.Type().Field(i).Name

			fieldValue := value.FieldByName(fieldName)
			fieldtype := fieldValue.Type()
			if fieldtype.Kind() != reflect.Ptr {
				//fmt.Println("struct_ptr")
				//如果内嵌不是指针，则声明已经用0值初始化了，此处只需要改写值就可以了
				fieldValuenew := value.Addr()
				if err := Unmarshal(data, fieldValuenew.Interface()); err != nil { //递归调用Unmarshal，给fieldValue的底层数据赋值
					return err
				}

			} else {
				//如果内嵌的是一个指针，则需要通过New()创建一个实例(申请内存空间)。不能给New()传指针类型的type，所以调用一下
				newValue := reflect.New(typ.Field(i).Type.Elem())
				if err := Unmarshal(data, newValue.Interface()); err != nil {
					return err
				}
				value.FieldByName(value.Type().Field(i).Name).Set(newValue)
			}
		}
	}

	return nil
}
func display(_v reflect.Value, b []byte) error {
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
				//fmt.Println("2222")
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
func extractItemData(b []byte) ([]byte, string, error) {
	for {
		_type := b[0] >> 2
		if _type == 0x00 { //list格式
			headerLength := getLengthByte(b) //获取其具体的长度
			b = b[1+headerLength:]

			continue
		}
		if _type != 0x00 {

			break
		}
	}

	headerLength := getLengthByte(b)

	bodyLength, err := getBodyLength(b)
	if err != nil {
		return b, "", err
	}

	data := b[1+headerLength : 1+headerLength+bodyLength]

	b = b[1+headerLength+bodyLength:]

	return b, string(data), nil
}
func getLengthByte(b []byte) int {
	headerLength := int(b[0] & 0x3)

	return headerLength
}

// getBodyLength 解析SECE-II消息中的Body Length，返回Item数据长度或List元素个数
func getBodyLength(b []byte) (int, error) {
	headerLength := getLengthByte(b)

	var bodyLength int

	switch headerLength {
	case 0:
		return 0, errors.New("invalid header length")
	case 1:
		bodyLength = int(b[1])
	case 2:
		bodyLength = int(binary.BigEndian.Uint16(b[1:3]))
	case 3:
		bodyLength = int(binary.BigEndian.Uint32(append([]byte{0x00}, b[1:4]...)))
	}

	return bodyLength, nil
}
