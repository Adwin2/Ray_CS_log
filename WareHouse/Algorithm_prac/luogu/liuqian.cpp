#include<iostream>
#include<vector>
using namespace std;

int joseph(int n, int m) {
	if (n == 1) return 0;
	return (joseph(n - 1, m) + m) % n;
}
void show_vector(vector<int>& v1) {
	for (vector<int>::iterator it = v1.begin(); it != v1.end(); it++) {
		cout << *it << " ";
	}
	cout<<'\n'<< endl;
}

inline void plus_vector(vector<int>& v1) {
	v1.push_back(1);
	v1.push_back(2);
	v1.push_back(3);
	v1.push_back(4);
}

void put_vector(vector<int>& v1) {
	int t = v1.at(0);
	v1.erase(v1.begin());
	v1.push_back(t);
}

void insert_vector(vector<int>& v1,int m) {
	vector<int> v2(v1.begin(), v1.begin()+m);
	v1.erase(v1.begin(),v1.begin()+m);
	int n = 0;
	show_vector(v1);
	cout << "����1~4���ֱ�ʾҪ����ļ��λ��" << endl;
	cin >> n;
	v1.insert(v1.begin() + n, v2.begin(), v2.end());
}
void throw_vector(vector<int>& v1) {
	cout << "��������������1��Ů��������2" << endl;
	int sex = 0;
	cin >> sex;
	v1.erase(v1.begin(), v1.begin() + sex);
}

int main() {
	//��ʼ��
	vector<int>v1;
	plus_vector(v1);
	cout << "�����������������Ʒֱ���" << endl;
	show_vector(v1);
	cout << "���ڽ�����˺��������~��\n" << endl;
	plus_vector(v1);
	cout << "���������е�����������" << endl;
	show_vector(v1);
	int n = 0;
	cout << "���������ֵ�����" << endl;
	cin >> n;
	for (; n > 0; n--) {
		put_vector(v1);
	}
	show_vector(v1);
	insert_vector(v1,3);

	show_vector(v1);
	int d = v1.at(0);
	v1.erase(v1.begin());
	cout << "���ڲ�����������\n" << d << endl;
	cout << "�Ϸ�������1������������2����ȷ��������3" << endl;
	int region = 0;
	cin >> region;
	insert_vector(v1, region-1);

	throw_vector(v1);
	show_vector(v1);
	cout << "���ΰ�ǰ���������ţ���7�Σ����Ϊ" << endl;
	for (int i = 7; i > 0; i--) {
		put_vector(v1);
	}
	show_vector(v1);
	int num = joseph(v1.size(), 2);
	cout << v1.at(num) << endl;;
	if (v1.at(num) == d)cout << "����� ��" << endl;
	else cout << "��ѽ����һ��С�������" << endl;
	return 0;
};
