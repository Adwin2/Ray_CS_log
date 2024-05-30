//单例模式 避免实例化 可重复使用\
本质上还是类  也可以是命名空间 道理相同 使用方法有差异

#include<iostream>

class Random {
public:
    Random(const Random&) = delete; //删除构造函数
    //Get() 相当于生成单例 不需要实例化\
    实现各种非静态方法(内部映射) 并用Get().访问它
    static Random& Get(){
        static Random Iinstance; //*静态*成员变量 避免实例化
        return Iinstance;
    }
    //同理，也可以不用 Get.直接使用静态函数进行内部映射
    static float Float() {
        return Get().IFloat();
    }
private:
    //类内部映射实现类的基本功能
    float IFloat(){ return m_instance; }
    Random(){};  
    float m_instance = 0.5f;
};

int main() {
    //auto a = Random::Get().Float();
    float number = Random::Float();
    std::cout<<number<<std::endl;
    std::cin.get();
}