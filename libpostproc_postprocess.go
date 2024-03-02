// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (C) 2001-2003 Michael Niedermayer (michaelni@gmx.at)
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

package ffmpeg

//#cgo pkg-config: libpostproc
//#include "libpostproc/version.h"
//#include <inttypes.h>
//#include "libpostproc/postprocess.h"
import "C"
import (
    "unsafe"
)

const PP_QUALITY_MAX = 6
const PP_CPU_CAPS_MMX = 0x80000000
const PP_CPU_CAPS_MMX2 = 0x20000000
const PP_CPU_CAPS_3DNOW = 0x40000000
const PP_CPU_CAPS_ALTIVEC = 0x10000000
const PP_CPU_CAPS_AUTO = 0x00080000
const PP_FORMAT = 0x00000008
const PP_FORMAT_420 = (0x00000011|PP_FORMAT)
const PP_FORMAT_422 = (0x00000001|PP_FORMAT)
const PP_FORMAT_411 = (0x00000002|PP_FORMAT)
const PP_FORMAT_444 = (0x00000000|PP_FORMAT)
const PP_FORMAT_440 = (0x00000010|PP_FORMAT)
const PP_PICT_TYPE_QP2 = 0x00000010


                              
                              

/**
 * @file
 * @ingroup lpp
 * external API header
 */

/**
 * @defgroup lpp libpostproc
 * Video postprocessing library.
 *
 * @{
 */

                                

/**
 * Return the LIBPOSTPROC_VERSION_INT constant.
 */
func Postproc_version() uint32 {
    return uint32(C.postproc_version())
}

/**
 * Return the libpostproc build-time configuration.
 */
func Postproc_configuration() string {
    return C.GoString(C.postproc_configuration())
}

/**
 * Return the libpostproc license.
 */
func Postproc_license() string {
    return C.GoString(C.postproc_license())
}



                     




                                      
                                
                          
                                                         
     
//extern const char pp_help[]; ///< a simple help text
      

func  Pp_postprocess(src[3] *uint8, srcStride[3] int32,
                     dst[3] *uint8, dstStride[3] int32,
                     horizontalSize int32, verticalSize int32,
                     QP_store *int8,  QP_stride int32,
                     mode unsafe.Pointer, ppContext unsafe.Pointer, pict_type int32)  {
    C.pp_postprocess((**C.uchar)(unsafe.Pointer(&src[0])), 
        (*C.int)(unsafe.Pointer(&srcStride[0])), 
        (**C.uchar)(unsafe.Pointer(&dst[0])), 
        (*C.int)(unsafe.Pointer(&dstStride[0])), C.int(horizontalSize), 
        C.int(verticalSize), (*C.schar)(unsafe.Pointer(QP_store)), 
        C.int(QP_stride), mode, ppContext, C.int(pict_type))
}


/**
 * Return a pp_mode or NULL if an error occurred.
 *
 * @param name    the string after "-pp" on the command line
 * @param quality a number from 0 to PP_QUALITY_MAX
 */
func Pp_get_mode_by_name_and_quality(name *byte, quality int32) unsafe.Pointer {
    return (unsafe.Pointer)(unsafe.Pointer(C.pp_get_mode_by_name_and_quality(
        (*C.char)(unsafe.Pointer(name)), C.int(quality))))
}
func Pp_free_mode(mode unsafe.Pointer)  {
    C.pp_free_mode(mode)
}

func Pp_get_context(width int32, height int32, flags int32) unsafe.Pointer {
    return (unsafe.Pointer)(unsafe.Pointer(C.pp_get_context(C.int(width), C.int(height), 
        C.int(flags))))
}
func Pp_free_context(ppContext unsafe.Pointer)  {
    C.pp_free_context(ppContext)
}














///< MPEG2 style QScale

/**
 * @}
 */

                                   
