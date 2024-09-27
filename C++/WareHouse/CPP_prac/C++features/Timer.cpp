#include<iostream>
#include <chrono>


//RAII resource acquisition is initialization
class Timer {
public:
    Timer(){
        startTime = std::chrono::high_resolution_clock::now();
    }
    ~Timer(){
        STOP();
    }
    void STOP(){
        auto endTime = std::chrono::high_resolution_clock::now();
        auto start = std::chrono::time_point_cast<std::chrono::microseconds>(startTime).time_since_epoch().count();
        auto end = std::chrono::time_point_cast<std::chrono::microseconds>(endTime).time_since_epoch().count();

        auto duration = end - start;
        double ms = duration * 0.001;
        
        std::cout<<duration<<"μs"<<"( "<<ms<<"ms )"<<std::endl;
    }
private:
    std::chrono::time_point<std::chrono::high_resolution_clock> startTime;
};

int main() {
    int value = 0;
    {
        Timer timer;
        for(int i = 0; i <= 1000000;i ++) 
            value += 2;
    }

    std::cin.get();
}