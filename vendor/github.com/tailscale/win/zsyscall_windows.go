// Code generated by 'go generate'; DO NOT EDIT.

package win

import (
	"syscall"
	"unsafe"

	"github.com/dblohm7/wingoes/com"
	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modcomctl32 = windows.NewLazySystemDLL("comctl32.dll")
	moddwmapi   = windows.NewLazySystemDLL("dwmapi.dll")
	modgdiplus  = windows.NewLazySystemDLL("gdiplus.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	modshcore   = windows.NewLazySystemDLL("shcore.dll")
	moduser32   = windows.NewLazySystemDLL("user32.dll")
	moduxtheme  = windows.NewLazySystemDLL("uxtheme.dll")

	procTaskDialogIndirect                    = modcomctl32.NewProc("TaskDialogIndirect")
	procDwmExtendFrameIntoClientArea          = moddwmapi.NewProc("DwmExtendFrameIntoClientArea")
	procDwmGetWindowAttribute                 = moddwmapi.NewProc("DwmGetWindowAttribute")
	procDwmInvalidateIconicBitmaps            = moddwmapi.NewProc("DwmInvalidateIconicBitmaps")
	procDwmQueryThumbnailSourceSize           = moddwmapi.NewProc("DwmQueryThumbnailSourceSize")
	procDwmRegisterThumbnail                  = moddwmapi.NewProc("DwmRegisterThumbnail")
	procDwmSetIconicLivePreviewBitmap         = moddwmapi.NewProc("DwmSetIconicLivePreviewBitmap")
	procDwmSetIconicThumbnail                 = moddwmapi.NewProc("DwmSetIconicThumbnail")
	procDwmSetWindowAttribute                 = moddwmapi.NewProc("DwmSetWindowAttribute")
	procDwmUnregisterThumbnail                = moddwmapi.NewProc("DwmUnregisterThumbnail")
	procDwmUpdateThumbnailProperties          = moddwmapi.NewProc("DwmUpdateThumbnailProperties")
	procGdipAddPathEllipseI                   = modgdiplus.NewProc("GdipAddPathEllipseI")
	procGdipCreateBitmapFromFile              = modgdiplus.NewProc("GdipCreateBitmapFromFile")
	procGdipCreateBitmapFromGraphics          = modgdiplus.NewProc("GdipCreateBitmapFromGraphics")
	procGdipCreateBitmapFromHBITMAP           = modgdiplus.NewProc("GdipCreateBitmapFromHBITMAP")
	procGdipCreateBitmapFromHICON             = modgdiplus.NewProc("GdipCreateBitmapFromHICON")
	procGdipCreateBitmapFromScan0             = modgdiplus.NewProc("GdipCreateBitmapFromScan0")
	procGdipCreateBitmapFromStream            = modgdiplus.NewProc("GdipCreateBitmapFromStream")
	procGdipCreateFontFamilyFromName          = modgdiplus.NewProc("GdipCreateFontFamilyFromName")
	procGdipCreateFromHDC                     = modgdiplus.NewProc("GdipCreateFromHDC")
	procGdipCreateHBITMAPFromBitmap           = modgdiplus.NewProc("GdipCreateHBITMAPFromBitmap")
	procGdipCreateHICONFromBitmap             = modgdiplus.NewProc("GdipCreateHICONFromBitmap")
	procGdipCreatePath                        = modgdiplus.NewProc("GdipCreatePath")
	procGdipCreateSolidFill                   = modgdiplus.NewProc("GdipCreateSolidFill")
	procGdipCreateStringFormat                = modgdiplus.NewProc("GdipCreateStringFormat")
	procGdipDeleteBrush                       = modgdiplus.NewProc("GdipDeleteBrush")
	procGdipDeleteFont                        = modgdiplus.NewProc("GdipDeleteFont")
	procGdipDeleteFontFamily                  = modgdiplus.NewProc("GdipDeleteFontFamily")
	procGdipDeleteGraphics                    = modgdiplus.NewProc("GdipDeleteGraphics")
	procGdipDeletePath                        = modgdiplus.NewProc("GdipDeletePath")
	procGdipDeleteStringFormat                = modgdiplus.NewProc("GdipDeleteStringFormat")
	procGdipDisposeImage                      = modgdiplus.NewProc("GdipDisposeImage")
	procGdipDrawImageRectI                    = modgdiplus.NewProc("GdipDrawImageRectI")
	procGdipDrawImageRectRectI                = modgdiplus.NewProc("GdipDrawImageRectRectI")
	procGdipDrawString                        = modgdiplus.NewProc("GdipDrawString")
	procGdipFillEllipseI                      = modgdiplus.NewProc("GdipFillEllipseI")
	procGdipGetCompositingMode                = modgdiplus.NewProc("GdipGetCompositingMode")
	procGdipGetGenericFontFamilyMonospace     = modgdiplus.NewProc("GdipGetGenericFontFamilyMonospace")
	procGdipGetGenericFontFamilySansSerif     = modgdiplus.NewProc("GdipGetGenericFontFamilySansSerif")
	procGdipGetGenericFontFamilySerif         = modgdiplus.NewProc("GdipGetGenericFontFamilySerif")
	procGdipGetImageDimension                 = modgdiplus.NewProc("GdipGetImageDimension")
	procGdipGetImageGraphicsContext           = modgdiplus.NewProc("GdipGetImageGraphicsContext")
	procGdipGetImageHeight                    = modgdiplus.NewProc("GdipGetImageHeight")
	procGdipGetImageHorizontalResolution      = modgdiplus.NewProc("GdipGetImageHorizontalResolution")
	procGdipGetImageVerticalResolution        = modgdiplus.NewProc("GdipGetImageVerticalResolution")
	procGdipGetImageWidth                     = modgdiplus.NewProc("GdipGetImageWidth")
	procGdipGraphicsClear                     = modgdiplus.NewProc("GdipGraphicsClear")
	procGdipResetClip                         = modgdiplus.NewProc("GdipResetClip")
	procGdipSetClipPath                       = modgdiplus.NewProc("GdipSetClipPath")
	procGdipSetCompositingMode                = modgdiplus.NewProc("GdipSetCompositingMode")
	procGdipSetCompositingQuality             = modgdiplus.NewProc("GdipSetCompositingQuality")
	procGdipSetInterpolationMode              = modgdiplus.NewProc("GdipSetInterpolationMode")
	procGdipSetPixelOffsetMode                = modgdiplus.NewProc("GdipSetPixelOffsetMode")
	procGdipSetSmoothingMode                  = modgdiplus.NewProc("GdipSetSmoothingMode")
	procGdipSetStringFormatAlign              = modgdiplus.NewProc("GdipSetStringFormatAlign")
	procGdipSetStringFormatLineAlign          = modgdiplus.NewProc("GdipSetStringFormatLineAlign")
	procGdipSetTextRenderingHint              = modgdiplus.NewProc("GdipSetTextRenderingHint")
	procGdiplusStartup                        = modgdiplus.NewProc("GdiplusStartup")
	procSwitchToThread                        = modkernel32.NewProc("SwitchToThread")
	procGetDpiForMonitor                      = modshcore.NewProc("GetDpiForMonitor")
	procGetDpiForShellUIComponent             = modshcore.NewProc("GetDpiForShellUIComponent")
	procGetScaleFactorForMonitor              = modshcore.NewProc("GetScaleFactorForMonitor")
	procCallMsgFilterW                        = moduser32.NewProc("CallMsgFilterW")
	procCreateDialogIndirectParamW            = moduser32.NewProc("CreateDialogIndirectParamW")
	procDefDlgProcW                           = moduser32.NewProc("DefDlgProcW")
	procGetNextDlgTabItem                     = moduser32.NewProc("GetNextDlgTabItem")
	procGetQueueStatus                        = moduser32.NewProc("GetQueueStatus")
	procGetWindowTextLengthW                  = moduser32.NewProc("GetWindowTextLengthW")
	procGetWindowTextW                        = moduser32.NewProc("GetWindowTextW")
	procMsgWaitForMultipleObjectsEx           = moduser32.NewProc("MsgWaitForMultipleObjectsEx")
	procSetWindowTextW                        = moduser32.NewProc("SetWindowTextW")
	procBeginBufferedPaint                    = moduxtheme.NewProc("BeginBufferedPaint")
	procBufferedPaintInit                     = moduxtheme.NewProc("BufferedPaintInit")
	procBufferedPaintSetAlpha                 = moduxtheme.NewProc("BufferedPaintSetAlpha")
	procCloseThemeData                        = moduxtheme.NewProc("CloseThemeData")
	procDrawThemeBackground                   = moduxtheme.NewProc("DrawThemeBackground")
	procDrawThemeParentBackground             = moduxtheme.NewProc("DrawThemeParentBackground")
	procDrawThemeTextEx                       = moduxtheme.NewProc("DrawThemeTextEx")
	procEndBufferedPaint                      = moduxtheme.NewProc("EndBufferedPaint")
	procGetThemeColor                         = moduxtheme.NewProc("GetThemeColor")
	procGetThemeEnumValue                     = moduxtheme.NewProc("GetThemeEnumValue")
	procGetThemeFont                          = moduxtheme.NewProc("GetThemeFont")
	procGetThemeInt                           = moduxtheme.NewProc("GetThemeInt")
	procGetThemeMargins                       = moduxtheme.NewProc("GetThemeMargins")
	procGetThemeMetric                        = moduxtheme.NewProc("GetThemeMetric")
	procGetThemePartSize                      = moduxtheme.NewProc("GetThemePartSize")
	procGetThemeSysFont                       = moduxtheme.NewProc("GetThemeSysFont")
	procGetThemeTextExtent                    = moduxtheme.NewProc("GetThemeTextExtent")
	procIsAppThemed                           = moduxtheme.NewProc("IsAppThemed")
	procIsThemeBackgroundPartiallyTransparent = moduxtheme.NewProc("IsThemeBackgroundPartiallyTransparent")
	procOpenThemeData                         = moduxtheme.NewProc("OpenThemeData")
	procSetWindowTheme                        = moduxtheme.NewProc("SetWindowTheme")
)

func TaskDialogIndirect(pTaskConfig *TASKDIALOGCONFIG, pnButton *int32, pnRadioButton *int32, pfVerificationFlagChecked *BOOL) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procTaskDialogIndirect.Addr(), 4, uintptr(unsafe.Pointer(pTaskConfig)), uintptr(unsafe.Pointer(pnButton)), uintptr(unsafe.Pointer(pnRadioButton)), uintptr(unsafe.Pointer(pfVerificationFlagChecked)), 0, 0)
	ret = HRESULT(r0)
	return
}

func DwmExtendFrameIntoClientArea(hwnd HWND, inset *MARGINS) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procDwmExtendFrameIntoClientArea.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(inset)), 0)
	ret = HRESULT(r0)
	return
}

func DwmGetWindowAttribute(hwnd HWND, attribute DWMWINDOWATTRIBUTE, attrVal unsafe.Pointer, attrValLen uint32) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procDwmGetWindowAttribute.Addr(), 4, uintptr(hwnd), uintptr(attribute), uintptr(attrVal), uintptr(attrValLen), 0, 0)
	ret = HRESULT(r0)
	return
}

func DwmInvalidateIconicBitmaps(hwnd HWND) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procDwmInvalidateIconicBitmaps.Addr(), 1, uintptr(hwnd), 0, 0)
	ret = HRESULT(r0)
	return
}

func DwmQueryThumbnailSourceSize(thumbnail HTHUMBNAIL, size *SIZE) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procDwmQueryThumbnailSourceSize.Addr(), 2, uintptr(thumbnail), uintptr(unsafe.Pointer(size)), 0)
	ret = HRESULT(r0)
	return
}

func DwmRegisterThumbnail(hwndDest HWND, hwndSrc HWND, handle *HTHUMBNAIL) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procDwmRegisterThumbnail.Addr(), 3, uintptr(hwndDest), uintptr(hwndSrc), uintptr(unsafe.Pointer(handle)))
	ret = HRESULT(r0)
	return
}

func DwmSetIconicLivePreviewBitmap(hwnd HWND, hbmp HBITMAP, clientRegionOffset *POINT, flags uint32) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procDwmSetIconicLivePreviewBitmap.Addr(), 4, uintptr(hwnd), uintptr(hbmp), uintptr(unsafe.Pointer(clientRegionOffset)), uintptr(flags), 0, 0)
	ret = HRESULT(r0)
	return
}

func DwmSetIconicThumbnail(hwnd HWND, hbmp HBITMAP, flags uint32) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procDwmSetIconicThumbnail.Addr(), 3, uintptr(hwnd), uintptr(hbmp), uintptr(flags))
	ret = HRESULT(r0)
	return
}

func DwmSetWindowAttribute(hwnd HWND, attribute DWMWINDOWATTRIBUTE, attrVal unsafe.Pointer, attrValLen uint32) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procDwmSetWindowAttribute.Addr(), 4, uintptr(hwnd), uintptr(attribute), uintptr(attrVal), uintptr(attrValLen), 0, 0)
	ret = HRESULT(r0)
	return
}

func DwmUnregisterThumbnail(thumbnail HTHUMBNAIL) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procDwmUnregisterThumbnail.Addr(), 1, uintptr(thumbnail), 0, 0)
	ret = HRESULT(r0)
	return
}

func DwmUpdateThumbnailProperties(thumbnail HTHUMBNAIL, props *DWM_THUMBNAIL_PROPERTIES) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procDwmUpdateThumbnailProperties.Addr(), 2, uintptr(thumbnail), uintptr(unsafe.Pointer(props)), 0)
	ret = HRESULT(r0)
	return
}

func GdipAddPathEllipseI(path *GpPath, x int32, y int32, width int32, height int32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipAddPathEllipseI.Addr(), 5, uintptr(unsafe.Pointer(path)), uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromFile(filename *uint16, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateBitmapFromFile.Addr(), 2, uintptr(unsafe.Pointer(filename)), uintptr(unsafe.Pointer(bitmap)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromGraphics(width int32, height int32, graphics *GpGraphics, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipCreateBitmapFromGraphics.Addr(), 4, uintptr(width), uintptr(height), uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(bitmap)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromHBITMAP(hbm HBITMAP, hpal HPALETTE, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateBitmapFromHBITMAP.Addr(), 3, uintptr(hbm), uintptr(hpal), uintptr(unsafe.Pointer(bitmap)))
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromHICON(hicon HICON, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateBitmapFromHICON.Addr(), 2, uintptr(hicon), uintptr(unsafe.Pointer(bitmap)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromScan0(width int32, height int32, stride int32, format PixelFormat, scan0 *byte, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipCreateBitmapFromScan0.Addr(), 6, uintptr(width), uintptr(height), uintptr(stride), uintptr(format), uintptr(unsafe.Pointer(scan0)), uintptr(unsafe.Pointer(bitmap)))
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromStream(stream *com.IStreamABI, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateBitmapFromStream.Addr(), 2, uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(bitmap)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateFontFamilyFromName(name *uint16, collection *GpFontCollection, family **GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateFontFamilyFromName.Addr(), 3, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(collection)), uintptr(unsafe.Pointer(family)))
	ret = GpStatus(r0)
	return
}

func GdipCreateFromHDC(hdc HDC, graphics **GpGraphics) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateFromHDC.Addr(), 2, uintptr(hdc), uintptr(unsafe.Pointer(graphics)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateHBITMAPFromBitmap(bitmap *GpBitmap, hbmReturn *HBITMAP, background ARGB) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateHBITMAPFromBitmap.Addr(), 3, uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(hbmReturn)), uintptr(background))
	ret = GpStatus(r0)
	return
}

func GdipCreateHICONFromBitmap(bitmap *GpBitmap, hbmReturn *HICON) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateHICONFromBitmap.Addr(), 2, uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(hbmReturn)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreatePath(fillMode FillMode, path **GpPath) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreatePath.Addr(), 2, uintptr(fillMode), uintptr(unsafe.Pointer(path)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateSolidFill(color ARGB, brush **GpSolidFill) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateSolidFill.Addr(), 2, uintptr(color), uintptr(unsafe.Pointer(brush)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateStringFormat(flags StringFormatFlags, language LANGID, format **GpStringFormat) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateStringFormat.Addr(), 3, uintptr(flags), uintptr(language), uintptr(unsafe.Pointer(format)))
	ret = GpStatus(r0)
	return
}

func GdipDeleteBrush(brush *GpBrush) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteBrush.Addr(), 1, uintptr(unsafe.Pointer(brush)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeleteFont(font *GpFont) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteFont.Addr(), 1, uintptr(unsafe.Pointer(font)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeleteFontFamily(family *GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteFontFamily.Addr(), 1, uintptr(unsafe.Pointer(family)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeleteGraphics(graphics *GpGraphics) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteGraphics.Addr(), 1, uintptr(unsafe.Pointer(graphics)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeletePath(path *GpPath) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeletePath.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeleteStringFormat(format *GpStringFormat) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteStringFormat.Addr(), 1, uintptr(unsafe.Pointer(format)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDisposeImage(image *GpImage) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDisposeImage.Addr(), 1, uintptr(unsafe.Pointer(image)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDrawImageRectI(graphics *GpGraphics, image *GpImage, x int32, y int32, width int32, height int32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipDrawImageRectI.Addr(), 6, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(image)), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	ret = GpStatus(r0)
	return
}

func GdipDrawImageRectRectI(graphics *GpGraphics, image *GpImage, dstX int32, dstY int32, dstWidth int32, dstHeight int32, srcX int32, srcY int32, srcWidth int32, srcHeight int32, srcUnit Unit, imgAttrs *GpImageAttributes, callback uintptr, callbackData uintptr) (ret GpStatus) {
	r0, _, _ := syscall.Syscall15(procGdipDrawImageRectRectI.Addr(), 14, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(image)), uintptr(dstX), uintptr(dstY), uintptr(dstWidth), uintptr(dstHeight), uintptr(srcX), uintptr(srcY), uintptr(srcWidth), uintptr(srcHeight), uintptr(srcUnit), uintptr(unsafe.Pointer(imgAttrs)), uintptr(callback), uintptr(callbackData), 0)
	ret = GpStatus(r0)
	return
}

func GdipDrawString(graphics *GpGraphics, text *uint16, textLength int32, font *GpFont, rectf *GpRectF, strFmt *GpStringFormat, brush *GpBrush) (ret GpStatus) {
	r0, _, _ := syscall.Syscall9(procGdipDrawString.Addr(), 7, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(text)), uintptr(textLength), uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(rectf)), uintptr(unsafe.Pointer(strFmt)), uintptr(unsafe.Pointer(brush)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipFillEllipseI(graphics *GpGraphics, brush *GpBrush, x int32, y int32, width int32, height int32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipFillEllipseI.Addr(), 6, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	ret = GpStatus(r0)
	return
}

func GdipGetCompositingMode(graphics *GpGraphics, compositingMode *CompositingMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetCompositingMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(compositingMode)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetGenericFontFamilyMonospace(family **GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetGenericFontFamilyMonospace.Addr(), 1, uintptr(unsafe.Pointer(family)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipGetGenericFontFamilySansSerif(family **GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetGenericFontFamilySansSerif.Addr(), 1, uintptr(unsafe.Pointer(family)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipGetGenericFontFamilySerif(family **GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetGenericFontFamilySerif.Addr(), 1, uintptr(unsafe.Pointer(family)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageDimension(image *GpImage, width *float32, height *float32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageDimension.Addr(), 3, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(width)), uintptr(unsafe.Pointer(height)))
	ret = GpStatus(r0)
	return
}

func GdipGetImageGraphicsContext(image *GpImage, graphics **GpGraphics) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageGraphicsContext.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(graphics)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageHeight(image *GpImage, height *uint32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageHeight.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(height)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageHorizontalResolution(image *GpImage, resolution *float32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageHorizontalResolution.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(resolution)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageVerticalResolution(image *GpImage, resolution *float32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageVerticalResolution.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(resolution)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageWidth(image *GpImage, width *uint32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageWidth.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(width)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGraphicsClear(graphics *GpGraphics, color ARGB) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGraphicsClear.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(color), 0)
	ret = GpStatus(r0)
	return
}

func GdipResetClip(graphics *GpGraphics) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipResetClip.Addr(), 1, uintptr(unsafe.Pointer(graphics)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipSetClipPath(graphics *GpGraphics, path *GpPath, combineMode CombineMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetClipPath.Addr(), 3, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(path)), uintptr(combineMode))
	ret = GpStatus(r0)
	return
}

func GdipSetCompositingMode(graphics *GpGraphics, compositingMode CompositingMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetCompositingMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(compositingMode), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetCompositingQuality(graphics *GpGraphics, compositingQuality CompositingQuality) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetCompositingQuality.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(compositingQuality), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetInterpolationMode(graphics *GpGraphics, interpolationMode InterpolationMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetInterpolationMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(interpolationMode), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetPixelOffsetMode(graphics *GpGraphics, pixelOffsetMode PixelOffsetMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetPixelOffsetMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(pixelOffsetMode), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetSmoothingMode(graphics *GpGraphics, smoothingMode SmoothingMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetSmoothingMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(smoothingMode), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetStringFormatAlign(format *GpStringFormat, align StringAlignment) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetStringFormatAlign.Addr(), 2, uintptr(unsafe.Pointer(format)), uintptr(align), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetStringFormatLineAlign(format *GpStringFormat, align StringAlignment) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetStringFormatLineAlign.Addr(), 2, uintptr(unsafe.Pointer(format)), uintptr(align), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetTextRenderingHint(graphics *GpGraphics, mode TextRenderingHint) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetTextRenderingHint.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(mode), 0)
	ret = GpStatus(r0)
	return
}

func GdiplusStartup(token *uintptr, input *GdiplusStartupInput, output *GdiplusStartupOutput) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdiplusStartup.Addr(), 3, uintptr(unsafe.Pointer(token)), uintptr(unsafe.Pointer(input)), uintptr(unsafe.Pointer(output)))
	ret = GpStatus(r0)
	return
}

func SwitchToThread() (ret bool) {
	r0, _, _ := syscall.Syscall(procSwitchToThread.Addr(), 0, 0, 0, 0)
	ret = r0 != 0
	return
}

func GetDpiForMonitor(hmonitor HMONITOR, dpiType MONITOR_DPI_TYPE, dpiX *uint32, dpiY *uint32) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procGetDpiForMonitor.Addr(), 4, uintptr(hmonitor), uintptr(dpiType), uintptr(unsafe.Pointer(dpiX)), uintptr(unsafe.Pointer(dpiY)), 0, 0)
	ret = HRESULT(r0)
	return
}

func GetDpiForShellUIComponent(component SHELL_UI_COMPONENT) (ret uint32) {
	r0, _, _ := syscall.Syscall(procGetDpiForShellUIComponent.Addr(), 1, uintptr(component), 0, 0)
	ret = uint32(r0)
	return
}

func GetScaleFactorForMonitor(hmonitor HMONITOR, scalePercent *uint32) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procGetScaleFactorForMonitor.Addr(), 2, uintptr(hmonitor), uintptr(unsafe.Pointer(scalePercent)), 0)
	ret = HRESULT(r0)
	return
}

func CallMsgFilter(msg *MSG, nCode int32) (ret bool) {
	r0, _, _ := syscall.Syscall(procCallMsgFilterW.Addr(), 2, uintptr(unsafe.Pointer(msg)), uintptr(nCode), 0)
	ret = r0 != 0
	return
}

func CreateDialogIndirectParam(instance HINSTANCE, template unsafe.Pointer, parent HWND, dlgProc uintptr, initParam uintptr) (ret HWND, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateDialogIndirectParamW.Addr(), 5, uintptr(instance), uintptr(template), uintptr(parent), uintptr(dlgProc), uintptr(initParam), 0)
	ret = HWND(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

func DefDlgProc(hdlg HWND, msg uint32, wParam uintptr, lParam uintptr) (ret uintptr) {
	r0, _, _ := syscall.Syscall6(procDefDlgProcW.Addr(), 4, uintptr(hdlg), uintptr(msg), uintptr(wParam), uintptr(lParam), 0, 0)
	ret = uintptr(r0)
	return
}

func GetNextDlgTabItem(hdlg HWND, hctl HWND, previous bool) (ret HWND, err error) {
	var _p0 uint32
	if previous {
		_p0 = 1
	}
	r0, _, e1 := syscall.Syscall(procGetNextDlgTabItem.Addr(), 3, uintptr(hdlg), uintptr(hctl), uintptr(_p0))
	ret = HWND(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetQueueStatus(flags uint32) (ret uint32) {
	r0, _, _ := syscall.Syscall(procGetQueueStatus.Addr(), 1, uintptr(flags), 0, 0)
	ret = uint32(r0)
	return
}

func GetWindowTextLength(hwnd HWND) (ret int32) {
	r0, _, _ := syscall.Syscall(procGetWindowTextLengthW.Addr(), 1, uintptr(hwnd), 0, 0)
	ret = int32(r0)
	return
}

func GetWindowText(hwnd HWND, text *uint16, maxCount int32) (ret int32) {
	r0, _, _ := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(text)), uintptr(maxCount))
	ret = int32(r0)
	return
}

func MsgWaitForMultipleObjectsEx(count uint32, handles *windows.Handle, timeoutMillis uint32, wakeMask uint32, flags uint32) (ret uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procMsgWaitForMultipleObjectsEx.Addr(), 5, uintptr(count), uintptr(unsafe.Pointer(handles)), uintptr(timeoutMillis), uintptr(wakeMask), uintptr(flags), 0)
	ret = uint32(r0)
	if ret == windows.WAIT_FAILED {
		err = errnoErr(e1)
	}
	return
}

func SetWindowText(hwnd HWND, text *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetWindowTextW.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(text)), 0)
	if int32(r1) == 0 {
		err = errnoErr(e1)
	}
	return
}

func BeginBufferedPaint(hdcTarget HDC, prcTarget *RECT, format BP_BUFFERFORMAT, paintParams *BP_PAINTPARAMS, phdc *HDC) (pb HPAINTBUFFER, err error) {
	r0, _, e1 := syscall.Syscall6(procBeginBufferedPaint.Addr(), 5, uintptr(hdcTarget), uintptr(unsafe.Pointer(prcTarget)), uintptr(format), uintptr(unsafe.Pointer(paintParams)), uintptr(unsafe.Pointer(phdc)), 0)
	pb = HPAINTBUFFER(r0)
	if pb == 0 {
		err = errnoErr(e1)
	}
	return
}

func BufferedPaintInit() (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procBufferedPaintInit.Addr(), 0, 0, 0, 0)
	ret = HRESULT(r0)
	return
}

func BufferedPaintSetAlpha(paintBuf HPAINTBUFFER, prc *RECT, alpha byte) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procBufferedPaintSetAlpha.Addr(), 3, uintptr(paintBuf), uintptr(unsafe.Pointer(prc)), uintptr(alpha))
	ret = HRESULT(r0)
	return
}

func CloseThemeData(hTheme HTHEME) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procCloseThemeData.Addr(), 1, uintptr(hTheme), 0, 0)
	ret = HRESULT(r0)
	return
}

func DrawThemeBackground(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pRect *RECT, pClipRect *RECT) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procDrawThemeBackground.Addr(), 6, uintptr(hTheme), uintptr(hdc), uintptr(iPartId), uintptr(iStateId), uintptr(unsafe.Pointer(pRect)), uintptr(unsafe.Pointer(pClipRect)))
	ret = HRESULT(r0)
	return
}

func DrawThemeParentBackground(hWnd HWND, hdc HDC, prc *RECT) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procDrawThemeParentBackground.Addr(), 3, uintptr(hWnd), uintptr(hdc), uintptr(unsafe.Pointer(prc)))
	ret = HRESULT(r0)
	return
}

func DrawThemeTextEx(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pszText *uint16, iCharCount int32, dwFlags uint32, pRect *RECT, pOptions *DTTOPTS) (ret HRESULT) {
	r0, _, _ := syscall.Syscall9(procDrawThemeTextEx.Addr(), 9, uintptr(hTheme), uintptr(hdc), uintptr(iPartId), uintptr(iStateId), uintptr(unsafe.Pointer(pszText)), uintptr(iCharCount), uintptr(dwFlags), uintptr(unsafe.Pointer(pRect)), uintptr(unsafe.Pointer(pOptions)))
	ret = HRESULT(r0)
	return
}

func EndBufferedPaint(paintBuf HPAINTBUFFER, updateTarget bool) (ret HRESULT) {
	var _p0 uint32
	if updateTarget {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall(procEndBufferedPaint.Addr(), 2, uintptr(paintBuf), uintptr(_p0), 0)
	ret = HRESULT(r0)
	return
}

func GetThemeColor(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, pColor *COLORREF) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procGetThemeColor.Addr(), 5, uintptr(hTheme), uintptr(iPartId), uintptr(iStateId), uintptr(iPropId), uintptr(unsafe.Pointer(pColor)), 0)
	ret = HRESULT(r0)
	return
}

func GetThemeEnumValue(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, piVal *int32) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procGetThemeEnumValue.Addr(), 5, uintptr(hTheme), uintptr(iPartId), uintptr(iStateId), uintptr(iPropId), uintptr(unsafe.Pointer(piVal)), 0)
	ret = HRESULT(r0)
	return
}

func GetThemeFont(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, iPropId int32, pFont *LOGFONT) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procGetThemeFont.Addr(), 6, uintptr(hTheme), uintptr(hdc), uintptr(iPartId), uintptr(iStateId), uintptr(iPropId), uintptr(unsafe.Pointer(pFont)))
	ret = HRESULT(r0)
	return
}

func GetThemeInt(hTheme HTHEME, iPartId int32, iStateId int32, iPropId int32, piVal *int32) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procGetThemeInt.Addr(), 5, uintptr(hTheme), uintptr(iPartId), uintptr(iStateId), uintptr(iPropId), uintptr(unsafe.Pointer(piVal)), 0)
	ret = HRESULT(r0)
	return
}

func GetThemeMargins(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, iPropId int32, prc *RECT, pMargins *MARGINS) (ret HRESULT) {
	r0, _, _ := syscall.Syscall9(procGetThemeMargins.Addr(), 7, uintptr(hTheme), uintptr(hdc), uintptr(iPartId), uintptr(iStateId), uintptr(iPropId), uintptr(unsafe.Pointer(prc)), uintptr(unsafe.Pointer(pMargins)), 0, 0)
	ret = HRESULT(r0)
	return
}

func GetThemeMetric(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, iPropId int32, piVal *int32) (ret HRESULT) {
	r0, _, _ := syscall.Syscall6(procGetThemeMetric.Addr(), 6, uintptr(hTheme), uintptr(hdc), uintptr(iPartId), uintptr(iStateId), uintptr(iPropId), uintptr(unsafe.Pointer(piVal)))
	ret = HRESULT(r0)
	return
}

func GetThemePartSize(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, prc *RECT, eSize THEMESIZE, psz *SIZE) (ret HRESULT) {
	r0, _, _ := syscall.Syscall9(procGetThemePartSize.Addr(), 7, uintptr(hTheme), uintptr(hdc), uintptr(iPartId), uintptr(iStateId), uintptr(unsafe.Pointer(prc)), uintptr(eSize), uintptr(unsafe.Pointer(psz)), 0, 0)
	ret = HRESULT(r0)
	return
}

func GetThemeSysFont(hTheme HTHEME, iFontId int32, plf *LOGFONT) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procGetThemeSysFont.Addr(), 3, uintptr(hTheme), uintptr(iFontId), uintptr(unsafe.Pointer(plf)))
	ret = HRESULT(r0)
	return
}

func GetThemeTextExtent(hTheme HTHEME, hdc HDC, iPartId int32, iStateId int32, pszText *uint16, iCharCount int32, dwTextFlags uint32, pBoundingRect *RECT, pExtentRect *RECT) (ret HRESULT) {
	r0, _, _ := syscall.Syscall9(procGetThemeTextExtent.Addr(), 9, uintptr(hTheme), uintptr(hdc), uintptr(iPartId), uintptr(iStateId), uintptr(unsafe.Pointer(pszText)), uintptr(iCharCount), uintptr(dwTextFlags), uintptr(unsafe.Pointer(pBoundingRect)), uintptr(unsafe.Pointer(pExtentRect)))
	ret = HRESULT(r0)
	return
}

func IsAppThemed() (ret bool) {
	r0, _, _ := syscall.Syscall(procIsAppThemed.Addr(), 0, 0, 0, 0)
	ret = r0 != 0
	return
}

func IsThemeBackgroundPartiallyTransparent(hTheme HTHEME, iPartId int32, iStateId int32) (ret bool) {
	r0, _, _ := syscall.Syscall(procIsThemeBackgroundPartiallyTransparent.Addr(), 3, uintptr(hTheme), uintptr(iPartId), uintptr(iStateId))
	ret = r0 != 0
	return
}

func OpenThemeData(hwnd HWND, pszClassList *uint16) (ret HTHEME) {
	r0, _, _ := syscall.Syscall(procOpenThemeData.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(pszClassList)), 0)
	ret = HTHEME(r0)
	return
}

func SetWindowTheme(hwnd HWND, pszSubAppName *uint16, pszSubIdList *uint16) (ret HRESULT) {
	r0, _, _ := syscall.Syscall(procSetWindowTheme.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(pszSubAppName)), uintptr(unsafe.Pointer(pszSubIdList)))
	ret = HRESULT(r0)
	return
}