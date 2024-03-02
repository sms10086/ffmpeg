// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
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

//#cgo pkg-config: libavcodec
//#include "libavcodec/avfft.h"
import "C"
import (
    "unsafe"
)



                       
                       

/**
 * @file
 * @ingroup lavc_fft
 * FFT functions
 */

/**
 * @defgroup lavc_fft FFT functions
 * @ingroup lavc_misc
 *
 * @{
 */

type FFTSample float32

type FFTComplex C.struct_FFTComplex

type FFTContext C.struct_FFTContext

/**
 * Set up a complex FFT.
 * @param nbits           log2 of the length of the input array
 * @param inverse         if 0 perform the forward transform, if 1 perform the inverse
 */
func Av_fft_init(nbits int32, inverse int32) *FFTContext {
    return (*FFTContext)(unsafe.Pointer(C.av_fft_init(C.int(nbits), C.int(inverse))))
}

/**
 * Do the permutation needed BEFORE calling ff_fft_calc().
 */
func Av_fft_permute(s *FFTContext, z *FFTComplex)  {
    C.av_fft_permute((*C.FFTContext)(unsafe.Pointer(s)), 
        (*C.FFTComplex)(unsafe.Pointer(z)))
}

/**
 * Do a complex FFT with the parameters defined in av_fft_init(). The
 * input data must be permuted before. No 1.0/sqrt(n) normalization is done.
 */
func Av_fft_calc(s *FFTContext, z *FFTComplex)  {
    C.av_fft_calc((*C.FFTContext)(unsafe.Pointer(s)), 
        (*C.FFTComplex)(unsafe.Pointer(z)))
}

func Av_fft_end(s *FFTContext)  {
    C.av_fft_end((*C.FFTContext)(unsafe.Pointer(s)))
}

func Av_mdct_init(nbits int32, inverse int32, scale float64) *FFTContext {
    return (*FFTContext)(unsafe.Pointer(C.av_mdct_init(C.int(nbits), C.int(inverse), 
        C.double(scale))))
}
func Av_imdct_calc(s *FFTContext, output *FFTSample, input *FFTSample)  {
    C.av_imdct_calc((*C.FFTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(output)), 
        (*C.FFTSample)(unsafe.Pointer(input)))
}
func Av_imdct_half(s *FFTContext, output *FFTSample, input *FFTSample)  {
    C.av_imdct_half((*C.FFTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(output)), 
        (*C.FFTSample)(unsafe.Pointer(input)))
}
func Av_mdct_calc(s *FFTContext, output *FFTSample, input *FFTSample)  {
    C.av_mdct_calc((*C.FFTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(output)), 
        (*C.FFTSample)(unsafe.Pointer(input)))
}
func Av_mdct_end(s *FFTContext)  {
    C.av_mdct_end((*C.FFTContext)(unsafe.Pointer(s)))
}

/* Real Discrete Fourier Transform */

type RDFTransformType C.enum_RDFTransformType

type RDFTContext C.struct_RDFTContext

/**
 * Set up a real FFT.
 * @param nbits           log2 of the length of the input array
 * @param trans           the type of transform
 */
func Av_rdft_init(nbits int32, trans RDFTransformType) *RDFTContext {
    return (*RDFTContext)(unsafe.Pointer(C.av_rdft_init(C.int(nbits), 
        C.enum_RDFTransformType(trans))))
}
func Av_rdft_calc(s *RDFTContext, data *FFTSample)  {
    C.av_rdft_calc((*C.RDFTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(data)))
}
func Av_rdft_end(s *RDFTContext)  {
    C.av_rdft_end((*C.RDFTContext)(unsafe.Pointer(s)))
}

/* Discrete Cosine Transform */

type DCTContext C.struct_DCTContext

type DCTTransformType C.enum_DCTTransformType

/**
 * Set up DCT.
 *
 * @param nbits           size of the input array:
 *                        (1 << nbits)     for DCT-II, DCT-III and DST-I
 *                        (1 << nbits) + 1 for DCT-I
 * @param type            the type of transform
 *
 * @note the first element of the input of DST-I is ignored
 */
func Av_dct_init(nbits int32, typex DCTTransformType) *DCTContext {
    return (*DCTContext)(unsafe.Pointer(C.av_dct_init(C.int(nbits), 
        C.enum_DCTTransformType(typex))))
}
func Av_dct_calc(s *DCTContext, data *FFTSample)  {
    C.av_dct_calc((*C.DCTContext)(unsafe.Pointer(s)), 
        (*C.FFTSample)(unsafe.Pointer(data)))
}
func Av_dct_end (s *DCTContext)  {
    C.av_dct_end((*C.DCTContext)(unsafe.Pointer(s)))
}

/**
 * @}
 */

                            
