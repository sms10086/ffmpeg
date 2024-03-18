// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * Copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
 * Copyright (c) 2008 Peter Ross
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
//#include <stdint.h>
//#include "libavutil/channel_layout.h"
import "C"
import (
    "unsafe"
)

const AV_CH_FRONT_LEFT =              0x00000001 
const AV_CH_FRONT_RIGHT =             0x00000002 
const AV_CH_FRONT_CENTER =            0x00000004 
const AV_CH_LOW_FREQUENCY =           0x00000008 
const AV_CH_BACK_LEFT =               0x00000010 
const AV_CH_BACK_RIGHT =              0x00000020 
const AV_CH_FRONT_LEFT_OF_CENTER =    0x00000040 
const AV_CH_FRONT_RIGHT_OF_CENTER =   0x00000080 
const AV_CH_BACK_CENTER =             0x00000100 
const AV_CH_SIDE_LEFT =               0x00000200 
const AV_CH_SIDE_RIGHT =              0x00000400 
const AV_CH_TOP_CENTER =              0x00000800 
const AV_CH_TOP_FRONT_LEFT =          0x00001000 
const AV_CH_TOP_FRONT_CENTER =        0x00002000 
const AV_CH_TOP_FRONT_RIGHT =         0x00004000 
const AV_CH_TOP_BACK_LEFT =           0x00008000 
const AV_CH_TOP_BACK_CENTER =         0x00010000 
const AV_CH_TOP_BACK_RIGHT =          0x00020000 
const AV_CH_STEREO_LEFT =             0x20000000   
const AV_CH_STEREO_RIGHT =            0x40000000   
const AV_CH_WIDE_LEFT =               uint64(0x0000000080000000) 
const AV_CH_WIDE_RIGHT =              uint64(0x0000000100000000) 
const AV_CH_SURROUND_DIRECT_LEFT =    uint64(0x0000000200000000) 
const AV_CH_SURROUND_DIRECT_RIGHT =   uint64(0x0000000400000000) 
const AV_CH_LOW_FREQUENCY_2 =         uint64(0x0000000800000000) 
const AV_CH_LAYOUT_NATIVE =           uint64(0x8000000000000000) 
const AV_CH_LAYOUT_MONO =               (AV_CH_FRONT_CENTER) 
const AV_CH_LAYOUT_STEREO =             (AV_CH_FRONT_LEFT|AV_CH_FRONT_RIGHT) 
const AV_CH_LAYOUT_2POINT1 =            (AV_CH_LAYOUT_STEREO|AV_CH_LOW_FREQUENCY) 
const AV_CH_LAYOUT_2_1 =                (AV_CH_LAYOUT_STEREO|AV_CH_BACK_CENTER) 
const AV_CH_LAYOUT_SURROUND =           (AV_CH_LAYOUT_STEREO|AV_CH_FRONT_CENTER) 
const AV_CH_LAYOUT_3POINT1 =            (AV_CH_LAYOUT_SURROUND|AV_CH_LOW_FREQUENCY) 
const AV_CH_LAYOUT_4POINT0 =            (AV_CH_LAYOUT_SURROUND|AV_CH_BACK_CENTER) 
const AV_CH_LAYOUT_4POINT1 =            (AV_CH_LAYOUT_4POINT0|AV_CH_LOW_FREQUENCY) 
const AV_CH_LAYOUT_2_2 =                (AV_CH_LAYOUT_STEREO|AV_CH_SIDE_LEFT|AV_CH_SIDE_RIGHT) 
const AV_CH_LAYOUT_QUAD =               (AV_CH_LAYOUT_STEREO|AV_CH_BACK_LEFT|AV_CH_BACK_RIGHT) 
const AV_CH_LAYOUT_5POINT0 =            (AV_CH_LAYOUT_SURROUND|AV_CH_SIDE_LEFT|AV_CH_SIDE_RIGHT) 
const AV_CH_LAYOUT_5POINT1 =            (AV_CH_LAYOUT_5POINT0|AV_CH_LOW_FREQUENCY) 
const AV_CH_LAYOUT_5POINT0_BACK =       (AV_CH_LAYOUT_SURROUND|AV_CH_BACK_LEFT|AV_CH_BACK_RIGHT) 
const AV_CH_LAYOUT_5POINT1_BACK =       (AV_CH_LAYOUT_5POINT0_BACK|AV_CH_LOW_FREQUENCY) 
const AV_CH_LAYOUT_6POINT0 =            (AV_CH_LAYOUT_5POINT0|AV_CH_BACK_CENTER) 
const AV_CH_LAYOUT_6POINT0_FRONT =      (AV_CH_LAYOUT_2_2|AV_CH_FRONT_LEFT_OF_CENTER|AV_CH_FRONT_RIGHT_OF_CENTER) 
const AV_CH_LAYOUT_HEXAGONAL =          (AV_CH_LAYOUT_5POINT0_BACK|AV_CH_BACK_CENTER) 
const AV_CH_LAYOUT_6POINT1 =            (AV_CH_LAYOUT_5POINT1|AV_CH_BACK_CENTER) 
const AV_CH_LAYOUT_6POINT1_BACK =       (AV_CH_LAYOUT_5POINT1_BACK|AV_CH_BACK_CENTER) 
const AV_CH_LAYOUT_6POINT1_FRONT =      (AV_CH_LAYOUT_6POINT0_FRONT|AV_CH_LOW_FREQUENCY) 
const AV_CH_LAYOUT_7POINT0 =            (AV_CH_LAYOUT_5POINT0|AV_CH_BACK_LEFT|AV_CH_BACK_RIGHT) 
const AV_CH_LAYOUT_7POINT0_FRONT =      (AV_CH_LAYOUT_5POINT0|AV_CH_FRONT_LEFT_OF_CENTER|AV_CH_FRONT_RIGHT_OF_CENTER) 
const AV_CH_LAYOUT_7POINT1 =            (AV_CH_LAYOUT_5POINT1|AV_CH_BACK_LEFT|AV_CH_BACK_RIGHT) 
const AV_CH_LAYOUT_7POINT1_WIDE =       (AV_CH_LAYOUT_5POINT1|AV_CH_FRONT_LEFT_OF_CENTER|AV_CH_FRONT_RIGHT_OF_CENTER) 
const AV_CH_LAYOUT_7POINT1_WIDE_BACK =  (AV_CH_LAYOUT_5POINT1_BACK|AV_CH_FRONT_LEFT_OF_CENTER|AV_CH_FRONT_RIGHT_OF_CENTER) 
const AV_CH_LAYOUT_OCTAGONAL =          (AV_CH_LAYOUT_5POINT0|AV_CH_BACK_LEFT|AV_CH_BACK_CENTER|AV_CH_BACK_RIGHT) 
const AV_CH_LAYOUT_HEXADECAGONAL =      (AV_CH_LAYOUT_OCTAGONAL|AV_CH_WIDE_LEFT|AV_CH_WIDE_RIGHT|AV_CH_TOP_BACK_LEFT|AV_CH_TOP_BACK_RIGHT|AV_CH_TOP_BACK_CENTER|AV_CH_TOP_FRONT_CENTER|AV_CH_TOP_FRONT_LEFT|AV_CH_TOP_FRONT_RIGHT) 
const AV_CH_LAYOUT_STEREO_DOWNMIX =     (AV_CH_STEREO_LEFT|AV_CH_STEREO_RIGHT) 



                               
                               

                   

/**
 * @file
 * audio channel layout utility functions
 */

/**
 * @addtogroup lavu_audio
 * @{
 */

/**
 * @defgroup channel_masks Audio channel masks
 *
 * A channel layout is a 64-bits integer with a bit set for every channel.
 * The number of bits set must be equal to the number of channels.
 * The value 0 means that the channel layout is not known.
 * @note this data structure is not powerful enough to handle channels
 * combinations that have the same channel multiple times, such as
 * dual-mono.
 *
 * @{
 */


















///< Stereo downmix.
///< See AV_CH_STEREO_LEFT.






/** Channel mask value used for AVCodecContext.request_channel_layout
    to indicate that the user requests the channel order of the decoder output
    to be the native codec channel order. */


/**
 * @}
 * @defgroup channel_mask_c Audio channel layouts
 * @{
 * */





























type AVMatrixEncoding int32
const (
    AV_MATRIX_ENCODING_NONE AVMatrixEncoding = iota
    AV_MATRIX_ENCODING_DOLBY
    AV_MATRIX_ENCODING_DPLII
    AV_MATRIX_ENCODING_DPLIIX
    AV_MATRIX_ENCODING_DPLIIZ
    AV_MATRIX_ENCODING_DOLBYEX
    AV_MATRIX_ENCODING_DOLBYHEADPHONE
    AV_MATRIX_ENCODING_NB
)


/**
 * Return a channel layout id that matches name, or 0 if no match is found.
 *
 * name can be one or several of the following notations,
 * separated by '+' or '|':
 * - the name of an usual channel layout (mono, stereo, 4.0, quad, 5.0,
 *   5.0(side), 5.1, 5.1(side), 7.1, 7.1(wide), downmix);
 * - the name of a single channel (FL, FR, FC, LFE, BL, BR, FLC, FRC, BC,
 *   SL, SR, TC, TFL, TFC, TFR, TBL, TBC, TBR, DL, DR);
 * - a number of channels, in decimal, followed by 'c', yielding
 *   the default channel layout for that number of channels (@see
 *   av_get_default_channel_layout);
 * - a channel layout mask, in hexadecimal starting with "0x" (see the
 *   AV_CH_* macros).
 *
 * Example: "stereo+FC" = "2c+FC" = "2c+1c" = "0x7"
 */
func Av_get_channel_layout(name *byte) uint64 {
    return uint64(C.av_get_channel_layout((*C.char)(unsafe.Pointer(name))))
}

/**
 * Return a channel layout and the number of channels based on the specified name.
 *
 * This function is similar to (@see av_get_channel_layout), but can also parse
 * unknown channel layout specifications.
 *
 * @param[in]  name             channel layout specification string
 * @param[out] channel_layout   parsed channel layout (0 if unknown)
 * @param[out] nb_channels      number of channels
 *
 * @return 0 on success, AVERROR(EINVAL) if the parsing fails.
 */
func Av_get_extended_channel_layout(name *byte, channel_layout *uint64, nb_channels *int32) int32 {
    return int32(C.av_get_extended_channel_layout((*C.char)(unsafe.Pointer(name)), 
        (*C.ulonglong)(unsafe.Pointer(channel_layout)), 
        (*C.int)(unsafe.Pointer(nb_channels))))
}

/**
 * Return a description of a channel layout.
 * If nb_channels is <= 0, it is guessed from the channel_layout.
 *
 * @param buf put here the string containing the channel layout
 * @param buf_size size in bytes of the buffer
 */
func Av_get_channel_layout_string(buf *byte, buf_size int32, nb_channels int32, channel_layout uint64)  {
    C.av_get_channel_layout_string((*C.char)(unsafe.Pointer(buf)), C.int(buf_size), 
        C.int(nb_channels), C.ulonglong(channel_layout))
}

// type AVBPrint
/**
 * Append a description of a channel layout to a bprint buffer.
 */
func Av_bprint_channel_layout(bp *AVBPrint, nb_channels int32, channel_layout uint64)  {
    C.av_bprint_channel_layout((*C.struct_AVBPrint)(unsafe.Pointer(bp)), 
        C.int(nb_channels), C.ulonglong(channel_layout))
}

/**
 * Return the number of channels in the channel layout.
 */
func Av_get_channel_layout_nb_channels(channel_layout uint64) int32 {
    return int32(C.av_get_channel_layout_nb_channels(C.ulonglong(channel_layout)))
}

/**
 * Return default channel layout for a given number of channels.
 */
func Av_get_default_channel_layout(nb_channels int32) int64 {
    return int64(C.av_get_default_channel_layout(C.int(nb_channels)))
}

/**
 * Get the index of a channel in channel_layout.
 *
 * @param channel a channel layout describing exactly one channel which must be
 *                present in channel_layout.
 *
 * @return index of channel in channel_layout on success, a negative AVERROR
 *         on error.
 */
func Av_get_channel_layout_channel_index(channel_layout uint64,
                                        channel uint64) int32 {
    return int32(C.av_get_channel_layout_channel_index(C.ulonglong(channel_layout), 
        C.ulonglong(channel)))
}

/**
 * Get the channel with the given index in channel_layout.
 */
func Av_channel_layout_extract_channel(channel_layout uint64, index int32) uint64 {
    return uint64(C.av_channel_layout_extract_channel(C.ulonglong(channel_layout), 
        C.int(index)))
}

/**
 * Get the name of a given channel.
 *
 * @return channel name on success, NULL on error.
 */
func Av_get_channel_name(channel uint64) string {
    return C.GoString(C.av_get_channel_name(C.ulonglong(channel)))
}

/**
 * Get the description of a given channel.
 *
 * @param channel  a channel layout with a single channel
 * @return  channel description on success, NULL on error
 */
func Av_get_channel_description(channel uint64) string {
    return C.GoString(C.av_get_channel_description(C.ulonglong(channel)))
}

/**
 * Get the value and name of a standard channel layout.
 *
 * @param[in]  index   index in an internal list, starting at 0
 * @param[out] layout  channel layout mask
 * @param[out] name    name of the layout
 * @return  0  if the layout exists,
 *          <0 if index is beyond the limits
 */
func Av_get_standard_channel_layout(index uint32, layout *uint64,
                                   name **byte) int32 {
    return int32(C.av_get_standard_channel_layout(C.unsigned(index), 
        (*C.ulonglong)(unsafe.Pointer(layout)), (**C.char)(unsafe.Pointer(name))))
}

/**
 * @}
 * @}
 */

                                    
