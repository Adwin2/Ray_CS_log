////BFS ʾ��
//
//#include<iostream>
//#include<vector>
//#include<queue>
//#include<algorithm>
//using namespace std;
//struct TreeNode {
//	TreeNode* left;
//	TreeNode* right;
//	int val;
//	TreeNode(int x) :val(x), left(nullptr), right(nullptr) {};
//};
//
//ostream& operator<<(ostream& out,TreeNode* a) {
//	out << a->val;
//	return out;
//}
//
//void printME(int x) {
//	cout << x << endl;
//}
//
//int main() {
//
//	//��������������
//	TreeNode* root = new TreeNode(1);
//	TreeNode* n1 = new TreeNode(2);
//	TreeNode* n2 = new TreeNode(3);
//	TreeNode* n3 = new TreeNode(4);
//	TreeNode* n4 = new TreeNode(5);
//	TreeNode* n5 = new TreeNode(6);
//	TreeNode* n6 = new TreeNode(7);
//
//	root->left = n1;
//	root->right = n2;
//	n1->left = n3;
//	n1->right = n4;
//	n2->left = n5;
//	n2->right = n6;
//
//	queue<TreeNode*>queue;
//	queue.push(root);
//	/*queue.pop();*/
//	cout << "���root->val���ԣ�" << endl;
//	cout << queue.front() << endl;//���� << ���Ų��� 
//
//	vector<int> record;//��¼�ڵ�
//
//	//���ṩ��������pop()��front()�����
//	while (!queue.empty())//empty ����Ϊ��ʱ ����1
//	{
//		TreeNode* node = queue.front();
//		queue.pop();
//		record.push_back(node->val);
//
//		if (node->left != nullptr)
//			queue.push(node->left);
//		if (node->right != nullptr)
//			queue.push(node->right);
//	}
//	cout << "�ڵ��¼ֵ��" << endl;
//	for_each(record.begin(),record.end(), printME);
//	return 0;
//}