
// CCalculatorDlg.cpp: 实现文件
//

#include "pch.h"
#include "framework.h"
#include "CCalculatorApp.h"
#include "CCalculatorDlg.h"
#include "afxdialogex.h"

#ifdef _DEBUG
#define new DEBUG_NEW
#endif


// 用于应用程序“关于”菜单项的 CAboutDlg 对话框

class CAboutDlg : public CDialogEx
{
public:
	CAboutDlg();

// 对话框数据
#ifdef AFX_DESIGN_TIME
	enum { IDD = IDD_ABOUTBOX };
#endif

	protected:
	virtual void DoDataExchange(CDataExchange* pDX);    // DDX/DDV 支持

// 实现
protected:
	DECLARE_MESSAGE_MAP()
};

CAboutDlg::CAboutDlg() : CDialogEx(IDD_ABOUTBOX)
{
}

void CAboutDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
}

BEGIN_MESSAGE_MAP(CAboutDlg, CDialogEx)
END_MESSAGE_MAP()


// CCalculatorDlg 对话框



CCalculatorDlg::CCalculatorDlg(CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_CALCULATOR_DIALOG, pParent)
	, m_Current(_T(""))
	, m_Histroy(_T(""))
{
	m_hIcon = AfxGetApp()->LoadIcon(IDR_MAINFRAME);
}

void CCalculatorDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	DDX_Text(pDX, IDC_EDIT_CURRENT, m_Current);
	DDX_Text(pDX, IDC_EDIT_HISTROY, m_Histroy);
}

BEGIN_MESSAGE_MAP(CCalculatorDlg, CDialogEx)
	ON_WM_SYSCOMMAND()
	ON_WM_PAINT()
	ON_WM_QUERYDRAGICON()
	ON_BN_CLICKED(IDC_NUM_0, &CCalculatorDlg::OnBnClickedNum0)
	ON_BN_CLICKED(IDC_NUM_1, &CCalculatorDlg::OnBnClickedNum1)
	ON_BN_CLICKED(IDC_NUM_2, &CCalculatorDlg::OnBnClickedNum2)
	ON_BN_CLICKED(IDC_NUM_3, &CCalculatorDlg::OnBnClickedNum3)
	ON_BN_CLICKED(IDC_NUM_4, &CCalculatorDlg::OnBnClickedNum4)
	ON_BN_CLICKED(IDC_NUM_5, &CCalculatorDlg::OnBnClickedNum5)
	ON_BN_CLICKED(IDC_NUM_6, &CCalculatorDlg::OnBnClickedNum6)
	ON_BN_CLICKED(IDC_NUM_7, &CCalculatorDlg::OnBnClickedNum7)
	ON_BN_CLICKED(IDC_NUM_8, &CCalculatorDlg::OnBnClickedNum8)
	ON_BN_CLICKED(IDC_NUM_9, &CCalculatorDlg::OnBnClickedNum9)
	ON_BN_CLICKED(IDC_CALC_DOT, &CCalculatorDlg::OnBnClickedCalcDot)
	ON_BN_CLICKED(IDC_PLUS, &CCalculatorDlg::OnBnClickedPlus)
	ON_BN_CLICKED(IDC_MINUS, &CCalculatorDlg::OnBnClickedMinus)
	ON_BN_CLICKED(IDC_MULTIPLY, &CCalculatorDlg::OnBnClickedMultiply)
	ON_BN_CLICKED(IDC_DIVIDE, &CCalculatorDlg::OnBnClickedDivide)
	ON_BN_CLICKED(IDC_POW, &CCalculatorDlg::OnBnClickedPow)
	ON_BN_CLICKED(IDC_MOD, &CCalculatorDlg::OnBnClickedMod)
	ON_BN_CLICKED(IDC_SIN, &CCalculatorDlg::OnBnClickedSin)
	ON_BN_CLICKED(IDC_COS, &CCalculatorDlg::OnBnClickedCos)
	ON_BN_CLICKED(IDC_TAN, &CCalculatorDlg::OnBnClickedTan)
	ON_BN_CLICKED(IDC_SQRT, &CCalculatorDlg::OnBnClickedSqrt)
	ON_BN_CLICKED(IDC_LOG, &CCalculatorDlg::OnBnClickedLog)
	ON_BN_CLICKED(IDC_LN, &CCalculatorDlg::OnBnClickedLn)
	ON_BN_CLICKED(IDC_RECIPROAL, &CCalculatorDlg::OnBnClickedReciproal)
	ON_BN_CLICKED(IDC_FACTORIAL, &CCalculatorDlg::OnBnClickedFactorial)
	ON_BN_CLICKED(IDC_NEGATIVE, &CCalculatorDlg::OnBnClickedNegative)
	ON_BN_CLICKED(IDC_E, &CCalculatorDlg::OnBnClickedE)
	ON_BN_CLICKED(IDC_PI, &CCalculatorDlg::OnBnClickedPi)
	ON_BN_CLICKED(IDC_LEFT, &CCalculatorDlg::OnBnClickedLeft)
	ON_BN_CLICKED(IDC_RIGHT, &CCalculatorDlg::OnBnClickedRight)
	ON_BN_CLICKED(IDC_ABS, &CCalculatorDlg::OnBnClickedAbs)
	ON_BN_CLICKED(IDC_EQUAL, &CCalculatorDlg::OnBnClickedEqual)
	ON_BN_CLICKED(IDC_ALLCLEAR, &CCalculatorDlg::OnBnClickedAllclear)
	ON_BN_CLICKED(IDC_DELETE, &CCalculatorDlg::OnBnClickedDelete)
END_MESSAGE_MAP()


// CCalculatorDlg 消息处理程序

BOOL CCalculatorDlg::OnInitDialog()
{
	CDialogEx::OnInitDialog();

	// 将“关于...”菜单项添加到系统菜单中。

	// IDM_ABOUTBOX 必须在系统命令范围内。
	ASSERT((IDM_ABOUTBOX & 0xFFF0) == IDM_ABOUTBOX);
	ASSERT(IDM_ABOUTBOX < 0xF000);

	CMenu* pSysMenu = GetSystemMenu(FALSE);
	if (pSysMenu != nullptr)
	{
		BOOL bNameValid;
		CString strAboutMenu;
		bNameValid = strAboutMenu.LoadString(IDS_ABOUTBOX);
		ASSERT(bNameValid);
		if (!strAboutMenu.IsEmpty())
		{
			pSysMenu->AppendMenu(MF_SEPARATOR);
			pSysMenu->AppendMenu(MF_STRING, IDM_ABOUTBOX, strAboutMenu);
		}
	}

	// 设置此对话框的图标。  当应用程序主窗口不是对话框时，框架将自动
	//  执行此操作
	SetIcon(m_hIcon, TRUE);			// 设置大图标
	SetIcon(m_hIcon, FALSE);		// 设置小图标

	// TODO: 在此添加额外的初始化代码

	return TRUE;  // 除非将焦点设置到控件，否则返回 TRUE
}

void CCalculatorDlg::OnSysCommand(UINT nID, LPARAM lParam)
{
	if ((nID & 0xFFF0) == IDM_ABOUTBOX)
	{
		CAboutDlg dlgAbout;
		dlgAbout.DoModal();
	}
	else
	{
		CDialogEx::OnSysCommand(nID, lParam);
	}
}

// 如果向对话框添加最小化按钮，则需要下面的代码
//  来绘制该图标。  对于使用文档/视图模型的 MFC 应用程序，
//  这将由框架自动完成。

void CCalculatorDlg::OnPaint()
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
HCURSOR CCalculatorDlg::OnQueryDragIcon()
{
	return static_cast<HCURSOR>(m_hIcon);
}

bool CCalculatorDlg::IsNewNum() {
	if (this->m_Expression.empty()) return true;  //空表达式
	if (this->m_Expression.back() == "INF") return true;  //前一个为INF，无法为其添加新数位
	if (SML::GET_MATH_OPERATOR(this->m_Expression.back()) != NULL) return true; //前一个占位符为运算符
	return false;
}

void CCalculatorDlg::OnBnClickedNum0()
{
	UpdateData();
	this->m_Current += "0";  //当前编辑框添加字符
	if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
		this->m_Expression.push_back("0");
	}
	else {  //非新数字-在表达式数组最后一个元素添加数字
		this->m_Expression[m_Expression.size() - 1] += "0";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedNum1()
{
	UpdateData();
	this->m_Current += "1";
	if (IsNewNum() == true) { 
		this->m_Expression.push_back("1");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "1";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedNum2()
{
	UpdateData();
	this->m_Current += "2";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("2");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "2";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedNum3()
{
	UpdateData();
	this->m_Current += "3";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("3");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "3";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedNum4()
{
	UpdateData();
	this->m_Current += "4";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("4");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "4";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedNum5()
{
	UpdateData();
	this->m_Current += "5";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("5");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "5";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedNum6()
{
	UpdateData();
	this->m_Current += "6";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("6");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "6";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedNum7()
{
	UpdateData();
	this->m_Current += "7";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("7");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "7";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedNum8()
{
	UpdateData();
	this->m_Current += "8";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("8");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "8";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedNum9()
{
	UpdateData();
	this->m_Current += "9";
	if (IsNewNum() == true) {
		this->m_Expression.push_back("9");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += "9";
	}
	UpdateData(false);
}


void CCalculatorDlg::OnBnClickedCalcDot()
{
	UpdateData();  //小数点也按照数字模式添加
	this->m_Current += ".";
	if (IsNewNum() == true) {
		this->m_Expression.push_back(".");
	}
	else {
		this->m_Expression[m_Expression.size() - 1] += ".";
	}
	UpdateData(false);
}


//运算符“+”
void CCalculatorDlg::OnBnClickedPlus()
{
	UpdateData();
	this->m_Current += "+";
	this->m_Expression.push_back("+");
	UpdateData(false);
}

//运算符“-”
void CCalculatorDlg::OnBnClickedMinus()
{
	UpdateData();
	//若减号前无元素或为括号，则需要在减号前加不可见0“zero”，方便如-3+（-7）之类的计算
	if (this->m_Expression.empty() || this->m_Expression.back() == "(") {
		this->m_Expression.push_back("zero");   //不可见0不会显示在编辑框中
	}
	this->m_Current += "-";
	this->m_Expression.push_back("-");
	UpdateData(false);
}

//运算符“*”
void CCalculatorDlg::OnBnClickedMultiply()
{
	UpdateData();
	this->m_Current += "*";
	this->m_Expression.push_back("*");
	UpdateData(false);
}

//运算符“/”
void CCalculatorDlg::OnBnClickedDivide()
{
	UpdateData();
	this->m_Current += "/";
	this->m_Expression.push_back("/");
	UpdateData(false);
}

//运算符“^”
void CCalculatorDlg::OnBnClickedPow()
{
	UpdateData();
	this->m_Current += "^";
	this->m_Expression.push_back("^");
	UpdateData(false);
}

//运算符“mod”
void CCalculatorDlg::OnBnClickedMod()
{
	UpdateData();
	this->m_Current += "mod";
	this->m_Expression.push_back("mod");
	UpdateData(false);
}

//运算符“sin”
void CCalculatorDlg::OnBnClickedSin()
{
	UpdateData();
	this->m_Current += "sin(";
	this->m_Expression.push_back("sin");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“cos”
void CCalculatorDlg::OnBnClickedCos()
{
	UpdateData();
	this->m_Current += "cos(";
	this->m_Expression.push_back("cos");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“tan”
void CCalculatorDlg::OnBnClickedTan()
{
	UpdateData();
	this->m_Current += "tan(";
	this->m_Expression.push_back("tan");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“sqrt”
void CCalculatorDlg::OnBnClickedSqrt()
{
	UpdateData();
	this->m_Current += "sqrt(";
	this->m_Expression.push_back("sqrt");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“log”
void CCalculatorDlg::OnBnClickedLog()
{
	UpdateData();
	this->m_Current += "log(";
	this->m_Expression.push_back("log");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“ln”
void CCalculatorDlg::OnBnClickedLn()
{
	UpdateData();
	this->m_Current += "ln(";
	this->m_Expression.push_back("ln");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“rec”
void CCalculatorDlg::OnBnClickedReciproal()
{
	UpdateData();
	this->m_Current += "rec(";
	this->m_Expression.push_back("rec");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“fact”
void CCalculatorDlg::OnBnClickedFactorial()
{
	UpdateData();
	this->m_Current += "fact(";
	this->m_Expression.push_back("fact");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“neg”
void CCalculatorDlg::OnBnClickedNegative()
{
	UpdateData();
	this->m_Current += "neg(";
	this->m_Expression.push_back("neg");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//数字“e”
void CCalculatorDlg::OnBnClickedE()
{
	UpdateData();
	this->m_Current += "e";
	this->m_Expression.push_back("e");
	UpdateData(false);
}

//数字“pi”
void CCalculatorDlg::OnBnClickedPi()
{
	UpdateData();
	this->m_Current += "pi";
	this->m_Expression.push_back("pi");
	UpdateData(false);
}

//运算符“(”
void CCalculatorDlg::OnBnClickedLeft()
{
	UpdateData();
	this->m_Current += "(";
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//运算符“)”
void CCalculatorDlg::OnBnClickedRight()
{
	UpdateData();
	this->m_Current += ")";
	this->m_Expression.push_back(")");
	UpdateData(false);
}

//运算符“abs”
void CCalculatorDlg::OnBnClickedAbs()
{
	UpdateData();
	this->m_Current += "abs(";
	this->m_Expression.push_back("abs");
	this->m_Expression.push_back("(");
	UpdateData(false);
}

//按下=开始计算表达式
void CCalculatorDlg::OnBnClickedEqual()
{
	UpdateData();
	
	double res = m_Calculator.Calculate(this->m_Expression);  //计算
	if (m_Calculator.IsSuccess() == false) { //计算失败弹窗
		AfxMessageBox(m_Calculator.GetError().c_str());
		return;
	}

	this->m_Current += "=";
	this->m_Histroy = this->m_Current;  //历史编辑框赋值

	this->m_Expression.clear();  //表达式清空
	this->m_Current.Empty(); //当前编辑框清空

	//为当前编辑框赋值
	if (res == INF) {
		this->m_Current = "INF";
	}
	else if (res == -INF) {
		this->m_Current = "-INF";
	}
	else {
		this->m_Current = std::to_string(res).c_str();
	}

	//为表达式赋值
	if (res < 0) {  //结果为负数需要添加zero
		this->m_Expression.push_back("zero");
		this->m_Expression.push_back("-");
		res = -res;
	}

	if (res == INF) this->m_Expression.push_back("INF");
	else this->m_Expression.push_back(std::to_string(res));

	UpdateData(false);
}

//清空所有编辑框
void CCalculatorDlg::OnBnClickedAllclear()
{
	UpdateData();
	this->m_Current.Empty();
	this->m_Histroy.Empty();
	this->m_Expression.clear();
	UpdateData(false);
}

//编辑框后退
void CCalculatorDlg::OnBnClickedDelete()
{
	UpdateData();
	if (this->m_Current.IsEmpty() == true) return; //当前编辑框为空，直接返回

	int length = 0;  //当前编辑框新长度
	if (SML::GET_MATH_OPERATOR(this->m_Expression.back()) != NULL || m_Expression.back() == "pi" || m_Expression.back() == "INF") {
		//表达式最后一个元素为运算符或者“pi”，删除该元素，当前编辑框长度更新
		length = m_Current.GetLength() - this->m_Expression.back().size();
		this->m_Expression.pop_back();
	}
	else {
		//表达式最后一个元素为数字，删除该数字的最后一位，当前编辑框长度更新
		length = m_Current.GetLength() - 1;
		this->m_Expression[m_Expression.size() - 1].pop_back();
		//表达式最后一个元素内容被清空，表达式需要删除该元素
		if (m_Expression[m_Expression.size() - 1].empty()) this->m_Expression.pop_back();;
	}

	//进行删除操作后，若剩下不可见0“zero”，则必须删除，保持表达式和编辑框的同步
	if ( !this->m_Expression.empty() && this->m_Expression.back() == "zero") {
		this->m_Expression.pop_back();
	}

	this->m_Current = this->m_Current.Left(length);  //当前编辑框截取相应长度，从左边算起

	UpdateData(false);
}
