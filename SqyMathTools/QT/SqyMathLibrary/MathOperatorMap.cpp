#include "./include/MathOperatorMap.h"

namespace SqyMathLibrary {

	//初始化类静态对象
	MathOperatorMap* MathOperatorMap::m_Instance = new MathOperatorMap();
	MathOperatorMap::ExitInstance MathOperatorMap::m_exit = MathOperatorMap::ExitInstance();

	MathOperatorMap* MathOperatorMap::GetInstance() {
		return m_Instance;
	}

	MathOperatorMap::MathOperatorMap() {
		//默认运算符注册
		/*this->Register("+", new Plus);
		this->Register("-", new Minus);
		this->Register("*", new Multiply);
		this->Register("/", new Divide);
		this->Register("(", new Left);
		this->Register(")", new Right);
		this->Register("^", new Pow);
		this->Register("%", new Mod);
		this->Register("sin", new Sin);
		this->Register("cos", new Cos);
		this->Register("tan", new Tan);
		this->Register("sqrt", new Sqrt);
		this->Register("log", new Log);
		this->Register("ln", new Ln);
		this->Register("rec", new Reciprocal);
		this->Register("fact", new Factorial);
		this->Register("neg", new Negative);
		this->Register("abs", new Abs);*/
		this->Register(new Plus);
		this->Register(new Minus);
		this->Register(new Multiply);
		this->Register(new Divide);
		this->Register(new Left);
		this->Register(new Right);
		this->Register(new Pow);
		this->Register(new Mod);
		this->Register(new Sin);
		this->Register(new Cos);
		this->Register(new Tan);
		this->Register(new Sqrt);
		this->Register(new Log);
		this->Register(new Ln);
		this->Register(new Reciprocal);
		this->Register(new Factorial);
		this->Register(new Negative);
		this->Register(new Abs);
	}

	MathOperatorMap::~MathOperatorMap() {
		//销毁对应的运算符类对象
		std::unordered_map< std::string, MathOperator* >::iterator it;
		for ( it = this->m_Map.begin(); it != this->m_Map.end(); it++) {
			delete it->second;
		}
	}

	MathOperator* MathOperatorMap::GetMathOperator(std::string key){
		//销毁对应的运算符类对象
		if (this->m_Map.find(key) != this->m_Map.end()) {
			return m_Map[key];
		}
		return NULL;
	}

	bool MathOperatorMap::Register(std::string key, MathOperator* val) {
		//销毁对应的运算符类对象
		if (this->m_Map.find(key) != this->m_Map.end()) {
			return false;
		}
		this->m_Map[key] = val;
		return true;
	}

	void MathOperatorMap::Destroy(std::string key) {
		//销毁对应的运算符类对象
		if (this->m_Map.find(key) == this->m_Map.end()) {
			return;
		}

		delete this->m_Map[key];
		this->m_Map[key] = NULL;

		this->m_Map.erase(key);
	}

}