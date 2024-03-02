// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2001 Fabrice Bellard
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

//#cgo pkg-config: libavcodec libavutil
//#include <errno.h>
//#include "libavutil/samplefmt.h"
//#include "libavutil/attributes.h"
//#include "libavutil/avutil.h"
//#include "libavutil/buffer.h"
//#include "libavutil/cpu.h"
//#include "libavutil/channel_layout.h"
//#include "libavutil/dict.h"
//#include "libavutil/frame.h"
//#include "libavutil/hwcontext.h"
//#include "libavutil/log.h"
//#include "libavutil/pixfmt.h"
//#include "libavutil/rational.h"
//#include "libavcodec/version.h"
//#include "libavcodec/avcodec.h"
import "C"
import (
    "syscall"
    "unsafe"
)

//const AV_CODEC_ID_IFF_BYTERUN1 = AV_CODEC_ID_IFF_ILBM
//const AV_CODEC_ID_H265 = AV_CODEC_ID_HEVC
const AV_CODEC_PROP_INTRA_ONLY = (1 << 0)
const AV_CODEC_PROP_LOSSY = (1 << 1)
const AV_CODEC_PROP_LOSSLESS = (1 << 2)
const AV_CODEC_PROP_REORDER = (1 << 3)
const AV_CODEC_PROP_BITMAP_SUB = (1 << 16)
const AV_CODEC_PROP_TEXT_SUB = (1 << 17)
const AV_INPUT_BUFFER_PADDING_SIZE = 64
const AV_INPUT_BUFFER_MIN_SIZE = 16384
const AV_CODEC_FLAG_UNALIGNED = (1 <<  0)
const AV_CODEC_FLAG_QSCALE = (1 <<  1)
const AV_CODEC_FLAG_4MV = (1 <<  2)
const AV_CODEC_FLAG_OUTPUT_CORRUPT = (1 <<  3)
const AV_CODEC_FLAG_QPEL = (1 <<  4)
const AV_CODEC_FLAG_PASS1 = (1 <<  9)
const AV_CODEC_FLAG_PASS2 = (1 << 10)
const AV_CODEC_FLAG_LOOP_FILTER = (1 << 11)
const AV_CODEC_FLAG_GRAY = (1 << 13)
const AV_CODEC_FLAG_PSNR = (1 << 15)
const AV_CODEC_FLAG_TRUNCATED = (1 << 16)
const AV_CODEC_FLAG_INTERLACED_DCT = (1 << 18)
const AV_CODEC_FLAG_LOW_DELAY = (1 << 19)
const AV_CODEC_FLAG_GLOBAL_HEADER = (1 << 22)
const AV_CODEC_FLAG_BITEXACT = (1 << 23)
const AV_CODEC_FLAG_AC_PRED = (1 << 24)
const AV_CODEC_FLAG_INTERLACED_ME = (1 << 29)
const AV_CODEC_FLAG_CLOSED_GOP = (uint32(1) << 31)
const AV_CODEC_FLAG2_FAST = (1 <<  0)
const AV_CODEC_FLAG2_NO_OUTPUT = (1 <<  2)
const AV_CODEC_FLAG2_LOCAL_HEADER = (1 <<  3)
const AV_CODEC_FLAG2_DROP_FRAME_TIMECODE = (1 << 13)
const AV_CODEC_FLAG2_CHUNKS = (1 << 15)
const AV_CODEC_FLAG2_IGNORE_CROP = (1 << 16)
const AV_CODEC_FLAG2_SHOW_ALL = (1 << 22)
const AV_CODEC_FLAG2_EXPORT_MVS = (1 << 28)
const AV_CODEC_FLAG2_SKIP_MANUAL = (1 << 29)
const AV_CODEC_FLAG2_RO_FLUSH_NOOP = (1 << 30)
const AV_CODEC_CAP_DRAW_HORIZ_BAND = (1 <<  0)
const AV_CODEC_CAP_DR1 = (1 <<  1)
const AV_CODEC_CAP_TRUNCATED = (1 <<  3)
const AV_CODEC_CAP_DELAY = (1 <<  5)
const AV_CODEC_CAP_SMALL_LAST_FRAME = (1 <<  6)
const AV_CODEC_CAP_SUBFRAMES = (1 <<  8)
const AV_CODEC_CAP_EXPERIMENTAL = (1 <<  9)
const AV_CODEC_CAP_CHANNEL_CONF = (1 << 10)
const AV_CODEC_CAP_FRAME_THREADS = (1 << 12)
const AV_CODEC_CAP_SLICE_THREADS = (1 << 13)
const AV_CODEC_CAP_PARAM_CHANGE = (1 << 14)
const AV_CODEC_CAP_AUTO_THREADS = (1 << 15)
const AV_CODEC_CAP_VARIABLE_FRAME_SIZE = (1 << 16)
const AV_CODEC_CAP_AVOID_PROBING = (1 << 17)
const AV_CODEC_CAP_INTRA_ONLY = 0x40000000
const AV_CODEC_CAP_LOSSLESS = 0x80000000
const AV_CODEC_CAP_HARDWARE = (1 << 18)
const AV_CODEC_CAP_HYBRID = (1 << 19)
const AV_GET_BUFFER_FLAG_REF = (1 << 0)
//const AV_PKT_DATA_QUALITY_FACTOR = AV_PKT_DATA_QUALITY_STATS
const AV_PKT_FLAG_KEY = 0x0001
const AV_PKT_FLAG_CORRUPT = 0x0002
const AV_PKT_FLAG_DISCARD = 0x0004
const AV_PKT_FLAG_TRUSTED = 0x0008
const AV_PKT_FLAG_DISPOSABLE = 0x0010
const FF_COMPRESSION_DEFAULT = -1
const FF_CMP_SAD = 0
const FF_CMP_SSE = 1
const FF_CMP_SATD = 2
const FF_CMP_DCT = 3
const FF_CMP_PSNR = 4
const FF_CMP_BIT = 5
const FF_CMP_RD = 6
const FF_CMP_ZERO = 7
const FF_CMP_VSAD = 8
const FF_CMP_VSSE = 9
const FF_CMP_NSSE = 10
const FF_CMP_W53 = 11
const FF_CMP_W97 = 12
const FF_CMP_DCTMAX = 13
const FF_CMP_DCT264 = 14
const FF_CMP_MEDIAN_SAD = 15
const FF_CMP_CHROMA = 256
const SLICE_FLAG_CODED_ORDER = 0x0001
const SLICE_FLAG_ALLOW_FIELD = 0x0002
const SLICE_FLAG_ALLOW_PLANE = 0x0004
const FF_MB_DECISION_SIMPLE = 0
const FF_MB_DECISION_BITS = 1
const FF_MB_DECISION_RD = 2
const FF_BUG_AUTODETECT = 1
const FF_BUG_XVID_ILACE = 4
const FF_BUG_UMP4 = 8
const FF_BUG_NO_PADDING = 16
const FF_BUG_AMV = 32
const FF_BUG_QPEL_CHROMA = 64
const FF_BUG_STD_QPEL = 128
const FF_BUG_QPEL_CHROMA2 = 256
const FF_BUG_DIRECT_BLOCKSIZE = 512
const FF_BUG_EDGE = 1024
const FF_BUG_HPEL_CHROMA = 2048
const FF_BUG_DC_CLIP = 4096
const FF_BUG_MS = 8192
const FF_BUG_TRUNCATED = 16384
const FF_BUG_IEDGE = 32768
const FF_COMPLIANCE_VERY_STRICT = 2
const FF_COMPLIANCE_STRICT = 1
const FF_COMPLIANCE_NORMAL = 0
const FF_COMPLIANCE_UNOFFICIAL = -1
const FF_COMPLIANCE_EXPERIMENTAL = -2
const FF_EC_GUESS_MVS = 1
const FF_EC_DEBLOCK = 2
const FF_EC_FAVOR_INTER = 256
const FF_DEBUG_PICT_INFO = 1
const FF_DEBUG_RC = 2
const FF_DEBUG_BITSTREAM = 4
const FF_DEBUG_MB_TYPE = 8
const FF_DEBUG_QP = 16
const FF_DEBUG_DCT_COEFF = 0x00000040
const FF_DEBUG_SKIP = 0x00000080
const FF_DEBUG_STARTCODE = 0x00000100
const FF_DEBUG_ER = 0x00000400
const FF_DEBUG_MMCO = 0x00000800
const FF_DEBUG_BUGS = 0x00001000
const FF_DEBUG_BUFFERS = 0x00008000
const FF_DEBUG_THREADS = 0x00010000
const FF_DEBUG_GREEN_MD = 0x00800000
const FF_DEBUG_NOMC = 0x01000000
const AV_EF_CRCCHECK = (1<<0)
const AV_EF_BITSTREAM = (1<<1)
const AV_EF_BUFFER = (1<<2)
const AV_EF_EXPLODE = (1<<3)
const AV_EF_IGNORE_ERR = (1<<15)
const AV_EF_CAREFUL = (1<<16)
const AV_EF_COMPLIANT = (1<<17)
const AV_EF_AGGRESSIVE = (1<<18)
const FF_DCT_AUTO = 0
const FF_DCT_FASTINT = 1
const FF_DCT_INT = 2
const FF_DCT_MMX = 3
const FF_DCT_ALTIVEC = 5
const FF_DCT_FAAN = 6
const FF_IDCT_AUTO = 0
const FF_IDCT_INT = 1
const FF_IDCT_SIMPLE = 2
const FF_IDCT_SIMPLEMMX = 3
const FF_IDCT_ARM = 7
const FF_IDCT_ALTIVEC = 8
const FF_IDCT_SIMPLEARM = 10
const FF_IDCT_XVID = 14
const FF_IDCT_SIMPLEARMV5TE = 16
const FF_IDCT_SIMPLEARMV6 = 17
const FF_IDCT_FAAN = 20
const FF_IDCT_SIMPLENEON = 22
const FF_IDCT_NONE = 24
const FF_IDCT_SIMPLEAUTO = 128
const FF_THREAD_FRAME = 1
const FF_THREAD_SLICE = 2
const FF_PROFILE_UNKNOWN = -99
const FF_PROFILE_RESERVED = -100
const FF_PROFILE_AAC_MAIN = 0
const FF_PROFILE_AAC_LOW = 1
const FF_PROFILE_AAC_SSR = 2
const FF_PROFILE_AAC_LTP = 3
const FF_PROFILE_AAC_HE = 4
const FF_PROFILE_AAC_HE_V2 = 28
const FF_PROFILE_AAC_LD = 22
const FF_PROFILE_AAC_ELD = 38
const FF_PROFILE_MPEG2_AAC_LOW = 128
const FF_PROFILE_MPEG2_AAC_HE = 131
const FF_PROFILE_DNXHD = 0
const FF_PROFILE_DNXHR_LB = 1
const FF_PROFILE_DNXHR_SQ = 2
const FF_PROFILE_DNXHR_HQ = 3
const FF_PROFILE_DNXHR_HQX = 4
const FF_PROFILE_DNXHR_444 = 5
const FF_PROFILE_DTS = 20
const FF_PROFILE_DTS_ES = 30
const FF_PROFILE_DTS_96_24 = 40
const FF_PROFILE_DTS_HD_HRA = 50
const FF_PROFILE_DTS_HD_MA = 60
const FF_PROFILE_DTS_EXPRESS = 70
const FF_PROFILE_MPEG2_422 = 0
const FF_PROFILE_MPEG2_HIGH = 1
const FF_PROFILE_MPEG2_SS = 2
const FF_PROFILE_MPEG2_SNR_SCALABLE = 3
const FF_PROFILE_MPEG2_MAIN = 4
const FF_PROFILE_MPEG2_SIMPLE = 5
const FF_PROFILE_H264_CONSTRAINED = (1<<9)
const FF_PROFILE_H264_INTRA = (1<<11)
const FF_PROFILE_H264_BASELINE = 66
const FF_PROFILE_H264_CONSTRAINED_BASELINE = (66|FF_PROFILE_H264_CONSTRAINED)
const FF_PROFILE_H264_MAIN = 77
const FF_PROFILE_H264_EXTENDED = 88
const FF_PROFILE_H264_HIGH = 100
const FF_PROFILE_H264_HIGH_10 = 110
const FF_PROFILE_H264_HIGH_10_INTRA = (110|FF_PROFILE_H264_INTRA)
const FF_PROFILE_H264_MULTIVIEW_HIGH = 118
const FF_PROFILE_H264_HIGH_422 = 122
const FF_PROFILE_H264_HIGH_422_INTRA = (122|FF_PROFILE_H264_INTRA)
const FF_PROFILE_H264_STEREO_HIGH = 128
const FF_PROFILE_H264_HIGH_444 = 144
const FF_PROFILE_H264_HIGH_444_PREDICTIVE = 244
const FF_PROFILE_H264_HIGH_444_INTRA = (244|FF_PROFILE_H264_INTRA)
const FF_PROFILE_H264_CAVLC_444 = 44
const FF_PROFILE_VC1_SIMPLE = 0
const FF_PROFILE_VC1_MAIN = 1
const FF_PROFILE_VC1_COMPLEX = 2
const FF_PROFILE_VC1_ADVANCED = 3
const FF_PROFILE_MPEG4_SIMPLE = 0
const FF_PROFILE_MPEG4_SIMPLE_SCALABLE = 1
const FF_PROFILE_MPEG4_CORE = 2
const FF_PROFILE_MPEG4_MAIN = 3
const FF_PROFILE_MPEG4_N_BIT = 4
const FF_PROFILE_MPEG4_SCALABLE_TEXTURE = 5
const FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION = 6
const FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE = 7
const FF_PROFILE_MPEG4_HYBRID = 8
const FF_PROFILE_MPEG4_ADVANCED_REAL_TIME = 9
const FF_PROFILE_MPEG4_CORE_SCALABLE = 10
const FF_PROFILE_MPEG4_ADVANCED_CODING = 11
const FF_PROFILE_MPEG4_ADVANCED_CORE = 12
const FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE = 13
const FF_PROFILE_MPEG4_SIMPLE_STUDIO = 14
const FF_PROFILE_MPEG4_ADVANCED_SIMPLE = 15
const FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0 = 1
const FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1 = 2
const FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION = 32768
const FF_PROFILE_JPEG2000_DCINEMA_2K = 3
const FF_PROFILE_JPEG2000_DCINEMA_4K = 4
const FF_PROFILE_VP9_0 = 0
const FF_PROFILE_VP9_1 = 1
const FF_PROFILE_VP9_2 = 2
const FF_PROFILE_VP9_3 = 3
const FF_PROFILE_HEVC_MAIN = 1
const FF_PROFILE_HEVC_MAIN_10 = 2
const FF_PROFILE_HEVC_MAIN_STILL_PICTURE = 3
const FF_PROFILE_HEVC_REXT = 4
const FF_PROFILE_AV1_MAIN = 0
const FF_PROFILE_AV1_HIGH = 1
const FF_PROFILE_AV1_PROFESSIONAL = 2
const FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT = 0xc0
const FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT = 0xc1
const FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT = 0xc2
const FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS = 0xc3
const FF_PROFILE_MJPEG_JPEG_LS = 0xf7
const FF_PROFILE_SBC_MSBC = 1
const FF_LEVEL_UNKNOWN = -99
const FF_SUB_CHARENC_MODE_DO_NOTHING = -1
const FF_SUB_CHARENC_MODE_AUTOMATIC = 0
const FF_SUB_CHARENC_MODE_PRE_DECODER = 1
const FF_SUB_CHARENC_MODE_IGNORE = 2
const FF_DEBUG_VIS_MV_P_FOR = 0x00000001
const FF_DEBUG_VIS_MV_B_FOR = 0x00000002
const FF_DEBUG_VIS_MV_B_BACK = 0x00000004
const FF_CODEC_PROPERTY_LOSSLESS = 0x00000001
const FF_CODEC_PROPERTY_CLOSED_CAPTIONS = 0x00000002
const FF_SUB_TEXT_FMT_ASS = 0
const AV_HWACCEL_CODEC_CAP_EXPERIMENTAL = 0x0200
const AV_HWACCEL_FLAG_IGNORE_LEVEL = (1 << 0)
const AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH = (1 << 1)
const AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH = (1 << 2)
const AV_SUBTITLE_FLAG_FORCED = 0x00000001
const AV_PARSER_PTS_NB = 4
const PARSER_FLAG_COMPLETE_FRAMES = 0x0001
const PARSER_FLAG_ONCE = 0x0002
const PARSER_FLAG_FETCHED_OFFSET = 0x0004
const PARSER_FLAG_USE_CODEC_TS = 0x1000


                         
                         

/**
 * @file
 * @ingroup libavc
 * Libavcodec external API header
 */

                  
                                
                                 
                             
                             
                          
                                     
                           
                            
                                
                          
                             
                               

                    

/**
 * @defgroup libavc libavcodec
 * Encoding/Decoding Library
 *
 * @{
 *
 * @defgroup lavc_decoding Decoding
 * @{
 * @}
 *
 * @defgroup lavc_encoding Encoding
 * @{
 * @}
 *
 * @defgroup lavc_codec Codecs
 * @{
 * @defgroup lavc_codec_native Native Codecs
 * @{
 * @}
 * @defgroup lavc_codec_wrappers External library wrappers
 * @{
 * @}
 * @defgroup lavc_codec_hwaccel Hardware Accelerators bridge
 * @{
 * @}
 * @}
 * @defgroup lavc_internal Internal
 * @{
 * @}
 * @}
 */

/**
 * @ingroup libavc
 * @defgroup lavc_encdec send/receive encoding and decoding API overview
 * @{
 *
 * The avcodec_send_packet()/avcodec_receive_frame()/avcodec_send_frame()/
 * avcodec_receive_packet() functions provide an encode/decode API, which
 * decouples input and output.
 *
 * The API is very similar for encoding/decoding and audio/video, and works as
 * follows:
 * - Set up and open the AVCodecContext as usual.
 * - Send valid input:
 *   - For decoding, call avcodec_send_packet() to give the decoder raw
 *     compressed data in an AVPacket.
 *   - For encoding, call avcodec_send_frame() to give the encoder an AVFrame
 *     containing uncompressed audio or video.
 *   In both cases, it is recommended that AVPackets and AVFrames are
 *   refcounted, or libavcodec might have to copy the input data. (libavformat
 *   always returns refcounted AVPackets, and av_frame_get_buffer() allocates
 *   refcounted AVFrames.)
 * - Receive output in a loop. Periodically call one of the avcodec_receive_*()
 *   functions and process their output:
 *   - For decoding, call avcodec_receive_frame(). On success, it will return
 *     an AVFrame containing uncompressed audio or video data.
 *   - For encoding, call avcodec_receive_packet(). On success, it will return
 *     an AVPacket with a compressed frame.
 *   Repeat this call until it returns AVERROR(EAGAIN) or an error. The
 *   AVERROR(EAGAIN) return value means that new input data is required to
 *   return new output. In this case, continue with sending input. For each
 *   input frame/packet, the codec will typically return 1 output frame/packet,
 *   but it can also be 0 or more than 1.
 *
 * At the beginning of decoding or encoding, the codec might accept multiple
 * input frames/packets without returning a frame, until its internal buffers
 * are filled. This situation is handled transparently if you follow the steps
 * outlined above.
 *
 * In theory, sending input can result in EAGAIN - this should happen only if
 * not all output was received. You can use this to structure alternative decode
 * or encode loops other than the one suggested above. For example, you could
 * try sending new input on each iteration, and try to receive output if that
 * returns EAGAIN.
 *
 * End of stream situations. These require "flushing" (aka draining) the codec,
 * as the codec might buffer multiple frames or packets internally for
 * performance or out of necessity (consider B-frames).
 * This is handled as follows:
 * - Instead of valid input, send NULL to the avcodec_send_packet() (decoding)
 *   or avcodec_send_frame() (encoding) functions. This will enter draining
 *   mode.
 * - Call avcodec_receive_frame() (decoding) or avcodec_receive_packet()
 *   (encoding) in a loop until AVERROR_EOF is returned. The functions will
 *   not return AVERROR(EAGAIN), unless you forgot to enter draining mode.
 * - Before decoding can be resumed again, the codec has to be reset with
 *   avcodec_flush_buffers().
 *
 * Using the API as outlined above is highly recommended. But it is also
 * possible to call functions outside of this rigid schema. For example, you can
 * call avcodec_send_packet() repeatedly without calling
 * avcodec_receive_frame(). In this case, avcodec_send_packet() will succeed
 * until the codec's internal buffer has been filled up (which is typically of
 * size 1 per output frame, after initial input), and then reject input with
 * AVERROR(EAGAIN). Once it starts rejecting input, you have no choice but to
 * read at least some output.
 *
 * Not all codecs will follow a rigid and predictable dataflow; the only
 * guarantee is that an AVERROR(EAGAIN) return value on a send/receive call on
 * one end implies that a receive/send call on the other end will succeed, or
 * at least will not fail with AVERROR(EAGAIN). In general, no codec will
 * permit unlimited buffering of input or output.
 *
 * This API replaces the following legacy functions:
 * - avcodec_decode_video2() and avcodec_decode_audio4():
 *   Use avcodec_send_packet() to feed input to the decoder, then use
 *   avcodec_receive_frame() to receive decoded frames after each packet.
 *   Unlike with the old video decoding API, multiple frames might result from
 *   a packet. For audio, splitting the input packet into frames by partially
 *   decoding packets becomes transparent to the API user. You never need to
 *   feed an AVPacket to the API twice (unless it is rejected with AVERROR(EAGAIN) - then
 *   no data was read from the packet).
 *   Additionally, sending a flush/draining packet is required only once.
 * - avcodec_encode_video2()/avcodec_encode_audio2():
 *   Use avcodec_send_frame() to feed input to the encoder, then use
 *   avcodec_receive_packet() to receive encoded packets.
 *   Providing user-allocated buffers for avcodec_receive_packet() is not
 *   possible.
 * - The new API does not handle subtitles yet.
 *
 * Mixing new and old function calls on the same AVCodecContext is not allowed,
 * and will result in undefined behavior.
 *
 * Some codecs might require using the new API; using the old API will return
 * an error when calling it. All codecs support the new API.
 *
 * A codec is not allowed to return AVERROR(EAGAIN) for both sending and receiving. This
 * would be an invalid state, which could put the codec user into an endless
 * loop. The API has no concept of time either: it cannot happen that trying to
 * do avcodec_send_packet() results in AVERROR(EAGAIN), but a repeated call 1 second
 * later accepts the packet (with no other receive/flush API calls involved).
 * The API is a strict state machine, and the passage of time is not supposed
 * to influence it. Some timing-dependent behavior might still be deemed
 * acceptable in certain cases. But it must never result in both send/receive
 * returning EAGAIN at the same time at any point. It must also absolutely be
 * avoided that the current state is "unstable" and can "flip-flop" between
 * the send/receive APIs allowing progress. For example, it's not allowed that
 * the codec randomly decides that it actually wants to consume a packet now
 * instead of returning a frame, after it just returned AVERROR(EAGAIN) on an
 * avcodec_send_packet() call.
 * @}
 */

/**
 * @defgroup lavc_core Core functions/structures.
 * @ingroup libavc
 *
 * Basic definitions, functions for querying libavcodec capabilities,
 * allocating core structures, etc.
 * @{
 */


/**
 * Identify the syntax and semantics of the bitstream.
 * The principle is roughly:
 * Two decoders with the same ID can decode the same streams.
 * Two encoders with the same ID can encode compatible streams.
 * There may be slight deviations from the principle due to implementation
 * details.
 *
 * If you add a codec ID to this list, add it so that
 * 1. no value of an existing codec ID changes (that would break ABI),
 * 2. it is as close as possible to similar codecs
 *
 * After adding new codec IDs, do not forget to add an entry to the codec
 * descriptor list and bump libavcodec minor version.
 */
type AVCodecID C.enum_AVCodecID

/**
 * This struct describes the properties of a single codec described by an
 * AVCodecID.
 * @see avcodec_descriptor_get()
 */
type AVCodecDescriptor C.struct_AVCodecDescriptor

/**
 * Codec uses only intra compression.
 * Video and audio codecs only.
 */

/**
 * Codec supports lossy compression. Audio and video codecs only.
 * @note a codec may support both lossy and lossless
 * compression modes
 */

/**
 * Codec supports lossless compression. Audio and video codecs only.
 */

/**
 * Codec supports frame reordering. That is, the coded order (the order in which
 * the encoded packets are output by the encoders / stored / input to the
 * decoders) may be different from the presentation order of the corresponding
 * frames.
 *
 * For codecs that do not have this property set, PTS and DTS should always be
 * equal.
 */

/**
 * Subtitle codec is bitmap based
 * Decoded AVSubtitle data can be read from the AVSubtitleRect->pict field.
 */

/**
 * Subtitle codec is text based.
 * Decoded AVSubtitle data can be read from the AVSubtitleRect->ass field.
 */


/**
 * @ingroup lavc_decoding
 * Required number of additionally allocated bytes at the end of the input bitstream for decoding.
 * This is mainly needed because some optimized bitstream readers read
 * 32 or 64 bit at once and could read over the end.<br>
 * Note: If the first 23 bits of the additional bytes are not 0, then damaged
 * MPEG bitstreams could cause overread and segfault.
 */


/**
 * @ingroup lavc_encoding
 * minimum encoding buffer size
 * Used to avoid some checks during header writing.
 */


/**
 * @ingroup lavc_decoding
 */
type AVDiscard C.enum_AVDiscard

type AVAudioServiceType C.enum_AVAudioServiceType

/**
 * @ingroup lavc_encoding
 */
type RcOverride C.struct_RcOverride

/* encoding support
   These flags can be passed in AVCodecContext.flags before initialization.
   Note: Not everything is supported yet.
*/

/**
 * Allow decoders to produce frames with data planes that are not aligned
 * to CPU requirements (e.g. due to cropping).
 */

/**
 * Use fixed qscale.
 */

/**
 * 4 MV per MB allowed / advanced prediction for H.263.
 */

/**
 * Output even those frames that might be corrupted.
 */

/**
 * Use qpel MC.
 */

/**
 * Use internal 2pass ratecontrol in first pass mode.
 */

/**
 * Use internal 2pass ratecontrol in second pass mode.
 */

/**
 * loop filter.
 */

/**
 * Only decode/encode grayscale.
 */

/**
 * error[?] variables will be set during encoding.
 */

/**
 * Input bitstream might be truncated at a random location
 * instead of only at frame boundaries.
 */

/**
 * Use interlaced DCT.
 */

/**
 * Force low delay.
 */

/**
 * Place global headers in extradata instead of every keyframe.
 */

/**
 * Use only bitexact stuff (except (I)DCT).
 */

/* Fx : Flag for H.263+ extra options */
/**
 * H.263 advanced intra coding / MPEG-4 AC prediction
 */

/**
 * interlaced motion estimation
 */



/**
 * Allow non spec compliant speedup tricks.
 */

/**
 * Skip bitstream encoding.
 */

/**
 * Place global headers at every keyframe instead of in extradata.
 */


/**
 * timecode is in drop frame format. DEPRECATED!!!!
 */


/**
 * Input bitstream might be truncated at a packet boundaries
 * instead of only at frame boundaries.
 */

/**
 * Discard cropping information from SPS.
 */


/**
 * Show all frames before the first keyframe
 */

/**
 * Export motion vectors through frame side data
 */

/**
 * Do not skip samples and export skip information as frame side data
 */

/**
 * Do not reset ASS ReadOrder field on flush (subtitles decoding)
 */


/* Unsupported options :
 *              Syntax Arithmetic coding (SAC)
 *              Reference Picture Selection
 *              Independent Segment Decoding */
/* /Fx */
/* codec capabilities */

/**
 * Decoder can use draw_horiz_band callback.
 */

/**
 * Codec uses get_buffer() for allocating buffers and supports custom allocators.
 * If not set, it might not use get_buffer() at all or use operations that
 * assume the buffer was allocated by avcodec_default_get_buffer.
 */


/**
 * Encoder or decoder requires flushing with NULL input at the end in order to
 * give the complete and correct output.
 *
 * NOTE: If this flag is not set, the codec is guaranteed to never be fed with
 *       with NULL data. The user can still send NULL data to the public encode
 *       or decode function, but libavcodec will not pass it along to the codec
 *       unless this flag is set.
 *
 * Decoders:
 * The decoder has a non-zero delay and needs to be fed with avpkt->data=NULL,
 * avpkt->size=0 at the end to get the delayed data until the decoder no longer
 * returns frames.
 *
 * Encoders:
 * The encoder needs to be fed with NULL data at the end of encoding until the
 * encoder no longer returns data.
 *
 * NOTE: For encoders implementing the AVCodec.encode2() function, setting this
 *       flag also means that the encoder must set the pts and duration for
 *       each output packet. If this flag is not set, the pts and duration will
 *       be determined by libavcodec from the input frame.
 */

/**
 * Codec can be fed a final frame with a smaller size.
 * This can be used to prevent truncation of the last audio samples.
 */


/**
 * Codec can output multiple frames per AVPacket
 * Normally demuxers return one frame at a time, demuxers which do not do
 * are connected to a parser to split what they return into proper frames.
 * This flag is reserved to the very rare category of codecs which have a
 * bitstream that cannot be split into frames without timeconsuming
 * operations like full decoding. Demuxers carrying such bitstreams thus
 * may return multiple frames in a packet. This has many disadvantages like
 * prohibiting stream copy in many cases thus it should only be considered
 * as a last resort.
 */

/**
 * Codec is experimental and is thus avoided in favor of non experimental
 * encoders
 */

/**
 * Codec should fill in channel configuration and samplerate instead of container
 */

/**
 * Codec supports frame-level multithreading.
 */

/**
 * Codec supports slice-based (or partition-based) multithreading.
 */

/**
 * Codec supports changed parameters at any point.
 */

/**
 * Codec supports avctx->thread_count == 0 (auto).
 */

/**
 * Audio encoder supports receiving a different number of samples in each call.
 */

/**
 * Decoder is not a preferred choice for probing.
 * This indicates that the decoder is not a good choice for probing.
 * It could for example be an expensive to spin up hardware decoder,
 * or it could simply not provide a lot of useful information about
 * the stream.
 * A decoder marked with this flag should only be used as last resort
 * choice for probing.
 */

/**
 * Codec is intra only.
 */

/**
 * Codec is lossless.
 */


/**
 * Codec is backed by a hardware implementation. Typically used to
 * identify a non-hwaccel hardware decoder. For information about hwaccels, use
 * avcodec_get_hw_config() instead.
 */


/**
 * Codec is potentially backed by a hardware implementation, but not
 * necessarily. This is used instead of AV_CODEC_CAP_HARDWARE, if the
 * implementation provides some sort of internal fallback.
 */


/**
 * Pan Scan area.
 * This specifies the area which should be displayed.
 * Note there may be multiple such areas for one frame.
 */
type AVPanScan C.struct_AVPanScan

/**
 * This structure describes the bitrate properties of an encoded bitstream. It
 * roughly corresponds to a subset the VBV parameters for MPEG-2 or HRD
 * parameters for H.264/HEVC.
 */
type AVCPBProperties C.struct_AVCPBProperties

/**
 * The decoder will keep a reference to the frame and may reuse it later.
 */


/**
 * @defgroup lavc_packet AVPacket
 *
 * Types and functions for working with AVPacket.
 * @{
 */
type AVPacketSideDataType C.enum_AVPacketSideDataType

//DEPRECATED

type AVPacketSideData C.struct_AVPacketSideData

/**
 * This structure stores compressed data. It is typically exported by demuxers
 * and then passed as input to decoders, or received as output from encoders and
 * then passed to muxers.
 *
 * For video, it should typically contain one compressed frame. For audio it may
 * contain several compressed frames. Encoders are allowed to output empty
 * packets, with no compressed data, containing only side data
 * (e.g. to update some stream parameters at the end of encoding).
 *
 * AVPacket is one of the few structs in FFmpeg, whose size is a part of public
 * ABI. Thus it may be allocated on stack and no new fields can be added to it
 * without libavcodec and libavformat major bump.
 *
 * The semantics of data ownership depends on the buf field.
 * If it is set, the packet data is dynamically allocated and is
 * valid indefinitely until a call to av_packet_unref() reduces the
 * reference count to 0.
 *
 * If the buf field is not set av_packet_ref() would make a copy instead
 * of increasing the reference count.
 *
 * The side data is always allocated with av_malloc(), copied by
 * av_packet_ref() and freed by av_packet_unref().
 *
 * @see av_packet_ref
 * @see av_packet_unref
 */
type AVPacket C.struct_AVPacket
///< The packet contains a keyframe
///< The packet content is corrupted
/**
 * Flag is used to discard packets which are required to maintain valid
 * decoder state but are not required for output and should be dropped
 * after decoding.
 **/

/**
 * The packet comes from a trusted source.
 *
 * Otherwise-unsafe constructs such as arbitrary pointers to data
 * outside the packet may be followed.
 */

/**
 * Flag is used to indicate packets that contain frames that can
 * be discarded by the decoder.  I.e. Non-reference frames.
 */



type AVSideDataParamChangeFlags C.enum_AVSideDataParamChangeFlags
/**
 * @}
 */

type AVCodecInternal C.struct_AVCodecInternal

type AVFieldOrder C.enum_AVFieldOrder

/**
 * main external API structure.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * You can use AVOptions (av_opt* / av_set/get*()) to access these fields from user
 * applications.
 * The name string for AVOptions options matches the associated command line
 * parameter name and can be found in libavcodec/options_table.h
 * The AVOption/command line parameter names differ in some cases from the C
 * structure field names for historic reasons or brevity.
 * sizeof(AVCodecContext) must not be used outside libav*.
 */
type AVCodecContext C.struct_AVCodecContext

                        
   
                                                                              
                                                     
   
                    
                                                                           
                    
                                                                                     

                    
                                                                                    
                    
                                                                                                             

                    
                                                                    

                 
                    
                                                      
                    
                                                         
      

                    
                                                            
                    
                                                               

                    
                                                                        
                    
                                                                            
      

/**
 * AVProfile.
 */
type AVProfile C.struct_AVProfile



type AVCodecHWConfig C.struct_AVCodecHWConfig

type AVCodecDefault C.struct_AVCodecDefault

type AVSubtitle C.struct_AVSubtitle

/**
 * AVCodec.
 */
type AVCodec C.struct_AVCodec

                        
                    
                                                  
      

type MpegEncContext C.struct_MpegEncContext

/**
 * Retrieve supported hardware configurations for a codec.
 *
 * Values of index from zero to some maximum return the indexed configuration
 * descriptor; all other values return NULL.  If the codec does not support
 * any hardware configurations then it will always return NULL.
 */
func Avcodec_get_hw_config(codec *AVCodec, index int32) *AVCodecHWConfig {
    return (*AVCodecHWConfig)(unsafe.Pointer(C.avcodec_get_hw_config(
        (*C.AVCodec)(unsafe.Pointer(codec)), C.int(index))))
}

/**
 * @defgroup lavc_hwaccel AVHWAccel
 *
 * @note  Nothing in this structure should be accessed by the user.  At some
 *        point in future it will not be externally visible at all.
 *
 * @{
 */
type AVHWAccel C.struct_AVHWAccel

/**
 * HWAccel is experimental and is thus avoided in favor of non experimental
 * codecs
 */


/**
 * Hardware acceleration should be used for decoding even if the codec level
 * used is unknown or higher than the maximum supported level reported by the
 * hardware driver.
 *
 * It's generally a good idea to pass this flag unless you have a specific
 * reason not to, as hardware tends to under-report supported levels.
 */


/**
 * Hardware acceleration can output YUV pixel formats with a different chroma
 * sampling than 4:2:0 and/or other than 8 bits per component.
 */


/**
 * Hardware acceleration should still be attempted for decoding when the
 * codec profile does not match the reported capabilities of the hardware.
 *
 * For example, this can be used to try to decode baseline profile H.264
 * streams in hardware - it will often succeed, because many streams marked
 * as baseline profile actually conform to constrained baseline profile.
 *
 * @warning If the stream is actually not supported then the behaviour is
 *          undefined, and may include returning entirely incorrect output
 *          while indicating success.
 */


/**
 * @}
 */

                    
   
                                   
  
                                       
     
   

   
                          
  
                                                                     
         
                                                        
   
                          
                        
                                                                                  
                        
                                                                         
            

   
     
   
      

type AVSubtitleType C.enum_AVSubtitleType



type AVSubtitleRect C.struct_AVSubtitleRect



/**
 * This struct describes the properties of an encoded stream.
 *
 * sizeof(AVCodecParameters) is not a part of the public ABI, this struct must
 * be allocated with avcodec_parameters_alloc() and freed with
 * avcodec_parameters_free().
 */
type AVCodecParameters C.struct_AVCodecParameters

/**
 * Iterate over all registered codecs.
 *
 * @param opaque a pointer where libavcodec will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered codec or NULL when the iteration is
 *         finished
 */
func Av_codec_iterate(opaque *unsafe.Pointer) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.av_codec_iterate(opaque)))
}

               
   
                                                    
                                                               
                                
   
                    
                                         
      

/**
 * Return the LIBAVCODEC_VERSION_INT constant.
 */
func Avcodec_version() uint32 {
    return uint32(C.avcodec_version())
}

/**
 * Return the libavcodec build-time configuration.
 */
func Avcodec_configuration() string {
    return C.GoString(C.avcodec_configuration())
}

/**
 * Return the libavcodec license.
 */
func Avcodec_license() string {
    return C.GoString(C.avcodec_license())
}

               
   
                                                      
  
                                                                         
                                         
  
                              
   
                    
                                      

   
                                                                               
                                                                              
                                                                          
             
  
                        
                                
                                    
   
                    
                                
      

/**
 * Allocate an AVCodecContext and set its fields to default values. The
 * resulting struct should be freed with avcodec_free_context().
 *
 * @param codec if non-NULL, allocate private data and initialize defaults
 *              for the given codec. It is illegal to then call avcodec_open2()
 *              with a different codec.
 *              If NULL, then the codec-specific defaults won't be initialized,
 *              which may result in suboptimal default settings (this is
 *              important mainly for encoders, e.g. libx264).
 *
 * @return An AVCodecContext filled with default values or NULL on failure.
 */
func Avcodec_alloc_context3(codec *AVCodec) *AVCodecContext {
    return (*AVCodecContext)(unsafe.Pointer(C.avcodec_alloc_context3(
        (*C.AVCodec)(unsafe.Pointer(codec)))))
}

/**
 * Free the codec context and everything associated with it and write NULL to
 * the provided pointer.
 */
func Avcodec_free_context(avctx **AVCodecContext)  {
    C.avcodec_free_context((**C.AVCodecContext)(unsafe.Pointer(avctx)))
}

                               
   
                                                                               
                                                                        
                              
   
                                                                           
      

/**
 * Get the AVClass for AVCodecContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func Avcodec_get_class() *AVClass {
    return (*AVClass)(unsafe.Pointer(C.avcodec_get_class()))
}

                       
   
                                                                  
                                                
  
                      
   
                                             

   
                                                                         
                                                
  
                      
   
                                                     

   
                                                                      
                                                                  
                                                                     
                                                                 
  
                                                               
                                                                        
                                  
                                                                          
  
                                                                               
                                                                                
                                                                     
                                                                      
             
   
                    
                                                                          
      

/**
 * Allocate a new AVCodecParameters and set its fields to default values
 * (unknown/invalid/0). The returned struct must be freed with
 * avcodec_parameters_free().
 */
func Avcodec_parameters_alloc() *AVCodecParameters {
    return (*AVCodecParameters)(unsafe.Pointer(C.avcodec_parameters_alloc()))
}

/**
 * Free an AVCodecParameters instance and everything associated with it and
 * write NULL to the supplied pointer.
 */
func Avcodec_parameters_free(par **AVCodecParameters)  {
    C.avcodec_parameters_free((**C.AVCodecParameters)(unsafe.Pointer(par)))
}

/**
 * Copy the contents of src to dst. Any allocated fields in dst are freed and
 * replaced with newly allocated duplicates of the corresponding fields in src.
 *
 * @return >= 0 on success, a negative AVERROR code on failure.
 */
func Avcodec_parameters_copy(dst *AVCodecParameters, src *AVCodecParameters) int32 {
    return int32(C.avcodec_parameters_copy(
        (*C.AVCodecParameters)(unsafe.Pointer(dst)), 
        (*C.AVCodecParameters)(unsafe.Pointer(src))))
}

/**
 * Fill the parameters struct based on the values from the supplied codec
 * context. Any allocated fields in par are freed and replaced with duplicates
 * of the corresponding fields in codec.
 *
 * @return >= 0 on success, a negative AVERROR code on failure
 */
func Avcodec_parameters_from_context(par *AVCodecParameters,
                                    codec *AVCodecContext) int32 {
    return int32(C.avcodec_parameters_from_context(
        (*C.AVCodecParameters)(unsafe.Pointer(par)), 
        (*C.AVCodecContext)(unsafe.Pointer(codec))))
}

/**
 * Fill the codec context based on the values from the supplied codec
 * parameters. Any allocated fields in codec that have a corresponding field in
 * par are freed and replaced with duplicates of the corresponding field in par.
 * Fields in codec that do not have a counterpart in par are not touched.
 *
 * @return >= 0 on success, a negative AVERROR code on failure.
 */
func Avcodec_parameters_to_context(codec *AVCodecContext,
                                  par *AVCodecParameters) int32 {
    return int32(C.avcodec_parameters_to_context(
        (*C.AVCodecContext)(unsafe.Pointer(codec)), 
        (*C.AVCodecParameters)(unsafe.Pointer(par))))
}

/**
 * Initialize the AVCodecContext to use the given AVCodec. Prior to using this
 * function the context has to be allocated with avcodec_alloc_context3().
 *
 * The functions avcodec_find_decoder_by_name(), avcodec_find_encoder_by_name(),
 * avcodec_find_decoder() and avcodec_find_encoder() provide an easy way for
 * retrieving a codec.
 *
 * @warning This function is not thread safe!
 *
 * @note Always call this function before using decoding routines (such as
 * @ref avcodec_receive_frame()).
 *
 * @code
 * avcodec_register_all();
 * av_dict_set(&opts, "b", "2.5M", 0);
 * codec = avcodec_find_decoder(AV_CODEC_ID_H264);
 * if (!codec)
 *     exit(1);
 *
 * context = avcodec_alloc_context3(codec);
 *
 * if (avcodec_open2(context, codec, opts) < 0)
 *     exit(1);
 * @endcode
 *
 * @param avctx The context to initialize.
 * @param codec The codec to open this context for. If a non-NULL codec has been
 *              previously passed to avcodec_alloc_context3() or
 *              for this context, then this parameter MUST be either NULL or
 *              equal to the previously passed codec.
 * @param options A dictionary filled with AVCodecContext and codec-private options.
 *                On return this object will be filled with options that were not found.
 *
 * @return zero on success, a negative value on error
 * @see avcodec_alloc_context3(), avcodec_find_decoder(), avcodec_find_encoder(),
 *      av_dict_set(), av_opt_find().
 */
func Avcodec_open2(avctx *AVCodecContext, codec *AVCodec, options **AVDictionary) int32 {
    return int32(C.avcodec_open2((*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVCodec)(unsafe.Pointer(codec)), 
        (**C.AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Close a given AVCodecContext and free all the data associated with it
 * (but not the AVCodecContext itself).
 *
 * Calling this function on an AVCodecContext that hasn't been opened will free
 * the codec-specific data allocated in avcodec_alloc_context3() with a non-NULL
 * codec. Subsequent calls will do nothing.
 *
 * @note Do not use this function. Use avcodec_free_context() to destroy a
 * codec context (either open or closed). Opening and closing a codec context
 * multiple times is not supported anymore -- use multiple codec contexts
 * instead.
 */
func Avcodec_close(avctx *AVCodecContext) int32 {
    return int32(C.avcodec_close((*C.AVCodecContext)(unsafe.Pointer(avctx))))
}

/**
 * Free all allocated data in the given subtitle struct.
 *
 * @param sub AVSubtitle to free.
 */
func Avsubtitle_free(sub *AVSubtitle)  {
    C.avsubtitle_free((*C.AVSubtitle)(unsafe.Pointer(sub)))
}

/**
 * @}
 */

/**
 * @addtogroup lavc_packet
 * @{
 */

/**
 * Allocate an AVPacket and set its fields to default values.  The resulting
 * struct must be freed using av_packet_free().
 *
 * @return An AVPacket filled with default values or NULL on failure.
 *
 * @note this only allocates the AVPacket itself, not the data buffers. Those
 * must be allocated through other means such as av_new_packet.
 *
 * @see av_new_packet
 */
func Av_packet_alloc() *AVPacket {
    return (*AVPacket)(unsafe.Pointer(C.av_packet_alloc()))
}

/**
 * Create a new packet that references the same data as src.
 *
 * This is a shortcut for av_packet_alloc()+av_packet_ref().
 *
 * @return newly created AVPacket on success, NULL on error.
 *
 * @see av_packet_alloc
 * @see av_packet_ref
 */
func Av_packet_clone(src *AVPacket) *AVPacket {
    return (*AVPacket)(unsafe.Pointer(C.av_packet_clone((*C.AVPacket)(unsafe.Pointer(src)))))
}

/**
 * Free the packet, if the packet is reference counted, it will be
 * unreferenced first.
 *
 * @param pkt packet to be freed. The pointer will be set to NULL.
 * @note passing NULL is a no-op.
 */
func Av_packet_free(pkt **AVPacket)  {
    C.av_packet_free((**C.AVPacket)(unsafe.Pointer(pkt)))
}

/**
 * Initialize optional fields of a packet with default values.
 *
 * Note, this does not touch the data and size members, which have to be
 * initialized separately.
 *
 * @param pkt packet
 */
func Av_init_packet(pkt *AVPacket)  {
    C.av_init_packet((*C.AVPacket)(unsafe.Pointer(pkt)))
}

/**
 * Allocate the payload of a packet and initialize its fields with
 * default values.
 *
 * @param pkt packet
 * @param size wanted payload size
 * @return 0 if OK, AVERROR_xxx otherwise
 */
func Av_new_packet(pkt *AVPacket, size int32) int32 {
    return int32(C.av_new_packet((*C.AVPacket)(unsafe.Pointer(pkt)), C.int(size)))
}

/**
 * Reduce packet size, correctly zeroing padding
 *
 * @param pkt packet
 * @param size new size
 */
func Av_shrink_packet(pkt *AVPacket, size int32)  {
    C.av_shrink_packet((*C.AVPacket)(unsafe.Pointer(pkt)), C.int(size))
}

/**
 * Increase packet size, correctly zeroing padding
 *
 * @param pkt packet
 * @param grow_by number of bytes by which to increase the size of the packet
 */
func Av_grow_packet(pkt *AVPacket, grow_by int32) int32 {
    return int32(C.av_grow_packet((*C.AVPacket)(unsafe.Pointer(pkt)), 
        C.int(grow_by)))
}

/**
 * Initialize a reference-counted packet from av_malloc()ed data.
 *
 * @param pkt packet to be initialized. This function will set the data, size,
 *        buf and destruct fields, all others are left untouched.
 * @param data Data allocated by av_malloc() to be used as packet data. If this
 *        function returns successfully, the data is owned by the underlying AVBuffer.
 *        The caller may not access the data through other means.
 * @param size size of data in bytes, without the padding. I.e. the full buffer
 *        size is assumed to be size + AV_INPUT_BUFFER_PADDING_SIZE.
 *
 * @return 0 on success, a negative AVERROR on error
 */
func Av_packet_from_data(pkt *AVPacket, data *uint8, size int32) int32 {
    return int32(C.av_packet_from_data((*C.AVPacket)(unsafe.Pointer(pkt)), 
        (*C.uchar)(unsafe.Pointer(data)), C.int(size)))
}

                           
   
                                                                              
                                                      
  
                                                             
   
                    
                                 
   
                                  
  
                                                 
  
                                
   
                    
                                                       

   
                        
  
                                                 
  
                                       
   
                    
                                                                 

   
                 
  
                                  
  
                            
   
                    
                                   
      
/**
 * Allocate new information of a packet.
 *
 * @param pkt packet
 * @param type side information type
 * @param size side information size
 * @return pointer to fresh allocated data or NULL otherwise
 */
func Av_packet_new_side_data(pkt *AVPacket, typex AVPacketSideDataType,
                                 size int32) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_packet_new_side_data((*C.AVPacket)(unsafe.Pointer(pkt)), 
        C.enum_AVPacketSideDataType(typex), C.int(size))))
}

/**
 * Wrap an existing array as a packet side data.
 *
 * @param pkt packet
 * @param type side information type
 * @param data the side data array. It must be allocated with the av_malloc()
 *             family of functions. The ownership of the data is transferred to
 *             pkt.
 * @param size side information size
 * @return a non-negative number on success, a negative AVERROR code on
 *         failure. On failure, the packet is unchanged and the data remains
 *         owned by the caller.
 */
func Av_packet_add_side_data(pkt *AVPacket, typex AVPacketSideDataType,
                            data *uint8, size uint64) int32 {
    return int32(C.av_packet_add_side_data((*C.AVPacket)(unsafe.Pointer(pkt)), 
        C.enum_AVPacketSideDataType(typex), (*C.uchar)(unsafe.Pointer(data)), 
        C.ulonglong(size)))
}

/**
 * Shrink the already allocated side data buffer
 *
 * @param pkt packet
 * @param type side information type
 * @param size new side information size
 * @return 0 on success, < 0 on failure
 */
func Av_packet_shrink_side_data(pkt *AVPacket, typex AVPacketSideDataType,
                               size int32) int32 {
    return int32(C.av_packet_shrink_side_data((*C.AVPacket)(unsafe.Pointer(pkt)), 
        C.enum_AVPacketSideDataType(typex), C.int(size)))
}

/**
 * Get side information from packet.
 *
 * @param pkt packet
 * @param type desired side information type
 * @param size pointer for side information size to store (optional)
 * @return pointer to data if present or NULL otherwise
 */
func Av_packet_get_side_data(pkt *AVPacket, typex AVPacketSideDataType,
                                 size *int32) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_packet_get_side_data((*C.AVPacket)(unsafe.Pointer(pkt)), 
        C.enum_AVPacketSideDataType(typex), (*C.int)(unsafe.Pointer(size)))))
}

                       
                    
                                             

                    
                                             
      

func Av_packet_side_data_name(typex AVPacketSideDataType) string {
    return C.GoString(C.av_packet_side_data_name(C.enum_AVPacketSideDataType(typex)))
}

/**
 * Pack a dictionary for use in side_data.
 *
 * @param dict The dictionary to pack.
 * @param size pointer to store the size of the returned data
 * @return pointer to data if successful, NULL otherwise
 */
func Av_packet_pack_dictionary(dict *AVDictionary, size *int32) *uint8 {
    return (*uint8)(unsafe.Pointer(C.av_packet_pack_dictionary(
        (*C.AVDictionary)(unsafe.Pointer(dict)), (*C.int)(unsafe.Pointer(size)))))
}
/**
 * Unpack a dictionary from side_data.
 *
 * @param data data from side_data
 * @param size size of the data
 * @param dict the metadata storage dictionary
 * @return 0 on success, < 0 on failure
 */
func Av_packet_unpack_dictionary(data *uint8, size int32, dict **AVDictionary) int32 {
    return int32(C.av_packet_unpack_dictionary((*C.uchar)(unsafe.Pointer(data)), 
        C.int(size), (**C.AVDictionary)(unsafe.Pointer(dict))))
}


/**
 * Convenience function to free all the side data stored.
 * All the other fields stay untouched.
 *
 * @param pkt packet
 */
func Av_packet_free_side_data(pkt *AVPacket)  {
    C.av_packet_free_side_data((*C.AVPacket)(unsafe.Pointer(pkt)))
}

/**
 * Setup a new reference to the data described by a given packet
 *
 * If src is reference-counted, setup dst as a new reference to the
 * buffer in src. Otherwise allocate a new buffer in dst and copy the
 * data from src into it.
 *
 * All the other fields are copied from src.
 *
 * @see av_packet_unref
 *
 * @param dst Destination packet
 * @param src Source packet
 *
 * @return 0 on success, a negative AVERROR on error.
 */
func Av_packet_ref(dst *AVPacket, src *AVPacket) int32 {
    return int32(C.av_packet_ref((*C.AVPacket)(unsafe.Pointer(dst)), 
        (*C.AVPacket)(unsafe.Pointer(src))))
}

/**
 * Wipe the packet.
 *
 * Unreference the buffer referenced by the packet and reset the
 * remaining packet fields to their default values.
 *
 * @param pkt The packet to be unreferenced.
 */
func Av_packet_unref(pkt *AVPacket)  {
    C.av_packet_unref((*C.AVPacket)(unsafe.Pointer(pkt)))
}

/**
 * Move every field in src to dst and reset src.
 *
 * @see av_packet_unref
 *
 * @param src Source packet, will be reset
 * @param dst Destination packet
 */
func Av_packet_move_ref(dst *AVPacket, src *AVPacket)  {
    C.av_packet_move_ref((*C.AVPacket)(unsafe.Pointer(dst)), 
        (*C.AVPacket)(unsafe.Pointer(src)))
}

/**
 * Copy only "properties" fields from src to dst.
 *
 * Properties for the purpose of this function are all the fields
 * beside those related to the packet data (buf, data, size)
 *
 * @param dst Destination packet
 * @param src Source packet
 *
 * @return 0 on success AVERROR on failure.
 */
func Av_packet_copy_props(dst *AVPacket, src *AVPacket) int32 {
    return int32(C.av_packet_copy_props((*C.AVPacket)(unsafe.Pointer(dst)), 
        (*C.AVPacket)(unsafe.Pointer(src))))
}

/**
 * Ensure the data described by a given packet is reference counted.
 *
 * @note This function does not ensure that the reference will be writable.
 *       Use av_packet_make_writable instead for that purpose.
 *
 * @see av_packet_ref
 * @see av_packet_make_writable
 *
 * @param pkt packet whose data should be made reference counted.
 *
 * @return 0 on success, a negative AVERROR on error. On failure, the
 *         packet is unchanged.
 */
func Av_packet_make_refcounted(pkt *AVPacket) int32 {
    return int32(C.av_packet_make_refcounted((*C.AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Create a writable reference for the data described by a given packet,
 * avoiding data copy if possible.
 *
 * @param pkt Packet whose data should be made writable.
 *
 * @return 0 on success, a negative AVERROR on failure. On failure, the
 *         packet is unchanged.
 */
func Av_packet_make_writable(pkt *AVPacket) int32 {
    return int32(C.av_packet_make_writable((*C.AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Convert valid timing fields (timestamps / durations) in a packet from one
 * timebase to another. Timestamps with unknown values (AV_NOPTS_VALUE) will be
 * ignored.
 *
 * @param pkt packet on which the conversion will be performed
 * @param tb_src source timebase, in which the timing fields in pkt are
 *               expressed
 * @param tb_dst destination timebase, to which the timing fields will be
 *               converted
 */
func Av_packet_rescale_ts(pkt *AVPacket, tb_src AVRational, tb_dst AVRational)  {
    C.av_packet_rescale_ts((*C.AVPacket)(unsafe.Pointer(pkt)), 
        C.AVRational(tb_src), C.AVRational(tb_dst))
}

/**
 * @}
 */

/**
 * @addtogroup lavc_decoding
 * @{
 */

/**
 * Find a registered decoder with a matching codec ID.
 *
 * @param id AVCodecID of the requested decoder
 * @return A decoder if one was found, NULL otherwise.
 */
func Avcodec_find_decoder(id AVCodecID) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.avcodec_find_decoder(C.enum_AVCodecID(id))))
}

/**
 * Find a registered decoder with the specified name.
 *
 * @param name name of the requested decoder
 * @return A decoder if one was found, NULL otherwise.
 */
func Avcodec_find_decoder_by_name(name *byte) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.avcodec_find_decoder_by_name(
        (*C.char)(unsafe.Pointer(name)))))
}

/**
 * The default callback for AVCodecContext.get_buffer2(). It is made public so
 * it can be called by custom get_buffer2() implementations for decoders without
 * AV_CODEC_CAP_DR1 set.
 */
func Avcodec_default_get_buffer2(s *AVCodecContext, frame *AVFrame, flags int32) int32 {
    return int32(C.avcodec_default_get_buffer2(
        (*C.AVCodecContext)(unsafe.Pointer(s)), 
        (*C.AVFrame)(unsafe.Pointer(frame)), C.int(flags)))
}

/**
 * Modify width and height values so that they will result in a memory
 * buffer that is acceptable for the codec if you do not use any horizontal
 * padding.
 *
 * May only be used if a codec with AV_CODEC_CAP_DR1 has been opened.
 */
func Avcodec_align_dimensions(s *AVCodecContext, width *int32, height *int32)  {
    C.avcodec_align_dimensions((*C.AVCodecContext)(unsafe.Pointer(s)), 
        (*C.int)(unsafe.Pointer(width)), (*C.int)(unsafe.Pointer(height)))
}

/**
 * Modify width and height values so that they will result in a memory
 * buffer that is acceptable for the codec if you also ensure that all
 * line sizes are a multiple of the respective linesize_align[i].
 *
 * May only be used if a codec with AV_CODEC_CAP_DR1 has been opened.
 */
func Avcodec_align_dimensions2(s *AVCodecContext, width *int32, height *int32,
                               linesize_align[AV_NUM_DATA_POINTERS] int32)  {
    C.avcodec_align_dimensions2((*C.AVCodecContext)(unsafe.Pointer(s)), 
        (*C.int)(unsafe.Pointer(width)), (*C.int)(unsafe.Pointer(height)), 
        (*C.int)(unsafe.Pointer(&linesize_align[0])))
}

/**
 * Converts AVChromaLocation to swscale x/y chroma position.
 *
 * The positions represent the chroma (0,0) position in a coordinates system
 * with luma (0,0) representing the origin and luma(1,1) representing 256,256
 *
 * @param xpos  horizontal chroma sample position
 * @param ypos  vertical   chroma sample position
 */
func Avcodec_enum_to_chroma_pos(xpos *int32, ypos *int32, pos AVChromaLocation) int32 {
    return int32(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(xpos)), 
        (*C.int)(unsafe.Pointer(ypos)), C.enum_AVChromaLocation(pos)))
}

/**
 * Converts swscale x/y chroma position to AVChromaLocation.
 *
 * The positions represent the chroma (0,0) position in a coordinates system
 * with luma (0,0) representing the origin and luma(1,1) representing 256,256
 *
 * @param xpos  horizontal chroma sample position
 * @param ypos  vertical   chroma sample position
 */
func Avcodec_chroma_pos_to_enum(xpos int32, ypos int32) AVChromaLocation {
    return AVChromaLocation(C.avcodec_chroma_pos_to_enum(C.int(xpos), 
        C.int(ypos)))
}

/**
 * Decode the audio frame of size avpkt->size from avpkt->data into frame.
 *
 * Some decoders may support multiple frames in a single AVPacket. Such
 * decoders would then just decode the first frame and the return value would be
 * less than the packet size. In this case, avcodec_decode_audio4 has to be
 * called again with an AVPacket containing the remaining data in order to
 * decode the second frame, etc...  Even if no frames are returned, the packet
 * needs to be fed to the decoder with remaining data until it is completely
 * consumed or an error occurs.
 *
 * Some decoders (those marked with AV_CODEC_CAP_DELAY) have a delay between input
 * and output. This means that for some packets they will not immediately
 * produce decoded output and need to be flushed at the end of decoding to get
 * all the decoded data. Flushing is done by calling this function with packets
 * with avpkt->data set to NULL and avpkt->size set to 0 until it stops
 * returning samples. It is safe to flush even those decoders that are not
 * marked with AV_CODEC_CAP_DELAY, then no samples will be returned.
 *
 * @warning The input buffer, avpkt->data must be AV_INPUT_BUFFER_PADDING_SIZE
 *          larger than the actual read bytes because some optimized bitstream
 *          readers read 32 or 64 bits at once and could read over the end.
 *
 * @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
 * before packets may be fed to the decoder.
 *
 * @param      avctx the codec context
 * @param[out] frame The AVFrame in which to store decoded audio samples.
 *                   The decoder will allocate a buffer for the decoded frame by
 *                   calling the AVCodecContext.get_buffer2() callback.
 *                   When AVCodecContext.refcounted_frames is set to 1, the frame is
 *                   reference counted and the returned reference belongs to the
 *                   caller. The caller must release the frame using av_frame_unref()
 *                   when the frame is no longer needed. The caller may safely write
 *                   to the frame if av_frame_is_writable() returns 1.
 *                   When AVCodecContext.refcounted_frames is set to 0, the returned
 *                   reference belongs to the decoder and is valid only until the
 *                   next call to this function or until closing or flushing the
 *                   decoder. The caller may not write to it.
 * @param[out] got_frame_ptr Zero if no frame could be decoded, otherwise it is
 *                           non-zero. Note that this field being set to zero
 *                           does not mean that an error has occurred. For
 *                           decoders with AV_CODEC_CAP_DELAY set, no given decode
 *                           call is guaranteed to produce a frame.
 * @param[in]  avpkt The input AVPacket containing the input buffer.
 *                   At least avpkt->data and avpkt->size should be set. Some
 *                   decoders might also require additional fields to be set.
 * @return A negative error code is returned if an error occurred during
 *         decoding, otherwise the number of bytes consumed from the input
 *         AVPacket is returned.
 *
* @deprecated Use avcodec_send_packet() and avcodec_receive_frame().
 */

func Avcodec_decode_audio4(avctx *AVCodecContext, frame *AVFrame,
                          got_frame_ptr *int32, avpkt *AVPacket) int32 {
    return int32(C.avcodec_decode_audio4(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVFrame)(unsafe.Pointer(frame)), 
        (*C.int)(unsafe.Pointer(got_frame_ptr)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt))))
}

/**
 * Decode the video frame of size avpkt->size from avpkt->data into picture.
 * Some decoders may support multiple frames in a single AVPacket, such
 * decoders would then just decode the first frame.
 *
 * @warning The input buffer must be AV_INPUT_BUFFER_PADDING_SIZE larger than
 * the actual read bytes because some optimized bitstream readers read 32 or 64
 * bits at once and could read over the end.
 *
 * @warning The end of the input buffer buf should be set to 0 to ensure that
 * no overreading happens for damaged MPEG streams.
 *
 * @note Codecs which have the AV_CODEC_CAP_DELAY capability set have a delay
 * between input and output, these need to be fed with avpkt->data=NULL,
 * avpkt->size=0 at the end to return the remaining frames.
 *
 * @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
 * before packets may be fed to the decoder.
 *
 * @param avctx the codec context
 * @param[out] picture The AVFrame in which the decoded video frame will be stored.
 *             Use av_frame_alloc() to get an AVFrame. The codec will
 *             allocate memory for the actual bitmap by calling the
 *             AVCodecContext.get_buffer2() callback.
 *             When AVCodecContext.refcounted_frames is set to 1, the frame is
 *             reference counted and the returned reference belongs to the
 *             caller. The caller must release the frame using av_frame_unref()
 *             when the frame is no longer needed. The caller may safely write
 *             to the frame if av_frame_is_writable() returns 1.
 *             When AVCodecContext.refcounted_frames is set to 0, the returned
 *             reference belongs to the decoder and is valid only until the
 *             next call to this function or until closing or flushing the
 *             decoder. The caller may not write to it.
 *
 * @param[in] avpkt The input AVPacket containing the input buffer.
 *            You can create such packet with av_init_packet() and by then setting
 *            data and size, some decoders might in addition need other fields like
 *            flags&AV_PKT_FLAG_KEY. All decoders are designed to use the least
 *            fields possible.
 * @param[in,out] got_picture_ptr Zero if no frame could be decompressed, otherwise, it is nonzero.
 * @return On error a negative value is returned, otherwise the number of bytes
 * used or zero if no frame could be decompressed.
 *
 * @deprecated Use avcodec_send_packet() and avcodec_receive_frame().
 */

func Avcodec_decode_video2(avctx *AVCodecContext, picture *AVFrame,
                         got_picture_ptr *int32,
                         avpkt *AVPacket) int32 {
    return int32(C.avcodec_decode_video2(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVFrame)(unsafe.Pointer(picture)), 
        (*C.int)(unsafe.Pointer(got_picture_ptr)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt))))
}

/**
 * Decode a subtitle message.
 * Return a negative value on error, otherwise return the number of bytes used.
 * If no subtitle could be decompressed, got_sub_ptr is zero.
 * Otherwise, the subtitle is stored in *sub.
 * Note that AV_CODEC_CAP_DR1 is not available for subtitle codecs. This is for
 * simplicity, because the performance difference is expect to be negligible
 * and reusing a get_buffer written for video codecs would probably perform badly
 * due to a potentially very different allocation pattern.
 *
 * Some decoders (those marked with AV_CODEC_CAP_DELAY) have a delay between input
 * and output. This means that for some packets they will not immediately
 * produce decoded output and need to be flushed at the end of decoding to get
 * all the decoded data. Flushing is done by calling this function with packets
 * with avpkt->data set to NULL and avpkt->size set to 0 until it stops
 * returning subtitles. It is safe to flush even those decoders that are not
 * marked with AV_CODEC_CAP_DELAY, then no subtitles will be returned.
 *
 * @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
 * before packets may be fed to the decoder.
 *
 * @param avctx the codec context
 * @param[out] sub The Preallocated AVSubtitle in which the decoded subtitle will be stored,
 *                 must be freed with avsubtitle_free if *got_sub_ptr is set.
 * @param[in,out] got_sub_ptr Zero if no subtitle could be decompressed, otherwise, it is nonzero.
 * @param[in] avpkt The input AVPacket containing the input buffer.
 */
func Avcodec_decode_subtitle2(avctx *AVCodecContext, sub *AVSubtitle,
                            got_sub_ptr *int32,
                            avpkt *AVPacket) int32 {
    return int32(C.avcodec_decode_subtitle2(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVSubtitle)(unsafe.Pointer(sub)), 
        (*C.int)(unsafe.Pointer(got_sub_ptr)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt))))
}

/**
 * Supply raw packet data as input to a decoder.
 *
 * Internally, this call will copy relevant AVCodecContext fields, which can
 * influence decoding per-packet, and apply them when the packet is actually
 * decoded. (For example AVCodecContext.skip_frame, which might direct the
 * decoder to drop the frame contained by the packet sent with this function.)
 *
 * @warning The input buffer, avpkt->data must be AV_INPUT_BUFFER_PADDING_SIZE
 *          larger than the actual read bytes because some optimized bitstream
 *          readers read 32 or 64 bits at once and could read over the end.
 *
 * @warning Do not mix this API with the legacy API (like avcodec_decode_video2())
 *          on the same AVCodecContext. It will return unexpected results now
 *          or in future libavcodec versions.
 *
 * @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
 *       before packets may be fed to the decoder.
 *
 * @param avctx codec context
 * @param[in] avpkt The input AVPacket. Usually, this will be a single video
 *                  frame, or several complete audio frames.
 *                  Ownership of the packet remains with the caller, and the
 *                  decoder will not write to the packet. The decoder may create
 *                  a reference to the packet data (or copy it if the packet is
 *                  not reference-counted).
 *                  Unlike with older APIs, the packet is always fully consumed,
 *                  and if it contains multiple frames (e.g. some audio codecs),
 *                  will require you to call avcodec_receive_frame() multiple
 *                  times afterwards before you can send a new packet.
 *                  It can be NULL (or an AVPacket with data set to NULL and
 *                  size set to 0); in this case, it is considered a flush
 *                  packet, which signals the end of the stream. Sending the
 *                  first flush packet will return success. Subsequent ones are
 *                  unnecessary and will return AVERROR_EOF. If the decoder
 *                  still has frames buffered, it will return them after sending
 *                  a flush packet.
 *
 * @return 0 on success, otherwise negative error code:
 *      AVERROR(EAGAIN):   input is not accepted in the current state - user
 *                         must read output with avcodec_receive_frame() (once
 *                         all output is read, the packet should be resent, and
 *                         the call will not fail with EAGAIN).
 *      AVERROR_EOF:       the decoder has been flushed, and no new packets can
 *                         be sent to it (also returned if more than 1 flush
 *                         packet is sent)
 *      AVERROR(EINVAL):   codec not opened, it is an encoder, or requires flush
 *      AVERROR(ENOMEM):   failed to add packet to internal queue, or similar
 *      other errors: legitimate decoding errors
 */
func Avcodec_send_packet(avctx *AVCodecContext, avpkt *AVPacket) int32 {
    return int32(C.avcodec_send_packet(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt))))
}

/**
 * Return decoded output data from a decoder.
 *
 * @param avctx codec context
 * @param frame This will be set to a reference-counted video or audio
 *              frame (depending on the decoder type) allocated by the
 *              decoder. Note that the function will always call
 *              av_frame_unref(frame) before doing anything else.
 *
 * @return
 *      0:                 success, a frame was returned
 *      AVERROR(EAGAIN):   output is not available in this state - user must try
 *                         to send new input
 *      AVERROR_EOF:       the decoder has been fully flushed, and there will be
 *                         no more output frames
 *      AVERROR(EINVAL):   codec not opened, or it is an encoder
 *      other negative values: legitimate decoding errors
 */
func Avcodec_receive_frame(avctx *AVCodecContext, frame *AVFrame) int32 {
    return int32(C.avcodec_receive_frame(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Supply a raw video or audio frame to the encoder. Use avcodec_receive_packet()
 * to retrieve buffered output packets.
 *
 * @param avctx     codec context
 * @param[in] frame AVFrame containing the raw audio or video frame to be encoded.
 *                  Ownership of the frame remains with the caller, and the
 *                  encoder will not write to the frame. The encoder may create
 *                  a reference to the frame data (or copy it if the frame is
 *                  not reference-counted).
 *                  It can be NULL, in which case it is considered a flush
 *                  packet.  This signals the end of the stream. If the encoder
 *                  still has packets buffered, it will return them after this
 *                  call. Once flushing mode has been entered, additional flush
 *                  packets are ignored, and sending frames will return
 *                  AVERROR_EOF.
 *
 *                  For audio:
 *                  If AV_CODEC_CAP_VARIABLE_FRAME_SIZE is set, then each frame
 *                  can have any number of samples.
 *                  If it is not set, frame->nb_samples must be equal to
 *                  avctx->frame_size for all frames except the last.
 *                  The final frame may be smaller than avctx->frame_size.
 * @return 0 on success, otherwise negative error code:
 *      AVERROR(EAGAIN):   input is not accepted in the current state - user
 *                         must read output with avcodec_receive_packet() (once
 *                         all output is read, the packet should be resent, and
 *                         the call will not fail with EAGAIN).
 *      AVERROR_EOF:       the encoder has been flushed, and no new frames can
 *                         be sent to it
 *      AVERROR(EINVAL):   codec not opened, refcounted_frames not set, it is a
 *                         decoder, or requires flush
 *      AVERROR(ENOMEM):   failed to add packet to internal queue, or similar
 *      other errors: legitimate decoding errors
 */
func Avcodec_send_frame(avctx *AVCodecContext, frame *AVFrame) int32 {
    return int32(C.avcodec_send_frame((*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Read encoded data from the encoder.
 *
 * @param avctx codec context
 * @param avpkt This will be set to a reference-counted packet allocated by the
 *              encoder. Note that the function will always call
 *              av_frame_unref(frame) before doing anything else.
 * @return 0 on success, otherwise negative error code:
 *      AVERROR(EAGAIN):   output is not available in the current state - user
 *                         must try to send input
 *      AVERROR_EOF:       the encoder has been fully flushed, and there will be
 *                         no more output packets
 *      AVERROR(EINVAL):   codec not opened, or it is an encoder
 *      other errors: legitimate decoding errors
 */
func Avcodec_receive_packet(avctx *AVCodecContext, avpkt *AVPacket) int32 {
    return int32(C.avcodec_receive_packet(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt))))
}

/**
 * Create and return a AVHWFramesContext with values adequate for hardware
 * decoding. This is meant to get called from the get_format callback, and is
 * a helper for preparing a AVHWFramesContext for AVCodecContext.hw_frames_ctx.
 * This API is for decoding with certain hardware acceleration modes/APIs only.
 *
 * The returned AVHWFramesContext is not initialized. The caller must do this
 * with av_hwframe_ctx_init().
 *
 * Calling this function is not a requirement, but makes it simpler to avoid
 * codec or hardware API specific details when manually allocating frames.
 *
 * Alternatively to this, an API user can set AVCodecContext.hw_device_ctx,
 * which sets up AVCodecContext.hw_frames_ctx fully automatically, and makes
 * it unnecessary to call this function or having to care about
 * AVHWFramesContext initialization at all.
 *
 * There are a number of requirements for calling this function:
 *
 * - It must be called from get_format with the same avctx parameter that was
 *   passed to get_format. Calling it outside of get_format is not allowed, and
 *   can trigger undefined behavior.
 * - The function is not always supported (see description of return values).
 *   Even if this function returns successfully, hwaccel initialization could
 *   fail later. (The degree to which implementations check whether the stream
 *   is actually supported varies. Some do this check only after the user's
 *   get_format callback returns.)
 * - The hw_pix_fmt must be one of the choices suggested by get_format. If the
 *   user decides to use a AVHWFramesContext prepared with this API function,
 *   the user must return the same hw_pix_fmt from get_format.
 * - The device_ref passed to this function must support the given hw_pix_fmt.
 * - After calling this API function, it is the user's responsibility to
 *   initialize the AVHWFramesContext (returned by the out_frames_ref parameter),
 *   and to set AVCodecContext.hw_frames_ctx to it. If done, this must be done
 *   before returning from get_format (this is implied by the normal
 *   AVCodecContext.hw_frames_ctx API rules).
 * - The AVHWFramesContext parameters may change every time time get_format is
 *   called. Also, AVCodecContext.hw_frames_ctx is reset before get_format. So
 *   you are inherently required to go through this process again on every
 *   get_format call.
 * - It is perfectly possible to call this function without actually using
 *   the resulting AVHWFramesContext. One use-case might be trying to reuse a
 *   previously initialized AVHWFramesContext, and calling this API function
 *   only to test whether the required frame parameters have changed.
 * - Fields that use dynamically allocated values of any kind must not be set
 *   by the user unless setting them is explicitly allowed by the documentation.
 *   If the user sets AVHWFramesContext.free and AVHWFramesContext.user_opaque,
 *   the new free callback must call the potentially set previous free callback.
 *   This API call may set any dynamically allocated fields, including the free
 *   callback.
 *
 * The function will set at least the following fields on AVHWFramesContext
 * (potentially more, depending on hwaccel API):
 *
 * - All fields set by av_hwframe_ctx_alloc().
 * - Set the format field to hw_pix_fmt.
 * - Set the sw_format field to the most suited and most versatile format. (An
 *   implication is that this will prefer generic formats over opaque formats
 *   with arbitrary restrictions, if possible.)
 * - Set the width/height fields to the coded frame size, rounded up to the
 *   API-specific minimum alignment.
 * - Only _if_ the hwaccel requires a pre-allocated pool: set the initial_pool_size
 *   field to the number of maximum reference surfaces possible with the codec,
 *   plus 1 surface for the user to work (meaning the user can safely reference
 *   at most 1 decoded surface at a time), plus additional buffering introduced
 *   by frame threading. If the hwaccel does not require pre-allocation, the
 *   field is left to 0, and the decoder will allocate new surfaces on demand
 *   during decoding.
 * - Possibly AVHWFramesContext.hwctx fields, depending on the underlying
 *   hardware API.
 *
 * Essentially, out_frames_ref returns the same as av_hwframe_ctx_alloc(), but
 * with basic frame parameters set.
 *
 * The function is stateless, and does not change the AVCodecContext or the
 * device_ref AVHWDeviceContext.
 *
 * @param avctx The context which is currently calling get_format, and which
 *              implicitly contains all state needed for filling the returned
 *              AVHWFramesContext properly.
 * @param device_ref A reference to the AVHWDeviceContext describing the device
 *                   which will be used by the hardware decoder.
 * @param hw_pix_fmt The hwaccel format you are going to return from get_format.
 * @param out_frames_ref On success, set to a reference to an _uninitialized_
 *                       AVHWFramesContext, created from the given device_ref.
 *                       Fields will be set to values required for decoding.
 *                       Not changed if an error is returned.
 * @return zero on success, a negative value on error. The following error codes
 *         have special semantics:
 *      AVERROR(ENOENT): the decoder does not support this functionality. Setup
 *                       is always manual, or it is a decoder which does not
 *                       support setting AVCodecContext.hw_frames_ctx at all,
 *                       or it is a software format.
 *      AVERROR(EINVAL): it is known that hardware decoding is not supported for
 *                       this configuration, or the device_ref is not supported
 *                       for the hwaccel referenced by hw_pix_fmt.
 */
func Avcodec_get_hw_frames_parameters(avctx *AVCodecContext,
                                     device_ref *AVBufferRef,
                                     hw_pix_fmt AVPixelFormat,
                                     out_frames_ref **AVBufferRef) int32 {
    return int32(C.avcodec_get_hw_frames_parameters(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVBufferRef)(unsafe.Pointer(device_ref)), 
        C.enum_AVPixelFormat(hw_pix_fmt), 
        (**C.AVBufferRef)(unsafe.Pointer(out_frames_ref))))
}



/**
 * @defgroup lavc_parsing Frame parsing
 * @{
 */

type AVPictureStructure C.enum_AVPictureStructure

type AVCodecParserContext C.struct_AVCodecParserContext

type AVCodecParser C.struct_AVCodecParser

/**
 * Iterate over all registered codec parsers.
 *
 * @param opaque a pointer where libavcodec will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered codec parser or NULL when the iteration is
 *         finished
 */
func Av_parser_iterate(opaque *unsafe.Pointer) *AVCodecParser {
    return (*AVCodecParser)(unsafe.Pointer(C.av_parser_iterate(opaque)))
}


func Av_parser_next(c *AVCodecParser) *AVCodecParser {
    return (*AVCodecParser)(unsafe.Pointer(C.av_parser_next((*C.AVCodecParser)(unsafe.Pointer(c)))))
}


func Av_register_codec_parser(parser *AVCodecParser)  {
    C.av_register_codec_parser((*C.AVCodecParser)(unsafe.Pointer(parser)))
}
func Av_parser_init(codec_id int32) *AVCodecParserContext {
    return (*AVCodecParserContext)(unsafe.Pointer(C.av_parser_init(C.int(codec_id))))
}

/**
 * Parse a packet.
 *
 * @param s             parser context.
 * @param avctx         codec context.
 * @param poutbuf       set to pointer to parsed buffer or NULL if not yet finished.
 * @param poutbuf_size  set to size of parsed buffer or zero if not yet finished.
 * @param buf           input buffer.
 * @param buf_size      buffer size in bytes without the padding. I.e. the full buffer
                        size is assumed to be buf_size + AV_INPUT_BUFFER_PADDING_SIZE.
                        To signal EOF, this should be 0 (so that the last frame
                        can be output).
 * @param pts           input presentation timestamp.
 * @param dts           input decoding timestamp.
 * @param pos           input byte position in stream.
 * @return the number of bytes of the input bitstream used.
 *
 * Example:
 * @code
 *   while(in_len){
 *       len = av_parser_parse2(myparser, AVCodecContext, &data, &size,
 *                                        in_data, in_len,
 *                                        pts, dts, pos);
 *       in_data += len;
 *       in_len  -= len;
 *
 *       if(size)
 *          decode_frame(data, size);
 *   }
 * @endcode
 */
func Av_parser_parse2(s *AVCodecParserContext,
                     avctx *AVCodecContext,
                     poutbuf **uint8, poutbuf_size *int32,
                     buf *uint8, buf_size int32,
                     pts int64, dts int64,
                     pos int64) int32 {
    return int32(C.av_parser_parse2((*C.AVCodecParserContext)(unsafe.Pointer(s)), 
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (**C.uchar)(unsafe.Pointer(poutbuf)), 
        (*C.int)(unsafe.Pointer(poutbuf_size)), (*C.uchar)(unsafe.Pointer(buf)), 
        C.int(buf_size), C.longlong(pts), C.longlong(dts), C.longlong(pos)))
}

/**
 * @return 0 if the output buffer is a subset of the input, 1 if it is allocated and must be freed
 * @deprecated use AVBitStreamFilter
 */
func Av_parser_change(s *AVCodecParserContext,
                     avctx *AVCodecContext,
                     poutbuf **uint8, poutbuf_size *int32,
                     buf *uint8, buf_size int32, keyframe int32) int32 {
    return int32(C.av_parser_change((*C.AVCodecParserContext)(unsafe.Pointer(s)), 
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (**C.uchar)(unsafe.Pointer(poutbuf)), 
        (*C.int)(unsafe.Pointer(poutbuf_size)), (*C.uchar)(unsafe.Pointer(buf)), 
        C.int(buf_size), C.int(keyframe)))
}
func Av_parser_close(s *AVCodecParserContext)  {
    C.av_parser_close((*C.AVCodecParserContext)(unsafe.Pointer(s)))
}

/**
 * @}
 * @}
 */

/**
 * @addtogroup lavc_encoding
 * @{
 */

/**
 * Find a registered encoder with a matching codec ID.
 *
 * @param id AVCodecID of the requested encoder
 * @return An encoder if one was found, NULL otherwise.
 */
func Avcodec_find_encoder(id AVCodecID) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.avcodec_find_encoder(C.enum_AVCodecID(id))))
}

/**
 * Find a registered encoder with the specified name.
 *
 * @param name name of the requested encoder
 * @return An encoder if one was found, NULL otherwise.
 */
func Avcodec_find_encoder_by_name(name *byte) *AVCodec {
    return (*AVCodec)(unsafe.Pointer(C.avcodec_find_encoder_by_name(
        (*C.char)(unsafe.Pointer(name)))))
}

/**
 * Encode a frame of audio.
 *
 * Takes input samples from frame and writes the next output packet, if
 * available, to avpkt. The output packet does not necessarily contain data for
 * the most recent frame, as encoders can delay, split, and combine input frames
 * internally as needed.
 *
 * @param avctx     codec context
 * @param avpkt     output AVPacket.
 *                  The user can supply an output buffer by setting
 *                  avpkt->data and avpkt->size prior to calling the
 *                  function, but if the size of the user-provided data is not
 *                  large enough, encoding will fail. If avpkt->data and
 *                  avpkt->size are set, avpkt->destruct must also be set. All
 *                  other AVPacket fields will be reset by the encoder using
 *                  av_init_packet(). If avpkt->data is NULL, the encoder will
 *                  allocate it. The encoder will set avpkt->size to the size
 *                  of the output packet.
 *
 *                  If this function fails or produces no output, avpkt will be
 *                  freed using av_packet_unref().
 * @param[in] frame AVFrame containing the raw audio data to be encoded.
 *                  May be NULL when flushing an encoder that has the
 *                  AV_CODEC_CAP_DELAY capability set.
 *                  If AV_CODEC_CAP_VARIABLE_FRAME_SIZE is set, then each frame
 *                  can have any number of samples.
 *                  If it is not set, frame->nb_samples must be equal to
 *                  avctx->frame_size for all frames except the last.
 *                  The final frame may be smaller than avctx->frame_size.
 * @param[out] got_packet_ptr This field is set to 1 by libavcodec if the
 *                            output packet is non-empty, and to 0 if it is
 *                            empty. If the function returns an error, the
 *                            packet can be assumed to be invalid, and the
 *                            value of got_packet_ptr is undefined and should
 *                            not be used.
 * @return          0 on success, negative error code on failure
 *
 * @deprecated use avcodec_send_frame()/avcodec_receive_packet() instead
 */

func Avcodec_encode_audio2(avctx *AVCodecContext, avpkt *AVPacket,
                          frame *AVFrame, got_packet_ptr *int32) int32 {
    return int32(C.avcodec_encode_audio2(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt)), 
        (*C.AVFrame)(unsafe.Pointer(frame)), 
        (*C.int)(unsafe.Pointer(got_packet_ptr))))
}

/**
 * Encode a frame of video.
 *
 * Takes input raw video data from frame and writes the next output packet, if
 * available, to avpkt. The output packet does not necessarily contain data for
 * the most recent frame, as encoders can delay and reorder input frames
 * internally as needed.
 *
 * @param avctx     codec context
 * @param avpkt     output AVPacket.
 *                  The user can supply an output buffer by setting
 *                  avpkt->data and avpkt->size prior to calling the
 *                  function, but if the size of the user-provided data is not
 *                  large enough, encoding will fail. All other AVPacket fields
 *                  will be reset by the encoder using av_init_packet(). If
 *                  avpkt->data is NULL, the encoder will allocate it.
 *                  The encoder will set avpkt->size to the size of the
 *                  output packet. The returned data (if any) belongs to the
 *                  caller, he is responsible for freeing it.
 *
 *                  If this function fails or produces no output, avpkt will be
 *                  freed using av_packet_unref().
 * @param[in] frame AVFrame containing the raw video data to be encoded.
 *                  May be NULL when flushing an encoder that has the
 *                  AV_CODEC_CAP_DELAY capability set.
 * @param[out] got_packet_ptr This field is set to 1 by libavcodec if the
 *                            output packet is non-empty, and to 0 if it is
 *                            empty. If the function returns an error, the
 *                            packet can be assumed to be invalid, and the
 *                            value of got_packet_ptr is undefined and should
 *                            not be used.
 * @return          0 on success, negative error code on failure
 *
 * @deprecated use avcodec_send_frame()/avcodec_receive_packet() instead
 */

func Avcodec_encode_video2(avctx *AVCodecContext, avpkt *AVPacket,
                          frame *AVFrame, got_packet_ptr *int32) int32 {
    return int32(C.avcodec_encode_video2(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.AVPacket)(unsafe.Pointer(avpkt)), 
        (*C.AVFrame)(unsafe.Pointer(frame)), 
        (*C.int)(unsafe.Pointer(got_packet_ptr))))
}

func Avcodec_encode_subtitle(avctx *AVCodecContext, buf *uint8, buf_size int32,
                            sub *AVSubtitle) int32 {
    return int32(C.avcodec_encode_subtitle(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), 
        (*C.uchar)(unsafe.Pointer(buf)), C.int(buf_size), 
        (*C.AVSubtitle)(unsafe.Pointer(sub))))
}


/**
 * @}
 */

                    
   
                           
     
   

   
                     
   
                    
                                                                                           

   
                     
   
                    
                                        

   
                                                  
   
                    
                                                          
                                                                      

   
                                                     
   
                    
                                                                      
                                           
                                                         

   
                                                      
   
                    
                                                                          

   
                                       
   
                    
                                                          
                                                                        

   
                     
   
                    
                                                         
                                                                             

   
                     
   
                    
                                                                                                           
                                                                              

   
     
   
      

/**
 * @defgroup lavc_misc Utility functions
 * @ingroup libavc
 *
 * Miscellaneous utility functions related to both encoding and decoding
 * (or neither).
 * @{
 */

/**
 * @defgroup lavc_misc_pixfmt Pixel formats
 *
 * Functions for working with pixel formats.
 * @{
 */

                    
   
                                                   
   

                    
                                                                                           
      

/**
 * Return a value representing the fourCC code associated to the
 * pixel format pix_fmt, or 0 if no associated fourCC code can be
 * found.
 */
func Avcodec_pix_fmt_to_codec_tag(pix_fmt AVPixelFormat) uint32 {
    return uint32(C.avcodec_pix_fmt_to_codec_tag(C.enum_AVPixelFormat(pix_fmt)))
}

/**
 * @deprecated see av_get_pix_fmt_loss()
 */
func Avcodec_get_pix_fmt_loss(dst_pix_fmt AVPixelFormat, src_pix_fmt AVPixelFormat,
                             has_alpha int32) int32 {
    return int32(C.avcodec_get_pix_fmt_loss(C.enum_AVPixelFormat(dst_pix_fmt), 
        C.enum_AVPixelFormat(src_pix_fmt), C.int(has_alpha)))
}

/**
 * Find the best pixel format to convert to given a certain source pixel
 * format.  When converting from one pixel format to another, information loss
 * may occur.  For example, when converting from RGB24 to GRAY, the color
 * information will be lost. Similarly, other losses occur when converting from
 * some formats to other formats. avcodec_find_best_pix_fmt_of_2() searches which of
 * the given pixel formats should be used to suffer the least amount of loss.
 * The pixel formats from which it chooses one, are determined by the
 * pix_fmt_list parameter.
 *
 *
 * @param[in] pix_fmt_list AV_PIX_FMT_NONE terminated array of pixel formats to choose from
 * @param[in] src_pix_fmt source pixel format
 * @param[in] has_alpha Whether the source pixel format alpha channel is used.
 * @param[out] loss_ptr Combination of flags informing you what kind of losses will occur.
 * @return The best pixel format to convert to or -1 if none was found.
 */
func Avcodec_find_best_pix_fmt_of_list(pix_fmt_list *AVPixelFormat,
                                            src_pix_fmt AVPixelFormat,
                                            has_alpha int32, loss_ptr *int32) AVPixelFormat {
    return AVPixelFormat(C.avcodec_find_best_pix_fmt_of_list(
        (*C.enum_AVPixelFormat)(unsafe.Pointer(pix_fmt_list)), 
        C.enum_AVPixelFormat(src_pix_fmt), C.int(has_alpha), 
        (*C.int)(unsafe.Pointer(loss_ptr))))
}

/**
 * @deprecated see av_find_best_pix_fmt_of_2()
 */
func Avcodec_find_best_pix_fmt_of_2(dst_pix_fmt1 AVPixelFormat, dst_pix_fmt2 AVPixelFormat,
                                            src_pix_fmt AVPixelFormat, has_alpha int32, loss_ptr *int32) AVPixelFormat {
    return AVPixelFormat(C.avcodec_find_best_pix_fmt_of_2(
        C.enum_AVPixelFormat(dst_pix_fmt1), C.enum_AVPixelFormat(dst_pix_fmt2), 
        C.enum_AVPixelFormat(src_pix_fmt), C.int(has_alpha), 
        (*C.int)(unsafe.Pointer(loss_ptr))))
}


func Avcodec_find_best_pix_fmt2(dst_pix_fmt1 AVPixelFormat, dst_pix_fmt2 AVPixelFormat,
                                            src_pix_fmt AVPixelFormat, has_alpha int32, loss_ptr *int32) AVPixelFormat {
    return AVPixelFormat(C.avcodec_find_best_pix_fmt2(
        C.enum_AVPixelFormat(dst_pix_fmt1), C.enum_AVPixelFormat(dst_pix_fmt2), 
        C.enum_AVPixelFormat(src_pix_fmt), C.int(has_alpha), 
        (*C.int)(unsafe.Pointer(loss_ptr))))
}

func Avcodec_default_get_format(s *AVCodecContext, fmt *AVPixelFormat) AVPixelFormat {
    return AVPixelFormat(C.avcodec_default_get_format(
        (*C.struct_AVCodecContext)(unsafe.Pointer(s)), 
        (*C.enum_AVPixelFormat)(unsafe.Pointer(fmt))))
}

/**
 * @}
 */

                     
   
                                                            
  
                                                
                                       
                                       
                                                                     
                                                               
  
                                                               
   
                    
                                                                                   
      

func Avcodec_string(buf *byte, buf_size int32, enc *AVCodecContext, encode int32)  {
    C.avcodec_string((*C.char)(unsafe.Pointer(buf)), C.int(buf_size), 
        (*C.AVCodecContext)(unsafe.Pointer(enc)), C.int(encode))
}

/**
 * Return a name for the specified profile, if available.
 *
 * @param codec the codec that is searched for the given profile
 * @param profile the profile value for which a name is requested
 * @return A name for the profile if found, NULL otherwise.
 */
func Av_get_profile_name(codec *AVCodec, profile int32) string {
    return C.GoString(C.av_get_profile_name((*C.AVCodec)(unsafe.Pointer(codec)), 
        C.int(profile)))
}

/**
 * Return a name for the specified profile, if available.
 *
 * @param codec_id the ID of the codec to which the requested profile belongs
 * @param profile the profile value for which a name is requested
 * @return A name for the profile if found, NULL otherwise.
 *
 * @note unlike av_get_profile_name(), which searches a list of profiles
 *       supported by a specific decoder or encoder implementation, this
 *       function searches the list of profiles from the AVCodecDescriptor
 */
func Avcodec_profile_name(codec_id AVCodecID, profile int32) string {
    return C.GoString(C.avcodec_profile_name(C.enum_AVCodecID(codec_id), 
        C.int(profile)))
}

func Avcodec_default_execute(c *AVCodecContext, funcx func(c2 *AVCodecContext, arg2 unsafe.Pointer) int32,arg unsafe.Pointer, ret *int32, count int32, size int32) int32 {
    cb1 := syscall.NewCallbackCDecl(funcx)
    return int32(C.avcodec_default_execute(
        (*C.AVCodecContext)(unsafe.Pointer(c)), (*[0]byte)(unsafe.Pointer(cb1)), 
        arg, (*C.int)(unsafe.Pointer(ret)), C.int(count), C.int(size)))
}
func Avcodec_default_execute2(c *AVCodecContext, funcx func(c2 *AVCodecContext, arg2 unsafe.Pointer, _var2 int32, _var3 int32) int32,arg unsafe.Pointer, ret *int32, count int32) int32 {
    cb1 := syscall.NewCallbackCDecl(funcx)
    return int32(C.avcodec_default_execute2(
        (*C.AVCodecContext)(unsafe.Pointer(c)), (*[0]byte)(unsafe.Pointer(cb1)), 
        arg, (*C.int)(unsafe.Pointer(ret)), C.int(count)))
}
//FIXME func typedef

/**
 * Fill AVFrame audio data and linesize pointers.
 *
 * The buffer buf must be a preallocated buffer with a size big enough
 * to contain the specified samples amount. The filled AVFrame data
 * pointers will point to this buffer.
 *
 * AVFrame extended_data channel pointers are allocated if necessary for
 * planar audio.
 *
 * @param frame       the AVFrame
 *                    frame->nb_samples must be set prior to calling the
 *                    function. This function fills in frame->data,
 *                    frame->extended_data, frame->linesize[0].
 * @param nb_channels channel count
 * @param sample_fmt  sample format
 * @param buf         buffer to use for frame data
 * @param buf_size    size of buffer
 * @param align       plane size sample alignment (0 = default)
 * @return            >=0 on success, negative error code on failure
 * @todo return the size in bytes required to store the samples in
 * case of success, at the next libavutil bump
 */
func Avcodec_fill_audio_frame(frame *AVFrame, nb_channels int32,
                             sample_fmt AVSampleFormat, buf *uint8,
                             buf_size int32, align int32) int32 {
    return int32(C.avcodec_fill_audio_frame((*C.AVFrame)(unsafe.Pointer(frame)), 
        C.int(nb_channels), C.enum_AVSampleFormat(sample_fmt), 
        (*C.uchar)(unsafe.Pointer(buf)), C.int(buf_size), C.int(align)))
}

/**
 * Reset the internal decoder state / flush internal buffers. Should be called
 * e.g. when seeking or when switching to a different stream.
 *
 * @note when refcounted frames are not used (i.e. avctx->refcounted_frames is 0),
 * this invalidates the frames previously returned from the decoder. When
 * refcounted frames are used, the decoder just releases any references it might
 * keep internally, but the caller's reference remains valid.
 */
func Avcodec_flush_buffers(avctx *AVCodecContext)  {
    C.avcodec_flush_buffers((*C.AVCodecContext)(unsafe.Pointer(avctx)))
}

/**
 * Return codec bits per sample.
 *
 * @param[in] codec_id the codec
 * @return Number of bits per sample or zero if unknown for the given codec.
 */
func Av_get_bits_per_sample(codec_id AVCodecID) int32 {
    return int32(C.av_get_bits_per_sample(C.enum_AVCodecID(codec_id)))
}

/**
 * Return the PCM codec associated with a sample format.
 * @param be  endianness, 0 for little, 1 for big,
 *            -1 (or anything else) for native
 * @return  AV_CODEC_ID_PCM_* or AV_CODEC_ID_NONE
 */
func Av_get_pcm_codec(fmt AVSampleFormat, be int32) AVCodecID {
    return AVCodecID(C.av_get_pcm_codec(C.enum_AVSampleFormat(fmt), C.int(be)))
}

/**
 * Return codec bits per sample.
 * Only return non-zero if the bits per sample is exactly correct, not an
 * approximation.
 *
 * @param[in] codec_id the codec
 * @return Number of bits per sample or zero if unknown for the given codec.
 */
func Av_get_exact_bits_per_sample(codec_id AVCodecID) int32 {
    return int32(C.av_get_exact_bits_per_sample(C.enum_AVCodecID(codec_id)))
}

/**
 * Return audio frame duration.
 *
 * @param avctx        codec context
 * @param frame_bytes  size of the frame, or 0 if unknown
 * @return             frame duration, in samples, if known. 0 if not able to
 *                     determine.
 */
func Av_get_audio_frame_duration(avctx *AVCodecContext, frame_bytes int32) int32 {
    return int32(C.av_get_audio_frame_duration(
        (*C.AVCodecContext)(unsafe.Pointer(avctx)), C.int(frame_bytes)))
}

/**
 * This function is the same as av_get_audio_frame_duration(), except it works
 * with AVCodecParameters instead of an AVCodecContext.
 */
func Av_get_audio_frame_duration2(par *AVCodecParameters, frame_bytes int32) int32 {
    return int32(C.av_get_audio_frame_duration2(
        (*C.AVCodecParameters)(unsafe.Pointer(par)), C.int(frame_bytes)))
}

                  
                                         
                    
                                           
                                 
                                          
       
                                                                                          
                                       
       
               
                           
      

type AVBSFInternal C.struct_AVBSFInternal

/**
 * The bitstream filter state.
 *
 * This struct must be allocated with av_bsf_alloc() and freed with
 * av_bsf_free().
 *
 * The fields in the struct will only be changed (by the caller or by the
 * filter) as described in their documentation, and are to be considered
 * immutable otherwise.
 */
type AVBSFContext C.struct_AVBSFContext

type AVBitStreamFilter C.struct_AVBitStreamFilter

                  
   
                                                                               
                                                                           
   
                    
                                                          
   
                                                                               
                                                                             
                                                             
   
                    
                                                                     
   
                                                                               
                                                                               
                                                    
   
                    
                                                              
                                                                       
                                                                    
                                                                               
   
                                                                               
                                                                               
                 
   
                    
                                                              
   
                                                                               
                                                                                  
                 
   
                    
                                                                              
      

/**
 * @return a bitstream filter with the specified name or NULL if no such
 *         bitstream filter exists.
 */
func Av_bsf_get_by_name(name *byte) *AVBitStreamFilter {
    return (*AVBitStreamFilter)(unsafe.Pointer(C.av_bsf_get_by_name(
        (*C.char)(unsafe.Pointer(name)))))
}

/**
 * Iterate over all registered bitstream filters.
 *
 * @param opaque a pointer where libavcodec will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered bitstream filter or NULL when the iteration is
 *         finished
 */
func Av_bsf_iterate(opaque *unsafe.Pointer) *AVBitStreamFilter {
    return (*AVBitStreamFilter)(unsafe.Pointer(C.av_bsf_iterate(opaque)))
}
               
                    
                                                    
      

/**
 * Allocate a context for a given bitstream filter. The caller must fill in the
 * context parameters as described in the documentation and then call
 * av_bsf_init() before sending any data to the filter.
 *
 * @param filter the filter for which to allocate an instance.
 * @param ctx a pointer into which the pointer to the newly-allocated context
 *            will be written. It must be freed with av_bsf_free() after the
 *            filtering is done.
 *
 * @return 0 on success, a negative AVERROR code on failure
 */
func Av_bsf_alloc(filter *AVBitStreamFilter, ctx **AVBSFContext) int32 {
    return int32(C.av_bsf_alloc((*C.AVBitStreamFilter)(unsafe.Pointer(filter)), 
        (**C.AVBSFContext)(unsafe.Pointer(ctx))))
}

/**
 * Prepare the filter for use, after all the parameters and options have been
 * set.
 */
func Av_bsf_init(ctx *AVBSFContext) int32 {
    return int32(C.av_bsf_init((*C.AVBSFContext)(unsafe.Pointer(ctx))))
}

/**
 * Submit a packet for filtering.
 *
 * After sending each packet, the filter must be completely drained by calling
 * av_bsf_receive_packet() repeatedly until it returns AVERROR(EAGAIN) or
 * AVERROR_EOF.
 *
 * @param pkt the packet to filter. The bitstream filter will take ownership of
 * the packet and reset the contents of pkt. pkt is not touched if an error occurs.
 * This parameter may be NULL, which signals the end of the stream (i.e. no more
 * packets will be sent). That will cause the filter to output any packets it
 * may have buffered internally.
 *
 * @return 0 on success, a negative AVERROR on error.
 */
func Av_bsf_send_packet(ctx *AVBSFContext, pkt *AVPacket) int32 {
    return int32(C.av_bsf_send_packet((*C.AVBSFContext)(unsafe.Pointer(ctx)), 
        (*C.AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Retrieve a filtered packet.
 *
 * @param[out] pkt this struct will be filled with the contents of the filtered
 *                 packet. It is owned by the caller and must be freed using
 *                 av_packet_unref() when it is no longer needed.
 *                 This parameter should be "clean" (i.e. freshly allocated
 *                 with av_packet_alloc() or unreffed with av_packet_unref())
 *                 when this function is called. If this function returns
 *                 successfully, the contents of pkt will be completely
 *                 overwritten by the returned data. On failure, pkt is not
 *                 touched.
 *
 * @return 0 on success. AVERROR(EAGAIN) if more packets need to be sent to the
 * filter (using av_bsf_send_packet()) to get more output. AVERROR_EOF if there
 * will be no further output from the filter. Another negative AVERROR value if
 * an error occurs.
 *
 * @note one input packet may result in several output packets, so after sending
 * a packet with av_bsf_send_packet(), this function needs to be called
 * repeatedly until it stops returning 0. It is also possible for a filter to
 * output fewer packets than were sent to it, so this function may return
 * AVERROR(EAGAIN) immediately after a successful av_bsf_send_packet() call.
 */
func Av_bsf_receive_packet(ctx *AVBSFContext, pkt *AVPacket) int32 {
    return int32(C.av_bsf_receive_packet((*C.AVBSFContext)(unsafe.Pointer(ctx)), 
        (*C.AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Reset the internal bitstream filter state / flush internal buffers.
 */
func Av_bsf_flush(ctx *AVBSFContext)  {
    C.av_bsf_flush((*C.AVBSFContext)(unsafe.Pointer(ctx)))
}

/**
 * Free a bitstream filter context and everything associated with it; write NULL
 * into the supplied pointer.
 */
func Av_bsf_free(ctx **AVBSFContext)  {
    C.av_bsf_free((**C.AVBSFContext)(unsafe.Pointer(ctx)))
}

/**
 * Get the AVClass for AVBSFContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func Av_bsf_get_class() *AVClass {
    return (*AVClass)(unsafe.Pointer(C.av_bsf_get_class()))
}

/**
 * Structure for chain/list of bitstream filters.
 * Empty list can be allocated by av_bsf_list_alloc().
 */
type AVBSFList C.struct_AVBSFList

/**
 * Allocate empty list of bitstream filters.
 * The list must be later freed by av_bsf_list_free()
 * or finalized by av_bsf_list_finalize().
 *
 * @return Pointer to @ref AVBSFList on success, NULL in case of failure
 */
func Av_bsf_list_alloc() *AVBSFList {
    return (*AVBSFList)(unsafe.Pointer(C.av_bsf_list_alloc()))
}

/**
 * Free list of bitstream filters.
 *
 * @param lst Pointer to pointer returned by av_bsf_list_alloc()
 */
func Av_bsf_list_free(lst **AVBSFList)  {
    C.av_bsf_list_free((**C.AVBSFList)(unsafe.Pointer(lst)))
}

/**
 * Append bitstream filter to the list of bitstream filters.
 *
 * @param lst List to append to
 * @param bsf Filter context to be appended
 *
 * @return >=0 on success, negative AVERROR in case of failure
 */
func Av_bsf_list_append(lst *AVBSFList, bsf *AVBSFContext) int32 {
    return int32(C.av_bsf_list_append((*C.AVBSFList)(unsafe.Pointer(lst)), 
        (*C.AVBSFContext)(unsafe.Pointer(bsf))))
}

/**
 * Construct new bitstream filter context given it's name and options
 * and append it to the list of bitstream filters.
 *
 * @param lst      List to append to
 * @param bsf_name Name of the bitstream filter
 * @param options  Options for the bitstream filter, can be set to NULL
 *
 * @return >=0 on success, negative AVERROR in case of failure
 */
func Av_bsf_list_append2(lst *AVBSFList, bsf_name *byte, options **AVDictionary) int32 {
    return int32(C.av_bsf_list_append2((*C.AVBSFList)(unsafe.Pointer(lst)), 
        (*C.char)(unsafe.Pointer(bsf_name)), 
        (**C.AVDictionary)(unsafe.Pointer(options))))
}
/**
 * Finalize list of bitstream filters.
 *
 * This function will transform @ref AVBSFList to single @ref AVBSFContext,
 * so the whole chain of bitstream filters can be treated as single filter
 * freshly allocated by av_bsf_alloc().
 * If the call is successful, @ref AVBSFList structure is freed and lst
 * will be set to NULL. In case of failure, caller is responsible for
 * freeing the structure by av_bsf_list_free()
 *
 * @param      lst Filter list structure to be transformed
 * @param[out] bsf Pointer to be set to newly created @ref AVBSFContext structure
 *                 representing the chain of bitstream filters
 *
 * @return >=0 on success, negative AVERROR in case of failure
 */
func Av_bsf_list_finalize(lst **AVBSFList, bsf **AVBSFContext) int32 {
    return int32(C.av_bsf_list_finalize((**C.AVBSFList)(unsafe.Pointer(lst)), 
        (**C.AVBSFContext)(unsafe.Pointer(bsf))))
}

/**
 * Parse string describing list of bitstream filters and create single
 * @ref AVBSFContext describing the whole chain of bitstream filters.
 * Resulting @ref AVBSFContext can be treated as any other @ref AVBSFContext freshly
 * allocated by av_bsf_alloc().
 *
 * @param      str String describing chain of bitstream filters in format
 *                 `bsf1[=opt1=val1:opt2=val2][,bsf2]`
 * @param[out] bsf Pointer to be set to newly created @ref AVBSFContext structure
 *                 representing the chain of bitstream filters
 *
 * @return >=0 on success, negative AVERROR in case of failure
 */
func Av_bsf_list_parse_str(str *byte, bsf **AVBSFContext) int32 {
    return int32(C.av_bsf_list_parse_str((*C.char)(unsafe.Pointer(str)), 
        (**C.AVBSFContext)(unsafe.Pointer(bsf))))
}

/**
 * Get null/pass-through bitstream filter.
 *
 * @param[out] bsf Pointer to be set to new instance of pass-through bitstream filter
 *
 * @return
 */
func Av_bsf_get_null_filter(bsf **AVBSFContext) int32 {
    return int32(C.av_bsf_get_null_filter(
        (**C.AVBSFContext)(unsafe.Pointer(bsf))))
}

/* memory */

/**
 * Same behaviour av_fast_malloc but the buffer has additional
 * AV_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
 *
 * In addition the whole buffer will initially and after resizes
 * be 0-initialized so that no uninitialized data will ever appear.
 */
func Av_fast_padded_malloc(ptr unsafe.Pointer, size *uint32, min_size uint64)  {
    C.av_fast_padded_malloc(ptr, (*C.uint)(unsafe.Pointer(size)), 
        C.ulonglong(min_size))
}

/**
 * Same behaviour av_fast_padded_malloc except that buffer will always
 * be 0-initialized after call.
 */
func Av_fast_padded_mallocz(ptr unsafe.Pointer, size *uint32, min_size uint64)  {
    C.av_fast_padded_mallocz(ptr, (*C.uint)(unsafe.Pointer(size)), 
        C.ulonglong(min_size))
}

/**
 * Encode extradata length to a buffer. Used by xiph codecs.
 *
 * @param s buffer to write to; must be at least (v/255+1) bytes long
 * @param v size of extradata in bytes
 * @return number of bytes written to the buffer.
 */
func Av_xiphlacing(s *uint8, v uint32) uint32 {
    return uint32(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), C.uint(v)))
}

                                 
   
                                             
  
                                                  
   
                    
                                             

   
                                                                         
                                                                           
                                                     
  
                                                                          
                                                 
   
                    
                                                     
      

                  
   
                                 
  
                                                              
   
               
                                      
                                      
                                        
                                            
  

   
                                                                  
                                                                     
                                                                       
                                                                   
                                                                        
                                                                      
                                                                    
                                                                       
                                                                    
                                                                    
                                                                     
                                                                   
                                                                     
  
                                                                     
                                                                     
                                                                    
                                                                       
                                                                 
                                                                   
                                          
  
                                                                           
                                                                    
   
                    
                                                                   
      

/**
 * Get the type of the given codec.
 */
func Avcodec_get_type(codec_id AVCodecID) AVMediaType {
    return AVMediaType(C.avcodec_get_type(C.enum_AVCodecID(codec_id)))
}

/**
 * Get the name of a codec.
 * @return  a static string identifying the codec; never NULL
 */
func Avcodec_get_name(id AVCodecID) string {
    return C.GoString(C.avcodec_get_name(C.enum_AVCodecID(id)))
}

/**
 * @return a positive value if s is open (i.e. avcodec_open2() was called on it
 * with no corresponding avcodec_close()), 0 otherwise.
 */
func Avcodec_is_open(s *AVCodecContext) int32 {
    return int32(C.avcodec_is_open((*C.AVCodecContext)(unsafe.Pointer(s))))
}

/**
 * @return a non-zero number if codec is an encoder, zero otherwise
 */
func Av_codec_is_encoder(codec *AVCodec) int32 {
    return int32(C.av_codec_is_encoder((*C.AVCodec)(unsafe.Pointer(codec))))
}

/**
 * @return a non-zero number if codec is a decoder, zero otherwise
 */
func Av_codec_is_decoder(codec *AVCodec) int32 {
    return int32(C.av_codec_is_decoder((*C.AVCodec)(unsafe.Pointer(codec))))
}

/**
 * @return descriptor for given codec ID or NULL if no descriptor exists.
 */
func Avcodec_descriptor_get(id AVCodecID) *AVCodecDescriptor {
    return (*AVCodecDescriptor)(unsafe.Pointer(C.avcodec_descriptor_get(C.enum_AVCodecID(id))))
}

/**
 * Iterate over all codec descriptors known to libavcodec.
 *
 * @param prev previous descriptor. NULL to get the first descriptor.
 *
 * @return next descriptor or NULL after the last descriptor
 */
func Avcodec_descriptor_next(prev *AVCodecDescriptor) *AVCodecDescriptor {
    return (*AVCodecDescriptor)(unsafe.Pointer(C.avcodec_descriptor_next(
        (*C.AVCodecDescriptor)(unsafe.Pointer(prev)))))
}

/**
 * @return codec descriptor with the given name or NULL if no such descriptor
 *         exists.
 */
func Avcodec_descriptor_get_by_name(name *byte) *AVCodecDescriptor {
    return (*AVCodecDescriptor)(unsafe.Pointer(C.avcodec_descriptor_get_by_name(
        (*C.char)(unsafe.Pointer(name)))))
}

/**
 * Allocate a CPB properties structure and initialize its fields to default
 * values.
 *
 * @param size if non-NULL, the size of the allocated struct will be written
 *             here. This is useful for embedding it in side data.
 *
 * @return the newly allocated struct or NULL on failure
 */
func Av_cpb_properties_alloc(size *uint64) *AVCPBProperties {
    return (*AVCPBProperties)(unsafe.Pointer(C.av_cpb_properties_alloc(
        (*C.ulonglong)(unsafe.Pointer(size)))))
}

/**
 * @}
 */

                              
