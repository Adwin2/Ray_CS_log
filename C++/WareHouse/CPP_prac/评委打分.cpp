//srand sort...

#include<iostream>
using namespace std;

#include<vector>
#include<deque>
#include<algorithm>

//ѡ����
class Person
{
public:
	Person(string name, int score) {
		this->m_name = name;
		this->m_score = score;
	}

	string m_name;
	int m_score;
};  //!!!classҪ�� ����β

//����ѡ��
void createPerson(vector<Person>& v) {
	string nameseed = "ABCDE";
	for (int i = 0; i < 5; i++) {
		string name = "ѡ��";
		name += nameseed[i];

		int score = 0;
		Person p(name, score);
		v.push_back(p);
	}
}
	/*
	���ֹ���
	һλѡ����10����ί����
	ɾ��ɾС ��ƽ�� ����ֵ
	*/
void setScore(vector<Person>&v) {
		for (vector<Person>::iterator it = v.begin(); it != v.end(); it++) {
			deque<int>d;
			for (int i = 0; i < 10; i++) {
				//��������ɡ���������������������������������������
				int score = rand() % 41 + 60;  // 60 ~ 100
				d.push_back(score);
			}

			sort(d.begin(), d.end());

			d.pop_front();
			d.pop_back();
			int sum = 0;
			for (deque<int>::const_iterator dit = d.begin(); dit != d.end(); dit++) {
				sum += *dit;// dit �����ǵ�������
			}
			int avg = sum / d.size();//or `sum / 8 `;

			it->m_score = avg;
		}
}

//չʾ���
void showScore(vector<Person>&v) {
	for (vector<Person>::iterator it = v.begin(); it != v.end(); it++) {
		cout << it->m_name << "ƽ���֣�" << it->m_score << "\n";
	}
	cout << endl;
}

int main()
{
	//��������ӡ����������������� srand( (unsigned int)time(NULL) );
	srand((unsigned int)time(NULL));
	vector<Person>v1;//���ѡ������
	//����ѡ��
	createPerson(v1);
	//¼��ƽ����
	setScore(v1);
	//��ʾƽ����
	showScore(v1);

	return 0;

}
