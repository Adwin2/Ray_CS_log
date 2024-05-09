////BFS 示例
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
//	//建立二叉树过程
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
//	cout << "输出root->val测试：" << endl;
//	cout << queue.front() << endl;//重载 << 符号测试 
//
//	vector<int> record;//记录节点
//
//	//不提供迭代器，pop()与front()相配合
//	while (!queue.empty())//empty 队列为空时 返回1
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
//	cout << "节点记录值：" << endl;
//	for_each(record.begin(),record.end(), printME);
//	return 0;
//}