////����һ��ֻ���� '('��')'��'{'��'}'��'['��']' ���ַ��� s ���ж��ַ����Ƿ���Ч��
////
////��Ч�ַ��������㣺
////�����ű�������ͬ���͵������űպϡ�
////�����ű�������ȷ��˳��պϡ�
////ÿ�������Ŷ���һ����Ӧ����ͬ���͵������š�
////
//#include<iostream>
//using namespace std;
//
//#include<string>
//#include<vector>
//#include<stack>
//
//
//class Solution {
//public:
//    bool ispair(char a, char b) {
//        if ((a == '(' && b == ')') || (a == '[' && b == ']') || (a == '{' && b == '}')) return true;
//        return false;
//    }
//    bool isValid(string s) {
//        stack<char> myStack;
//
//        for (auto element : s) {
//            if (element == '(' || element == '[' || element == '{') {
//                myStack.push(element);
//            }
//            else {
//                if ((!myStack.empty())&&ispair(myStack.top(), element)) myStack.pop();
//                else return false;
//            }
//        }
//        return myStack.empty();
//    }
//};
//
//
//int main() {
//    Solution a;
//
//    string p = "(){}[]";
//    string p1 = "{[]}";
//    string p2 = "(([]){})";
//    
//    if (a.isValid(p2)) cout << "true" << endl;
//    else cout << "false" << endl;
//
//
//
//	return 0;
//}
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
