package main

import (
	"fmt"
	_ "fmt"
)

/*
思路
在一个a*b的范围内客户需要向右移动a步,向左移动b步,
把数组看成有权图,每个结点的后继为下和右,
从上到下从左到右标号0-a*b-1
走到繁忙区域的路径权重为1,其他为0
考察已经确定距离的点(未确定的距离设为a*b)
找到权重最小考察他的后继把距离更新,更新到他最快的前驱
然后不断重复直到每个点都被考察到
然后从最后(a-1,b-1)回溯
如果最短路径权重大于m则不可到达
否则可以到达
*/

func main() {
	var M = [5][3]int{
		{0,0,0},//1  2  3
		{1,1,0},//4  5  6
		{0,0,0},//7  8  9
		{0,1,1},//10 11 12
		{0,0,0},//13 14 15
	}
	var dist[14]int
	var path[14]int

	for i,_ := range dist {
		dist[i]=15
		path[i]=-1
	}
	//没时间做了,思路写在注释里了
	fmt.Println(dist,path,M)
	
}

func Min(values []int) (min int) {
	//找到最小值返回下标的函数
	min = values[0]
	var i int
	for i,_ = range values {
		if values[i] < min {
			min = values[i]
		}
	}
	return i
}