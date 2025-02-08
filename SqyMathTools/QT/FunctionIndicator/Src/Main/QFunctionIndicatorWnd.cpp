#include "Main/QFunctionIndicatorWnd.h"
#include "ui_QFunctionIndicatorWnd.h"
#include "Dialog/QNormalFuncDlg.h"
#include "Dialog/QPolarFuncDlg.h"
#include "Dialog/QTwoFuncDlg.h"
#include "Dialog/QFuncNumDlg.h"
#include "Dialog/QSetXYRangeDlg.h"
#include <QDebug>
#include <QLabel>
#include <QMessageBox>
#include <QWheelEvent>

QFunctionIndicatorWnd::QFunctionIndicatorWnd(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::QFunctionIndicatorWnd)
{
    //初始化
    ui->setupUi(this);
    this->m_ShowGrid = true;
    this->m_ShowAxis = true;
    this->m_ShowEdge = true;
    this->m_Single = false;
    this->m_ShowFuncInfo = true;
    this->m_ShowFuncList = true;
    this -> m_ShowStatusBar = true;
    this->m_MinX = -10;
    this->m_MaxX = 10;
    this->m_MinY = -5;
    this->m_MaxY = 5;
    this->m_MoveMode = BAN;

    //设置控件初始状态
    ui -> m_Show_Edge -> setChecked(true);
    ui -> m_Show_Axis -> setChecked(true);
    ui -> m_Show_Grid -> setChecked(true);
    ui -> m_Show_FuncInfo -> setChecked(true);
    ui -> m_Show_FuncList -> setChecked(true);
    ui -> m_View_StatusBar -> setChecked(true);

    //设置状态栏
    this -> m_StatusBarItem1 = new QLabel(this);
    this -> m_StatusBarItem2 = new QLabel(this);

    this -> m_StatusBarItem1 -> setFrameStyle(QFrame::Box | QFrame::Sunken);
    this -> m_StatusBarItem2 -> setFrameStyle(QFrame::Box | QFrame::Sunken);

    this -> m_StatusBarItem1 -> setText("当前光标位置");
    this -> m_StatusBarItem2 -> setText("当前函数序号");

    //设置状态栏控件固定大小
    m_StatusBarItem2 -> setFixedSize(m_StatusBarItem2 -> size());

    ui-> statusbar -> addPermanentWidget(this->m_StatusBarItem1);
    ui-> statusbar -> addPermanentWidget(this->m_StatusBarItem2);

    //设置鼠标移动时自动捕获
    this -> centralWidget() -> setMouseTracking(true);
    setMouseTracking(true);

}

QFunctionIndicatorWnd::~QFunctionIndicatorWnd()
{
    delete ui;
}

void QFunctionIndicatorWnd::paintEvent(QPaintEvent *){
    QPainter *p = new QPainter(this);

    //绘画前必须设置绘画区域
    this->SetShowWindow();

    //画边框
    if ( m_ShowEdge == true ) this->DrawEdge(p);

    //标注坐标信息
    this->MarkCoordinateValue(p);

    //画网格
    if ( m_ShowGrid == true ) this->DrawGrid(p);

    //画坐标轴
    if ( m_ShowAxis == true ) this->DrawAxis(p);

    //画函数图像
    this->DrawFunction(p);

    //在视图中展示函数表达式信息
    if ( m_ShowFuncList == true ) ShowFuncExpression(p);

    //显示状态栏
    if ( this -> m_ShowStatusBar == true ){
        ui -> statusbar -> setVisible(true);
    }else{
        ui -> statusbar -> setVisible(false);
    }

    //绘制函数点数据，非空则绘制
    if ( m_FuncPointDrawData.msg.isEmpty() == false ){
        p -> drawText(m_FuncPointDrawData.textRect, Qt::AlignCenter, m_FuncPointDrawData.msg);
    }

    delete  p;
}

void QFunctionIndicatorWnd::wheelEvent(QWheelEvent *event){
     //向前滚动滚轮
    if ( event -> delta() > 0 ){
         this->AmplifyImage();  //放大图像
     }else{//向后滚动滚轮
         this->ShrinkImage();  //缩小图像
     }
}

void QFunctionIndicatorWnd::mousePressEvent(QMouseEvent *event){
    //左键被按下
    if ( event -> button() == Qt::LeftButton ){
        //当前允许移动
        if ( m_MoveMode == MOVE ){
            //设置移动模式为正在移动
            m_MoveMode = MOVING;

            //保存移动初始状态
            m_MoveStart.point = event -> pos();
            m_MoveStart.minX = m_MinX;
            m_MoveStart.maxX = m_MaxX;
            m_MoveStart.minY = m_MinY;
            m_MoveStart.maxY = m_MaxY;

            //设置光标
            setCursor(QCursor(Qt::ClosedHandCursor));
        }
    }
}

void QFunctionIndicatorWnd::mouseReleaseEvent(QMouseEvent *event){
    //左键被按下
    if ( event -> button() == Qt::LeftButton ){
        //当前正在移动
        if ( m_MoveMode == MOVING ){
            //设置移动模式为允许移动
            m_MoveMode = MOVE;

            //计算当前光标位置与初始状态光标位置的相对位移，需要将数值转换为函数坐标系
            double valueX = this->TransformX(event->pos().x(), true) - this->TransformX(this->m_MoveStart.point.x(), true);
            double valueY = this->TransformY(event->pos().y() , true) - this->TransformY(this->m_MoveStart.point.y(), true);

            //设置新的x-y范围
            SetRange(this->m_MoveStart.minX - valueX, this->m_MoveStart.maxX - valueX,
            this->m_MoveStart.minY - valueY, this->m_MoveStart.maxY - valueY);

            //设置光标
            setCursor(QCursor(Qt::OpenHandCursor));

            //刷新绘画
            update();
        }
    }
}

void QFunctionIndicatorWnd::mouseMoveEvent(QMouseEvent *event){

    //显示当前光标图像位置
    this->ShowImagePoint(event->pos());

    //显示函数信息
    this->ShowFunctionInformation(event->pos());

    //当前移动模式为正在移动
    if ( m_MoveMode == MOVING ){
        //计算当前光标位置与初始状态光标位置的相对位移，需要将数值转换为函数坐标系
        double valueX = this->TransformX(event->pos().x(), true) - this->TransformX(this->m_MoveStart.point.x(), true);
        double valueY = this->TransformY(event->pos().y() , true) - this->TransformY(this->m_MoveStart.point.y(), true);

        //设置新的x-y范围
        SetRange(this->m_MoveStart.minX - valueX, this->m_MoveStart.maxX - valueX,
        this->m_MoveStart.minY - valueY, this->m_MoveStart.maxY - valueY);

        //刷新绘画
        update();
    }
}

void QFunctionIndicatorWnd::SetRange(double minX, double maxX, double minY, double maxY) {
    this->m_MinX = minX;
    this->m_MaxX = maxX;
    this->m_MinY = minY;
    this->m_MaxY = maxY;
}

void QFunctionIndicatorWnd::DelFunction(int num) {
    //序号超过范围，直接返回
    if (num <= 0 || num > this->m_FunctionList.size()) return;

    //遍历链表，搜索对应位置函数
    std::list<SML::MathFunction*>::iterator itFunc = this->m_FunctionList.begin();
    std::list<DrawFuncData*>::iterator itDraw = this->m_DrawDataList.begin();
    while ( num > 1 ) {
        itFunc++;
        itDraw++;
        num--;
    }

    //释放资源
    delete *itFunc;
    delete *itDraw;

    //删除链表结点
    this->m_FunctionList.erase(itFunc);
    this->m_DrawDataList.erase(itDraw);

}

void QFunctionIndicatorWnd::ClearFunction() {
    std::list<SML::MathFunction*>::iterator itFunc = this->m_FunctionList.begin();
    std::list<DrawFuncData*>::iterator itDraw = this->m_DrawDataList.begin();

    //遍历所有链表结点，并释放资源
    while (itFunc != this->m_FunctionList.end()) {
        delete* itFunc;
        delete* itDraw;
        itFunc++;
        itDraw++;
    }

    //链表清空
    this->m_FunctionList.clear();
    this->m_DrawDataList.clear();
}

void QFunctionIndicatorWnd::AlterFunction(int num) {
    //序号超过范围，直接返回
    if (num <= 0 || num > this->m_FunctionList.size()) return;

    //遍历链表，搜索对应位置函数
    std::list<SML::MathFunction*>::iterator itFunc = this->m_FunctionList.begin();
    std::list<DrawFuncData*>::iterator itDraw = this->m_DrawDataList.begin();
    while (num > 1) {
        itFunc++;
        itDraw++;
        num--;
    }

    switch ((*itFunc)->GetType()) {
    case SML::Normal:
        this->AlterNormalFunc(itFunc, itDraw);
        break;
    case SML::Polar:
        this->AlterPolarFunc(itFunc, itDraw);
        break;
    case SML::Two:
        this->AlterTwoFunc(itFunc, itDraw);
        break;
    default:
        break;
    }
}

void QFunctionIndicatorWnd::AlterNormalFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw) {
    //获得普通函数表达式
    SML::FunctionExpression funcExp = *((*itFunc)->GetExpression().begin());

    //创建对话框，并传入特殊参数
    QNormalFuncDlg dlg("修改普通函数", funcExp.m_Expression,

        //表达式字符串前5个字符是辅助用的，需要去掉
        (*itDraw)->expressionStr[0].right((*itDraw)->expressionStr[0].size() - 5),

        funcExp.m_Left, funcExp.m_Right, (*itDraw)->precision,
        (*itDraw)->lineWidth, (*itDraw)->lineType, (*itDraw)->lineColor);

    //运行对话框
    while (dlg.exec() == QDialog::Accepted) {

        //获得函数类对象
        SML::MathFunction* pFunction = dlg.GetMathFunction();

        //新建绘画数据，传入必要参数
        DrawFuncData* dfd = new DrawFuncData;
        dfd->precision = dlg.GetPrecision();

        //计算函数图像数据，范围为X-Y轴显示范围，并将结果保存至绘画数据中
        //dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);
        //该方式虽然更为精准，但会导致图像移动时的卡顿，故直接采用下方以定义域范围绘制图像


        //计算函数图像数据，范围为该函数定义域，并将结果保存至绘画数据中
        //该方式与极坐标和参数方程的计算方式保持一致，图像移动会比较顺畅
        dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);

        //计算失败，弹出提示弹窗，释放相应资源并重新显示对话框
        if (pFunction->IsSuccess() == false) {
            QMessageBox msg(QMessageBox::Critical, "Function-Indicator", pFunction->GetError().c_str());
            msg.exec();
            delete pFunction;
            delete dfd;
            continue;
         }

        //将其他数据添加至绘画数据中
        dfd->expressionStr.push_back(QString("f(x)=") + dlg.GetExpressionStr());
        dfd->lineWidth = dlg.GetLineWidth();
        dfd->lineType = dlg.GetLineType();
        dfd->lineColor = dlg.GetLineColor();

        //释放旧的资源
        delete* itFunc;
        delete* itDraw;

        //指针指向新的结点
        (*itFunc) = pFunction;
        (*itDraw) = dfd;

        break;
    }

}

void QFunctionIndicatorWnd::AlterPolarFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw) {
    //获得极坐标函数表达式
    SML::FunctionExpression funcExp = *((*itFunc)->GetExpression().begin());

    //创建对话框，并传入特殊参数
    QPolarFuncDlg dlg("修改极坐标函数", funcExp.m_Expression,

        //表达式字符串前5个字符是辅助用的，需要去掉
        (*itDraw)->expressionStr[0].right((*itDraw)->expressionStr[0].size() - 5),

        funcExp.m_Left, funcExp.m_Right, (*itDraw)->precision,
        (*itDraw)->lineWidth, (*itDraw)->lineType, (*itDraw)->lineColor);

    while ( dlg.exec() == QDialog::Accepted ){
        //获得函数类对象
        SML::MathFunction* pFunction = dlg.GetMathFunction();

        //新建绘画数据，传入必要参数
        DrawFuncData* dfd = new DrawFuncData;
        dfd->precision = dlg.GetPrecision();

        //计算函数图像数据，范围为X-Y轴显示范围，并将结果保存至绘画数据中
        //dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);
        //该方式虽然更为精准，但会导致图像移动时的卡顿，故直接采用下方以定义域范围绘制图像


        //计算函数图像数据，范围为该函数定义域，并将结果保存至绘画数据中
        //该方式与极坐标和参数方程的计算方式保持一致，图像移动会比较顺畅
        dfd->drawPoint = pFunction->Calculate(-PI, PI, dfd->precision);

        //计算失败，弹出提示弹窗，释放相应资源并重新显示对话框
        if (pFunction->IsSuccess() == false) {
            QMessageBox msg(QMessageBox::Critical, "Function-Indicator", pFunction->GetError().c_str());
            msg.exec();
            delete pFunction;
            delete dfd;
            continue;
         }

        //将其他数据添加至绘画数据中
        dfd->expressionStr.push_back(QString("r(a)=") + dlg.GetExpressionStr());
        dfd->lineWidth = dlg.GetLineWidth();
        dfd->lineType = dlg.GetLineType();
        dfd->lineColor = dlg.GetLineColor();

        //释放旧的资源
        delete* itFunc;
        delete* itDraw;

        //指针指向新的结点
        (*itFunc) = pFunction;
        (*itDraw) = dfd;

        break;
    }
}

void QFunctionIndicatorWnd::AlterTwoFunc(std::list<SML::MathFunction*>::iterator itFunc, std::list<DrawFuncData*>::iterator itDraw) {
    //获得极坐标函数表达式
    std::list<SML::FunctionExpression> funcExpList = (*itFunc)->GetExpression();
    SML::FunctionExpression funcExpX = *(funcExpList.begin());
    SML::FunctionExpression funcExpY = *(++funcExpList.begin());

    //创建对话框，并传入特殊参数
    QTwoFuncDlg dlg("修改参数方程函数", funcExpX.m_Expression,

        //表达式字符串前5个字符是辅助用的，需要去掉
        (*itDraw)->expressionStr[0].right((*itDraw)->expressionStr[0].size() - 5),

        funcExpY.m_Expression,

        //表达式字符串前5个字符是辅助用的，需要去掉
        (*itDraw)->expressionStr[1].right((*itDraw)->expressionStr[1].size() - 5),

        funcExpX.m_Left, funcExpX.m_Right, (*itDraw)->precision,
        (*itDraw)->lineWidth, (*itDraw)->lineType, (*itDraw)->lineColor);

    while ( dlg.exec() == QDialog::Accepted ){
        //获得函数类对象
        SML::MathFunction* pFunction = dlg.GetMathFunction();

        //新建绘画数据，传入必要参数
        DrawFuncData* dfd = new DrawFuncData;
        dfd->precision = dlg.GetPrecision();

        //计算函数图像数据，范围为X-Y轴显示范围，并将结果保存至绘画数据中
        //dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);
        //该方式虽然更为精准，但会导致图像移动时的卡顿，故直接采用下方以定义域范围绘制图像


        //计算函数图像数据，范围为该函数定义域，并将结果保存至绘画数据中
        //该方式与极坐标和参数方程的计算方式保持一致，图像移动会比较顺畅
        dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);

        //计算失败，弹出提示弹窗，释放相应资源并重新显示对话框
        if (pFunction->IsSuccess() == false) {
            QMessageBox msg(QMessageBox::Critical, "Function-Indicator", pFunction->GetError().c_str());
            msg.exec();
            delete pFunction;
            delete dfd;
            continue;
         }

        //单函数模式下链表仅保存一个结点
        if (this->m_Single == true) {
            this->ClearFunction();
         }

        //将其他数据添加至绘画数据中
        dfd->expressionStr.push_back(QString("x(t)=") + dlg.GetExpressionStrX());
        dfd->expressionStr.push_back(QString("y(t)=") + dlg.GetExpressionStrY());
        dfd->lineWidth = dlg.GetLineWidth();
        dfd->lineType = dlg.GetLineType();
        dfd->lineColor = dlg.GetLineColor();

        //释放旧的资源
        delete* itFunc;
        delete* itDraw;

        //指针指向新的结点
        (*itFunc) = pFunction;
        (*itDraw) = dfd;

        break;
    }

}

double QFunctionIndicatorWnd::TransformX(double x, bool mode) {
    if (mode == true) {  //视图坐标系->函数坐标系
        double unit = (m_MaxX - m_MinX) / (this->m_Right - this->m_Left); //转换单元
        return m_MinX + (x - this->m_Left) * unit;
    }
    else { //函数坐标系->视图坐标系
        double unit = (this->m_Right - this->m_Left) / (m_MaxX - m_MinX); //转换单元
        return this->m_Left + (x - m_MinX) * unit;
    }
}

double QFunctionIndicatorWnd::TransformY(double y, bool mode) {

    if (mode == true) {  //视图坐标系->函数坐标系
        double unit = (m_MaxY - m_MinY) / (this->m_Bottom - this->m_Top);  //转换单元
        return m_MinY + (this->m_Bottom - y) * unit;
    }
    else {  //函数坐标系->视图坐标系
        double unit = (this->m_Bottom - this->m_Top) / (m_MaxY - m_MinY);  //转换单元
        return this->m_Bottom - (y - m_MinY) * unit;
    }

}

void QFunctionIndicatorWnd::SetShowWindow() {
    QRect rect = this -> rect();//获得当前视图区坐标范围大小数据

    //设定绘画有效区域，不会占据整个视图窗口
    if (m_ShowFuncList == true) {
        //如果要显示函数列表，则需要视图右侧一块区域用于显示函数列表，相应有效区域减小
        this->m_Top = rect.bottom() * 0.1;
        this->m_Bottom = rect.bottom() * 0.9;
        this->m_Left = rect.right() * 0.05;
        this->m_Right = rect.right() * 0.8;
    }
    else {
        this->m_Top = rect.bottom() * 0.1;
        this->m_Bottom = rect.bottom() * 0.95;
        this->m_Left = rect.right() * 0.05;
        this->m_Right = rect.right() * 0.95;
    }

}

void QFunctionIndicatorWnd::DrawEdge(QPainter* p) {
    //沿着坐标范围绘制普通直线
    QRect r = QRect( this->m_Left, this->m_Top, this -> m_Right - this -> m_Left, this -> m_Bottom - this -> m_Top);
    p -> drawRect(r);
//    p -> drawLine(this->m_Left, this->m_Top, this->m_Left, this->m_Bottom);
//    p -> drawLine(this->m_Left, this->m_Bottom, this->m_Right, this->m_Bottom);
//    p -> drawLine(this->m_Right, this->m_Bottom, this->m_Right, this->m_Top);
//    p -> drawLine(this->m_Right, this->m_Top, this->m_Left, this->m_Top);
}

void QFunctionIndicatorWnd::MarkCoordinateValue(QPainter* p) {

    //标注x坐标信息
    double xPoint;

    //标注间的单元距离
    //当x坐标值比较大时，文本的显示内容更长，为了得标注文本清晰可见，需要使各个标注的单元距离更长，标注分布更稀
    int xUnit = m_MinX > 100 ? 100 : 50;

    for (xPoint = this->m_Left; xPoint < this->m_Right; xPoint += xUnit ) { //每隔一个xUnit大小的像素标注一个坐标值
        //标注的显示区域
        QRect textRect(xPoint - (xUnit / 2), this->m_Bottom + 1, xUnit, 20);

        //设置标注文本内容，需要进行x坐标值转换，仅显示小数点后两位
        QString xInfo = QString::number(this->TransformX(xPoint, true), 'f', 2);

        //绘制，单行,上下左右居中显示
        p -> drawText(textRect, Qt::AlignCenter, xInfo);
    }

    //标注最后一个x坐标，与绘制框太远则不显示
    if (xPoint - this->m_Right <= (xUnit / 2) ) {
        QRect textRect(this->m_Right, this->m_Bottom + 1, 50, 20);
        QString xInfo = QString::number(m_MaxX, 'f', 2);
        p -> drawText(textRect, Qt::AlignLeft | Qt::AlignTop, xInfo);
    }


    //标注y坐标信息
    int yPoint;

    //y轴为横向绘制标注，故设置固定单元距离，不会影响文本内容显示
    for (yPoint = this->m_Bottom - 50; yPoint > this->m_Top; yPoint -= 50) {
        //标注的显示区域
        QRect textRect(this->m_Left - 200, yPoint - 10, 197, 20);

        //设置标注文本内容，需要进行y坐标值转换，仅显示小数点后两位
        QString yInfo = QString::number(this->TransformY(yPoint, true), 'f', 2);

        //绘制，单行,靠右显示
        p -> drawText(textRect, Qt::AlignRight, yInfo);
    }

    //标注最后一个y坐标，与绘制框太远则不显示
    if (this->m_Top - yPoint <= 25) {
        QRect textRect(this->m_Left - 200, this->m_Top - 10, 197, 20);
        QString yInfo = QString::number(m_MaxY, 'f', 2);
        p -> drawText(textRect, Qt::AlignRight | Qt::AlignBottom, yInfo);
    }
}

void QFunctionIndicatorWnd::DrawGrid(QPainter* p) {
    //创建笔，虚线，灰色
    QPen pen;
    pen.setStyle(Qt::DotLine);
    pen.setColor(QColor(100, 100, 100));
    pen.setWidth(1);

    //设置画笔
    QPen oldPen = p -> pen();
    p -> setPen(pen);

    //从左往右每隔50个像素点绘制一条垂直方向虚线
    for (double xPoint = this->m_Left + 50; xPoint < this->m_Right; xPoint += 50) {
        p -> drawLine(xPoint, this -> m_Top, xPoint, this -> m_Bottom);
    }

    //从下至上每隔50个像素点绘制一条水平方向虚线
    for (double yPoint = this->m_Bottom - 50; yPoint > this->m_Top; yPoint -= 50) {
        p -> drawLine(this -> m_Left, yPoint, this -> m_Right, yPoint);
    }

    p -> setPen(oldPen);
}

void QFunctionIndicatorWnd::DrawAxis(QPainter *p) {
    //创建笔，实线，黑色
    QPen pen;
    pen.setStyle(Qt::SolidLine);
    pen.setColor(QColor(0, 0, 0));
    pen.setWidth(2);

    //设置画笔
    QPen oldPen = p -> pen();
    p -> setPen(pen);

    //获得函数坐标系上的(0, 0)点在视图坐标系上的具体坐标
    int zeroX = this->TransformX(0, false);
    int zeroY = this->TransformY(0, false);

    //判断x,y轴是否在当前视图显示范围内
    bool showX = zeroY >= this->m_Top && zeroY <= this->m_Bottom;
    bool showY = zeroX >= this->m_Left && zeroX <= this->m_Right;

    //X轴绘制
    if (showX) {
        //绘制X轴线
        p -> drawLine(this->m_Left - 10, zeroY, this->m_Right + 10, zeroY);

        //绘制X轴标注
        p -> drawLine(this->m_Right + 10, zeroY, this->m_Right + 5, zeroY + 5);
        p -> drawLine(this->m_Right + 10, zeroY, this->m_Right + 5, zeroY - 5);
        p -> drawText(this->m_Right + 10, zeroY, "X轴");
    }

    //Y轴绘制
    if (showY) {
        //绘制Y轴线
        p -> drawLine(zeroX, this->m_Bottom + 10, zeroX, this->m_Top - 10);

        //绘制Y轴标注
        p -> drawLine(zeroX, this->m_Top - 10, zeroX - 5, this->m_Top - 5);
        p -> drawLine(zeroX, this->m_Top - 10, zeroX + 5, this->m_Top - 5);
        p -> drawText(zeroX + 5, this->m_Top - 10, "Y轴");
    }

    //X轴与Y轴均在显示范围内，需标注坐标原点
    if (showX && showY) {
        p -> drawText(zeroX + 3, zeroY + 15, "0");
    }

    p -> setPen(oldPen);
}

void QFunctionIndicatorWnd::DrawFunction(QPainter *p) {

    std::list<DrawFuncData*>::iterator it;  //遍历图像绘制数据链表
    for (it = m_DrawDataList.begin(); it != m_DrawDataList.end(); it++) {

        DrawFuncData* data = *it;  //获取图像数据

        //根据设置创建笔
        QPen pen;
        pen.setStyle(Qt::PenStyle(data -> lineType));
        pen.setWidth(data->lineWidth);
        pen.setColor(data -> lineColor);

        QPen oldPen = p -> pen();
        p -> setPen(pen);

        bool move = true;  //移动布尔值，遇到连续点时需要绘制直线，不连续点时仅移动当前绘制点
        QPoint start;  //当前绘制点

        //遍历函数点
        for (uint i = 0; i < data->drawPoint->size(); i++) {
            SML::FunctionPoint point = (*(data->drawPoint))[i];

            //函数点为无穷或超出当前坐标范围时不进行绘制
            if (point.first == INF || point.second == INF
                ||point.first == -INF || point.second == -INF
                || point.first < m_MinX || point.first > m_MaxX
                || point.second < m_MinY || point.second > m_MaxY) {
                move = true;  //遇到下个可标记点时不绘制直线，只移动当前绘制点
                continue;
            }

            //移动与绘制点间直线
            if (move) {
                //移动当前绘制点，坐标需要进行坐标系转换
                start = QPoint(this->TransformX(point.first,false), this->TransformY(point.second, false));
                move = false;
            }
            else {
                //绘制点间直线
                QPoint end = QPoint(this->TransformX(point.first, false), this->TransformY(point.second, false));
                p -> drawLine(start, end);

                //移动当前绘制点
                start = end;
            }
        }

        p -> setPen(oldPen);
    }


}

void QFunctionIndicatorWnd::ShowFuncExpression(QPainter *p) {

    QRect rect = this -> rect();//获得当前视图区坐标范围大小数据

    //绘制边框和提醒文字
    p -> drawText(this->m_Right + 150, this->m_Top - 10, "函数表达式");
    p -> drawLine(rect.right(), this->m_Top, this->m_Right + 50, this->m_Top);
    p -> drawLine(this->m_Right + 50, this->m_Top, this->m_Right + 50, this -> m_Bottom );

    //获得绘制数据
    std::list<DrawFuncData*> drawDataList = this -> m_DrawDataList;

    int top = this->m_Top + 15; //开始标注函数表达式的起始点
    int order = 1;  //函数表达式的序号

    QPen pen;
    QPen oldPen = p -> pen();

    std::list<DrawFuncData*>::iterator it;  //遍历函数图像数据
    for (it = drawDataList.begin(); it != drawDataList.end(); it++, top += 15, order++) {
        DrawFuncData* data = *it;  //获取图像数据

        pen.setColor(data->lineColor);  //设置文本颜色
        p -> setPen(pen);

        QString orderStr = QString("%1:").arg(order);//序号

        p -> drawText(this->m_Right + 60, top, orderStr);

        //标注该函数的函数表达式
        for (uint j = 0; j < data->expressionStr.size(); j++, top += 20) {
            //标注函数表达式文本
            p -> drawText(this->m_Right + 80, top, data->expressionStr[j]);
        }

        p -> setPen(oldPen);  //还原文本颜色

        //标记分隔线
        top -= 10;
        p -> drawLine(rect.right(), top, this->m_Right + 50, top);

    }

}

void QFunctionIndicatorWnd::AmplifyImage() {

    //计算放大值，固定为0.1
    double valueX = (m_MaxX - m_MinX) * 0.1;
    double valueY = (m_MaxY - m_MinY) * 0.1;

    //设置新xy范围
    this -> SetRange(m_MinX + valueX, m_MaxX - valueX, m_MinY + valueY, m_MaxY - valueY);

    //刷新绘画
    /*this->Invalidate();
    this->UpdateWindow();*/
    update();

}

void QFunctionIndicatorWnd::ShrinkImage() {

    //计算缩小值，固定为0.1
    double valueX = (m_MaxX - m_MinX) * 0.1;
    double valueY = (m_MaxY - m_MinY) * 0.1;

    //设置新xy范围
    this -> SetRange(m_MinX - valueX, m_MaxX + valueX, m_MinY - valueY, m_MaxY + valueY);

    //刷新绘画
    /*this->Invalidate();
    this->UpdateWindow();*/
    update();  //双缓冲绘画不会发送闪烁现象
}

double QFunctionIndicatorWnd::GetDistacne(SML::FunctionPoint& a, SML::FunctionPoint& b) {
    //平面上两点间的距离公式
    OPERAND x = (a.first - b.first) * (a.first - b.first);
    OPERAND y = (a.second - b.second) * (a.second - b.second);
    return sqrt(x + y);
}

SML::FunctionPoint QFunctionIndicatorWnd::GetClosestPoint(QPoint point) {

    //转换光标点参数，转换为函数图像点
    SML::FunctionPoint targetPoint = std::make_pair(this->TransformX(point.x(), true), this->TransformY(point.y(), true));

    //初始化结果函数图像点和最近距离
    SML::FunctionPoint res = std::make_pair(INF, INF);
    double minDistance = INF;


    //遍历图像绘制数据链表
    std::list<DrawFuncData*>::iterator it;
    int num = 1;  //图像数据链表结点序号，以1为起始
    for (it = m_DrawDataList.begin(); it != m_DrawDataList.end(); it++, num++) {

        DrawFuncData* data = *it;  //获取图像数据

        //遍历该函数的所有函数点
        for (uint i = 0; i < data->drawPoint->size(); i++) {

            SML::FunctionPoint curPoint = (*(data->drawPoint))[i];

            //当前点与光标点的x距离过长，忽视该点
            if (abs(curPoint.first - targetPoint.first) > 0.1) continue;

            //获得当前点与光标点的具体距离
            double curDistance = this->GetDistacne(curPoint, targetPoint);

            //该点离光标点更近，更新相关数据
            if ( curDistance < minDistance) {
                minDistance = curDistance;
                res = curPoint;
                this->m_CurFuncNum = num;
            }
        }

    }

    return res;
}

bool QFunctionIndicatorWnd::GetFunctionPoint(QPoint point) {

    //数据清空
    m_FuncPointDrawData.msg.clear();

    //获得里该点最近的函数图像点
    SML::FunctionPoint closestPoint = this->GetClosestPoint(point);

    //该点无效则不显示
    if (closestPoint.first == INF || closestPoint.second == INF) return false;

    //获得最近函数点的视图坐标
    double outputX = this->TransformX(closestPoint.first, false);
    double outputY = this->TransformY(closestPoint.second, false);

    //必须在视图坐标系上也足够近才显示函数信息
    if (abs(outputX - point.x()) > 5 || abs(outputY - point.y()) > 5) return false;

    //数据填充，留待绘图事件进行绘画
    m_FuncPointDrawData.msg = QString::number(closestPoint.first, 'f', 2) + "，" + QString::number(closestPoint.second, 'f', 2);
    m_FuncPointDrawData.textRect = QRect(outputX, outputY - 10, 100, 20);


    //QT在其他事件绘图会被绘画事件所覆盖，导致该绘画无效
//    //在函数图像上显示最近函数点坐标
//    QPainter *p = new QPainter(this);
//    QString msg = QString::number(closestPoint.first, 'g', 2) + "," + QString::number(closestPoint.second, 'g', 2);
//    QRect textrect(outputX, outputY - 10, 100, 20);
//    p -> drawText(textrect, Qt::AlignCenter, msg);

//    delete p;

    return true;
}

void QFunctionIndicatorWnd::ShowFunctionNum() {
    //状态栏显示函数点对应的函数序号
    QString message = QString("当前函数: %1").arg(this->m_CurFuncNum);
    this -> m_StatusBarItem2 -> setText(message);
}

void QFunctionIndicatorWnd::ShowFunctionInformation(QPoint point) {

    //移动模式为正在移动或不允许显示函数信息，则不显示函数信息
    if (m_MoveMode != MOVING  && m_ShowFuncInfo == true) {

        //如果函数点坐标显示成功，则显示函数序号
        if (this->GetFunctionPoint(point)) {
            this->ShowFunctionNum();
        }else{
            this -> m_StatusBarItem2 -> setText("当前函数序号");
        }

        //刷新绘图
        update();
    }
}

void QFunctionIndicatorWnd::ShowImagePoint(QPoint point) {
    QString msg;

    if ( point.x() >= this->m_Left && point.x() <= this->m_Right && point.y() >= this->m_Top && point.y() <= this->m_Bottom){
        //msg = QString("(%1, %2)").arg(this->TransformX(point.x(), true)).arg(this->TransformY(point.y(), true));
        msg = "(" + QString::number(this->TransformX(point.x(), true), 'f', 2) + "，" + QString::number(this->TransformY(point.y(), true), 'f', 2) + ")";
    }else{
        msg = "当前光标位置";
    }

    this -> m_StatusBarItem1 -> setText(msg);
}

void QFunctionIndicatorWnd::on_m_Add_NormalFunc_triggered()
{
    QNormalFuncDlg dlg("添加普通函数", SML::MathExpression(), "", this->m_MinX, this->m_MaxX);
    while ( dlg.exec() == QDialog::Accepted ){
        //获得函数类对象
        SML::MathFunction* pFunction = dlg.GetMathFunction();

        //新建绘画数据，传入必要参数
        DrawFuncData* dfd = new DrawFuncData;
        dfd->precision = dlg.GetPrecision();

        //计算函数图像数据，范围为X-Y轴显示范围，并将结果保存至绘画数据中
        //dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);
        //该方式虽然更为精准，但会导致图像移动时的卡顿，故直接采用下方以定义域范围绘制图像


        //计算函数图像数据，范围为该函数定义域，并将结果保存至绘画数据中
        //该方式与极坐标和参数方程的计算方式保持一致，图像移动会比较顺畅
        dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);

        //计算失败，弹出提示弹窗，释放相应资源并重新显示对话框
        if (pFunction->IsSuccess() == false) {
            QMessageBox msg(QMessageBox::Critical, "Function-Indicator", pFunction->GetError().c_str());
            msg.exec();
            delete pFunction;
            delete dfd;
            continue;
         }

        //单函数模式下链表仅保存一个结点
        if (this->m_Single == true) {
            this->ClearFunction();
         }

        //将其他数据添加至绘画数据中
        dfd->expressionStr.push_back(QString("f(x)=") + dlg.GetExpressionStr());
        dfd->lineWidth = dlg.GetLineWidth();
        dfd->lineType = dlg.GetLineType();
        dfd->lineColor = dlg.GetLineColor();

        //将函数类对象和绘画数据添加至链表中保存
        this->m_FunctionList.push_back(pFunction);
        this->m_DrawDataList.push_back(dfd);

        update();

        break;
    }
}

void QFunctionIndicatorWnd::on_m_Add_PolarFunc_triggered()
{
    QPolarFuncDlg dlg("添加极坐标函数", SML::MathExpression(), "",  -3.14, 3.14);
    while ( dlg.exec() == QDialog::Accepted ){
        //获得函数类对象
        SML::MathFunction* pFunction = dlg.GetMathFunction();

        //新建绘画数据，传入必要参数
        DrawFuncData* dfd = new DrawFuncData;
        dfd->precision = dlg.GetPrecision();

        //计算函数图像数据，范围为X-Y轴显示范围，并将结果保存至绘画数据中
        //dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);
        //该方式虽然更为精准，但会导致图像移动时的卡顿，故直接采用下方以定义域范围绘制图像


        //计算函数图像数据，范围为该函数定义域，并将结果保存至绘画数据中
        //该方式与极坐标和参数方程的计算方式保持一致，图像移动会比较顺畅
        dfd->drawPoint = pFunction->Calculate(-PI, PI, dfd->precision);

        //计算失败，弹出提示弹窗，释放相应资源并重新显示对话框
        if (pFunction->IsSuccess() == false) {
            QMessageBox msg(QMessageBox::Critical, "Function-Indicator", pFunction->GetError().c_str());
            msg.exec();
            delete pFunction;
            delete dfd;
            continue;
         }

        //单函数模式下链表仅保存一个结点
        if (this->m_Single == true) {
            this->ClearFunction();
         }

        //将其他数据添加至绘画数据中
        dfd->expressionStr.push_back(QString("r(a)=") + dlg.GetExpressionStr());
        dfd->lineWidth = dlg.GetLineWidth();
        dfd->lineType = dlg.GetLineType();
        dfd->lineColor = dlg.GetLineColor();

        //将函数类对象和绘画数据添加至链表中保存
        this->m_FunctionList.push_back(pFunction);
        this->m_DrawDataList.push_back(dfd);

        update();
        break;
    }
}

void QFunctionIndicatorWnd::on_m_Add_TwoFuc_triggered()
{
    QTwoFuncDlg dlg("添加参数方程函数", SML::MathExpression(), "", SML::MathExpression(), "", -5, 5);
    while ( dlg.exec() == QDialog::Accepted ){
        //获得函数类对象
        SML::MathFunction* pFunction = dlg.GetMathFunction();

        //新建绘画数据，传入必要参数
        DrawFuncData* dfd = new DrawFuncData;
        dfd->precision = dlg.GetPrecision();

        //计算函数图像数据，范围为X-Y轴显示范围，并将结果保存至绘画数据中
        //dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);
        //该方式虽然更为精准，但会导致图像移动时的卡顿，故直接采用下方以定义域范围绘制图像


        //计算函数图像数据，范围为该函数定义域，并将结果保存至绘画数据中
        //该方式与极坐标和参数方程的计算方式保持一致，图像移动会比较顺畅
        dfd->drawPoint = pFunction->Calculate(dlg.GetMin(), dlg.GetMax(), dfd->precision);

        //计算失败，弹出提示弹窗，释放相应资源并重新显示对话框
        if (pFunction->IsSuccess() == false) {
            QMessageBox msg(QMessageBox::Critical, "Function-Indicator", pFunction->GetError().c_str());
            msg.exec();
            delete pFunction;
            delete dfd;
            continue;
         }

        //单函数模式下链表仅保存一个结点
        if (this->m_Single == true) {
            this->ClearFunction();
         }

        //将其他数据添加至绘画数据中
        dfd->expressionStr.push_back(QString("x(t)=") + dlg.GetExpressionStrX());
        dfd->expressionStr.push_back(QString("y(t)=") + dlg.GetExpressionStrY());
        dfd->lineWidth = dlg.GetLineWidth();
        dfd->lineType = dlg.GetLineType();
        dfd->lineColor = dlg.GetLineColor();

        //将函数类对象和绘画数据添加至链表中保存
        this->m_FunctionList.push_back(pFunction);
        this->m_DrawDataList.push_back(dfd);

        update();
        break;
    }
}

void QFunctionIndicatorWnd::on_m_Alter_Func_triggered()
{
    QFuncNumDlg dlg("修改函数参数", "请输入要修改的函数序号");
    while ( dlg.exec() == QDialog::Accepted){
        int num = dlg.GetNum();
        if (num <= 0 || num > this->m_FunctionList.size()) {
            QMessageBox msg(QMessageBox::Critical, "Function-Indicator", "请输入有效范围的函数序号");
            msg.exec();
            continue;
        }
        this->AlterFunction(num);  //修改函数
        update();
        break;
    }

}

void QFunctionIndicatorWnd::on_m_Del_Func_triggered()
{
    QFuncNumDlg dlg("删除函数", "请输入要删除的函数序号");
    while ( dlg.exec() == QDialog::Accepted){
        int num = dlg.GetNum();
        if (num <= 0 || num > this->m_FunctionList.size()) {
            QMessageBox msg(QMessageBox::Critical, "Function-Indicator", "请输入有效范围的函数序号");
            msg.exec();
            continue;
        }
        this->DelFunction(num);  //删除函数
        update();
        break;
    }
}

void QFunctionIndicatorWnd::on_m_Del_AllFunc_triggered()
{
     this->ClearFunction(); //清空函数
     update();
}

void QFunctionIndicatorWnd::on_m_Set_XY_Range_triggered()
{
    QSetXYRangeDlg dlg(m_MinX, m_MaxX, m_MinY, m_MaxY);
    while ( dlg.exec() == QDialog::Accepted ){
        if (dlg.GetMinX() >= dlg.GetMaxX() || dlg.GetMinY() >= dlg.GetMaxY()){
            QMessageBox msg(QMessageBox::Critical, "Function-Indicator", "输入不合法，请检查！");
            msg.exec();
            continue;
        }
        this->SetRange(dlg.GetMinX(), dlg.GetMaxX(), dlg.GetMinY(), dlg.GetMaxY());
        update();
        break;
    }
}

void QFunctionIndicatorWnd::on_m_Set_Move_triggered()
{
    if ( this -> m_MoveMode == BAN ){
        this -> m_MoveMode = MOVE;
        setCursor(QCursor(Qt::OpenHandCursor));
    }
    else{
        this -> m_MoveMode = BAN;
        setCursor(QCursor(Qt::ArrowCursor));
    }
}

void QFunctionIndicatorWnd::on_m_Set_Single_triggered()
{
    this->m_Single = !this->m_Single;
    //由多函数转单函数，且当前存在多个函数，需清除所有函数
    if ( this->m_Single == true && this->m_FunctionList.size() > 1 ){
        this->ClearFunction();
        update();
    }
}

void QFunctionIndicatorWnd::on_m_Show_Edge_triggered()
{
    this->m_ShowEdge = !this->m_ShowEdge;
    update();
}

void QFunctionIndicatorWnd::on_m_Show_Axis_triggered()
{
    this->m_ShowAxis = !this->m_ShowAxis;
    update();
}

void QFunctionIndicatorWnd::on_m_Show_Grid_triggered()
{
    this->m_ShowGrid = !this->m_ShowGrid;
    update();
}

void QFunctionIndicatorWnd::on_m_Show_FuncInfo_triggered()
{
    this->m_ShowFuncInfo = !this->m_ShowFuncInfo;
    update();
}

void QFunctionIndicatorWnd::on_m_Show_FuncList_triggered()
{
    this->m_ShowFuncList = !this->m_ShowFuncList;
    update();
}

void QFunctionIndicatorWnd::on_m_View_StatusBar_triggered()
{
    this -> m_ShowStatusBar = !this -> m_ShowStatusBar;
    update();
}
