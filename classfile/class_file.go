package classfile
import "fmt"
//ClassFile 结构体 用来表示class文件
type ClassFile struct{
   
   magic uint32 //魔数
   minorVersion uint16 //次版本号
   majorVersion uint32 //主版本号
   constantPool ConstantPool //常量池  待实现
   accessFlags uint16 //访问控制标志
   thisClass   uint16 //
   superClass uint16 //主类
   interfaces []uint16 //实现的接口
}

