/***************************************
     文件名：MathOperator.h
     作者：苏青羽
     功能描述：数学运算符类
     创建日期：2023-02-07
     版本：v1.0
     最后修改日期：2022-02-08
***************************************/

#ifndef _MATHOPERATOR_H_
#define _MATHOPERATOR_H_

#include <string>
#include <cmath>
#include "MathDef.h"

namespace SqyMathLibrary {

    //数学运算符基类,具体运算符都继承该类
    class MathOperator {

    public: //类默认方法
        MathOperator( const std::string symbol, const int optCount, const int priority );
        virtual ~MathOperator() {};
        //以下函数由编译器自动生成即可
        //MathOperator(const MathOperator&);
        //MathOperator& operator=(const MathOperator&);

    public:  //基本接口
        //运算,方法交由子类实现，默认为双目运算格式，单目运算时只对b进行运算
        virtual OPERAND Operate(const OPERAND a, const OPERAND b) = 0;

        bool IsSuccess() const;  //运算是否成功
        std::string GetError() const;  //获取运算失败原因

        //获取类数据成员
        std::string GetSymbol() const;
        int GetCount() const;
        int GetPriority() const;

    protected:
        void SetResult( bool res, std::string reason = "" );  //设置运算成功与否

    private:  //类数据成员
        const std::string m_Symbol;  //运算符符号
        const size_t m_OptCount;  //运算符目数，现有单目、双目
        const size_t m_Priority;  //运算符优先级
       
        bool m_Success;  //上一次运算的结果
        std::string m_Error;  //若上一次运算失败，会保存上一次运算失败的原因

    };

    //运算符“+”
    class Plus :public MathOperator {
    public:
        Plus();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“-”
    class Minus :public MathOperator {
    public:
        Minus();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“*”
    class Multiply :public MathOperator {
    public:
        Multiply();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“/”
    class Divide :public MathOperator {
    public:
        Divide();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“(”
    class Left :public MathOperator {
    public:
        Left();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“)”
    class Right :public MathOperator {
    public:
        Right();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“^”
    class Pow :public MathOperator {
    public:
        Pow();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“%”
    class Mod :public MathOperator {
    public:
        Mod();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“sin”
    class Sin :public MathOperator {
    public:
        Sin();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“cos”
    class Cos :public MathOperator {
    public:
        Cos();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“tan”
    class Tan :public MathOperator {
    public:
        Tan();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“sqrt”
    class Sqrt :public MathOperator {
    public:
        Sqrt();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“log”
    class Log :public MathOperator {
    public:
        Log();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“ln”
    class Ln :public MathOperator {
    public:
        Ln();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“rec”
    class Reciprocal :public MathOperator {
    public:
        Reciprocal();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“fact”
    class Factorial :public MathOperator {
    public:
        Factorial();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“neg”
    class Negative :public MathOperator {
    public:
        Negative();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //运算符“abs”
    class Abs :public MathOperator {
    public:
        Abs();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

}

#endif
