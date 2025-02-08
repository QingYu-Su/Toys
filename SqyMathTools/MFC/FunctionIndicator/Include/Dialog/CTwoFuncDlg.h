#pragma once
#include "afxdialogex.h"
#include "../../../SqyMathLibrary/include/MathFunction.h"
#pragma comment(lib, "../x64/Debug/SqyMathLibrary.lib")
namespace SML = SqyMathLibrary;  //命名空间重命名

// CTwoFuncDlg 对话框

class CTwoFuncDlg : public CDialogEx
{
	DECLARE_DYNAMIC(CTwoFuncDlg)

public:
	CTwoFuncDlg(CString title, SML::MathExpression expressionX,
		CString expressionStrX, SML::MathExpression expressionY,
		CString expressionStrY, double left, double right, size_t precision = 1000,
		int lineWidth = 1, int lineType = PS_SOLID, COLORREF color = RGB(0, 0, 0),
		CWnd* pParent = nullptr);   // 标准构造函数
	virtual ~CTwoFuncDlg();

// 对话框数据
#ifdef AFX_DESIGN_TIME
	enum { IDD = IDD_TWOFUNC_DIALOG};
#endif

protected:
	virtual void DoDataExchange(CDataExchange* pDX);    // DDX/DDV 支持

public:
	BOOL OnInitDialog();
	SML::MathFunction* GetMathFunction();  //获得函数类对象
	CString GetExpressionStrX();  //返回X(t)函数表达式字符串
	CString GetExpressionStrY();  //返回X(t)函数表达式字符串
	size_t GetPrecision();  //获得函数精度
	double GetMin();  //获得函数定义域左端点
	double GetMax();  //获得函数定义域右端点
	int GetLineWidth();  //获得线宽
	int GetLineType();  //获得线型
	COLORREF GetLineColor();  //获得线颜色

private:
	CString m_Title;  //对话框标题
	SML::MathExpression m_ExpressionX;  //X(t)函数表达式
	SML::MathExpression m_ExpressionY;  //Y(t)函数表达式
	CString m_ExpressionStrX;  //X(t)函数表达式字符串
	CString m_ExpressionStrY;  //Y(t)函数表达式字符串
	double m_Left, m_Right;  //函数定义域
	size_t m_Precision;  //图像精细度
	int m_LineWidth;  //线宽
	int m_LineType;  //线型
	COLORREF m_LineColor;  //线颜色

	DECLARE_MESSAGE_MAP()
public:
	afx_msg void OnBnClickedButtonEditfuncX();
	afx_msg void OnBnClickedButtonEditfuncY();
	afx_msg void OnBnClickedButtonLineColor();
};
