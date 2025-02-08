
// FunctionIndicatorDoc.h: CFunctionIndicatorDoc 类的接口
//


#pragma once
#include <list>
#include <string>
#include <vector>
#include <algorithm>
#include "afxdialogex.h"
#include "../../../SqyMathLibrary/include/MathFunction.h"
#pragma comment(lib, "../x64/Debug/SqyMathLibrary.lib")
namespace SML = SqyMathLibrary;  //命名空间重命名

//绘制所需的必要函数数据，简化Doc与View的数据传输，保证View只绘画而不计算
struct DrawFuncData { 
	std::vector<CString> expressionStr;  //函数表达式字符串，由于可能表达式不止一个，故用一个数组存放
	SML::FunctionMap *drawPoint;  //函数绘制点
	size_t precision; //函数精细度
	int lineWidth;  //线宽
	int lineType;  //线型
	COLORREF lineColor;  //线颜色
};

//移动模式
enum MoveMode {
	BAN, //禁止移动
	MOVE,  //允许移动
	MOVING  //正在移动
};


class CFunctionIndicatorDoc : public CDocument
{
protected: // 仅从序列化创建
	CFunctionIndicatorDoc() noexcept;
	DECLARE_DYNCREATE(CFunctionIndicatorDoc)

// 特性
public:

// 操作
public:

// 重写
public:
	virtual BOOL OnNewDocument();
	virtual void Serialize(CArchive& ar);
#ifdef SHARED_HANDLERS
	virtual void InitializeSearchContent();
	virtual void OnDrawThumbnail(CDC& dc, LPRECT lprcBounds);
#endif // SHARED_HANDLERS

// 实现
public:
	virtual ~CFunctionIndicatorDoc();
#ifdef _DEBUG
	virtual void AssertValid() const;
	virtual void Dump(CDumpContext& dc) const;
#endif

protected:

// 生成的消息映射函数
protected:
	DECLARE_MESSAGE_MAP()

#ifdef SHARED_HANDLERS
	// 用于为搜索处理程序设置搜索内容的 Helper 函数
	void SetSearchContent(const CString& value);
#endif // SHARED_HANDLERS

public:
	double GetMinX();  //获得x范围最小值
	double GetMaxX();  //获得x范围最大值
	double GetMinY();  //获得y范围最小值
	double GetMaxY();  //获得y范围最大值
	std::list<DrawFuncData*> GetDrawDataList();  //获得绘制函数所需的所有数据
	void SetRange(double minX, double maxX, double minY, double maxY);  //设置当前X-Y轴范围
	void UpdateFunction();  //更新函数，重新计算函数图像绘制点，仅对普通函数更新
	void DelFunction(int num); //删除特定序号函数，序号由1开始
	void ClearFunction(); //清空所有函数

	void AlterFunction(int num); //修改特定序号函数的参数，序号由1开始

	//修改链表中的普通函数
	void AlterNormalFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw); 
	
	//修改链表中的极坐标函数
	void AlterPolarFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw);  
	
	//修改链表中的参数方程函数
	void AlterTwoFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw);  
	
	MoveMode GetMoveMode();  //获得当前移动模式
	void SetMoveMode(MoveMode mode);  //设置当前移动模式
	bool IsShowEdge(); //是否显示边框
	bool IsShowAxis(); //是否显示坐标轴
	bool IsShowGrid();  //是否显示网格
	bool IsShowFuncInfo(); //是否显示函数坐标点
	bool IsShowFuncList();  //是否显示函数列表

private:
	bool m_ShowEdge; //显示边框
	bool m_ShowAxis; //显示坐标轴
	bool m_ShowGrid; //显示网格
	bool m_Single;   //显示单个函数
	bool m_ShowFuncInfo;//鼠标接近函数线时显示其函数信息，包括函数点坐标和函数序号
	bool m_ShowFuncList; //在函数图像右侧显示函数表达式列表

	double m_MinX, m_MaxX; //当前x轴范围 
	double m_MinY, m_MaxY;//当前y轴范围
	MoveMode m_MoveMode;  //当前移动模式

	//函数链表，保存了所有添加的函数以及一系列可供调用的方法
	std::list<SML::MathFunction*> m_FunctionList;

	//绘制函数数据链表，仅保存绘制图像所需要的必要数据，用来提供View绘图，与函数链表一一对应
	std::list<DrawFuncData*> m_DrawDataList;  

public:
	afx_msg void OnAddNormalFunc();
	afx_msg void OnAddPolarFunc();
	afx_msg void OnAddTwoFunc();
	afx_msg void OnDelFunc();
	afx_msg void OnDelAllFunc();
	afx_msg void OnSetXyRange();
	afx_msg void OnEdge();
	afx_msg void OnAxis();
	afx_msg void OnGrid();
	afx_msg void OnUpdateEdge(CCmdUI* pCmdUI);
	afx_msg void OnUpdateAxis(CCmdUI* pCmdUI);
	afx_msg void OnUpdateGrid(CCmdUI* pCmdUI);
	afx_msg void OnMove();
	afx_msg void OnUpdateMove(CCmdUI* pCmdUI);
	afx_msg void OnFuncInfo();
	afx_msg void OnUpdateFunctionInfo(CCmdUI* pCmdUI);
	afx_msg void OnFuncList();
	afx_msg void OnUpdateFuncList(CCmdUI* pCmdUI);
	afx_msg void OnAlterFunc();
	afx_msg void OnSingle();
	afx_msg void OnUpdateSingle(CCmdUI* pCmdUI);
};
