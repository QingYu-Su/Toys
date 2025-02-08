#include "./include/MathOperator.h"

namespace SqyMathLibrary {

	
	MathOperator::MathOperator(const std::string symbol, const int optCount, const int priority) 
		:m_Symbol(symbol), m_OptCount(optCount), m_Priority(priority)
	{
		this->m_Success = true;
		this->m_Error = "";
	}

	bool MathOperator::IsSuccess() const {
		return this->m_Success;
	}

	std::string MathOperator::GetError() const {
		return this->m_Error;
	}

	std::string MathOperator::GetSymbol() const {
		return this->m_Symbol;
	}
	int MathOperator::GetCount() const {
		return this->m_OptCount;
	}
	int MathOperator::GetPriority() const {
		return this->m_Priority;
	}

	void MathOperator::SetResult(bool res, std::string reason) {
		this->m_Success = res;
		this->m_Error = reason;
	}


	//�������+��
	Plus::Plus() 
		: MathOperator("+", 2, 2)
	{}
	OPERAND Plus::Operate(const OPERAND a, const OPERAND b){
		this -> SetResult(true);
		return a + b;
	}

	//�������-��
	Minus::Minus()
		: MathOperator("-", 2, 2)
	{}
	OPERAND Minus::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return a - b;
	}

	//�������*��
	Multiply::Multiply()
		: MathOperator("*", 2, 3)
	{}
	OPERAND Multiply::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return a * b;
	}

	//�������/��
	Divide::Divide()
		: MathOperator("/", 2, 3)
	{}
	OPERAND Divide::Operate(const OPERAND a, const OPERAND b){
		if (b == 0) { 
			this->SetResult(false, CALC_ERROR_DIV0);
			return INF;
		}
		this->SetResult(true);
		return a / b;
	}

	//�������(��
	Left::Left()
		: MathOperator("(", 1, 0)
	{}
	OPERAND Left::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return b;  //���������������
	}

	//�������)��
	Right::Right()
		: MathOperator(")", 1, 0)
	{}
	OPERAND Right::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return b;  //���������������
	}

	//�������^��
	Pow::Pow()
		: MathOperator("^", 2, 4)
	{}
	OPERAND Pow::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return pow(a, b);
	}

	//�������%��
	Mod::Mod()
		: MathOperator("mod", 2, 3)
	{}
	OPERAND Mod::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return fmod(a, b);
	}

	//�������sin��
	Sin::Sin()
		: MathOperator("sin", 1, 5)
	{}
	OPERAND Sin::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return sin(b);  //��Ŀ�������ֻ�Ե���ֵ����
	}

	//�������cos��
	Cos::Cos()
		: MathOperator("cos", 1, 5)
	{}
	OPERAND Cos::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return cos(b);  //��Ŀ�������ֻ�Ե���ֵ����
	}

	//�������tan��
	Tan::Tan()
		: MathOperator("tan", 1, 5)
	{}
	OPERAND Tan::Operate(const OPERAND a, const OPERAND b){
		if (fmod(b, PI / 2) == 0 && (int)( b / (PI / 2) ) % 2 == 1 ) {
			this->SetResult(false, CALC_ERROR_TAN);
			return INF;
		}
		this->SetResult(true);
		return tan(b);  //��Ŀ�������ֻ�Ե���ֵ����
	}

	//�������sqrt��
	Sqrt::Sqrt()
		: MathOperator("sqrt", 1, 5)
	{}
	OPERAND Sqrt::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return sqrt(b);  //��Ŀ�������ֻ�Ե���ֵ����
	}

	//�������log��
	Log::Log()
		: MathOperator("log", 1, 5)
	{}
	OPERAND Log::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return log10(b);  //��Ŀ�������ֻ�Ե���ֵ����
	}

	//�������ln��
	Ln::Ln()
		: MathOperator("ln", 1, 5)
	{}
	OPERAND Ln::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return log(b);
	}

	//�������rec��
	Reciprocal::Reciprocal()
		: MathOperator("rec", 1, 5)
	{}
	OPERAND Reciprocal::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return 1.0/b;  //��Ŀ�������ֻ�Ե���ֵ����
	}

	//�������fact��
	Factorial::Factorial()
		: MathOperator("fact", 1, 5)
	{}
	OPERAND Factorial::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		int i = 0, sum = 1;
		for (i = 1; i <= (int)b; i++) {
			sum *= i;
		}
		return sum;  //��Ŀ�������ֻ�Ե���ֵ����
	}

	//�������neg��
	Negative::Negative()
		: MathOperator("neg", 1, 5)
	{}
	OPERAND Negative::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return -b;  //��Ŀ�������ֻ�Ե���ֵ����
	}

	//�������abs��
	Abs::Abs()
		: MathOperator("abs", 1, 5)
	{}
	OPERAND Abs::Operate(const OPERAND a, const OPERAND b){
		this->SetResult(true);
		return abs(b);  //��Ŀ�������ֻ�Ե���ֵ����
	}
}