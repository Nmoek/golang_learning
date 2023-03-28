#include <iostream>


using namespace std;

/*
 *@brief: 虚基类ZiMu
*/
class Zimu {

public:

virtual void m_func() = 0;

};


/*
 *@brief: 字母A类 继承字母基类
*/
class A: public Zimu {

public:

void m_func() override {
    std::cout << "i'm A!!!" << std::endl;
}

};

/*
 *@brief: 字母B类 继承字母基类
*/
class B: public Zimu {

public:

void m_func() override {
    std::cout << "i'm B!!!" << std::endl;
}

};



int main() {

    Zimu *z1 = new A;
    Zimu *z2 = new B;
    z1->m_func();
    z2->m_func();

    return 0;
}