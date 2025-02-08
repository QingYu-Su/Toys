// CPolarFuncDlg.cpp: 实现文件
//

#include "../../pch.h"
#include "../../Include/Main/FunctionIndicator.h"
#include "afxdialogex.h"
#include "../../Include/Dialog/CPolarFuncDlg.h"
#include "../../Include/Dialog/CFunctionDlg.h"


// CPolarFuncDlg 对话框

IMPLEMENT_DYNAMIC(CPolarFuncDlg, CDialogEx)

CPolarFuncDlg::CPolarFuncDlg(CString title, SML::MathExpression expression,
	CString expressionStr, double left, double right, size_t precision,
	int lineWidth, int lineType, COLORREF color, CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_POLARFUNC_DIALOG, pParent)
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

CPolarFuncDlg::~CPolarFuncDlg()
{
}

void CPolarFuncDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	DDX_Text(pDX, IDC_EDIT_FUNCTION, m_ExpressionStr);
	DDX_Text(pDX, IDC_EDIT_MIN, m_Left);
	DDX_Text(pDX, IDC_EDIT_MAX, m_Right);
	DDX_Text(pDX, IDC_EDIT_PRECISION, m_Precision);
	DDX_Text(pDX, IDC_EDIT_LINE_WIDTH, m_LineWidth);
	DDX_CBIndex(pDX, IDC_COMBO_LINE_TYPE, m_LineType);
}

BOOL CPolarFuncDlg::OnInitDialog() {
	CDialogEx::OnInitDialog();
	this->SetWindowTextA(m_Title);  //设置对话框标题
	return TRUE;
}

SML::MathFunction* CPolarFuncDlg::GetMathFunction() {
	//创建极坐标函数表达式
	SML::FunctionExpression funcExpression(this->m_Expression, 'r', 'a', this->m_Left, this->m_Right);

	return new SML::PolarFunction(funcExpression); //返回极坐标函数对象

}

CString CPolarFuncDlg::GetExpressionStr() {
	return this->m_ExpressionStr;
}

size_t CPolarFuncDlg::GetPrecision() {
	return this->m_Precision;
}

double CPolarFuncDlg::GetMin() {
	return this->m_Left;
}

double CPolarFuncDlg::GetMax() {
	return this->m_Right;
}

int CPolarFuncDlg::GetLineWidth() {
	return this->m_LineWidth;
}

int CPolarFuncDlg::GetLineType() {
	return this->m_LineType;
}

COLORREF CPolarFuncDlg::GetLineColor() {
	return this->m_LineColor;
}



BEGIN_MESSAGE_MAP(CPolarFuncDlg, CDialogEx)
	ON_BN_CLICKED(IDC_BUTTON_EDITFUNC, &CPolarFuncDlg::OnBnClickedButtonEditfunc)
	ON_BN_CLICKED(IDC_BUTTON_LINE_COLOR, &CPolarFuncDlg::OnBnClickedButtonPenColor)
END_MESSAGE_MAP()

void CPolarFuncDlg::OnBnClickedButtonEditfunc()
{
	UpdateData();

	//弹出函数表达式编辑对话框
	CFunctionDlg dlg('a', this->m_Expression, this->m_ExpressionStr);
	if (dlg.DoModal() == IDOK) {
		this->m_Expression = dlg.GetExpression();  //同步函数表达式
		this->m_ExpressionStr = dlg.GetEditStr();  //同步函数表达式字符串
	}
	UpdateData(false);
}


void CPolarFuncDlg::OnBnClickedButtonPenColor()
{
	//弹出设置颜色对话框
	CColorDialog dlg;
	if (dlg.DoModal() == IDOK) {
		m_LineColor = dlg.GetColor();
	}
}


// CPolarFuncDlg 消息处理程序
