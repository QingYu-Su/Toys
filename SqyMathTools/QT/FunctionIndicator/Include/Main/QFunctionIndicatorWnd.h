#ifndef QFUNCTIONINDICATORWND_H
#define QFUNCTIONINDICATORWND_H

#include <QMainWindow>
#include <QPainter>
#include <QLabel>
#include "MathFunction.h"
namespace SML = SqyMathLibrary;  //命名空间重命名



QT_BEGIN_NAMESPACE
namespace Ui { class QFunctionIndicatorWnd; }
QT_END_NAMESPACE

class QFunctionIndicatorWnd : public QMainWindow
{
    Q_OBJECT

    //绘制所需的函数数据
    struct DrawFuncData {
        std::vector<QString> expressionStr;  //函数表达式字符串，由于可能表达式不止一个，故用一个数组存放
        SML::FunctionMap *drawPoint;  //函数绘制点
        size_t precision; //函数精细度
        int lineWidth;  //线宽
        int lineType;  //线型
        QColor lineColor;  //线颜色
    };

    //移动模式
    enum MoveMode {
        BAN, //禁止移动
        MOVE,  //允许移动
        MOVING  //正在移动
    };

    //画面移动时的状态数据
    struct MoveStatus {
        QPoint point;   //当前鼠标点
        double minX, maxX;  //x轴范围
        double minY, maxY;  //y轴范围
    };

    //函数点绘制数据
    struct FuncPointDrawData{
        QRect textRect;  //绘制矩形
        QString msg;  //绘制字符串
    };

public:
    QFunctionIndicatorWnd(QWidget *parent = nullptr);
    ~QFunctionIndicatorWnd();
    void paintEvent(QPaintEvent *);
    void wheelEvent(QWheelEvent *event);
    void mousePressEvent(QMouseEvent *event);
    void mouseReleaseEvent(QMouseEvent *event);
    void mouseMoveEvent(QMouseEvent *event);

private slots:
    void on_m_Add_NormalFunc_triggered();

    void on_m_Add_PolarFunc_triggered();

    void on_m_Add_TwoFuc_triggered();

    void on_m_Alter_Func_triggered();

    void on_m_Del_Func_triggered();

    void on_m_Del_AllFunc_triggered();

    void on_m_Set_XY_Range_triggered();

    void on_m_Set_Move_triggered();

    void on_m_Set_Single_triggered();

    void on_m_Show_Edge_triggered();

    void on_m_Show_Axis_triggered();

    void on_m_Show_Grid_triggered();

    void on_m_Show_FuncInfo_triggered();

    void on_m_Show_FuncList_triggered();

    void on_m_View_StatusBar_triggered();

private:
    Ui::QFunctionIndicatorWnd *ui;

public:
    void SetRange(double minX, double maxX, double minY, double maxY);  //设置当前X-Y轴范围
    void DelFunction(int num); //删除特定序号函数，序号由1开始
    void ClearFunction(); //清空所有函数

    void AlterFunction(int num); //修改特定序号函数的参数，序号由1开始

    //修改链表中的普通函数
    void AlterNormalFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw);

    //修改链表中的极坐标函数
    void AlterPolarFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw);

    //修改链表中的参数方程函数
    void AlterTwoFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw);

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
    void DrawEdge(QPainter* p);  //画绘画区域的边框
    void MarkCoordinateValue(QPainter* p); //标注坐标系的各个坐标值
    void DrawGrid(QPainter* p);  //画绘画区域内的网格
    void DrawAxis(QPainter* p); //画绘画区域内的X-Y坐标轴
    void DrawFunction(QPainter* p);  //绘制函数图像
    void ShowFuncExpression(QPainter* p);  //显示函数表达式
    void AmplifyImage();  //放大坐标轴图像
    void ShrinkImage();  //缩小坐标轴图像
    double GetDistacne(SML::FunctionPoint &a, SML::FunctionPoint &b);  //获得两个函数图像点的距离
    SML::FunctionPoint GetClosestPoint(QPoint point);  //获得离光标点最近的函数图像点，没有则返回INF点
    bool GetFunctionPoint(QPoint point);  //获得图像中光标所在的函数坐标点信息,false表示光标不在函数图像上
    void ShowFunctionNum();  //在状态栏中显示光标处的函数序号
    void ShowFunctionInformation(QPoint point);  //显示离光标点最近的函数的相关信息,包括函数点坐标和函数序号
    void ShowImagePoint(QPoint point);	//在状态栏显示光标所在的图像坐标

private:
    bool m_ShowEdge; //显示边框
    bool m_ShowAxis; //显示坐标轴
    bool m_ShowGrid; //显示网格
    bool m_Single;   //显示单个函数
    bool m_ShowFuncInfo;//鼠标接近函数线时显示其函数信息，包括函数点坐标和函数序号
    bool m_ShowFuncList; //在函数图像右侧显示函数表达式列表
    bool m_ShowStatusBar; //显示状态栏

    //函数链表，保存了所有添加的函数以及一系列可供调用的方法
    std::list<SML::MathFunction*> m_FunctionList;

    //绘制函数数据链表，仅保存绘制图像所需要的必要数据，用来提供绘图，与函数链表一一对应
    std::list<DrawFuncData*> m_DrawDataList;

    double m_MinX, m_MaxX; //当前x轴范围
    double m_MinY, m_MaxY;//当前y轴范围
    MoveMode m_MoveMode;  //当前移动模式

    //在视图坐标系上绘制函数图像的区域坐标范围(真实坐标，非函数坐标)
    double m_Top, m_Bottom, m_Left, m_Right;
    MoveStatus m_MoveStart;  //移动模式下，鼠标一开始点击时的初始状态
    int m_CurFuncNum;  //当前鼠标位置指向的函数的序号

    //状态栏控件
    QLabel *m_StatusBarItem1;
    QLabel *m_StatusBarItem2;

    FuncPointDrawData m_FuncPointDrawData;  //函数点绘制数据
};
#endif // QFUNCTIONINDICATORWND_H
