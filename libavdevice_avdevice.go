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

//#cgo pkg-config: libavutil libavdevice libavformat
//#include "libavdevice/version.h"
//#include "libavutil/log.h"
//#include "libavutil/opt.h"
//#include "libavutil/dict.h"
//#include "libavformat/avformat.h"
//#include "libavdevice/avdevice.h"
import "C"
import (
    "unsafe"
)




                           
                           

                    

/**
 * @file
 * @ingroup lavd
 * Main libavdevice API header
 */

/**
 * @defgroup lavd libavdevice
 * Special devices muxing/demuxing library.
 *
 * Libavdevice is a complementary library to @ref libavf "libavformat". It
 * provides various "special" platform-specific muxers and demuxers, e.g. for
 * grabbing devices, audio capture and playback etc. As a consequence, the
 * (de)muxers in libavdevice are of the AVFMT_NOFILE type (they use their own
 * I/O functions). The filename passed to avformat_open_input() often does not
 * refer to an actually existing file, but has some special device-specific
 * meaning - e.g. for xcbgrab it is the display name.
 *
 * To use libavdevice, simply call avdevice_register_all() to register all
 * compiled muxers and demuxers. They all use standard libavformat API.
 *
 * @{
 */

                          
                          
                           
                                 

/**
 * Return the LIBAVDEVICE_VERSION_INT constant.
 */
func Avdevice_version() uint32 {
    return uint32(C.avdevice_version())
}

/**
 * Return the libavdevice build-time configuration.
 */
func Avdevice_configuration() string {
    return C.GoString(C.avdevice_configuration())
}

/**
 * Return the libavdevice license.
 */
func Avdevice_license() string {
    return C.GoString(C.avdevice_license())
}

/**
 * Initialize libavdevice and register all the input and output devices.
 */
func Avdevice_register_all()  {
    C.avdevice_register_all()
}

/**
 * Audio input devices iterator.
 *
 * If d is NULL, returns the first registered input audio/video device,
 * if d is non-NULL, returns the next registered input audio/video device after d
 * or NULL if d is the last one.
 */
func Av_input_audio_device_next(d *AVInputFormat) *AVInputFormat {
    return (*AVInputFormat)(unsafe.Pointer(C.av_input_audio_device_next(
        (*C.struct_AVInputFormat)(unsafe.Pointer(d)))))
}

/**
 * Video input devices iterator.
 *
 * If d is NULL, returns the first registered input audio/video device,
 * if d is non-NULL, returns the next registered input audio/video device after d
 * or NULL if d is the last one.
 */
func Av_input_video_device_next(d *AVInputFormat) *AVInputFormat {
    return (*AVInputFormat)(unsafe.Pointer(C.av_input_video_device_next(
        (*C.struct_AVInputFormat)(unsafe.Pointer(d)))))
}

/**
 * Audio output devices iterator.
 *
 * If d is NULL, returns the first registered output audio/video device,
 * if d is non-NULL, returns the next registered output audio/video device after d
 * or NULL if d is the last one.
 */
func Av_output_audio_device_next(d *AVOutputFormat) *AVOutputFormat {
    return (*AVOutputFormat)(unsafe.Pointer(C.av_output_audio_device_next(
        (*C.struct_AVOutputFormat)(unsafe.Pointer(d)))))
}

/**
 * Video output devices iterator.
 *
 * If d is NULL, returns the first registered output audio/video device,
 * if d is non-NULL, returns the next registered output audio/video device after d
 * or NULL if d is the last one.
 */
func Av_output_video_device_next(d *AVOutputFormat) *AVOutputFormat {
    return (*AVOutputFormat)(unsafe.Pointer(C.av_output_video_device_next(
        (*C.struct_AVOutputFormat)(unsafe.Pointer(d)))))
}

type AVDeviceRect struct {
    X int32
    Y int32
    Width int32
    Height int32
}


/**
 * Message types used by avdevice_app_to_dev_control_message().
 */
type AVAppToDevMessageType int32


/**
 * Message types used by avdevice_dev_to_app_control_message().
 */
type AVDevToAppMessageType int32


/**
 * Send control message from application to device.
 *
 * @param s         device context.
 * @param type      message type.
 * @param data      message data. Exact type depends on message type.
 * @param data_size size of message data.
 * @return >= 0 on success, negative on error.
 *         AVERROR(ENOSYS) when device doesn't implement handler of the message.
 */
func Avdevice_app_to_dev_control_message(s *AVFormatContext,
                                        typex AVAppToDevMessageType,
                                        data unsafe.Pointer, data_size uint64) int32 {
    return int32(C.avdevice_app_to_dev_control_message(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        C.enum_AVAppToDevMessageType(typex), data, C.ulonglong(data_size)))
}

/**
 * Send control message from device to application.
 *
 * @param s         device context.
 * @param type      message type.
 * @param data      message data. Can be NULL.
 * @param data_size size of message data.
 * @return >= 0 on success, negative on error.
 *         AVERROR(ENOSYS) when application doesn't implement handler of the message.
 */
func Avdevice_dev_to_app_control_message(s *AVFormatContext,
                                        typex AVDevToAppMessageType,
                                        data unsafe.Pointer, data_size uint64) int32 {
    return int32(C.avdevice_dev_to_app_control_message(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        C.enum_AVDevToAppMessageType(typex), data, C.ulonglong(data_size)))
}

/**
 * Following API allows user to probe device capabilities (supported codecs,
 * pixel formats, sample formats, resolutions, channel counts, etc).
 * It is build on top op AVOption API.
 * Queried capabilities make it possible to set up converters of video or audio
 * parameters that fit to the device.
 *
 * List of capabilities that can be queried:
 *  - Capabilities valid for both audio and video devices:
 *    - codec:          supported audio/video codecs.
 *                      type: AV_OPT_TYPE_INT (AVCodecID value)
 *  - Capabilities valid for audio devices:
 *    - sample_format:  supported sample formats.
 *                      type: AV_OPT_TYPE_INT (AVSampleFormat value)
 *    - sample_rate:    supported sample rates.
 *                      type: AV_OPT_TYPE_INT
 *    - channels:       supported number of channels.
 *                      type: AV_OPT_TYPE_INT
 *    - channel_layout: supported channel layouts.
 *                      type: AV_OPT_TYPE_INT64
 *  - Capabilities valid for video devices:
 *    - pixel_format:   supported pixel formats.
 *                      type: AV_OPT_TYPE_INT (AVPixelFormat value)
 *    - window_size:    supported window sizes (describes size of the window size presented to the user).
 *                      type: AV_OPT_TYPE_IMAGE_SIZE
 *    - frame_size:     supported frame sizes (describes size of provided video frames).
 *                      type: AV_OPT_TYPE_IMAGE_SIZE
 *    - fps:            supported fps values
 *                      type: AV_OPT_TYPE_RATIONAL
 *
 * Value of the capability may be set by user using av_opt_set() function
 * and AVDeviceCapabilitiesQuery object. Following queries will
 * limit results to the values matching already set capabilities.
 * For example, setting a codec may impact number of formats or fps values
 * returned during next query. Setting invalid value may limit results to zero.
 *
 * Example of the usage basing on opengl output device:
 *
 * @code
 *  AVFormatContext *oc = NULL;
 *  AVDeviceCapabilitiesQuery *caps = NULL;
 *  AVOptionRanges *ranges;
 *  int ret;
 *
 *  if ((ret = avformat_alloc_output_context2(&oc, NULL, "opengl", NULL)) < 0)
 *      goto fail;
 *  if (avdevice_capabilities_create(&caps, oc, NULL) < 0)
 *      goto fail;
 *
 *  //query codecs
 *  if (av_opt_query_ranges(&ranges, caps, "codec", AV_OPT_MULTI_COMPONENT_RANGE)) < 0)
 *      goto fail;
 *  //pick codec here and set it
 *  av_opt_set(caps, "codec", AV_CODEC_ID_RAWVIDEO, 0);
 *
 *  //query format
 *  if (av_opt_query_ranges(&ranges, caps, "pixel_format", AV_OPT_MULTI_COMPONENT_RANGE)) < 0)
 *      goto fail;
 *  //pick format here and set it
 *  av_opt_set(caps, "pixel_format", AV_PIX_FMT_YUV420P, 0);
 *
 *  //query and set more capabilities
 *
 * fail:
 *  //clean up code
 *  avdevice_capabilities_free(&query, oc);
 *  avformat_free_context(oc);
 * @endcode
 */

/**
 * Structure describes device capabilities.
 *
 * It is used by devices in conjunction with av_device_capabilities AVOption table
 * to implement capabilities probing API based on AVOption API. Should not be used directly.
 */
type AVDeviceCapabilitiesQuery struct {
    Av_class *AVClass
    Device_context *AVFormatContext
    Codec AVCodecID
    Sample_format AVSampleFormat
    Pixel_format AVPixelFormat
    Sample_rate int32
    Channels int32
    Channel_layout int64
    Window_width int32
    Window_height int32
    Frame_width int32
    Frame_height int32
    Fps AVRational
}


/**
 * AVOption table used by devices to implement device capabilities API. Should not be used by a user.
 */
//extern const AVOption av_device_capabilities[];

/**
 * Initialize capabilities probing API based on AVOption API.
 *
 * avdevice_capabilities_free() must be called when query capabilities API is
 * not used anymore.
 *
 * @param[out] caps      Device capabilities data. Pointer to a NULL pointer must be passed.
 * @param s              Context of the device.
 * @param device_options An AVDictionary filled with device-private options.
 *                       On return this parameter will be destroyed and replaced with a dict
 *                       containing options that were not found. May be NULL.
 *                       The same options must be passed later to avformat_write_header() for output
 *                       devices or avformat_open_input() for input devices, or at any other place
 *                       that affects device-private options.
 *
 * @return >= 0 on success, negative otherwise.
 */
func Avdevice_capabilities_create(caps **AVDeviceCapabilitiesQuery, s *AVFormatContext,
                                 device_options **AVDictionary) int32 {
    return int32(C.avdevice_capabilities_create(
        (**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(caps)), 
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (**C.struct_AVDictionary)(unsafe.Pointer(device_options))))
}

/**
 * Free resources created by avdevice_capabilities_create()
 *
 * @param caps Device capabilities data to be freed.
 * @param s    Context of the device.
 */
func Avdevice_capabilities_free(caps **AVDeviceCapabilitiesQuery, s *AVFormatContext)  {
    C.avdevice_capabilities_free(
        (**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(caps)), 
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)))
}

/**
 * Structure describes basic parameters of the device.
 */
type AVDeviceInfo struct {
    Device_name *byte
    Device_description *byte
}


/**
 * List of devices.
 */
type AVDeviceInfoList struct {
    Devices **AVDeviceInfo
    Nb_devices int32
    Default_device int32
}


/**
 * List devices.
 *
 * Returns available device names and their parameters.
 *
 * @note: Some devices may accept system-dependent device names that cannot be
 *        autodetected. The list returned by this function cannot be assumed to
 *        be always completed.
 *
 * @param s                device context.
 * @param[out] device_list list of autodetected devices.
 * @return count of autodetected devices, negative on error.
 */
func Avdevice_list_devices(s *AVFormatContext, device_list **AVDeviceInfoList) int32 {
    return int32(C.avdevice_list_devices(
        (*C.struct_AVFormatContext)(unsafe.Pointer(s)), 
        (**C.struct_AVDeviceInfoList)(unsafe.Pointer(device_list))))
}

/**
 * Convenient function to free result of avdevice_list_devices().
 *
 * @param devices device list to be freed.
 */
func Avdevice_free_list_devices(device_list **AVDeviceInfoList)  {
    C.avdevice_free_list_devices(
        (**C.struct_AVDeviceInfoList)(unsafe.Pointer(device_list)))
}

/**
 * List devices.
 *
 * Returns available device names and their parameters.
 * These are convinient wrappers for avdevice_list_devices().
 * Device context is allocated and deallocated internally.
 *
 * @param device           device format. May be NULL if device name is set.
 * @param device_name      device name. May be NULL if device format is set.
 * @param device_options   An AVDictionary filled with device-private options. May be NULL.
 *                         The same options must be passed later to avformat_write_header() for output
 *                         devices or avformat_open_input() for input devices, or at any other place
 *                         that affects device-private options.
 * @param[out] device_list list of autodetected devices
 * @return count of autodetected devices, negative on error.
 * @note device argument takes precedence over device_name when both are set.
 */
func Avdevice_list_input_sources(device *AVInputFormat, device_name *byte,
                                device_options *AVDictionary, device_list **AVDeviceInfoList) int32 {
    return int32(C.avdevice_list_input_sources(
        (*C.struct_AVInputFormat)(unsafe.Pointer(device)), 
        (*C.char)(unsafe.Pointer(device_name)), 
        (*C.struct_AVDictionary)(unsafe.Pointer(device_options)), 
        (**C.struct_AVDeviceInfoList)(unsafe.Pointer(device_list))))
}
func Avdevice_list_output_sinks(device *AVOutputFormat, device_name *byte,
                               device_options *AVDictionary, device_list **AVDeviceInfoList) int32 {
    return int32(C.avdevice_list_output_sinks(
        (*C.struct_AVOutputFormat)(unsafe.Pointer(device)), 
        (*C.char)(unsafe.Pointer(device_name)), 
        (*C.struct_AVDictionary)(unsafe.Pointer(device_options)), 
        (**C.struct_AVDeviceInfoList)(unsafe.Pointer(device_list))))
}

/**
 * @}
 */

                                
