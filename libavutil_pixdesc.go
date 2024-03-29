// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * pixel format descriptor
 * Copyright (c) 2009 Michael Niedermayer <michaelni@gmx.at>
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

package ffmpeg

//#cgo pkg-config: libavutil
//#include <inttypes.h>
//#include "libavutil/attributes.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/version.h"
//#include "libavutil/pixdesc.h"
import "C"
import (
    "unsafe"
)

const AV_PIX_FMT_FLAG_BE =            (1 << 0) 
const AV_PIX_FMT_FLAG_PAL =           (1 << 1) 
const AV_PIX_FMT_FLAG_BITSTREAM =     (1 << 2) 
const AV_PIX_FMT_FLAG_HWACCEL =       (1 << 3) 
const AV_PIX_FMT_FLAG_PLANAR =        (1 << 4) 
const AV_PIX_FMT_FLAG_RGB =           (1 << 5) 
const AV_PIX_FMT_FLAG_PSEUDOPAL =     (1 << 6) 
const AV_PIX_FMT_FLAG_ALPHA =         (1 << 7) 
const AV_PIX_FMT_FLAG_BAYER =         (1 << 8) 
const AV_PIX_FMT_FLAG_FLOAT =         (1 << 9) 
const FF_LOSS_RESOLUTION =   0x0001  
const FF_LOSS_DEPTH =        0x0002  
const FF_LOSS_COLORSPACE =   0x0004  
const FF_LOSS_ALPHA =        0x0008  
const FF_LOSS_COLORQUANT =   0x0010  
const FF_LOSS_CHROMA =       0x0020  



                        
                        

                     

                       
                   
                    

type AVComponentDescriptor struct {
    Plane int32
    Step int32
    Offset int32
    Shift int32
    Depth int32
    Step_minus1 int32
    Depth_minus1 int32
    Offset_plus1 int32
}


/**
 * Descriptor that unambiguously describes how the bits of a pixel are
 * stored in the up to 4 data planes of an image. It also stores the
 * subsampling factors and number of components.
 *
 * @note This is separate of the colorspace (RGB, YCbCr, YPbPr, JPEG-style YUV
 *       and all the YUV variants) AVPixFmtDescriptor just stores how values
 *       are stored not what these values represent.
 */
type AVPixFmtDescriptor struct {
    Name *byte
    Nb_components uint8
    Log2_chroma_w uint8
    Log2_chroma_h uint8
    Flags uint64
    Comp [4]AVComponentDescriptor
    Alias *byte
}


/**
 * Pixel format is big-endian.
 */

/**
 * Pixel format has a palette in data[1], values are indexes in this palette.
 */

/**
 * All values of a component are bit-wise packed end to end.
 */

/**
 * Pixel format is an HW accelerated format.
 */

/**
 * At least one pixel component is not in the first data plane.
 */

/**
 * The pixel format contains RGB-like data (as opposed to YUV/grayscale).
 */


/**
 * The pixel format is "pseudo-paletted". This means that it contains a
 * fixed palette in the 2nd plane but the palette is fixed/constant for each
 * PIX_FMT. This allows interpreting the data as if it was PAL8, which can
 * in some cases be simpler. Or the data can be interpreted purely based on
 * the pixel format without using the palette.
 * An example of a pseudo-paletted format is AV_PIX_FMT_GRAY8
 *
 * @deprecated This flag is deprecated, and will be removed. When it is removed,
 * the extra palette allocation in AVFrame.data[1] is removed as well. Only
 * actual paletted formats (as indicated by AV_PIX_FMT_FLAG_PAL) will have a
 * palette. Starting with FFmpeg versions which have this flag deprecated, the
 * extra "pseudo" palette is already ignored, and API users are not required to
 * allocate a palette for AV_PIX_FMT_FLAG_PSEUDOPAL formats (it was required
 * before the deprecation, though).
 */


/**
 * The pixel format has an alpha channel. This is set on all formats that
 * support alpha in some way, including AV_PIX_FMT_PAL8. The alpha is always
 * straight, never pre-multiplied.
 *
 * If a codec or a filter does not support alpha, it should set all alpha to
 * opaque, or use the equivalent pixel formats without alpha component, e.g.
 * AV_PIX_FMT_RGB0 (or AV_PIX_FMT_RGB24 etc.) instead of AV_PIX_FMT_RGBA.
 */


/**
 * The pixel format is following a Bayer pattern
 */


/**
 * The pixel format contains IEEE-754 floating point values. Precision (double,
 * single, or half) should be determined by the pixel size (64, 32, or 16 bits).
 */


/**
 * Return the number of bits per pixel used by the pixel format
 * described by pixdesc. Note that this is not the same as the number
 * of bits per sample.
 *
 * The returned number of bits refers to the number of bits actually
 * used for storing the pixel information, that is padding bits are
 * not counted.
 */
func Av_get_bits_per_pixel(pixdesc *AVPixFmtDescriptor) int32 {
    return int32(C.av_get_bits_per_pixel(
        (*C.struct_AVPixFmtDescriptor)(unsafe.Pointer(pixdesc))))
}

/**
 * Return the number of bits per pixel for the pixel format
 * described by pixdesc, including any padding or unused bits.
 */
func Av_get_padded_bits_per_pixel(pixdesc *AVPixFmtDescriptor) int32 {
    return int32(C.av_get_padded_bits_per_pixel(
        (*C.struct_AVPixFmtDescriptor)(unsafe.Pointer(pixdesc))))
}

/**
 * @return a pixel format descriptor for provided pixel format or NULL if
 * this pixel format is unknown.
 */
func Av_pix_fmt_desc_get(pix_fmt AVPixelFormat) *AVPixFmtDescriptor {
    return (*AVPixFmtDescriptor)(unsafe.Pointer(C.av_pix_fmt_desc_get(
        C.enum_AVPixelFormat(pix_fmt))))
}

/**
 * Iterate over all pixel format descriptors known to libavutil.
 *
 * @param prev previous descriptor. NULL to get the first descriptor.
 *
 * @return next descriptor or NULL after the last descriptor
 */
func Av_pix_fmt_desc_next(prev *AVPixFmtDescriptor) *AVPixFmtDescriptor {
    return (*AVPixFmtDescriptor)(unsafe.Pointer(C.av_pix_fmt_desc_next(
        (*C.struct_AVPixFmtDescriptor)(unsafe.Pointer(prev)))))
}

/**
 * @return an AVPixelFormat id described by desc, or AV_PIX_FMT_NONE if desc
 * is not a valid pointer to a pixel format descriptor.
 */
func Av_pix_fmt_desc_get_id(desc *AVPixFmtDescriptor) AVPixelFormat {
    return AVPixelFormat(C.av_pix_fmt_desc_get_id(
        (*C.struct_AVPixFmtDescriptor)(unsafe.Pointer(desc))))
}

/**
 * Utility function to access log2_chroma_w log2_chroma_h from
 * the pixel format AVPixFmtDescriptor.
 *
 * @param[in]  pix_fmt the pixel format
 * @param[out] h_shift store log2_chroma_w (horizontal/width shift)
 * @param[out] v_shift store log2_chroma_h (vertical/height shift)
 *
 * @return 0 on success, AVERROR(ENOSYS) on invalid or unknown pixel format
 */
func Av_pix_fmt_get_chroma_sub_sample(pix_fmt AVPixelFormat,
                                     h_shift *int32, v_shift *int32) int32 {
    return int32(C.av_pix_fmt_get_chroma_sub_sample(C.enum_AVPixelFormat(pix_fmt), 
        (*C.int)(unsafe.Pointer(h_shift)), (*C.int)(unsafe.Pointer(v_shift))))
}

/**
 * @return number of planes in pix_fmt, a negative AVERROR if pix_fmt is not a
 * valid pixel format.
 */
func Av_pix_fmt_count_planes(pix_fmt AVPixelFormat) int32 {
    return int32(C.av_pix_fmt_count_planes(C.enum_AVPixelFormat(pix_fmt)))
}

/**
 * @return the name for provided color range or NULL if unknown.
 */
func Av_color_range_name(rangex AVColorRange) string {
    return C.GoString(C.av_color_range_name(C.enum_AVColorRange(rangex)))
}

/**
 * @return the AVColorRange value for name or an AVError if not found.
 */
func Av_color_range_from_name(name *byte) int32 {
    return int32(C.av_color_range_from_name((*C.char)(unsafe.Pointer(name))))
}

/**
 * @return the name for provided color primaries or NULL if unknown.
 */
func Av_color_primaries_name(primaries AVColorPrimaries) string {
    return C.GoString(C.av_color_primaries_name(C.enum_AVColorPrimaries(primaries)))
}

/**
 * @return the AVColorPrimaries value for name or an AVError if not found.
 */
func Av_color_primaries_from_name(name *byte) int32 {
    return int32(C.av_color_primaries_from_name((*C.char)(unsafe.Pointer(name))))
}

/**
 * @return the name for provided color transfer or NULL if unknown.
 */
func Av_color_transfer_name(transfer AVColorTransferCharacteristic) string {
    return C.GoString(C.av_color_transfer_name(
        C.enum_AVColorTransferCharacteristic(transfer)))
}

/**
 * @return the AVColorTransferCharacteristic value for name or an AVError if not found.
 */
func Av_color_transfer_from_name(name *byte) int32 {
    return int32(C.av_color_transfer_from_name((*C.char)(unsafe.Pointer(name))))
}

/**
 * @return the name for provided color space or NULL if unknown.
 */
func Av_color_space_name(space AVColorSpace) string {
    return C.GoString(C.av_color_space_name(C.enum_AVColorSpace(space)))
}

/**
 * @return the AVColorSpace value for name or an AVError if not found.
 */
func Av_color_space_from_name(name *byte) int32 {
    return int32(C.av_color_space_from_name((*C.char)(unsafe.Pointer(name))))
}

/**
 * @return the name for provided chroma location or NULL if unknown.
 */
func Av_chroma_location_name(location AVChromaLocation) string {
    return C.GoString(C.av_chroma_location_name(C.enum_AVChromaLocation(location)))
}

/**
 * @return the AVChromaLocation value for name or an AVError if not found.
 */
func Av_chroma_location_from_name(name *byte) int32 {
    return int32(C.av_chroma_location_from_name((*C.char)(unsafe.Pointer(name))))
}

/**
 * Return the pixel format corresponding to name.
 *
 * If there is no pixel format with name name, then looks for a
 * pixel format with the name corresponding to the native endian
 * format of name.
 * For example in a little-endian system, first looks for "gray16",
 * then for "gray16le".
 *
 * Finally if no pixel format has been found, returns AV_PIX_FMT_NONE.
 */
func Av_get_pix_fmt(name *byte) AVPixelFormat {
    return AVPixelFormat(C.av_get_pix_fmt((*C.char)(unsafe.Pointer(name))))
}

/**
 * Return the short name for a pixel format, NULL in case pix_fmt is
 * unknown.
 *
 * @see av_get_pix_fmt(), av_get_pix_fmt_string()
 */
func Av_get_pix_fmt_name(pix_fmt AVPixelFormat) string {
    return C.GoString(C.av_get_pix_fmt_name(C.enum_AVPixelFormat(pix_fmt)))
}

/**
 * Print in buf the string corresponding to the pixel format with
 * number pix_fmt, or a header if pix_fmt is negative.
 *
 * @param buf the buffer where to write the string
 * @param buf_size the size of buf
 * @param pix_fmt the number of the pixel format to print the
 * corresponding info string, or a negative value to print the
 * corresponding header.
 */
func Av_get_pix_fmt_string(buf *byte, buf_size int32,
                            pix_fmt AVPixelFormat) string {
    return C.GoString(C.av_get_pix_fmt_string((*C.char)(unsafe.Pointer(buf)), 
        C.int(buf_size), C.enum_AVPixelFormat(pix_fmt)))
}

/**
 * Read a line from an image, and write the values of the
 * pixel format component c to dst.
 *
 * @param data the array containing the pointers to the planes of the image
 * @param linesize the array containing the linesizes of the image
 * @param desc the pixel format descriptor for the image
 * @param x the horizontal coordinate of the first pixel to read
 * @param y the vertical coordinate of the first pixel to read
 * @param w the width of the line to read, that is the number of
 * values to write to dst
 * @param read_pal_component if not zero and the format is a paletted
 * format writes the values corresponding to the palette
 * component c in data[1] to dst, rather than the palette indexes in
 * data[0]. The behavior is undefined if the format is not paletted.
 * @param dst_element_size size of elements in dst array (2 or 4 byte)
 */
func Av_read_image_line2(dst unsafe.Pointer, data [4]*uint8,
                        linesize [4]int32, desc *AVPixFmtDescriptor,
                        x int32, y int32, c int32, w int32, read_pal_component int32,
                        dst_element_size int32)  {
    C.av_read_image_line2(dst, (**C.uchar)(unsafe.Pointer(&data[0])), 
        (*C.int)(unsafe.Pointer(&linesize[0])), 
        (*C.struct_AVPixFmtDescriptor)(unsafe.Pointer(desc)), C.int(x), C.int(y), 
        C.int(c), C.int(w), C.int(read_pal_component), C.int(dst_element_size))
}

func Av_read_image_line(dst *uint16, data [4]*uint8,
                        linesize [4]int32, desc *AVPixFmtDescriptor,
                        x int32, y int32, c int32, w int32, read_pal_component int32)  {
    C.av_read_image_line((*C.ushort)(unsafe.Pointer(dst)), 
        (**C.uchar)(unsafe.Pointer(&data[0])), 
        (*C.int)(unsafe.Pointer(&linesize[0])), 
        (*C.struct_AVPixFmtDescriptor)(unsafe.Pointer(desc)), C.int(x), C.int(y), 
        C.int(c), C.int(w), C.int(read_pal_component))
}

/**
 * Write the values from src to the pixel format component c of an
 * image line.
 *
 * @param src array containing the values to write
 * @param data the array containing the pointers to the planes of the
 * image to write into. It is supposed to be zeroed.
 * @param linesize the array containing the linesizes of the image
 * @param desc the pixel format descriptor for the image
 * @param x the horizontal coordinate of the first pixel to write
 * @param y the vertical coordinate of the first pixel to write
 * @param w the width of the line to write, that is the number of
 * values to write to the image line
 * @param src_element_size size of elements in src array (2 or 4 byte)
 */
func Av_write_image_line2(src unsafe.Pointer, data [4]*uint8,
                         linesize [4]int32, desc *AVPixFmtDescriptor,
                         x int32, y int32, c int32, w int32, src_element_size int32)  {
    C.av_write_image_line2(src, (**C.uchar)(unsafe.Pointer(&data[0])), 
        (*C.int)(unsafe.Pointer(&linesize[0])), 
        (*C.struct_AVPixFmtDescriptor)(unsafe.Pointer(desc)), C.int(x), C.int(y), 
        C.int(c), C.int(w), C.int(src_element_size))
}

func Av_write_image_line(src *uint16, data [4]*uint8,
                         linesize [4]int32, desc *AVPixFmtDescriptor,
                         x int32, y int32, c int32, w int32)  {
    C.av_write_image_line((*C.ushort)(unsafe.Pointer(src)), 
        (**C.uchar)(unsafe.Pointer(&data[0])), 
        (*C.int)(unsafe.Pointer(&linesize[0])), 
        (*C.struct_AVPixFmtDescriptor)(unsafe.Pointer(desc)), C.int(x), C.int(y), 
        C.int(c), C.int(w))
}

/**
 * Utility function to swap the endianness of a pixel format.
 *
 * @param[in]  pix_fmt the pixel format
 *
 * @return pixel format with swapped endianness if it exists,
 * otherwise AV_PIX_FMT_NONE
 */
func Av_pix_fmt_swap_endianness(pix_fmt AVPixelFormat) AVPixelFormat {
    return AVPixelFormat(C.av_pix_fmt_swap_endianness(C.enum_AVPixelFormat(pix_fmt)))
}

/**< loss due to resolution change */
/**< loss due to color depth change */
/**< loss due to color space conversion */
/**< loss of alpha bits */
/**< loss due to color quantization */
/**< loss of chroma (e.g. RGB to gray conversion) */

/**
 * Compute what kind of losses will occur when converting from one specific
 * pixel format to another.
 * When converting from one pixel format to another, information loss may occur.
 * For example, when converting from RGB24 to GRAY, the color information will
 * be lost. Similarly, other losses occur when converting from some formats to
 * other formats. These losses can involve loss of chroma, but also loss of
 * resolution, loss of color depth, loss due to the color space conversion, loss
 * of the alpha bits or loss due to color quantization.
 * av_get_fix_fmt_loss() informs you about the various types of losses
 * which will occur when converting from one pixel format to another.
 *
 * @param[in] dst_pix_fmt destination pixel format
 * @param[in] src_pix_fmt source pixel format
 * @param[in] has_alpha Whether the source pixel format alpha channel is used.
 * @return Combination of flags informing you what kind of losses will occur
 * (maximum loss for an invalid dst_pix_fmt).
 */
func Av_get_pix_fmt_loss(dst_pix_fmt AVPixelFormat,
                        src_pix_fmt AVPixelFormat,
                        has_alpha int32) int32 {
    return int32(C.av_get_pix_fmt_loss(C.enum_AVPixelFormat(dst_pix_fmt), 
        C.enum_AVPixelFormat(src_pix_fmt), C.int(has_alpha)))
}

/**
 * Compute what kind of losses will occur when converting from one specific
 * pixel format to another.
 * When converting from one pixel format to another, information loss may occur.
 * For example, when converting from RGB24 to GRAY, the color information will
 * be lost. Similarly, other losses occur when converting from some formats to
 * other formats. These losses can involve loss of chroma, but also loss of
 * resolution, loss of color depth, loss due to the color space conversion, loss
 * of the alpha bits or loss due to color quantization.
 * av_get_fix_fmt_loss() informs you about the various types of losses
 * which will occur when converting from one pixel format to another.
 *
 * @param[in] dst_pix_fmt destination pixel format
 * @param[in] src_pix_fmt source pixel format
 * @param[in] has_alpha Whether the source pixel format alpha channel is used.
 * @return Combination of flags informing you what kind of losses will occur
 * (maximum loss for an invalid dst_pix_fmt).
 */
func Av_find_best_pix_fmt_of_2(dst_pix_fmt1 AVPixelFormat, dst_pix_fmt2 AVPixelFormat,
                                             src_pix_fmt AVPixelFormat, has_alpha int32, loss_ptr *int32) AVPixelFormat {
    return AVPixelFormat(C.av_find_best_pix_fmt_of_2(
        C.enum_AVPixelFormat(dst_pix_fmt1), C.enum_AVPixelFormat(dst_pix_fmt2), 
        C.enum_AVPixelFormat(src_pix_fmt), C.int(has_alpha), 
        (*C.int)(unsafe.Pointer(loss_ptr))))
}

                             
