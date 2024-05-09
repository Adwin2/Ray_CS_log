////deque 容器  sort 排序算法 ! ! !
//
//#include<iostream>
//using namespace std;
//#include <deque>
//
//#include <algorithm>//标准算法头文件
//
//void printDeque(const deque<int>& d)//以引用的方式传入数据
//{
//	for (deque<int>::const_iterator it = d.begin(); it != d.end(); it++) //deque容器的迭代器用const_iterator
//	{
//		cout << *it << " ";
//	}
//	cout << endl;
//}
//
////deque 容器排序
//void test01() {
//	deque<int>d;
//	d.push_back(10);
//	d.push_back(20);
//	d.push_back(30);
//	d.push_front(100);
//
//
//	printDeque(d);
//
//	/*排序  默认 从小到大 升序
//	  降序——sort(d.rbegin(),d.rend());即可
//	  对于支持随机访问的迭代器的容器都适用
//	  会有时间超限的情况，不建议竞赛使用
//	*/
//	sort(d.begin(), d.end());
//	cout << "排序后：" << endl;
//	printDeque(d);
//	
//}
//
//int main() {
//
//	test01();
//
//	return 0;
//
//}