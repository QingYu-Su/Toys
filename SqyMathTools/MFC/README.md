# SqyMathTools
---

����Ŀ��ʹ��C++��д�ģ���Windowsƽ̨�����еģ�����MFC�Ŀ��ӻ���ѧ�����䣬����һ����ѧ��������һ������ͼ����ʾ�����������׶�����ش�����������ע��

## ��װ����
---
- ����ϵͳ��Windows10 1903

- IDE��Microsoft Visual Studio Community 2022

## �ڱ�������
---

- ���ر���Ŀ��ʹ��Visual Studio��.sln�ļ����ڽ�������д�������С��Ŀ��Calculator������������FunctionIndicator������ͼ����ʾ������SqyMathLibrary�������ѧ����ʵ�ֿ⣩

- ��ע�⣺����SqyMathLibraryΪ��̬�⣬��������Ŀ�����ã��޷���������

- �û���ѡ��Calculator��FunctionIndicatorΪ���������ڱ������У����Ǳ����ȹ���SqyMathLibrary���ɿ��ļ��ſ��ԡ�

## ʹ�÷���/ʾ��
---
### Calculator����ѧ��������

1. ֧���������㡢���Ǻ������׳ˡ�����ֵ��ƽ�������ȳ�����㣬����֧�ֶԸ��ӱ��ʽ�ļ��㣬�纬�ж�����ţ���ͬ���ȼ������֮�����ϼ��㡣

![Calculator Demo1](Picture/Calculator-Demo1.png)


2. ���ڲ���double���ʹ�����ݣ�Ϊ�˱���һ���ľ��ȣ��ü��������֧��1e+15���µļ��㣬���ڵ���1e+15�����ֻ�ֱ����ʾINF����ʾ���㳬����Χ

![Calculator Demo2](Picture/Calculator-Demo2.png)


3. ������������������ʽʱ�ἰʱ�����������ṩ��Ӧ�Ĵ���ԭ��

![Calculator Demo3](Picture/Calculator-Demo3.png)

![Calculator Demo4](Picture/Calculator-Demo4.png)

![Calculator Demo5](Picture/Calculator-Demo5.png)


### FunctionIndicator����ѧ����ͼ����ʾ����

1. ֧�ֶ���ͨX-Y�����������꺯�����������̵Ⱥ����Ĳ���������������ͼ�����Ƴ���Ӧ�ĺ���ͼ��

![FunctionIndicator Demo1](Picture/FunctionIndicator-Demo1.png)


2. ֧����Ӻ������޸ĺ�����ɾ�������ȳ��湦�ܣ����������ú���������ͼ��ϸ�ȡ�ͼ����ɫ�����Ͷ��ֲ�����

![FunctionIndicator Demo2](Picture/FunctionIndicator-Demo2.png)

![FunctionIndicator Demo3](Picture/FunctionIndicator-Demo3.png)

![FunctionIndicator Demo4](Picture/FunctionIndicator-Demo4.png)

![FunctionIndicator Demo5](Picture/FunctionIndicator-Demo5.png)


3. ��ǰ�������ĺ������ʽ����ʾ������ͼ�ҷ������У��ڹ����������ͼ��ʱ����ʾ��Ӧ�ĺ���ͼ������㣬�������·���״̬������ʾ�����ָ�ĵ�ǰ������š�

![FunctionIndicator Demo6](Picture/FunctionIndicator-Demo6.png)


4. ����֧�������ֽ���ͼ��ķ��������������������ƶ�ģʽ���û��������ƶ�ģʽ��Ȼ�������קͼ������ƶ���

![FunctionIndicator Demo7](Picture/FunctionIndicator-Demo7.png)

![FunctionIndicator Demo8](Picture/FunctionIndicator-Demo8.png)


5. ��������ʾ�����ᣬ���񣬿̶�ֵ�������б�������Ϣ

![FunctionIndicator Demo9](Picture/FunctionIndicator-Demo9.png)

6. ֧�ֵ�/�ຯ��ģʽ���ڵ�����ģʽ�½���ʾһ������

![FunctionIndicator Demo10](Picture/FunctionIndicator-Demo10.png)


## ��Ŀ�������
---
### SqyMathLibrary
1. ����Ŀ��һ����̬�⣬��Calculator��FunctionIndicator��ʹ�ã��޷��������С��ÿ�ǳ��ؼ����������������ӻ���Ŀ����ĺ������֧�֣����е�ʵ�����ݼ��㶼�ڿ���ʵ��

2. ����Ŀ�Ĵ����ֹʹ��Windows�µ�API������C++98��׼���д����д����ʹ��STL��ȷ�����Ĵ������ֲ�ԣ����������ֲ��linuxƽ̨��

3. ��������Ŀ������ģ��

 ![SqyMathLibrary1](Picture/SqyMathLibrary-1.png)

4. �����Ƕ�ÿ��ģ��ļ򵥽��ܣ������˳��鿴Դ�������������

- MathDef:һЩ���õ�Ԥ���壬����һЩ��ѧ������Ҫ����ֵ����e���еȣ�����һЩ�����ַ�����������Դ�ļ���������

- MathOperator: ��������ѧ�����������࣬���еľ����������Ӽ��˳��ȶ����������࣬��ʵ����Ӧ�����㷽����

- MathOperatorMap: ��������ѧ������������������ַ�����ӳ�䣬���ṩ����Ӧ����ȥע�ᡢ��ȡ��������ѧ�������

- Calculator: ʵ����һ����̨��������������ѧ���ʽ���㡣

- UNIFunction�������˺������ʽ�ͺ������ߣ�ʵ�ֺ������ʽ��ֵ������

- MathFunction��������һ��������X-Yƽ������ʾ����ͼ��ĳ�����࣬�Լ������ľ���ĺ����࣬�ɽ��ɸ�����㺯��ͼ��㡣

5. ������ԭ��

- ��������ʵ�ֱ����Ͼ�������׺���ʽת��׺���ʽ�Ĺ����У�����������ջ�������ջ�߼����ת���Ĺ���,ԭ�����ɲο���[��׺���ʽ ת�� ��׺���ʽ](https://blog.csdn.net/zm_miner/article/details/115324206)


- ��������ԭ�������ӣ�����ʵ�������м����ѵ㣺
��1����������Ͷ��ֶ����������ȼ������㷽ʽ������ͬ����2����Щ��������ܻ��ڼ�������ֵ����ִ������0��tan����/2��
��3��������������û�����������ԣ��ᵼ�¸��ִ�����������bug������Ԥ��


- Ϊ�˽���������ͳһ�����⣬����Ŀ�ж�����һ���������**������ࣨMathOperator��**�����߱����������㶼�߱��Ĳ��������ȼ��������Ŀ������������ţ�
���Ҷ�����һ�����������������Operate�����÷�����һ�����麯�������巽���ɸ��������������ʵ�֡�

``` C++
public:
//����,������������ʵ�֣�Ĭ��Ϊ˫Ŀ�����ʽ����Ŀ����ʱֻ��b��������
        virtual OPERAND Operate(const OPERAND a, const OPERAND b) = 0;
private:  //�����ݳ�Ա
        const std::string m_Symbol;  //���������
        const size_t m_OptCount;  //�����Ŀ�������е�Ŀ��˫Ŀ
        const size_t m_Priority;  //��������ȼ�
```

- �������̳���ϵ������C++��������˼ά�������������Ͷ����������⣬����Ϊÿ����������岻ͬ�����㷽����ֻ����ѭ��ͬ�����ӿڼ��ɣ��ҿ���չ�Ժ�ά��������
  ![SqyMathLibrary2](Picture/SqyMathLibrary-2.png)

 - ����û��������ǻ����ַ����ı��ʽ�����ǻ��������һ�����ַ������������ӳ�䣬����Ϳ���ֱ��ͨ���ַ���ȡ����Ӧ�������������в������÷�����MathOperatorMapʵ�֡�


 - MathOperatorMap����Ʋ�ȡ����ģʽ����Ϊ����ӳ����ϵͳ�в�Ӧ���ڵڶ���ʵ����Ϊ�˱�֤�̰߳�ȫ����ȡ�˶���ʽ���ڳ����ʼ��ʱ�ʹ�������
 **ע�⣺����ȫ�ֶ���;�̬�����ʼ��˳��Ĳ�ȷ���ԣ���ò�Ҫ������ȫ�ֶ����캯����ȥʹ�ø����ṩ�ķ�����������ܳ������⡣**

 ``` C++
 private:  //����ģʽ�²�����ʽ����
        MathOperatorMap();
        MathOperatorMap(const MathOperatorMap&) {};
        MathOperatorMap& operator=(const MathOperatorMap&) {};
        ~MathOperatorMap();  //��������˽�У��ⲿ����ֱ��delete�ö���
        
        
    private:
        static MathOperatorMap* m_Instance;  //����ʵ��ָ��
        std::unordered_map< std::string, MathOperator* > m_Map; //������������������ַ�����ӳ��
```

 - Ϊ�˼��ٸ��ӵı��룬��MathOperatorMap�������ṩ�������꣺ע���������������������������������������Ҫ�������������������޸�Դ���룬
 ֻ��Ҫ������̳���������ࣨMathOperator����������ʵ��������ı�Ҫ���ܣ��ٵ�����Щ�꼴�ɡ�

``` C++
//ע����ѧ�������keyֵĬ��Ϊ�������Symbol������ʹ�÷�Ĭ��Symbol���벻Ҫʹ�ú�
#define REGISTER_MATH_OPERATOR(T) MathOperatorMap::GetInstance() -> Register<T>()

//��ȡ��ѧ����������
#define GET_MATH_OPERATOR(key) MathOperatorMap::GetInstance() -> GetMathOperator(key)

//������ѧ�����
#define DESTROY_MATH_OPERATOR(key) MathOperatorMap::GetInstance() -> Destroy(key)
```

 - ͨ��MathOperatorMap�����Ǿͳɹ���������������ַ�����˫��ӳ�䣬�����һ��������������������ǾͿ���ͨ����Ա������ö�Ӧ���ַ�����
 �����һ��������ַ��������ǿ���ͨ�����ȡ��ϵͳ��ע�������������Ӷ��õ����㷽�����Ժ����ı���൱���㡣
  ![SqyMathLibrary3](Picture/SqyMathLibrary-3.png)

 - Calculator����ǰ��Щģ��Ļ�����������[��׺���ʽ ת�� ��׺���ʽ](https://blog.csdn.net/zm_miner/article/details/115324206)ʵ����һ����̨���������ڼ���ʱһ�����ִ���Ὣ����ԭ����б��档
 ������ʵ�ֲ�����ȫ���������ܵı��ʽ��MathExpression�����ǳ�����ַ���������һ���ַ������飨std::vector<std::string>������Ҫ�ⲿ����ʱ���Ƚ�����������������з��룬�ú�̨������ֻ������Ӧ�ļ��㲿�֡�
 
 ``` C++
OPERAND Calculator::Calculate(const MathExpression& expression) {
		//�����ʽ�Ƿ�Ϊ��
		if (expression.empty()) {
			this->SetResult(false, CALC_ERROR_EMPTY_EXPR);
			return 0;
		}

		this->SetResult(true);  //��ʼ��������
		//����
		for (int i = 0; i < expression.size(); i++) {

			//����������Ŵ��ڣ�����ᵼ�¼���bug
			if (i > 0 && expression[i - 1] == "(" && expression[i] == ")") {
				this->SetResult(false, CALC_ERROR_EMPTY_BRA);
				this->Reset();
				return 0;
			}

			if (this->IsNum(expression[i]) == true) {  //����
				this->m_Nums.push(this->GetNum(expression[i])); //ֱ��ѹջ
			}
			else if (this->IsOperator(expression[i]) == true) { //�����
				MathOperator* opt = this->GetMathOperator(expression[i]);
				this->OptTackle(opt);
				if (this->IsSuccess() == false) {  //���������ʧ�ܣ���ռ��������˳�
					this->Reset();
					return 0;
				}
			}
			else {  //�����ַ������
				this->SetResult(false, CALC_ERROR_EXPRESSION);
				this->Reset();
				return 0;
			}

		}

		//������������������ս��
		while (!this->m_MathOpts.empty()) {
			this->Operate();
			if (this->IsSuccess() == false) {
				this->Reset();
				return 0;
			}
		}

		//����ջ�ռ�����Ӧֻʣһ������
		if (this->m_Nums.size() != 1) {
			this->SetResult(false, CALC_ERROR_EXCE_NUM);
			this->Reset();
			return 0;
		}

		OPERAND res = this->m_Nums.top();
		this->m_Nums.pop();

		return res;
	}
```

 6. ����

 -  ��ѧ������ʽ��������ͬ�����Ա������ܲ�ͬ�������Ҳ���ܲ�ͬ�������������Ʒֶκ������ǲ������̵ĺ��������߱��˶���������ʽ����Щ�ڴ����б���������ֺ����ϡ�
 
 - �ڿ��е�**UniFunction**���Ѿ����ȶ�����һ��ͨ�õĺ������ʽ�ṹ�壨FunctionExpression�������߱����ʽ���Ա������������������Ȳ�����
 
 ``` C++
  //ͨ�ú������ʽ
    struct FunctionExpression {
        MathExpression m_Expression;  //������ʽ�����к����Ա������ţ��޷�ֱ�Ӽ���
        char m_IV;  //����� 
        char m_DV;  //�Ա���
        OPERAND m_Left, m_Right;  //�Ա������������䣬���Ҷ˵㲻�ɳ���INF��Ĭ��Ϊ-INF~INF
        FunctionExpression(MathExpression& expression, char iv, char dv, OPERAND left = -INF, OPERAND right = INF);
    };
 ```

 - �������ʽ����Ʋ������ݺͷ����ֿ����������߱��κη�����Ҫ�������Щ���ʽ�����Խ��ɺ������߶���Щ����ʵ����صı��ʽ�Ա�����ֵ���㡣
  
  ``` C++
 //���������࣬�ṩһϵ�еķ���ר�����Դ��������ʽ
    class FunctionTool {
    public:  //��Ĭ�Ϸ���
        FunctionTool();

        //���º����ɱ������Զ����ɼ���
        //~FunctionTool() {};
        //FunctionTool(const FunctionTool&) {};
        //FunctionTool& operator=(const FunctionTool&) {};
    public:
        //���뺯�����ʽ�������е��Ա��������ض�ֵ������һ���������ѧ���ʽ
        MathExpression TranslateExpression(FunctionExpression& fe, OPERAND value);

        //���ض�ֵ���뺯�����ʽ�����㣬���ؽ��ֵ��INV��ʾ����ʧ��
        //ע�⣬����ֵINF��һ����Чֵ������ʾ����󣬶����Ǵ˴μ���ʧ��
        OPERAND GetValue(FunctionExpression& fe, OPERAND parameter);

        bool IsSuccess(); //�����Ƿ�ɹ�
        std::string GetError();  //��������ʧ��ԭ��
    private:
        void SetResult(bool res, std::string reason = ""); //���ò����ɹ����
    private:
        bool m_Success;  //��һ�β����Ľ��
        std::string m_Error;   //����һ�β���ʧ�ܣ��ᱣ����һ�β���ʧ�ܵ�ԭ��
        Calculator m_Calc;  //������
    }; 
```

 - �ڿ��е�MathFunction�ж�����һ��������X-Yƽ������ʾͼ��ĺ����࣬����һ��**�������**�������ṩ�˻��������ӿڣ�����Ҫ����ʵ��,�����Ǹ���Ĳ��ֽӿڡ�
  
  ``` C++
class MathFunction{
 public:
        //�����ض������ϵ�����X-Yƽ���ϵĺ����㡣����������Ϊ����������������Ӧ�˺�����ľ�ϸ��
        //ע�⣬��������Ķ��岢���̶��������Ա����Ķ�����ı䣬����ͨ����������x�ϵ����䣬�������̺���������t�ϵ����䣬�����꺯�����ڽǶ��ϵ�����
        //����ֵ�п���ΪINF������һ����Чֵ������ʾ�õ���y�����޶���
        //Ϊ�����Ч�ʣ�����ֵΪFunctionMap��ָ�룬��ָ�벻���ⲿ�ͷţ�������ʧ�ܣ��򷵻�NULL
        FunctionMap* Calculate(OPERAND left, OPERAND right, size_t precision);

        FunctionType GetType();  //��ú�������

        bool IsSuccess();  //�����Ƿ�ɹ� 
        std::string GetError();   //��������ʧ��ԭ��

        //��øú����ĺ������ʽ�����ڱ��ʽ���ܲ�ֹһ�����ʷ���һ������������ʵ�־������
        virtual std::list<FunctionExpression> GetExpression() = 0; 

        //���øú����ĺ������ʽ�����ڱ��ʽ���ܲ�ֹһ�����ʲ���Ϊһ������������ʵ�־������
        virtual void SetExpression(std::list<FunctionExpression> expressionList) = 0;
}
```

- ��Ϊһ�������߱��ı��ʽ���ܻ��ж��������Ҳ��һ�������ڳ�������в��趨�ڲ��ĺ������ʽ�����������趨�봦�������Ǹ���ļ̳���ϵ��
 
 ![SqyMathLibrary4](Picture/SqyMathLibrary-4.png)


- ���������е�Calculate������ͨ���������ļ��㷶Χ�ͼ��㾫��ȥ���㷶Χ�����еĺ���ͼ��㣬�÷������������ģʽ��ģ�巽�����������Ѿ��ڻ�����ʵ�֣������еĲ�����GetX��GetY����Ҫ����ȥʵ�֣���ͬ�������в�ͬ��ʵ�֡�

``` C++
class MathFunction{
 protected://����ʹ����ģ�巽�������·���������ʵ�֣�������������ʹ��
        virtual bool IsValid() = 0;  //�����Ƿ���Ч������������ʵ��
        virtual void PreProcess() {}  //����֮ǰ��Ԥ����û���򲻱�ʵ��
        virtual OPERAND GetX(OPERAND parameter) = 0;  //��ȡ�����Ա�����ĳ��ֵ��x���ϵĶ�Ӧֵ������������ʵ�֣�����INV��ʾ����ʧ��
        virtual OPERAND GetY(OPERAND parameter) = 0;  //��ȡ�����Ա�����ĳ��ֵ��y���ϵĶ�Ӧֵ������������ʵ�֣�����INV��ʾ����ʧ��
        virtual void PostProcess() {} //����֮��ĺ���û���򲻱�ʵ��
}

 FunctionMap* MathFunction::Calculate(OPERAND left, OPERAND right, size_t precision) {
		//FunctionMap res(precision+1);
		this->m_FM.clear();

		if (this->IsValid() == false) return NULL;  //������Ч��ֱ�ӷ���

		if (left >= right) {
			this->SetResult(false, FUNC_ERROR_DEF_RANGE);  //���㷶Χ����
			return NULL;
		}

		this->PreProcess();  //Ԥ����

		//x-y��ֵ��ʼ��
		this->m_MinX = INF;
		this->m_MaxX = -INF;
		this->m_MinY = INF;
		this->m_MaxY = -INF;

		OPERAND unit = (right - left) / precision;  //�������㵥Ԫ

		int i = 0;
		for (OPERAND parameter = left; parameter <= right; parameter += unit) {

			OPERAND x = this->GetX(parameter);  //��ȡxֵ
			if (x == INV) {
				//res.clear();  
				return NULL;  //xֵ��Ч��ֱ�ӷ���
			}

			OPERAND y = this->GetY(parameter);  //��ȡyֵ
			if (y == INV) {
				//res.clear();  
				return NULL; //yֵ��Ч��ֱ�ӷ���
			}

			//����X��Y�ļ�ֵ
			this->m_MaxX = std::max(this->m_MaxX, x);
			this->m_MinX = std::min(this->m_MinX, x);
			this->m_MaxY = std::max(this->m_MaxY, y);
			this->m_MinY = std::min(this->m_MinY, y);

			//��Ӻ�����
			FunctionPoint fp;
			fp.first = x;
			fp.second = y;
			this->m_FM.push_back(fp);
		}

		this->PostProcess();  //����

		return &(this->m_FM);
	}
```

### CalCulator
1. ����Ŀ��һ������**MFC�Ի���ܹ�**����ѧ��������������ָ�����ѧ���ʽ�ļ��㡣

2. ��������Ŀ������ģ��

 ![Calculator1](Picture/Calculator-1.png)

3. �����Ƕ�ÿ��ģ��ļ򵥵Ľ��ܣ������˳��鿴Դ�������������

- CCalculatorDlg:����̳���MFC�Ļ����Ի�����(DiaLogEx),�ṩ����ѧ�������ĶԻ�����棬�Լ���صİ�ť����

- CCalculatorApp:����ΪӦ�ó����࣬û���ر��ʵ�֣���Ϊ�Ի����������ж����ڡ�

4. ��Ҫ�ļ��㹦�ܽ���SqyMathLibrary��Calculator��������Ŀ��Ҫ�������û��Ľ����ͶԻ���������ʾ���Լ���֤�����ı��ʽ�ַ�������ȷת��Ϊ����������Ҫ����ѧ���ʽ��MathLibrary����

### Function-Indicator

1. ����Ŀ��һ������**MFC����ͼ�ܹ�**����ѧ����ͼ����ʾ����������ʾ���ֺ���ͼ�񣬰������ֹ��ܣ������򵥣��������֡�

2. ��������Ŀ������ģ��
 
 ![FunctionIndicator1](Picture/FunctionIndicator-1.png)

3. �����Ƕ�ÿ��ģ��ļ򵥵Ľ��ܣ������˳��鿴Դ�������������

- CFunctionDlg:�����Ǽ������е�CCalculator�ĶԻ�������ֲ�����ģ������ṩһ�����뺯�����ʽ�ĶԻ���ҳ�档���������Ų��˰�ť���֣�ɾ���˵��ڰ�ť���������Ա�����ť��������ݵ�ǰ����ĺ������͸����Ա������š���������ⲿ�ṩ��������ѧ���ʽ��������ַ�����

- CNormalFuncDlg:����Ϊ�Ի����࣬���ػ�ר��Ϊ�༭��ͨ�����ĶԻ����ڸöԻ����У��û��ɱ༭��ͨ�������ʽ�����ú����������Լ�ͼ��ϸ�ȵ������������ã����ڲ��Ὣ��Щ�������б��棬������ʹ�á�

- CPolarFuncDlg:����Ϊ�Ի����࣬���ػ�ר��Ϊ�༭�����꺯���ĶԻ��򡣹�����CNormalFuncDlg���ƣ����ǻ����һЩ������趨�������Ա������ţ�ͼ����ʾ��Χ���ڸöԻ�������о�������ʾ

- CTwoFuncDlg:����Ϊ�Ի����࣬���ػ�ר��Ϊ�༭�������̺����ĶԻ��򡣹�����CNormalFuncDlg���ƣ������ʽ����������Ϊ������

- CSetXYRangeDlg:����Ϊ�Ի����࣬�ṩһ������ͼ����ʾ��Χ�ĶԻ��򣬹��ܽ�Ϊ�򵥡�

- CFuncNumDlg:����Ϊ�Ի����࣬�ṩһ�����뺯����ŵļ򵥶Ի��򣬺����Ĺ��ܻ��õ�����

- FunctionIndicatorDoc:����ΪFunction-Indicator���ĵ��࣬��������е����ݴ洢��������ת�ȹ��ܣ��Լ����в˵���ID����Ӧ��

- FunctionIndicatorView:����ΪFunction-Indicator����ͼ�࣬ר�Ÿ�������ͼ���ڽ��л�ͼ�������ĵ���(Doc)�����Ƶ�������ݽ�����

- MainFrm��������࣬����MFC�����ܹ���û�д�ĸĶ�

- FunctionIndicator���������࣬�����������У�û�д�ĸĶ�

4. �ĵ��ࣨFunctionIndicatorDoc�����
- �ĵ����ڸ���Ŀ����Ϊ��Ҫ���ǳ��������̺��ģ����еĳ������ݶ�Ҫ���������д���ͷַ���


- ���б����˳���������Ҫ���ݣ��纯������������������

``` C++
class CFunctionIndicatorDoc : public CDocument{
private:
	bool m_ShowEdge; //��ʾ�߿�
	bool m_ShowAxis; //��ʾ������
	bool m_ShowGrid; //��ʾ����
	bool m_Single;   //��ʾ��������
	bool m_ShowFuncInfo;//���ӽ�������ʱ��ʾ�亯����Ϣ����������������ͺ������
	bool m_ShowFuncList; //�ں���ͼ���Ҳ���ʾ�������ʽ�б�

	double m_MinX, m_MaxX; //��ǰx�᷶Χ 
	double m_MinY, m_MaxY;//��ǰy�᷶Χ
	MoveMode m_MoveMode;  //��ǰ�ƶ�ģʽ

	//��������������������ӵĺ����Լ�һϵ�пɹ����õķ���
	std::list<SML::MathFunction*> m_FunctionList;

	//���ƺ��������������������ͼ������Ҫ�ı�Ҫ���ݣ������ṩView��ͼ���뺯������һһ��Ӧ
	std::list<DrawFuncData*> m_DrawDataList;  
}
```
- �������ݣ�DrawFuncData�������ĵ����ж���Ľṹ�壬�䶨�����£���ֻ������ͼ����Ҫ�ı�Ҫ���ݣ�ȷ����ͼ���ĵ��ĸ��ھۺ͵����

``` C++
//��������ı�Ҫ�������ݣ���Doc��View�����ݴ��䣬��֤Viewֻ�滭��������
struct DrawFuncData { 
	std::vector<CString> expressionStr;  //�������ʽ�ַ��������ڿ��ܱ��ʽ��ֹһ��������һ��������
	SML::FunctionMap *drawPoint;  //�������Ƶ�
	size_t precision; //������ϸ��
	int lineWidth;  //�߿�
	int lineType;  //����
	COLORREF lineColor;  //����ɫ
};
```

- ���ĵ����ж�����ö������-�ƶ�ģʽ��MoveMode����ר�����ڴ��ݸ���ͼ��ȷ����ǰ������ƶ�״̬���ƶ�ģʽ������ͼ��Ľ�������ϸ������

 ``` C++
//�ƶ�ģʽ
enum MoveMode {
	BAN, //��ֹ�ƶ�
	MOVE,  //�����ƶ�
	MOVING  //�����ƶ�
};
```

- ��Ҳ�ṩ����෽�����ⲿ���ã�ɾ����������պ������޸ĺ��������ȡ�

``` C++
class CFunctionIndicatorDoc : public CDocument{
public:
	double GetMinX();  //���x��Χ��Сֵ
	double GetMaxX();  //���x��Χ���ֵ
	double GetMinY();  //���y��Χ��Сֵ
	double GetMaxY();  //���y��Χ���ֵ
	std::list<DrawFuncData*> GetDrawDataList();  //��û��ƺ����������������
	void SetRange(double minX, double maxX, double minY, double maxY);  //���õ�ǰX-Y�᷶Χ
	void UpdateFunction();  //���º��������¼��㺯��ͼ����Ƶ㣬������ͨ��������
	void DelFunction(int num); //ɾ���ض���ź����������1��ʼ
	void ClearFunction(); //������к���

	void AlterFunction(int num); //�޸��ض���ź����Ĳ����������1��ʼ

	//�޸������е���ͨ����
	void AlterNormalFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw); 
	
	//�޸������еļ����꺯��
	void AlterPolarFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw);  
	
	//�޸������еĲ������̺���
	void AlterTwoFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw);  
	
	MoveMode GetMoveMode();  //��õ�ǰ�ƶ�ģʽ
	void SetMoveMode(MoveMode mode);  //���õ�ǰ�ƶ�ģʽ
	bool IsShowEdge(); //�Ƿ���ʾ�߿�
	bool IsShowAxis(); //�Ƿ���ʾ������
	bool IsShowGrid();  //�Ƿ���ʾ����
	bool IsShowFuncInfo(); //�Ƿ���ʾ���������
	bool IsShowFuncList();  //�Ƿ���ʾ�����б�
}
```

- ͨ����Щ���ݺͷ������ĵ���ʵ����������������ݴ���������Ӧ���в˵������Ϣ���������ݣ����ݸ��º�ἰʱ������ͼ�������ͼ���ǳ��ؼ���

- �ĵ��಻������庯���ļ��㣬���ǽ���SqyMathLibrary�еĺ����ദ��

4. ��ͼ�ࣨFunctionIndicatorDoc�����

- ����������ͼ��صı�Ҫ���ݣ����ͼ�ķ�Χ����ǰ���ָ��ĺ�����ŵȣ�������Ҫ���ݾ������ĵ��ౣ�棬���ֳ���ĸ��ھۺ͵����

``` C++
class CFunctionIndicatorView : public CView{
private:
	//����ͼ����ϵ�ϻ��ƺ���ͼ����������귶Χ(��ʵ���꣬�Ǻ�������)
	double m_Top, m_Bottom, m_Left, m_Right;
	MoveStatus m_MoveStart;  //�ƶ�ģʽ�£����һ��ʼ���ʱ�ĳ�ʼ״̬
	int m_CurFuncNum;  //��ǰ���λ��ָ��ĺ��������
}
```

- �����ṩ�˶��ֻ�ͼ�������Զ��������ʽʵ�֣��ܹ�����

``` C++
class CFunctionIndicatorView : public CView{
private:
	//ʵ��x���ϵ���ֵ�ڲ�ͬ����ϵ��ת�����ڶ�����Ϊת��ģʽ
	//true-��ʾ��ͼ����ϵ����������ϵ��ת��
	//false-��ʾ��������ϵ����ͼ����ϵ��ת��
	double TransformX(double x, bool mode);

	//ʵ��y���ϵ���ֵ�ڲ�ͬ����ϵ��ת�����ڶ�����Ϊת��ģʽ
	//true-��ʾ��ͼ����ϵ����������ϵ��ת��
	//false-��ʾ��������ϵ����ͼ����ϵ��ת��
	double TransformY(double y, bool mode);

	void SetShowWindow(); //���ú���ͼ�����Ч�滭����
	void DrawEdge(CDC* pDC);  //���滭����ı߿�
	void MarkCoordinateValue(CDC* pDC); //��ע����ϵ�ĸ�������ֵ
	void DrawGrid(CDC* pDC);  //���滭�����ڵ�����
	void DrawAxis(CDC* pDC); //���滭�����ڵ�X-Y������
	void DrawFunction(CDC* pDC);  //���ƺ���ͼ��
	void ShowFuncExpression(CDC* pDC);  //��ʾ�������ʽ 
	void AmplifyImage();  //�Ŵ�������ͼ��
	void ShrinkImage();  //��С������ͼ��
	void DoubleBufferDraw();  //����˫������ƽ��л滭������Ƶ���滭ʱ��ͼ����˸
	double GetDistacne(SML::FunctionPoint &a, SML::FunctionPoint &b);  //�����������ͼ���ľ���
	SML::FunctionPoint GetClosestPoint(CPoint point);  //������������ĺ���ͼ��㣬û���򷵻�INF��
	bool ShowFunctionPoint(CPoint point);  //��ͼ������ʾ������ڵĺ����������Ϣ,false��ʾ��겻�ں���ͼ����
	void ShowFunctionNum();  //��״̬������ʾ��괦�ĺ������
	void ShowFunctionInformation(CPoint point);  //��ʾ���������ĺ����������Ϣ,��������������ͺ������
	void ShowImagePoint(CPoint point);	//��״̬����ʾ������ڵ�ͼ������
}
```

- ��ͼ���ṩ��˫����滭������DoubleBufferDraw����[˫����滭](https://blog.csdn.net/oceanlucy/article/details/46827809?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-7.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-7.nonecase)
��һ�ֻ�����ͼ��ͼ���ͼ���������ȣ������ڴ��д���һ������Ļ��ͼ����һ�µĶ���Ȼ��ͼ�λ��Ƶ��ڴ��е���������ϣ�������������ϵ�ͼ������һ���Եؿ�������ʾ����Ļ�ϡ����ּ����ܹ�������߻�ͼ���ٶȣ����ٿ��ٺ����������⡣

``` C++
void CFunctionIndicatorView::DoubleBufferDraw() {
	//�����ȡ��˫������ƽ��ж�̬��ͼ����ֹͼ����˸

	CDC* pDC = GetDC();

	//����һ���ڴ��е���ʾ�豸
	CDC MemDC;
	MemDC.CreateCompatibleDC(NULL);

	//����һ���ڴ��е�ͼ��
	CBitmap MemBitmap;
	CRect rect;
	GetClientRect(&rect);
	MemBitmap.CreateCompatibleBitmap(pDC, rect.right, rect.bottom);

	//ָ���ڴ���ʾ�豸���ڴ��е�ͼ���ϻ�ͼ
	MemDC.SelectObject(&MemBitmap);

	//����һ����ɫ��Ϊ�ڴ���ʾ�豸�ı���ɫ
	MemDC.FillSolidRect(rect.left, rect.top, rect.right, rect.bottom, RGB(255, 255, 255));
	this->OnDraw(&MemDC);

	//���ڴ��л��õ�ͼ��ֱ�ӿ�������Ļָ��������
	pDC->BitBlt(rect.left, rect.top, rect.right, rect.bottom, &MemDC, 0, 0, SRCCOPY);

	//�ͷ������Դ
	ReleaseDC(pDC);
}
```

- ��ͼ�໹�ṩͼ���ƶ����ơ����ƶ�ģʽ���ƶ�ģʽ��Ҫ�������������¡��ſ����ƶ���Ϣʹ�ã�����Ϊ�˱����ƶ������е�ͼ����˸������Ҫʹ�õ�˫����滭��

``` C++
//�����ƶ�ʱ��״̬����
struct MoveStatus {
	CPoint point;   //��ǰ����
	double minX, maxX;  //x�᷶Χ
	double minY, maxY;  //y�᷶Χ
};

void CFunctionIndicatorView::OnLButtonDown(UINT nFlags, CPoint point)
{
	CFunctionIndicatorDoc* pDoc = GetDocument();
	SetCapture();  //��겶��,��֤֮��Ĺ���ƶ���ʹ�뿪��ǰ��������Ҳ����Ӧ�����Ϣ

	//��ǰ�����ƶ�
	if (pDoc->GetMoveMode() == MOVE) {
		pDoc->SetMoveMode(MOVING); //�����ƶ�ģʽΪ�����ƶ�

		//�����ƶ���ʼ״̬
		this->m_MoveStart.point = point;
		this->m_MoveStart.maxX = pDoc->GetMaxX();
		this->m_MoveStart.minX = pDoc->GetMinX();
		this->m_MoveStart.maxY = pDoc->GetMaxY();
		this->m_MoveStart.minY = pDoc->GetMinY();
		
		::SetCursor(LoadCursor(NULL, IDC_SIZEALL));  //���ù��
	}

	CView::OnLButtonDown(nFlags, point);
}


void CFunctionIndicatorView::OnLButtonUp(UINT nFlags, CPoint point)
{
	CFunctionIndicatorDoc* pDoc = GetDocument();

	//��ǰ�����ƶ�����������ſ�
	if (pDoc->GetMoveMode() == MOVING) {
		pDoc->SetMoveMode(MOVE); //�����ƶ�ģʽΪ�����ƶ�

		::SetCursor(LoadCursor(NULL, IDC_HAND));  //���ù��

		//���㵱ǰ���λ�����ʼ״̬���λ�õ����λ�ƣ���Ҫ����ֵת��Ϊ��������ϵ
		double valueX = this->TransformX(point.x, true) - this->TransformX(this->m_MoveStart.point.x, true);
		double valueY = this->TransformY(point.y, true) - this->TransformY(this->m_MoveStart.point.y, true);

		//�����µ�x-y��Χ
		pDoc->SetRange(this->m_MoveStart.minX - valueX, this->m_MoveStart.maxX - valueX,
			this->m_MoveStart.minY - valueY, this->m_MoveStart.maxY - valueY);

		//���º���
		pDoc->UpdateFunction();
		
		//ˢ�»滭
		this->Invalidate();
		this->UpdateWindow();
	}

	ReleaseCapture(); //�ͷ���겶��
	CView::OnLButtonUp(nFlags, point);
}


void CFunctionIndicatorView::OnMouseMove(UINT nFlags, CPoint point)
{

	//��ʾ��ǰ���ͼ��λ��
	this->ShowImagePoint(point);

	CFunctionIndicatorDoc* pDoc = GetDocument();

	//��ǰ�ƶ�ģʽΪ�����ƶ�
	if (pDoc->GetMoveMode() == MOVING) {
		::SetCursor(LoadCursor(NULL, IDC_SIZEALL));  //���ù��

		//���㵱ǰ���λ�����ʼ״̬���λ�õ����λ�ƣ���Ҫ����ֵת��Ϊ��������ϵ
		double valueX = this->TransformX(point.x, true) - this->TransformX(this->m_MoveStart.point.x, true);
		double valueY = this->TransformY(point.y, true) - this->TransformY(this->m_MoveStart.point.y, true);

		////�����µ�x-y��Χ
		pDoc->SetRange(this->m_MoveStart.minX - valueX, this->m_MoveStart.maxX - valueX,
			this->m_MoveStart.minY - valueY, this->m_MoveStart.maxY - valueY);

		//���º���
		pDoc->UpdateFunction();

		//˫����滭
		this->DoubleBufferDraw();
	}
	
	//��ǰΪ�����ƶ�ģʽ�������ù��
	if (pDoc->GetMoveMode() == MOVE) {
		::SetCursor(LoadCursor(NULL, IDC_HAND));
	}

	//��ʾ������Ϣ
	this->ShowFunctionInformation(point);

	CView::OnMouseMove(nFlags, point);
}
```

## �ο�����

- [Calculator_Qt](https://github.com/QiTianyu-0403/Calculator_Qt)

- [mfc��ѧ�������߻��Ƴ���](https://gitee.com/rewine/mfcplot#https://gitee.com/link?target=https%3A%2F%2Fblog.csdn.net%2Foceanlucy%2Farticle%2Fdetails%2F46827809%3Futm_medium%3Ddistribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-7.nonecase%26depth_1-utm_source%3Ddistribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-7.nonecase)

- [��׺���ʽ ת�� ��׺���ʽ](https://blog.csdn.net/zm_miner/article/details/115324206)

## ����

- [@������](https://github.com/suqingyu123456)

## ֤��

[MIT](https://choosealicense.com/licenses/mit/)