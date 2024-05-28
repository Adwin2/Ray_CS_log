#include <iostream>

class Entity 
{
public:
    virtual void PrintName(){}; //polymorphic(多态)
};
class Player:public Entity
{

};

class Enemy:public Entity
{

};

int main() {
    Player* player = new Player();
    Enemy* actualenemy = new Enemy();
    Entity* actualPlayer = player;
    
    Player* p1 = dynamic_cast<Player*>(actualPlayer);

    Player* p2 = dynamic_cast<Player*>(actualenemy);


    std::cin.get();
}

//dynamic_cast return the right ptr ,or NULL  \
it store RTTI(runtime type infomation) .add the overhead
