
// FunctionIndicatorView.h: CFunctionIndicatorView 类的接口
//

#pragma once

#include "afxdialogex.h"
#include "FunctionIndicatorDoc.h"

//画面移动时的状态数据
struct MoveStatus {
	CPoint point;   //当前鼠标点
	double minX, maxX;  //x轴范围
	double minY, maxY;  //y轴范围
};


class CFunctionIndicatorView : public CView
{
protected: // 仅从序列化创建
	CFunctionIndicatorView() noexcept;
	DECLARE_DYNCREATE(CFunctionIndicatorView)

// 特性
public:
	CFunctionIndicatorDoc* GetDocument() const;

// 操作
public:

// 重写
public:
	virtual void OnDraw(CDC* pDC);  // 重写以绘制该视图
	virtual BOOL PreCreateWindow(CREATESTRUCT& cs);
protected:

// 实现
public:
	virtual ~CFunctionIndicatorView();
#ifdef _DEBUG
	virtual void AssertValid() const;
	virtual void Dump(CDumpContext& dc) const;
#endif

protected:

private:
	//在视图坐标系上绘制函数图像的区域坐标范围(真实坐标，非函数坐标)
	double m_Top, m_Bottom, m_Left, m_Right;
	MoveStatus m_MoveStart;  //移动模式下，鼠标一开始点击时的初始状态
	int m_CurFuncNum;  //当前鼠标位置指向的函数的序号

private:
	//实现x轴上的数值在不同坐标系的转换，第二参数为转换模式
	//true-表示视图坐标系到函数坐标系的转换
	//false-表示函数坐标系到视图坐标系的转换
	double TransformX(double x, bool mode);

	//实现y轴上的数值在不同坐标系的转换，第二参数为转换模式
	//true-表示视图坐标系到函数坐标系的转换
	//false-表示函数坐标系到视图坐标系的转换
	double TransformY(double y, bool mode);

	void SetShowWindow(); //设置函数图像的有效绘画区域
	void DrawEdge(CDC* pDC);  //画绘画区域的边框
	void MarkCoordinateValue(CDC* pDC); //标注坐标系的各个坐标值
	void DrawGrid(CDC* pDC);  //画绘画区域内的网格
	void DrawAxis(CDC* pDC); //画绘画区域内的X-Y坐标轴
	void DrawFunction(CDC* pDC);  //绘制函数图像
	void ShowFuncExpression(CDC* pDC);  //显示函数表达式 
	void AmplifyImage();  //放大坐标轴图像
	void ShrinkImage();  //缩小坐标轴图像
	void DoubleBufferDraw();  //利用双缓冲机制进行绘画，反正频繁绘画时的图像闪烁
	double GetDistacne(SML::FunctionPoint &a, SML::FunctionPoint &b);  //获得两个函数图像点的距离
	SML::FunctionPoint GetClosestPoint(CPoint point);  //获得离光标点最近的函数图像点，没有则返回INF点
	bool ShowFunctionPoint(CPoint point);  //在图像中显示光标所在的函数坐标点信息,false表示光标不在函数图像上
	void ShowFunctionNum();  //在状态栏中显示光标处的函数序号
	void ShowFunctionInformation(CPoint point);  //显示离光标点最近的函数的相关信息,包括函数点坐标和函数序号
	void ShowImagePoint(CPoint point);	//在状态栏显示光标所在的图像坐标
	

// 生成的消息映射函数
protected:
	DECLARE_MESSAGE_MAP()
public:
	afx_msg BOOL OnMouseWheel(UINT nFlags, short zDelta, CPoint pt);
	afx_msg void OnLButtonDown(UINT nFlags, CPoint point);
	afx_msg void OnLButtonUp(UINT nFlags, CPoint point);
	afx_msg void OnMouseMove(UINT nFlags, CPoint point);
};

#ifndef _DEBUG  // FunctionIndicatorView.cpp 中的调试版本
inline CFunctionIndicatorDoc* CFunctionIndicatorView::GetDocument() const
   { return reinterpret_cast<CFunctionIndicatorDoc*>(m_pDocument); }
#endif

