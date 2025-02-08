
// CFunctionDlg.cpp: 实现文件
//

#include "../../pch.h"
#include "../../framework.h"
#include "../../Include/Main/FunctionIndicator.h"
#include "../../Include/Dialog/CFunctionDlg.h"
#include "afxdialogex.h"

#ifdef _DEBUG
#define new DEBUG_NEW
#endif


// CFunctionDlg 对话框



CFunctionDlg::CFunctionDlg(char dv, SML::MathExpression &expression, CString &editStr, CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_FUNCTION_DIALOG, pParent)
	, m_DV(dv)
	, m_Expression(expression)
	, m_Edit(editStr)
{
	m_hIcon = AfxGetApp()->LoadIcon(IDR_MAINFRAME);
}

void CFunctionDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	DDX_Text(pDX, IDC_EDIT_FUNCTION, m_Edit);
}

BEGIN_MESSAGE_MAP(CFunctionDlg, CDialogEx)
	ON_WM_SYSCOMMAND()
	ON_WM_PAINT()
	ON_WM_QUERYDRAGICON()
	ON_BN_CLICKED(IDC_NUM_0, &CFunctionDlg::OnBnClickedNum0)
	ON_BN_CLICKED(IDC_NUM_1, &CFunctionDlg::OnBnClickedNum1)
	ON_BN_CLICKED(IDC_NUM_2, &CFunctionDlg::OnBnClickedNum2)
	ON_BN_CLICKED(IDC_NUM_3, &CFunctionDlg::OnBnClickedNum3)
	ON_BN_CLICKED(IDC_NUM_4, &CFunctionDlg::OnBnClickedNum4)
	ON_BN_CLICKED(IDC_NUM_5, &CFunctionDlg::OnBnClickedNum5)
	ON_BN_CLICKED(IDC_NUM_6, &CFunctionDlg::OnBnClickedNum6)
	ON_BN_CLICKED(IDC_NUM_7, &CFunctionDlg::OnBnClickedNum7)
	ON_BN_CLICKED(IDC_NUM_8, &CFunctionDlg::OnBnClickedNum8)
	ON_BN_CLICKED(IDC_NUM_9, &CFunctionDlg::OnBnClickedNum9)
	ON_BN_CLICKED(IDC_CALC_DOT, &CFunctionDlg::OnBnClickedCalcDot)
	ON_BN_CLICKED(IDC_PLUS, &CFunctionDlg::OnBnClickedPlus)
	ON_BN_CLICKED(IDC_MINUS, &CFunctionDlg::OnBnClickedMinus)
	ON_BN_CLICKED(IDC_MULTIPLY, &CFunctionDlg::OnBnClickedMultiply)
	ON_BN_CLICKED(IDC_DIVIDE, &CFunctionDlg::OnBnClickedDivide)
	ON_BN_CLICKED(IDC_POW, &CFunctionDlg::OnBnClickedPow)
	ON_BN_CLICKED(IDC_MOD, &CFunctionDlg::OnBnClickedMod)
	ON_BN_CLICKED(IDC_SIN, &CFunctionDlg::OnBnClickedSin)
	ON_BN_CLICKED(IDC_COS, &CFunctionDlg::OnBnClickedCos)
	ON_BN_CLICKED(IDC_TAN, &CFunctionDlg::OnBnClickedTan)
	ON_BN_CLICKED(IDC_SQRT, &CFunctionDlg::OnBnClickedSqrt)
	ON_BN_CLICKED(IDC_LOG, &CFunctionDlg::OnBnClickedLog)
	ON_BN_CLICKED(IDC_LN, &CFunctionDlg::OnBnClickedLn)
	ON_BN_CLICKED(IDC_RECIPROAL, &CFunctionDlg::OnBnClickedReciproal)
	ON_BN_CLICKED(IDC_FACTORIAL, &CFunctionDlg::OnBnClickedFactorial)
	ON_BN_CLICKED(IDC_NEGATIVE, &CFunctionDlg::OnBnClickedNegative)
	ON_BN_CLICKED(IDC_E, &CFunctionDlg::OnBnClickedE)
	ON_BN_CLICKED(IDC_PI, &CFunctionDlg::OnBnClickedPi)
	ON_BN_CLICKED(IDC_LEFT, &CFunctionDlg::OnBnClickedLeft)
	ON_BN_CLICKED(IDC_RIGHT, &CFunctionDlg::OnBnClickedRight)
	ON_BN_CLICKED(IDC_ABS, &CFunctionDlg::OnBnClickedAbs)
	ON_BN_CLICKED(IDC_ALLCLEAR, &CFunctionDlg::OnBnClickedAllclear)
	ON_BN_CLICKED(IDC_DELETE, &CFunctionDlg::OnBnClickedDelete)
	ON_BN_CLICKED(IDC_DV, &CFunctionDlg::OnBnClickedDv)
END_MESSAGE_MAP()


// CFunctionDlg 消息处理程序

BOOL CFunctionDlg::OnInitDialog()
{
	CDialogEx::OnInitDialog();

	// 设置此对话框的图标。  当应用程序主窗口不是对话框时，框架将自动
	//  执行此操作
	SetIcon(m_hIcon, TRUE);			// 设置大图标
	SetIcon(m_hIcon, FALSE);		// 设置小图标

	// TODO: 在此添加额外的初始化代码

	return TRUE;  // 除非将焦点设置到控件，否则返回 TRUE
}

void CFunctionDlg::OnSysCommand(UINT nID, LPARAM lParam)
{
		CDialogEx::OnSysCommand(nID, lParam);

}

// 如果向对话框添加最小化按钮，则需要下面的代码
//  来绘制该图标。  对于使用文档/视图模型的 MFC 应用程序，
//  这将由框架自动完成。

void CFunctionDlg::OnPaint()
{
	if (IsIconic())
	{
		CPaintDC dc(this); // 用于绘制的设备上下文

		SendMessage(WM_ICONERASEBKGND, reinterpret_cast<WPARAM>(dc.GetSafeHdc()), 0);

		// 使图标在工作区矩形中居中
		int cxIcon = GetSystemMetrics(SM_CXICON);
		int cyIcon = GetSystemMetrics(SM_CYICON);
		CRect rect;
		GetClientRect(&rect);
		int x = (rect.Width() - cxIcon + 1) / 2;
		int y = (rect.Height() - cyIcon + 1) / 2;

		// 绘制图标
		dc.DrawIcon(x, y, m_hIcon);
	}
	else
	{
		CDialogEx::OnPaint();
	}
}

//当用户拖动最小化窗口时系统调用此函数取得光标
//显示。
HCURSOR CFunctionDlg::OnQueryDragIcon()
{
	return static_cast<HCURSOR>(m_hIcon);
}

SML::MathExpression CFunctionDlg::GetExpression() {
	return this->m_Expression;
}

CString CFunctionDlg::GetEditStr() {
	return this->m_Edit;
}

bool CFunctionDlg::IsNewNum() {
	if (this->m_Expression.empty()) return true;  //空表达式

	//前一个为自变量，无法为其添加新数位
	if (this->m_Expression.back().size() == 0 && 
		this->m_Expression.back()[0] == this->m_DV ) return true;

	//前一个占位符为运算符
	if (SML::GET_MATH_OPERATOR(this->m_Expression.back()) != NULL) return true; 
	return false;
}

void CFunctionDlg::OnBnClickedNum0()
{
	UpdateData();
	this->m_Edit += "0";  //当前编辑框添加字符
	if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
		this->m_Expression.push_back("0");
	}
	else {  //非新数字-在表达式数组最后一个元素添加数字
		this->m_Expression[m_Expression.size() - 1] += "0";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedNum1()
{
	UpdateData();
	this->m_Edit += "1";
	if (IsNewNum() == true) { 
		this->m_Expression.push_back("1");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "1";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedNum2()
{
	UpdateData();
	this->m_Edit += "2";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("2");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "2";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedNum3()
{
	UpdateData();
	this->m_Edit += "3";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("3");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "3";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedNum4()
{
	UpdateData();
	this->m_Edit += "4";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("4");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "4";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedNum5()
{
	UpdateData();
	this->m_Edit += "5";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("5");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "5";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedNum6()
{
	UpdateData();
	this->m_Edit += "6";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("6");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "6";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedNum7()
{
	UpdateData();
	this->m_Edit += "7";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("7");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "7";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedNum8()
{
	UpdateData();
	this->m_Edit += "8";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("8");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "8";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedNum9()
{
	UpdateData();
	this->m_Edit += "9";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("9");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "9";
	}
	UpdateData(false);
}


void CFunctionDlg::OnBnClickedCalcDot()
{
	UpdateData();  //小数点也按照数字模式添加
	this->m_Edit += ".";
	if (IsNewNum() == true) {
		this->m_Expression.push_back(".");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += ".";
	}
	UpdateData(false);
}


//运算符“+”
void CFunctionDlg::OnBnClickedPlus()
{
	UpdateData();
	this->m_Edit += "+";
	this->m_Expression.push_back("+");
	UpdateData(false);
}

//运算符“-”
void CFunctionDlg::OnBnClickedMinus()
{
	UpdateData();
	//若减号前无元素或为括号，则需要在减号前加不可见0“zero”，方便如-3+（-7）之类的计算
	if (this->m_Expression.empty() || this->m_Expression.back() == "(") {
		this->m_Expression.push_back("zero");   //不可见0不会显示在编辑框中
	}
	this->m_Edit += "-";
	this->m_Expression.push_back("-");
	UpdateData(false);
}

//运算符“*”
void CFunctionDlg::OnBnClickedMultiply()
{
	UpdateData();
	this->m_Edit += "*";
	this->m_Expression.push_back("*");
	UpdateData(false);
}

//运算符“/”
void CFunctionDlg::OnBnClickedDivide()
{
	UpdateData();
	this->m_Edit += "/";
	this->m_Expression.push_back("/");
	UpdateData(false);
}

//运算符“^”
void CFunctionDlg::OnBnClickedPow()
{
	UpdateData();
	this->m_Edit += "^";
	this->m_Expression.push_back("^");
	UpdateData(false);
}

//运算符“mod”
void CFunctionDlg::OnBnClickedMod()
{
	UpdateData();
	this->m_Edit += "mod";
	this->m_Expression.push_back("mod");
	UpdateData(false);
}

//运算符“sin”
void CFunctionDlg::OnBnClickedSin()
{
	UpdateData();
	this->m_Edit += "sin(";
	this->m_Expression.push_back("sin");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“cos”
void CFunctionDlg::OnBnClickedCos()
{
	UpdateData();
	this->m_Edit += "cos(";
	this->m_Expression.push_back("cos");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“tan”
void CFunctionDlg::OnBnClickedTan()
{
	UpdateData();
	this->m_Edit += "tan(";
	this->m_Expression.push_back("tan");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“sqrt”
void CFunctionDlg::OnBnClickedSqrt()
{
	UpdateData();
	this->m_Edit += "sqrt(";
	this->m_Expression.push_back("sqrt");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“log”
void CFunctionDlg::OnBnClickedLog()
{
	UpdateData();
	this->m_Edit += "log(";
	this->m_Expression.push_back("log");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“ln”
void CFunctionDlg::OnBnClickedLn()
{
	UpdateData();
	this->m_Edit += "ln(";
	this->m_Expression.push_back("ln");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“rec”
void CFunctionDlg::OnBnClickedReciproal()
{
	UpdateData();
	this->m_Edit += "rec(";
	this->m_Expression.push_back("rec");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“fact”
void CFunctionDlg::OnBnClickedFactorial()
{
	UpdateData();
	this->m_Edit += "fact(";
	this->m_Expression.push_back("fact");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“neg”
void CFunctionDlg::OnBnClickedNegative()
{
	UpdateData();
	this->m_Edit += "neg(";
	this->m_Expression.push_back("neg");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//数字“e”
void CFunctionDlg::OnBnClickedE()
{
	UpdateData();
	this->m_Edit += "e";
	this->m_Expression.push_back("e");
	UpdateData(false);
}

//数字“pi”
void CFunctionDlg::OnBnClickedPi()
{
	UpdateData();
	this->m_Edit += "pi";
	this->m_Expression.push_back("pi");
	UpdateData(false);
}

//运算符“(”
void CFunctionDlg::OnBnClickedLeft()
{
	UpdateData();
	this->m_Edit += "(";
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“)”
void CFunctionDlg::OnBnClickedRight()
{
	UpdateData();
	this->m_Edit += ")";
	this->m_Expression.push_back(")");
	UpdateData(false);
}

//运算符“abs”
void CFunctionDlg::OnBnClickedAbs()
{
	UpdateData();
	this->m_Edit += "abs(";
	this->m_Expression.push_back("abs");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//清空所有编辑框
void CFunctionDlg::OnBnClickedAllclear()
{
	UpdateData();
	this->m_Edit.Empty();
	this->m_Expression.clear();
	UpdateData(false);
}

//编辑框后退
void CFunctionDlg::OnBnClickedDelete()
{
	UpdateData();
	if (this->m_Edit.IsEmpty() == true) return; //当前编辑框为空，直接返回

	int length = 0;  //当前编辑框新长度
	if (SML::GET_MATH_OPERATOR(this->m_Expression.back()) != NULL || m_Expression.back() == "pi" ) {
		//表达式最后一个元素为运算符或者“pi”，删除该元素，当前编辑框长度更新
		length = m_Edit.GetLength() - this->m_Expression.back().size();
		this->m_Expression.pop_back();
	}
	else {
		//表达式最后一个元素为数字，删除该数字的最后一位，当前编辑框长度更新
		length = m_Edit.GetLength() - 1;
		this->m_Expression[m_Expression.size() - 1].pop_back();
		//表达式最后一个元素内容被清空，表达式需要删除该元素
		if (m_Expression[m_Expression.size() - 1].empty()) this->m_Expression.pop_back();;
	}

	//进行删除操作后，若剩下不可见0“zero”，则必须删除，保持表达式和编辑框的同步
	if ( !this->m_Expression.empty() && this->m_Expression.back() == "zero") {
		this->m_Expression.pop_back();
	}

	this->m_Edit = this->m_Edit.Left(length);  //当前编辑框截取相应长度，从左边算起

	UpdateData(false);
}


void CFunctionDlg::OnBnClickedDv()
{
	UpdateData();

	//设置自变量，具体符号由外部调用者决定
	std::string dv;
	dv.push_back(this->m_DV);
	this->m_Edit += dv.c_str();
	this->m_Expression.push_back(dv);


	UpdateData(false);
}
