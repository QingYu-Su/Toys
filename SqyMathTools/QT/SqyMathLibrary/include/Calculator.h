/***************************************
     文件名：Calculator.h
     作者：苏青羽
     功能描述：数学表达式计算
     创建日期：2023-02-06
     版本：v1.0
     最后修改日期：2022-02-06
***************************************/

#ifndef _CALCULATOR_H_
#define _CALCULATOR_H_

#include <string>
#include <stack>
#include <vector>
#include "MathOperatorMap.h"

namespace SqyMathLibrary {

    typedef std::vector<std::string> MathExpression;

    class Calculator {

    public:  //类默认方法
        Calculator();

        //以下函数由编译器自动生成即可
        //~Calculator() {};
        //Calculator(const Calculator&) {};
        //Calculator& operator=(const Calculator&) {};

    public:  //基本接口
        //计算综合表达式，参数为运算符与运算数分离的字符串数组，返回运算结果
        OPERAND Calculate(const MathExpression &expression);

        //计算综合表达式，参数为表达式字符串，且各元素之间以“#”号隔开，返回运算结果
        //如3+(5*4)，其表达式字符串应该为3#+#(#5#*#4#)
        OPERAND Calculate(const std::string expression);

        bool IsSuccess() const;  //运算是否成功
        std::string GetError() const;  //获取运算失败原因

    private:  //私有接口
        bool IsNum( const std::string &s) const;  //是否为数字
        bool IsOperator(const std::string &s) const;  //是否为运算符

        OPERAND GetNum(const std::string &s) const;  //将字符串转化为数字
        MathOperator* GetMathOperator(const std::string &s) const;  //获取运算符类对象

        void SetResult(bool res, std::string reason = "");  //设置表达式运算成功与否
        void Reset(); //清空计算器栈中内容
        
        void Operate();  //进行单次运算符运算，结果会存放在栈中
        void OptTackle( MathOperator* opt); //识别运算符并做相应处理
        

        
    private:
        std::stack<OPERAND> m_Nums;  //数字栈
        std::stack<MathOperator*> m_MathOpts;  //运算符栈
        bool m_Success;  //上一次运算的结果
        std::string m_Error;  //若上一次运算失败，会保存上一次运算失败的原因
    
    };

}




#endif


