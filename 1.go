package main

import (
	"fmt"
	"time"
)

func main() {
	/*s:="1223333#33333#wwwwww"
	s1:=strings.Split(s,"#")
	fmt.Println(len(s1))
	fmt.Println(s1)
	sendMg:=strings.Join(s1[1:],"#")
	fmt.Println(sendMg)
	fmt.Println(s1[1])*/

}




// func main() {
//    var a int = 10

//  // %x 16进制  %d 10进制   %0 8进制

//    fmt.Printf("变量的地址: %d\n", &a  )// 1163c068
//    fmt.Println("变量的地址", &a  )    //0x1163c068
// }

/*


 匿名函数 闭包 计数器思想

//php 闭包文章 http://www.thinkphp.cn/topic/13624.html

js:	闭包是可访问上一层函数作用域里变量的函数，即便上一层函数已经关闭。

一：
func bibao(a,b int) func() int {
   i:=0
   return func() int {
   	 i++
     return i+a+b
   }
}

func main(){

   nextNumber :=  bibao(2,3)

   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
}

-------------------------------------------------------------
二：
 func bibao() func(a,b int) int {
   i:=0
   return func(a,b int) int {
   	 i++
     return i+a+b
   }
}

func main(){

   nextNumber :=  bibao()

   fmt.Println(nextNumber(2,3))
   fmt.Println(nextNumber(2,3))
}
*/

/*  求素数
  func main() {
    var a, b int//声明变量
    a=1
    A: for a < 100 {
           a++
           for b=2; b < a ; b++ {
               if a%b==0 {
                   goto A //若发现因子则不是素数
               }
           }
           fmt.Println(a,"是素数")
    }
}*/

/*func main() {
   var dengji string;
   var fenshu int = 90;

   switch fenshu {
      case 90: dengji = "A"
      case 80: dengji = "B"
      case 50,60,70 : dengji = "C"
      default: dengji = "D"
   }

   switch {
      case dengji == "A" :
         fmt.Printf("优秀!\n" )
      case dengji == "B", dengji == "C" :
         fmt.Printf("良好\n" )
      case dengji == "D" :
         fmt.Printf("及格\n" )
      case dengji == "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" );
   }
   fmt.Printf("你的等级是 %s\n", dengji );
}*/

/*func main() {
	var a = 4
	var ptr *int
	 // fmt.Printf("dddd %t\n",a)
	 fmt.Printf("%T\n", a )
	 fmt.Printf("%T\n", a )
	 ptr = &a

	 fmt.Printf( "Printf    a 的值为 %d\n",a) //先格式化数字（针对变量），并换行
	 fmt.Println("Println   a 的值为 ",a) //直接连接字符串，并输出

	 fmt.Printf("Printf   指针结果 %d\n",*ptr) //指针变量的结果
	 fmt.Printf("Printf   指针变量 %d\n",ptr) //指针变量的地址 2923398160

	 fmt.Println("Println   指针结果 ",*ptr) //指针变量的结果
	 fmt.Println("Println   指针地址 ",ptr) //指针变量的地址     0x116cc068

}
*/
// var a uint = 60    /* 60 = 0011 1100 */
// var b uint = 13    /* 13 = 0000 1101 */
// var c uint = 0

// c = a & b       /* 12 = 0000 1100 */
// fmt.Printf("第一行 - c 的值为 %d\n", c )

// c = a | b       /* 61 = 0011 1101 */
// fmt.Printf("第二行 - c 的值为 %d\n", c )

// c = a ^ b       /* 49 = 0011 0001 */
// fmt.Printf("第三行 - c 的值为 %d\n", c )

// c = a << 2     /* 240 = 1111 0000 */
// fmt.Printf("第四行 - c 的值为 %d\n", c )

// c = a >> 2     /* 15 = 0000 1111 */
// fmt.Printf("第五行 - c 的值为 %d\n", c )

/*const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a) //我的答案8，文档等于16
)

func main(){
    println(a, b, c)
}*/

/*const (
	a = "ddddddd"
	c = unsafe.Sizeof(a)
    i=40>>iota    //左移乘2，右移除2
    j=40>>iota
    k
    l
)

func main() {
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)

    fmt.Println(c)
}*/
/*func main() {
	var a = 1
	var g  string= "dddd"
	var b = 1.2
	const c = "dddd单独"

	var balance [5] float32
	var balance =  [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	balance[4]

	var a = [5][2]int
	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	a[3][2]



	d,e,f := 2,4,"ddd"
	// fmt.Println("hello,,li")
	fmt.Println(a,b,c,d,e,f,g)



}
*/
// var b bool = true
// 整型 int 和浮点型 float32、float64，
// 字符串Go语言的字符串的字节使用UTF-8编码标识Unicode文本。
/*
	派生类型:

	(a) 指针类型（Pointer）
	(b) 数组类型
	(c) 结构化类型(struct)
	(d) Channel 类型
	(e) 函数类型
	(f) 切片类型
	(g) 接口类型（interface）
	(h) Map 类型
*/
