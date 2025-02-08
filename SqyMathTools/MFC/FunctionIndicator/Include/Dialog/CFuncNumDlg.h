#pragma once
#include "afxdialogex.h"


// CFuncNumDlg 对话框

class CFuncNumDlg : public CDialogEx
{
	DECLARE_DYNAMIC(CFuncNumDlg)

public:
	CFuncNumDlg(CString title, CString tip, CWnd* pParent = nullptr);   // 标准构造函数
	virtual ~CFuncNumDlg();

// 对话框数据
#ifdef AFX_DESIGN_TIME
	enum { IDD = IDD_FUNC_NUM_DIALOG };
#endif

protected:
	virtual void DoDataExchange(CDataExchange* pDX);    // DDX/DDV 支持

public:
	BOOL OnInitDialog();
	int GetNum();

	DECLARE_MESSAGE_MAP()
private:
	int m_Num;  //输入框数字
	CString m_Tip;  //提示文本
	CString m_Title;  //对话框标题
	CStatic m_StaticTip;  //提示文本控件
};
