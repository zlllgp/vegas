// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *UserRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UserRequest[number], err)
}

func (x *UserRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *UserRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Ua, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Idfa, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Imei, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DrawRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DrawRequest[number], err)
}

func (x *DrawRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v UserRequest
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Uq = &v
	return offset, nil
}

func (x *DrawRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Eh, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RightsDTO) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RightsDTO[number], err)
}

func (x *RightsDTO) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RightsDTO) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Num, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RightsDTO) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Amt, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DrawResult) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DrawResult[number], err)
}

func (x *DrawResult) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DrawResult) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DrawResult) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v RightsDTO
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Rights = &v
	return offset, nil
}

func (x *UserRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *UserRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *UserRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUserName())
	return offset
}

func (x *UserRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Ua == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUa())
	return offset
}

func (x *UserRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Idfa == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetIdfa())
	return offset
}

func (x *UserRequest) fastWriteField5(buf []byte) (offset int) {
	if x.Imei == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetImei())
	return offset
}

func (x *DrawRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DrawRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Uq == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUq())
	return offset
}

func (x *DrawRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Eh == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetEh())
	return offset
}

func (x *RightsDTO) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RightsDTO) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *RightsDTO) fastWriteField2(buf []byte) (offset int) {
	if x.Num == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetNum())
	return offset
}

func (x *RightsDTO) fastWriteField3(buf []byte) (offset int) {
	if x.Amt == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAmt())
	return offset
}

func (x *DrawResult) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DrawResult) fastWriteField1(buf []byte) (offset int) {
	if x.Code == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *DrawResult) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *DrawResult) fastWriteField3(buf []byte) (offset int) {
	if x.Rights == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetRights())
	return offset
}

func (x *UserRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *UserRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *UserRequest) sizeField2() (n int) {
	if x.UserName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUserName())
	return n
}

func (x *UserRequest) sizeField3() (n int) {
	if x.Ua == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUa())
	return n
}

func (x *UserRequest) sizeField4() (n int) {
	if x.Idfa == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetIdfa())
	return n
}

func (x *UserRequest) sizeField5() (n int) {
	if x.Imei == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetImei())
	return n
}

func (x *DrawRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DrawRequest) sizeField1() (n int) {
	if x.Uq == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUq())
	return n
}

func (x *DrawRequest) sizeField2() (n int) {
	if x.Eh == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetEh())
	return n
}

func (x *RightsDTO) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *RightsDTO) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetId())
	return n
}

func (x *RightsDTO) sizeField2() (n int) {
	if x.Num == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetNum())
	return n
}

func (x *RightsDTO) sizeField3() (n int) {
	if x.Amt == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAmt())
	return n
}

func (x *DrawResult) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DrawResult) sizeField1() (n int) {
	if x.Code == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetCode())
	return n
}

func (x *DrawResult) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *DrawResult) sizeField3() (n int) {
	if x.Rights == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetRights())
	return n
}

var fieldIDToName_UserRequest = map[int32]string{
	1: "UserId",
	2: "UserName",
	3: "Ua",
	4: "Idfa",
	5: "Imei",
}

var fieldIDToName_DrawRequest = map[int32]string{
	1: "Uq",
	2: "Eh",
}

var fieldIDToName_RightsDTO = map[int32]string{
	1: "Id",
	2: "Num",
	3: "Amt",
}

var fieldIDToName_DrawResult = map[int32]string{
	1: "Code",
	2: "Msg",
	3: "Rights",
}
