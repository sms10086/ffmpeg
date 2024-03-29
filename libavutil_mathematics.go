// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2005-2012 Michael Niedermayer <michaelni@gmx.at>
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
//#include <math.h>
//#include "libavutil/attributes.h"
//#include "libavutil/rational.h"
//#include "libavutil/intfloat.h"
//#include "libavutil/mathematics.h"
import "C"
import (
    "unsafe"
)

//const NAN =             av_int2float(0x7fc00000) 
//const INFINITY =        av_int2float(0x7f800000) 



/**
 * @file
 * @addtogroup lavu_math
 * Mathematical utilities for working with timestamp and time base.
 */

                            
                            

                   
                 
                       
                     
                     

           
                                                      
      
             
                                                            
      
              
                                                             
      
                 
                                                             
      
             
                                                                        
      
            
                                                       
      
              
                                                         
      
                 
                                                              
      
               
                                                            
      
           

      
                

      

/**
 * @addtogroup lavu_math
 *
 * @{
 */

/**
 * Rounding methods.
 */
type AVRounding int32
const (
    AV_ROUND_ZERO AVRounding = 0 + iota
    AV_ROUND_INF = 1
    AV_ROUND_DOWN = 2
    AV_ROUND_UP = 3
    AV_ROUND_NEAR_INF = 5
    AV_ROUND_PASS_MINMAX = 8192
)


/**
 * Compute the greatest common divisor of two integer operands.
 *
 * @param a,b Operands
 * @return GCD of a and b up to sign; if a >= 0 and b >= 0, return value is >= 0;
 * if a == 0 and b == 0, returns 0.
 */
func  Av_gcd(a int64, b int64) int64 {
    return int64(C.av_gcd(C.longlong(a), C.longlong(b)))
}

/**
 * Rescale a 64-bit integer with rounding to nearest.
 *
 * The operation is mathematically equivalent to `a * b / c`, but writing that
 * directly can overflow.
 *
 * This function is equivalent to av_rescale_rnd() with #AV_ROUND_NEAR_INF.
 *
 * @see av_rescale_rnd(), av_rescale_q(), av_rescale_q_rnd()
 */
func Av_rescale(a int64, b int64, c int64)  int64 {
    return int64(C.av_rescale(C.longlong(a), C.longlong(b), C.longlong(c)))
}

/**
 * Rescale a 64-bit integer with specified rounding.
 *
 * The operation is mathematically equivalent to `a * b / c`, but writing that
 * directly can overflow, and does not support different rounding methods.
 *
 * @see av_rescale(), av_rescale_q(), av_rescale_q_rnd()
 */
func Av_rescale_rnd(a int64, b int64, c int64, rnd AVRounding)  int64 {
    return int64(C.av_rescale_rnd(C.longlong(a), C.longlong(b), C.longlong(c), 
        C.enum_AVRounding(rnd)))
}

/**
 * Rescale a 64-bit integer by 2 rational numbers.
 *
 * The operation is mathematically equivalent to `a * bq / cq`.
 *
 * This function is equivalent to av_rescale_q_rnd() with #AV_ROUND_NEAR_INF.
 *
 * @see av_rescale(), av_rescale_rnd(), av_rescale_q_rnd()
 */
func Av_rescale_q(a int64, bq AVRational, cq AVRational)  int64 {
    return int64(C.av_rescale_q(C.longlong(a), 
        *(*C.struct_AVRational)(unsafe.Pointer(&bq)), 
        *(*C.struct_AVRational)(unsafe.Pointer(&cq))))
}

/**
 * Rescale a 64-bit integer by 2 rational numbers with specified rounding.
 *
 * The operation is mathematically equivalent to `a * bq / cq`.
 *
 * @see av_rescale(), av_rescale_rnd(), av_rescale_q()
 */
func Av_rescale_q_rnd(a int64, bq AVRational, cq AVRational,
                         rnd AVRounding)  int64 {
    return int64(C.av_rescale_q_rnd(C.longlong(a), 
        *(*C.struct_AVRational)(unsafe.Pointer(&bq)), 
        *(*C.struct_AVRational)(unsafe.Pointer(&cq)), C.enum_AVRounding(rnd)))
}

/**
 * Compare two timestamps each in its own time base.
 *
 * @return One of the following values:
 *         - -1 if `ts_a` is before `ts_b`
 *         - 1 if `ts_a` is after `ts_b`
 *         - 0 if they represent the same position
 *
 * @warning
 * The result of the function is undefined if one of the timestamps is outside
 * the `int64_t` range when represented in the other's timebase.
 */
func Av_compare_ts(ts_a int64, tb_a AVRational, ts_b int64, tb_b AVRational) int32 {
    return int32(C.av_compare_ts(C.longlong(ts_a), 
        *(*C.struct_AVRational)(unsafe.Pointer(&tb_a)), C.longlong(ts_b), 
        *(*C.struct_AVRational)(unsafe.Pointer(&tb_b))))
}

/**
 * Compare the remainders of two integer operands divided by a common divisor.
 *
 * In other words, compare the least significant `log2(mod)` bits of integers
 * `a` and `b`.
 *
 * @code{.c}
 * av_compare_mod(0x11, 0x02, 0x10) < 0 // since 0x11 % 0x10  (0x1) < 0x02 % 0x10  (0x2)
 * av_compare_mod(0x11, 0x02, 0x20) > 0 // since 0x11 % 0x20 (0x11) > 0x02 % 0x20 (0x02)
 * @endcode
 *
 * @param a,b Operands
 * @param mod Divisor; must be a power of 2
 * @return
 *         - a negative value if `a % mod < b % mod`
 *         - a positive value if `a % mod > b % mod`
 *         - zero             if `a % mod == b % mod`
 */
func Av_compare_mod(a uint64, b uint64, mod uint64) int64 {
    return int64(C.av_compare_mod(C.ulonglong(a), C.ulonglong(b), C.ulonglong(mod)))
}

/**
 * Rescale a timestamp while preserving known durations.
 *
 * This function is designed to be called per audio packet to scale the input
 * timestamp to a different time base. Compared to a simple av_rescale_q()
 * call, this function is robust against possible inconsistent frame durations.
 *
 * The `last` parameter is a state variable that must be preserved for all
 * subsequent calls for the same stream. For the first call, `*last` should be
 * initialized to #AV_NOPTS_VALUE.
 *
 * @param[in]     in_tb    Input time base
 * @param[in]     in_ts    Input timestamp
 * @param[in]     fs_tb    Duration time base; typically this is finer-grained
 *                         (greater) than `in_tb` and `out_tb`
 * @param[in]     duration Duration till the next call to this function (i.e.
 *                         duration of the current packet/frame)
 * @param[in,out] last     Pointer to a timestamp expressed in terms of
 *                         `fs_tb`, acting as a state variable
 * @param[in]     out_tb   Output timebase
 * @return        Timestamp expressed in terms of `out_tb`
 *
 * @note In the context of this function, "duration" is in term of samples, not
 *       seconds.
 */
func Av_rescale_delta(in_tb AVRational, in_ts int64,  fs_tb AVRational, duration int32, last *int64, out_tb AVRational) int64 {
    return int64(C.av_rescale_delta(*(*C.struct_AVRational)(unsafe.Pointer(&in_tb)), 
        C.longlong(in_ts), *(*C.struct_AVRational)(unsafe.Pointer(&fs_tb)), 
        C.int(duration), (*C.longlong)(unsafe.Pointer(last)), 
        *(*C.struct_AVRational)(unsafe.Pointer(&out_tb))))
}

/**
 * Add a value to a timestamp.
 *
 * This function guarantees that when the same value is repeatly added that
 * no accumulation of rounding errors occurs.
 *
 * @param[in] ts     Input timestamp
 * @param[in] ts_tb  Input timestamp time base
 * @param[in] inc    Value to be added
 * @param[in] inc_tb Time base of `inc`
 */
func Av_add_stable(ts_tb AVRational, ts int64, inc_tb AVRational, inc int64) int64 {
    return int64(C.av_add_stable(*(*C.struct_AVRational)(unsafe.Pointer(&ts_tb)), 
        C.longlong(ts), *(*C.struct_AVRational)(unsafe.Pointer(&inc_tb)), 
        C.longlong(inc)))
}


/**
 * @}
 */

                                 
