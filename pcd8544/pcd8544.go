package pcd8544

/*
#cgo LDFLAGS: -L . -lPCD8544 -lwiringPi

#include <wiringPi.h>
#include <PCD8544.h>
#include <stdio.h>
#include <stdlib.h>


void Test(char *c, char *c2) {
	LCDclear();
	LCDdrawstring(0, 24, c);
	LCDdrawstring(0, 34, c2);
	LCDdisplay();
}
*/
import "C"

const (
	BLACK = 1
	WHITE = 0
)

func Test(s1 string, s2 string) {
	C.Test(C.CString(s1), C.CString(s2))
}

func LCDInit(_sclk int, _din int, _dc int, _cs int, _rst int, contrast int) {
	C.LCDInit(C.uint8_t(_sclk), C.uint8_t(_din), C.uint8_t(_dc), C.uint8_t(_cs), C.uint8_t(_rst), C.uint8_t(contrast))
}

func LCDcommand(c int) {
	C.LCDcommand(C.uint8_t(c))
}

func LCDdata(c uint8) {
	C.LCDdata(C.uint8_t(c))
}

func LCDsetContrast(val int) {
	C.LCDsetContrast(C.uint8_t(val))
}

func LCDclear() {
	C.LCDclear()
}

func LCDdisplay() {
	C.LCDdisplay()
}

func LCDsetPixel(x uint8, y uint8, color uint8) {
	C.LCDsetPixel(C.uint8_t(x), C.uint8_t(y), C.uint8_t(color))
}

func LCDgetPixel(x uint8, y uint8) int {
	n, _ := C.LCDgetPixel(C.uint8_t(x), C.uint8_t(y))
	return int(n)
}

func LCDfillcircle(x0 uint8, y0 uint8, r uint8, color uint8) {
	C.LCDfillcircle(C.uint8_t(x0), C.uint8_t(y0), C.uint8_t(r), C.uint8_t(color))
}

func LCDdrawcircle(x0 uint8, y0 uint8, r uint8, color uint8) {
	C.LCDdrawcircle(C.uint8_t(x0), C.uint8_t(y0), C.uint8_t(r), C.uint8_t(color))
}

func LCDdrawrect(x uint8, y uint8, w uint8, h uint8, color uint8) {
	C.LCDdrawrect(C.uint8_t(x), C.uint8_t(y), C.uint8_t(w), C.uint8_t(h), C.uint8_t(color))
}

func LCDfillrect(x uint8, y uint8, w uint8, h uint8, color uint8) {
	C.LCDfillrect(C.uint8_t(x), C.uint8_t(y), C.uint8_t(w), C.uint8_t(h), C.uint8_t(color))
}

func LCDdrawline(x0 int, y0 int, x1 int, y1 int, color int) {
	C.LCDdrawline(C.uint8_t(x0), C.uint8_t(y0), C.uint8_t(x1), C.uint8_t(y1), C.uint8_t(color))
}

func LCDsetCursor(x uint8, y uint8) {
	C.LCDsetCursor(C.uint8_t(x), C.uint8_t(y))
}

//func LCDsetTextSize(s uint8) {
//	C.LCDsetTextSize(C.uint8_t(s))
//}
//
//func LCDsetTextColor(c uint8) {
//	C.LCDsetTextColor(C.uint8_t(c))
//}

func LCDwrite(c uint8) {
	C.LCDwrite(C.uint8_t(c))
}

func LCDshowLogo() {
	C.LCDshowLogo()
}

func LCDdrawchar(x uint8, line uint8, c byte) {
	C.LCDdrawchar(C.uint8_t(x), C.uint8_t(line), C.char(c))
}

func LCDdrawstring(x int8, y int8, str string) {
	C.LCDdrawstring(C.uint8_t(x), C.uint8_t(y), C.CString(str))
}

