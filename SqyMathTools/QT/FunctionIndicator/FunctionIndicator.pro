QT       += core gui

greaterThan(QT_MAJOR_VERSION, 4): QT += widgets

CONFIG += c++11

# The following define makes your compiler emit warnings if you use
# any Qt feature that has been marked deprecated (the exact warnings
# depend on your compiler). Please consult the documentation of the
# deprecated API in order to know how to port your code away from it.
DEFINES += QT_DEPRECATED_WARNINGS

# You can also make your code fail to compile if it uses deprecated APIs.
# In order to do so, uncomment the following line.
# You can also select to disable deprecated APIs only up to a certain version of Qt.
#DEFINES += QT_DISABLE_DEPRECATED_BEFORE=0x060000    # disables all the APIs deprecated before Qt 6.0.0

SOURCES += \
    Src/Dialog/QFuncNumDlg.cpp \
    Src/Dialog/QFunctionDlg.cpp \
    Src/Dialog/QNormalFuncDlg.cpp \
    Src/Dialog/QPolarFuncDlg.cpp \
    Src/Dialog/QSetXYRangeDlg.cpp \
    Src/Dialog/QTwoFuncDlg.cpp \
    Src/Main/QFunctionIndicatorWnd.cpp \
    Src/Main/main.cpp

HEADERS += \
    Include/Dialog/QFuncNumDlg.h \
    Include/Dialog/QFunctionDlg.h \
    Include/Dialog/QNormalFuncDlg.h \
    Include/Dialog/QPolarFuncDlg.h \
    Include/Dialog/QSetXYRangeDlg.h \
    Include/Dialog/QTwoFuncDlg.h \
    Include/Main/QFunctionIndicatorWnd.h


FORMS += \
    Ui/QFuncNumDlg.ui \
    Ui/QFunctionDlg.ui \
    Ui/QFunctionIndicatorWnd.ui \
    Ui/QNormalFuncDlg.ui \
    Ui/QPolarFuncDlg.ui \
    Ui/QSetXYRangeDlg.ui \
    Ui/QTwoFuncDlg.ui

INCLUDEPATH += $$PWD/Include
INCLUDEPATH += $$PWD/Include/Dialog
INCLUDEPATH += $$PWD/Include/Main

# Default rules for deployment.
qnx: target.path = /tmp/$${TARGET}/bin
else: unix:!android: target.path = /opt/$${TARGET}/bin
!isEmpty(target.path): INSTALLS += target

win32:CONFIG(release, debug|release): LIBS += -L$$PWD/../build-SqyMathLibrary-Desktop_Qt_5_14_2_MinGW_32_bit-Debug/release/ -lSqyMathLibrary
else:win32:CONFIG(debug, debug|release): LIBS += -L$$PWD/../build-SqyMathLibrary-Desktop_Qt_5_14_2_MinGW_32_bit-Debug/debug/ -lSqyMathLibrary
else:unix:!macx: LIBS += -L$$PWD/../build-SqyMathLibrary-Desktop_Qt_5_14_2_MinGW_32_bit-Debug/ -lSqyMathLibrary

INCLUDEPATH += $$PWD/../SqyMathLibrary/include
DEPENDPATH += $$PWD/../SqyMathLibrary/include

win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/../build-SqyMathLibrary-Desktop_Qt_5_14_2_MinGW_32_bit-Debug/release/libSqyMathLibrary.a
else:win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/../build-SqyMathLibrary-Desktop_Qt_5_14_2_MinGW_32_bit-Debug/debug/libSqyMathLibrary.a
else:win32:!win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/../build-SqyMathLibrary-Desktop_Qt_5_14_2_MinGW_32_bit-Debug/release/SqyMathLibrary.lib
else:win32:!win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/../build-SqyMathLibrary-Desktop_Qt_5_14_2_MinGW_32_bit-Debug/debug/SqyMathLibrary.lib
else:unix:!macx: PRE_TARGETDEPS += $$PWD/../build-SqyMathLibrary-Desktop_Qt_5_14_2_MinGW_32_bit-Debug/libSqyMathLibrary.a
