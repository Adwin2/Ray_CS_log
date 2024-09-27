////优先比较距离，其次比较编号
//#include<iostream>
//#include<vector>
//#include<cmath>
//#include<algorithm>
//using namespace std;
//class person {
//public:
//	int x = 0;
//	int y = 0;
//};
//class how_far {
//public:
//	int far = 0;
//	int num = 0;
//};
//bool compareByfar(const how_far& f_1, const how_far& f_2) {
//	if (f_1.far == f_2.far) return f_1.num < f_2.num;
//	return f_1.far < f_2.far;
//}
//int main() {
//	person p_1;
//	int n = 0;
//	cin >> n >> p_1.x >> p_1.y;
//	vector<how_far>v_res;
//	for (int i = 0; i < n; i++) {
//		person pp;
//		how_far ff;
//		cin >> pp.x >> pp.y;
//		ff.far = pow(p_1.x-pp.x, 2) + pow(p_1.y-pp.y, 2);
//		ff.num = i + 1;
//		v_res.push_back(ff);
//	}
//	sort(v_res.begin(), v_res.end(),compareByfar);
//	cout << v_res.at(0).num << endl;
//	cout << v_res.at(1).num << endl;
//	cout << v_res.at(2).num << endl;
//	return 0;
//}