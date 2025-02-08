// CSetXYRangeDlg.cpp: 实现文件
//

#include "../../pch.h"
#include "../../Include/Main/FunctionIndicator.h"
#include "afxdialogex.h"
#include "../../Include/Dialog/CSetXYRangeDlg.h"


// CSetXYRangeDlg 对话框

IMPLEMENT_DYNAMIC(CSetXYRangeDlg, CDialogEx)

CSetXYRangeDlg::CSetXYRangeDlg(CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_SET_RANGE_DIALOG, pParent)
	, m_MinX(0)
	, m_MaxX(0)
	, m_MinY(0)
	, m_MaxY(0)
{

}

CSetXYRangeDlg::~CSetXYRangeDlg()
{
}

void CSetXYRangeDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	DDX_Text(pDX, IDC_EDIT_MIN_X, m_MinX);
	DDX_Text(pDX, IDC_EDIT_MAX_X, m_MaxX);
	DDX_Text(pDX, IDC_EDIT_MIN_Y, m_MinY);
	DDX_Text(pDX, IDC_EDIT_MAX_Y, m_MaxY);
}

double CSetXYRangeDlg::GetMinX() {
	return this->m_MinX;
}

double CSetXYRangeDlg::GetMaxX() {
	return this->m_MaxX;
}

double CSetXYRangeDlg::GetMinY() {
	return this->m_MinY;
}

double CSetXYRangeDlg::GetMaxY() {
	return this->m_MaxY;
}


BEGIN_MESSAGE_MAP(CSetXYRangeDlg, CDialogEx)
END_MESSAGE_MAP()


// CSetXYRangeDlg 消息处理程序
