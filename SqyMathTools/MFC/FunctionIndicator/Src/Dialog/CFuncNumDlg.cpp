// CFuncNumDlg.cpp: 实现文件
//

#include "../../pch.h"
#include "../../Include/Main/FunctionIndicator.h"
#include "afxdialogex.h"
#include "../../Include/Dialog/CFuncNumDlg.h"


// CFuncNumDlg 对话框

IMPLEMENT_DYNAMIC(CFuncNumDlg, CDialogEx)

CFuncNumDlg::CFuncNumDlg(CString title, CString tip, CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_FUNC_NUM_DIALOG, pParent)
	, m_Num(1)
	, m_Title(title)
	, m_Tip(tip)
{

}

CFuncNumDlg::~CFuncNumDlg()
{
}

void CFuncNumDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	DDX_Text(pDX, IDC_EDIT_NUM, m_Num);
	DDX_Control(pDX, IDC_TIP, m_StaticTip);
}

BOOL CFuncNumDlg::OnInitDialog() {
	CDialogEx::OnInitDialog();
	this->SetWindowTextA(this->m_Title); //设置标题
	this->m_StaticTip.SetWindowTextA(this->m_Tip);  //设置提示文本
	return TRUE;
}

int CFuncNumDlg::GetNum() {
	return this->m_Num;
}


BEGIN_MESSAGE_MAP(CFuncNumDlg, CDialogEx)
END_MESSAGE_MAP()


// CFuncNumDlg 消息处理程序
