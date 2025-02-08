#pragma once
#include "afxdialogex.h"


// CSetXYRangeDlg 对话框

class CSetXYRangeDlg : public CDialogEx
{
	DECLARE_DYNAMIC(CSetXYRangeDlg)

public:
	CSetXYRangeDlg(CWnd* pParent = nullptr);   // 标准构造函数
	virtual ~CSetXYRangeDlg();

// 对话框数据
#ifdef AFX_DESIGN_TIME
	enum { IDD = IDD_SET_RANGE_DIALOG };
#endif

protected:
	virtual void DoDataExchange(CDataExchange* pDX);    // DDX/DDV 支持

public:
	double GetMinX();
	double GetMaxX();
	double GetMinY();
	double GetMaxY();

	DECLARE_MESSAGE_MAP()
private:
	double m_MinX;
	double m_MaxX;
	double m_MinY;
	double m_MaxY;
};
