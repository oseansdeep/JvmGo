package classfile
import "encoding/binary"

// ClassReader 结构体 辅助读取字节流
type ClassReader struct{
    data []byte  //包装字节流
}

//读取u1类型数据
func(self *ClassReader) readUint8() uint8{
    val := self.data[0]
    self.data = self.data[1:]
    return val
}

//读取u2类型数据
func(self *ClassReader) readUint16() uint16{
    val := binary.BigEndian.Uint16(self.data)
    self.data = self.data[2:]
    return val
}

//读取u4类型数据
func(self *ClassReader) readUint32() uint32{
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[4:]
	return val
}

//读取u8类型数据(Java虚拟机规范并没有定义u8类型数据)
func(self *ClassReader) readUint64() uint64{
	val := binary.BigEndian.uint64(self.data)
	self.data = self.data[8:]
	return val
}

//读取表 表的大小由开头的uint16数据给出
func(self *ClassReader) readUint16s() []uint16{
	 n:= self.readUint16()
	 s := make([]uint16,n)
	 for i:= range s{
	 	s[i] = self.readUint16()
	 }
	 return s
}

//读取指定数量的字节
func(self *ClassReader) readBytes(n uint32) []bytes{

	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}