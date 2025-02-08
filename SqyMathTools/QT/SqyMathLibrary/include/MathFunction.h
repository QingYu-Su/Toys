/***************************************
     文件名：MathFunction.h
     作者：苏青羽
     功能描述：数学函数
     创建日期：2023-02-11
     版本：v1.0
     最后修改日期：2022-02-13
***************************************/
#ifndef _MATHFUNCTION_H_
#define _MATHFUNCTION_H_

#include "UniFunction.h"
#include <list>

namespace SqyMathLibrary {
    typedef std::pair<OPERAND, OPERAND> FunctionPoint;  //函数点

    typedef std::vector<FunctionPoint> FunctionMap;  //函数图像

    enum FunctionType {  //函数类型
        Normal,  //普通函数
        Polar,  //极坐标函数
        Two  //参数方程函数
    };

    //数学函数类，它表示了一个可以在XY平面上显示对应图像的函数
    class MathFunction {
    public: //类默认方法
        MathFunction(FunctionType type);
        virtual ~MathFunction(){}

        //以下函数由编译器自动生成即可
        //MathFunction(const MathFunction&) {};
        //MathFunction& operator=(const MathFunction&) {};
    public:
        //计算特定区间上的所有X-Y平面上的函数点。第三个参数为函数点数量，它反应了函数点的精细度
        //注意，这里区间的定义并不固定，它视自变量的定义而改变，如普通函数则是在x上的区间，参数方程函数则是在t上的区间，极坐标函数是在角度上的区间
        //函数值有可能为INF，它是一个有效值，它表示该点在y轴上无定义
        //为了提高效率，返回值为FunctionMap的指针，该指针不用外部释放，若计算失败，则返回NULL
        FunctionMap* Calculate(OPERAND left, OPERAND right, size_t precision);

        FunctionType GetType();  //获得函数类型

        bool IsSuccess();  //操作是否成功 
        std::string GetError();   //操作运算失败原因

        //获得该函数的函数表达式，由于表达式可能不止一个，故返回一个链表，由子类实现具体代码
        virtual std::list<FunctionExpression> GetExpression() = 0; 

        //设置该函数的函数表达式，由于表达式可能不止一个，故参数为一个链表，由子类实现具体代码
        virtual void SetExpression(std::list<FunctionExpression> expressionList) = 0;

    protected://本类使用了模板方法，以下方法由子类实现，但无需由子类使用
        virtual bool IsValid() = 0;  //函数是否有效，必须由子类实现
        virtual void PreProcess() {}  //计算之前的预处理，没有则不必实现
        virtual OPERAND GetX(OPERAND parameter) = 0;  //获取函数自变量的某个值在x轴上的对应值，必须由子类实现，返回INV表示计算失败
        virtual OPERAND GetY(OPERAND parameter) = 0;  //获取函数自变量的某个值在y轴上的对应值，必须由子类实现，返回INV表示计算失败
        virtual void PostProcess() {} //计算之后的后处理，没有则不必实现

    protected:
        void SetResult(bool res, std::string reason = "");  //设置操作成功与否
        
    private:
        FunctionType m_Type;  //函数类型
        OPERAND m_MinX, m_MaxX;  //函数在X轴上的极值，会随着Calculate的使用而更新
        OPERAND m_MinY, m_MaxY;  //函数在Y轴上的极值，会随着CalCulate的使用而更新
        bool m_Success;  //上一次操作的结果
        std::string m_Error; //若上一次操作失败，会保存上一次操作失败的原因
        FunctionMap m_FM;  //函数图像坐标点
    };

    //普通函数
    class NormalFunction :public MathFunction {
    public://类默认方法
        NormalFunction(FunctionExpression& expression);

        //以下函数由编译器自动生成即可
        // ~NormalFunction() {};
        //MathFunction(const MathFunction&) {};
        //MathFunction& operator=(const MathFunction&) {};
    public:
        std::list<FunctionExpression> GetExpression();
        void SetExpression(std::list<FunctionExpression> expressionList);
    protected:  //基类方法的子类实现
        bool IsValid();
        OPERAND GetX(OPERAND parameter);
        OPERAND GetY(OPERAND parameter);
    private:
        FunctionExpression m_Expression;  //具体函数表达式，表示F(x) = y;
        FunctionTool m_Tool;  //处理函数表达式的工具
    };

    //极坐标函数
    class PolarFunction :public MathFunction {
    public://类默认方法
        PolarFunction(FunctionExpression& expression);

        //以下函数由编译器自动生成即可
        // ~PolarFunction() {};
        //PolarFunction(const PolarFunction&) {};
        //PolarFunction& operator=(const PolarFunction&) {};
    public:
        std::list<FunctionExpression> GetExpression();
        void SetExpression(std::list<FunctionExpression> expressionList);
    protected: //基类方法的子类实现
        bool IsValid();
        OPERAND GetX(OPERAND parameter);
        OPERAND GetY(OPERAND parameter);
    private:
        OPERAND GetR(OPERAND parameter);  //获取某个特定角度对应的r值
        FunctionExpression m_Expression;  //具体函数表达式，表示F(a) = r;
        FunctionTool m_Tool;  //处理函数表达式的工具
    };

    class TwoFunction :public MathFunction {
    public://类默认方法
        TwoFunction(FunctionExpression& expressionX, FunctionExpression& expressionY);
        //以下函数由编译器自动生成即可
        // ~TwoFunction() {};
        //TwoFunction(const TwoFunction&) {};
        //TwoFunction& operator=(const TwoFunction&) {};
    public:
        std::list<FunctionExpression> GetExpression();
        void SetExpression(std::list<FunctionExpression> expressionList);
    protected://基类方法的子类实现
        bool IsValid();
        OPERAND GetX(OPERAND parameter);
        OPERAND GetY(OPERAND parameter);
    private:
        FunctionExpression m_ExpressionX;  //具体函数表达式,表示F(t) = x;
        FunctionExpression m_ExpressionY;  //具体函数表达式,表示F(t) = y;
        FunctionTool m_Tool;  //处理函数表达式的工具
    };
}


#endif
