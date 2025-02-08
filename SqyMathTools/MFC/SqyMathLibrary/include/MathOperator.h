/***************************************
     �ļ�����MathOperator.h
     ���ߣ�������
     ������������ѧ�������
     �������ڣ�2023-02-07
     �汾��v1.0
     ����޸����ڣ�2022-02-08
***************************************/

#ifndef _MATHOPERATOR_H_
#define _MATHOPERATOR_H_

#include <string>
#include <cmath>
#include "MathDef.h"

namespace SqyMathLibrary {

    //��ѧ���������,������������̳и���
    class MathOperator {

    public: //��Ĭ�Ϸ���
        MathOperator( const std::string symbol, const int optCount, const int priority );
        virtual ~MathOperator() {};
        //���º����ɱ������Զ����ɼ���
        //MathOperator(const MathOperator&);
        //MathOperator& operator=(const MathOperator&);

    public:  //�����ӿ�
        //����,������������ʵ�֣�Ĭ��Ϊ˫Ŀ�����ʽ����Ŀ����ʱֻ��b��������
        virtual OPERAND Operate(const OPERAND a, const OPERAND b) = 0;

        bool IsSuccess() const;  //�����Ƿ�ɹ�
        std::string GetError() const;  //��ȡ����ʧ��ԭ��

        //��ȡ�����ݳ�Ա
        std::string GetSymbol() const;
        int GetCount() const;
        int GetPriority() const;

    protected:
        void SetResult( bool res, std::string reason = "" );  //��������ɹ����

    private:  //�����ݳ�Ա
        const std::string m_Symbol;  //���������
        const size_t m_OptCount;  //�����Ŀ�������е�Ŀ��˫Ŀ
        const size_t m_Priority;  //��������ȼ�
       
        bool m_Success;  //��һ������Ľ��
        std::string m_Error;  //����һ������ʧ�ܣ��ᱣ����һ������ʧ�ܵ�ԭ��

    };

    //�������+��
    class Plus :public MathOperator {
    public:
        Plus();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������-��
    class Minus :public MathOperator {
    public:
        Minus();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������*��
    class Multiply :public MathOperator {
    public:
        Multiply();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������/��
    class Divide :public MathOperator {
    public:
        Divide();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������(��
    class Left :public MathOperator {
    public:
        Left();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������)��
    class Right :public MathOperator {
    public:
        Right();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������^��
    class Pow :public MathOperator {
    public:
        Pow();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������%��
    class Mod :public MathOperator {
    public:
        Mod();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������sin��
    class Sin :public MathOperator {
    public:
        Sin();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������cos��
    class Cos :public MathOperator {
    public:
        Cos();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������tan��
    class Tan :public MathOperator {
    public:
        Tan();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������sqrt��
    class Sqrt :public MathOperator {
    public:
        Sqrt();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������log��
    class Log :public MathOperator {
    public:
        Log();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������ln��
    class Ln :public MathOperator {
    public:
        Ln();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������rec��
    class Reciprocal :public MathOperator {
    public:
        Reciprocal();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������fact��
    class Factorial :public MathOperator {
    public:
        Factorial();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������neg��
    class Negative :public MathOperator {
    public:
        Negative();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

    //�������abs��
    class Abs :public MathOperator {
    public:
        Abs();
        OPERAND Operate(const OPERAND a, const OPERAND b);
    };

}

#endif
