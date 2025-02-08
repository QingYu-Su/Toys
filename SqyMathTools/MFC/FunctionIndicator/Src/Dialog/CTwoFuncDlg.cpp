// CTwoFuncDlg.cpp: 实现文件
//

#include "../../pch.h"
#include "../../Include/Main/FunctionIndicator.h"
#include "afxdialogex.h"
#include "../../Include/Dialog/CTwoFuncDlg.h"
#include "../../Include/Dialog/CFunctionDlg.h"


// CTwoFuncDlg 对话框

IMPLEMENT_DYNAMIC(CTwoFuncDlg, CDialogEx)

CTwoFuncDlg::CTwoFuncDlg(CString title, SML::MathExpression expressionX,
	CString expressionStrX, SML::MathExpression expressionY,
	CString expressionStrY, double left, double right, size_t precision,
	int lineWidth, int lineType, COLORREF color, CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_TWOFUNC_DIALOG, pParent)
	, m_Title(title)
	, m_ExpressionX(expressionX)
	, m_ExpressionStrX(expressionStrX)
	, m_ExpressionY(expressionY)
	, m_ExpressionStrY(expressionStrY)
	, m_Precision(precision)
	, m_LineWidth(lineWidth)
	, m_LineType(lineType)
	, m_LineColor(color)
	, m_Left(left)
	, m_Right(right)
{

}


CTwoFuncDlg::~CTwoFuncDlg()
{
}

void CTwoFuncDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	DDX_Text(pDX, IDC_EDIT_MIN, m_Left);
	DDX_Text(pDX, IDC_EDIT_MAX, m_Right);
	DDX_Text(pDX, IDC_EDIT_PRECISION, m_Precision);
	DDX_Text(pDX, IDC_EDIT_LINE_WIDTH, m_LineWidth);
	DDX_CBIndex(pDX, IDC_COMBO_LINE_TYPE, m_LineType);
	DDX_Text(pDX, IDC_EDIT_FUNCTION_X, m_ExpressionStrX);
	DDX_Text(pDX, IDC_EDIT_FUNCTION_Y, m_ExpressionStrY);
}

BOOL CTwoFuncDlg::OnInitDialog() {
	CDialogEx::OnInitDialog();
	this->SetWindowTextA(m_Title);  //设置对话框标题
	return TRUE;
}

SML::MathFunction* CTwoFuncDlg::GetMathFunction() {
	//创建X(t)函数表达式
	SML::FunctionExpression funcExpressionX(this->m_ExpressionX, 'x', 't', this->m_Left, this->m_Right);

	//创建Y(t)函数表达式
	SML::FunctionExpression funcExpressionY(this->m_ExpressionY, 'y', 't', this->m_Left, this->m_Right);

	return new SML::TwoFunction(funcExpressionX, funcExpressionY); //返回参数方程函数对象

}

CString CTwoFuncDlg::GetExpressionStrX() {
	return this->m_ExpressionStrX;
}

CString CTwoFuncDlg::GetExpressionStrY() {
	return this->m_ExpressionStrY;
}

size_t CTwoFuncDlg::GetPrecision() {
	return this->m_Precision;
}

double CTwoFuncDlg::GetMin() {
	return this->m_Left;
}

double CTwoFuncDlg::GetMax() {
	return this->m_Right;
}

int CTwoFuncDlg::GetLineWidth() {
	return this->m_LineWidth;
}

int CTwoFuncDlg::GetLineType() {
	return this->m_LineType;
}

COLORREF CTwoFuncDlg::GetLineColor() {
	return this->m_LineColor;
}



BEGIN_MESSAGE_MAP(CTwoFuncDlg, CDialogEx)
	ON_BN_CLICKED(IDC_BUTTON_EDITFUNC_X, &CTwoFuncDlg::OnBnClickedButtonEditfuncX)
	ON_BN_CLICKED(IDC_BUTTON_EDITFUNC_Y, &CTwoFuncDlg::OnBnClickedButtonEditfuncY)
	ON_BN_CLICKED(IDC_BUTTON_LINE_COLOR, &CTwoFuncDlg::OnBnClickedButtonLineColor)
END_MESSAGE_MAP()


// CTwoFuncDlg 消息处理程序


void CTwoFuncDlg::OnBnClickedButtonEditfuncX()
{
	UpdateData();

	//弹出函数表达式编辑对话框
	CFunctionDlg dlg('t', this->m_ExpressionX, this->m_ExpressionStrX);
	if (dlg.DoModal() == IDOK) {
		this->m_ExpressionX = dlg.GetExpression();  //同步函数表达式
		this->m_ExpressionStrX = dlg.GetEditStr();  //同步函数表达式字符串
	}
	UpdateData(false);
}


void CTwoFuncDlg::OnBnClickedButtonEditfuncY()
{
	UpdateData();

	//弹出函数表达式编辑对话框
	CFunctionDlg dlg('t', this->m_ExpressionY, this->m_ExpressionStrY);
	if (dlg.DoModal() == IDOK) {
		this->m_ExpressionY = dlg.GetExpression();  //同步函数表达式
		this->m_ExpressionStrY = dlg.GetEditStr();  //同步函数表达式字符串
	}
	UpdateData(false);
}


void CTwoFuncDlg::OnBnClickedButtonLineColor()
{
	//弹出设置颜色对话框
	CColorDialog dlg;
	if (dlg.DoModal() == IDOK) {
		m_LineColor = dlg.GetColor();
	}
}
