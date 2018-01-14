package classfile
import "fmt"
//ClassFile 结构体 用来表示class文件
type ClassFile struct{
   
   magic uint32 //魔数 u4
   minorVersion uint16 //次版本号 u2
   majorVersion uint16 //主版本号 u2
   constantPool ConstantPool //常量池  待实现
   accessFlags uint16 //访问控制标志
   thisClass   uint16 //
   superClass uint16 //主类
   interfaces []uint16 //实现的接口
}

//magic(魔数)的唯一作用就是确定这个文件是否为一个能被虚拟机所接受的class文件 固定值为0XCAFEBABE
func(self *ClassFile) readAndCheckMagic(reader *ClassReader){

	magic := reader.readUint32()
	if magic != 0XCAFEBABE{
		panic("java.lang.ClassFormatError: magic!")//Java虚拟机规范法规定 如果加载的class文件不符合要求的格式 应当抛出java.lang.ClassFormatError异常

	}
}

//class文件的格式版本号M.m  Java虚拟机规范固定只能支持某个范围的class文件
func(self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion{
	case 45:
		return
	case 46,47,48,49,50,51,52:
		if self.minorVersion == 0{
			return
		}
      panic("java.lang.UnsupportedClassVersionError!")//如果版本号不支持 则抛出java.lang.UnsupportedClassVersionError异常

	}
}
