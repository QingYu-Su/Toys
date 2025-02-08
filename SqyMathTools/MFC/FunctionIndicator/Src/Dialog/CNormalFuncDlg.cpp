// CNormalFuncDlg.cpp: 实现文件
//

#include "../../pch.h"
#include "../../Include/Main/FunctionIndicator.h"
#include "afxdialogex.h"
#include "../../Include/Dialog/CNormalFuncDlg.h"
#include "../../Include/Dialog/CFunctionDlg.h"


// CNormalFuncDlg 对话框

IMPLEMENT_DYNAMIC(CNormalFuncDlg, CDialogEx)

CNormalFuncDlg::CNormalFuncDlg(CString title, SML::MathExpression expression,
	CString expressionStr, double left, double right, size_t precision,
	int lineWidth, int lineType, COLORREF color, CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_NORMALFUNC_DIALOG, pParent)
	, m_Title(title)
	, m_Expression(expression)
	, m_ExpressionStr(expressionStr)
	, m_Precision(precision)
	, m_LineWidth(lineWidth)
	, m_LineType(lineType)
	, m_LineColor(color)
	, m_Left(left)
	, m_Right(right)
{
}

CNormalFuncDlg::~CNormalFuncDlg()
{
}

void CNormalFuncDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	DDX_Text(pDX, IDC_EDIT_FUNCTION, m_ExpressionStr);
	DDX_Text(pDX, IDC_EDIT_MIN, m_Left);
	DDX_Text(pDX, IDC_EDIT_MAX, m_Right);
	DDX_Text(pDX, IDC_EDIT_PRECISION, m_Precision);
	DDX_Text(pDX, IDC_EDIT_LINE_WIDTH, m_LineWidth);
	DDX_CBIndex(pDX, IDC_COMBO_LINE_TYPE, m_LineType);
}

BOOL CNormalFuncDlg::OnInitDialog() {
	CDialogEx::OnInitDialog();
	this->SetWindowTextA(m_Title);  //设置对话框标题
	return TRUE;
}

SML::FunctionExpression CNormalFuncDlg::GetFunctionExpression() {
	//返回一个封装好了的函数表达式
	return SML::FunctionExpression(this->m_Expression, 'y', 'x', this->m_Left, this->m_Right);
}

SML::MathFunction* CNormalFuncDlg::GetMathFunction() {
	
	return new SML::NormalFunction(this->GetFunctionExpression()); //返回普通函数对象

}

CString CNormalFuncDlg::GetExpressionStr() {
	return this->m_ExpressionStr;
}

size_t CNormalFuncDlg::GetPrecision() {
	return this->m_Precision;
}

double CNormalFuncDlg::GetMin() {
	return this->m_Left;
}

double CNormalFuncDlg::GetMax() {
	return this->m_Right;
}

int CNormalFuncDlg::GetLineWidth() {
	return this->m_LineWidth;
}

int CNormalFuncDlg::GetLineType() {
	return this->m_LineType;
}

COLORREF CNormalFuncDlg::GetLineColor() {
	return this->m_LineColor;
}



BEGIN_MESSAGE_MAP(CNormalFuncDlg, CDialogEx)
	ON_BN_CLICKED(IDC_BUTTON_EDITFUNC, &CNormalFuncDlg::OnBnClickedButtonEditfunc)
	ON_BN_CLICKED(IDC_BUTTON_LINE_COLOR, &CNormalFuncDlg::OnBnClickedButtonPenColor)
END_MESSAGE_MAP()


// CNormalFuncDlg 消息处理程序


void CNormalFuncDlg::OnBnClickedButtonEditfunc()
{
	UpdateData();
	
	//弹出函数表达式编辑对话框
	CFunctionDlg dlg('x', this->m_Expression, this->m_ExpressionStr);
	if (dlg.DoModal() == IDOK) {
		this->m_Expression = dlg.GetExpression();  //同步函数表达式
		this->m_ExpressionStr = dlg.GetEditStr();  //同步函数表达式字符串
	}
	UpdateData(false);
}


void CNormalFuncDlg::OnBnClickedButtonPenColor()
{
	//弹出设置颜色对话框
	CColorDialog dlg;
	if (dlg.DoModal() == IDOK) {
		m_LineColor = dlg.GetColor();
	}
}
