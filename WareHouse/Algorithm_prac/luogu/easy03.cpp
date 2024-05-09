////给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
////
////有效字符串需满足：
////左括号必须用相同类型的右括号闭合。
////左括号必须以正确的顺序闭合。
////每个右括号都有一个对应的相同类型的左括号。
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
