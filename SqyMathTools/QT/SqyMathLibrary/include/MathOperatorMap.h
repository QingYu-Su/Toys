/***************************************
     文件名：MathOperatorMap.h
     作者：苏青羽
     功能描述：数学运算符映射类
     创建日期：2023-02-08
     版本：v1.0
     最后修改日期：2022-02-10
***************************************/

#ifndef _MATHOPERATORMAP_H_
#define _MATHOPERATORMAP_H_

#include "MathOperator.h"
#include<unordered_map>
#include <string>

//注册数学运算符，key值默认为运算符的Symbol，若想使用非默认Symbol，请不要使用宏
#define REGISTER_MATH_OPERATOR(T) MathOperatorMap::GetInstance() -> Register<T>()

//获取数学运算符类对象
#define GET_MATH_OPERATOR(key) MathOperatorMap::GetInstance() -> GetMathOperator(key)

//销毁数学运算符
#define DESTROY_MATH_OPERATOR(key) MathOperatorMap::GetInstance() -> Destroy(key)

namespace SqyMathLibrary {

    //数学运算符映射类，实现运算符类对象与运算符字符串的映射
    //该类为单例模式，只存在一个实例对象
    class MathOperatorMap {

    public: //类默认方法
        static MathOperatorMap* GetInstance(); //获取类对象实例
        

    public:  //基本接口
        bool Register(std::string key, MathOperator* val);  //注册运算符，实现key-value的映射，运算符已存在-返回false

        //以下的注册运算符函数，key值默认为运算符类对象的Symbol,因此不用传入string
        bool Register(MathOperator* val) { //注册运算符，实现key-value的映射，运算符已存在-返回false
            return this->Register(val->GetSymbol(), val);
        }
        template<typename T>  //注册运算符的函数模板，方便外部调用
        bool Register() {
            return this->Register(new T);
        }

        void Destroy(std::string key); //销毁运算符
        MathOperator* GetMathOperator(std::string key);  //获取运算符
        
    private:  //单例模式下不可显式构造
        MathOperatorMap();
        MathOperatorMap(const MathOperatorMap&) {};
        MathOperatorMap& operator=(const MathOperatorMap&) {};
        ~MathOperatorMap();  //析构函数私有，外部不能直接delete该对象
        
        
    private:
        static MathOperatorMap* m_Instance;  //对象实例指针
        std::unordered_map< std::string, MathOperator* > m_Map; //运算符类对象与运算符字符串的映射
        

    //该单例对象在堆中创建，必须显式delete才能调用析构函数并释放内存
    private: //保证单例对象在程序结束时可以从堆中释放,声明为私有，保证不会被外部释放
        //该类仅作释放单例对象的辅助作用
        class ExitInstance {
        public:
            ~ExitInstance() {
                if (MathOperatorMap::m_Instance != NULL) {
                    delete MathOperatorMap::m_Instance;  //析构函数中析构单例对象
                }
            }
        };
        static ExitInstance m_exit;  //类静态变量，程序结束自动析构
        
    };
}

#endif
