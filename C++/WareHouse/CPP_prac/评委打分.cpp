//srand sort...

#include<iostream>
using namespace std;

#include<vector>
#include<deque>
#include<algorithm>

//选手类
class Person
{
public:
	Person(string name, int score) {
		this->m_name = name;
		this->m_score = score;
	}

	string m_name;
	int m_score;
};  //!!!class要用 ；结尾

//创建选手
void createPerson(vector<Person>& v) {
	string nameseed = "ABCDE";
	for (int i = 0; i < 5; i++) {
		string name = "选手";
		name += nameseed[i];

		int score = 0;
		Person p(name, score);
		v.push_back(p);
	}
}
	/*
	评分规则：
	一位选手有10个评委评分
	删大删小 求平均 并赋值
	*/
void setScore(vector<Person>&v) {
		for (vector<Person>::iterator it = v.begin(); it != v.end(); it++) {
			deque<int>d;
			for (int i = 0; i < 10; i++) {
				//随机数生成————————！！！！！！！！！！！！
				int score = rand() % 41 + 60;  // 60 ~ 100
				d.push_back(score);
			}

			sort(d.begin(), d.end());

			d.pop_front();
			d.pop_back();
			int sum = 0;
			for (deque<int>::const_iterator dit = d.begin(); dit != d.end(); dit++) {
				sum += *dit;// dit 本身是迭代器！
			}
			int avg = sum / d.size();//or `sum / 8 `;

			it->m_score = avg;
		}
}

//展示结果
void showScore(vector<Person>&v) {
	for (vector<Person>::iterator it = v.begin(); it != v.end(); it++) {
		cout << it->m_name << "平均分：" << it->m_score << "\n";
	}
	cout << endl;
}

int main()
{
	//随机数种子——————！！！ srand( (unsigned int)time(NULL) );
	srand((unsigned int)time(NULL));
	vector<Person>v1;//存放选手容器
	//创建选手
	createPerson(v1);
	//录入平均分
	setScore(v1);
	//显示平均分
	showScore(v1);

	return 0;

}
