// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// sms10086@hotmail.com

/*
 * copyright (c) 2003 Fabrice Bellard
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


//const LIBAVUTIL_VERSION_MAJOR = 56
//const LIBAVUTIL_VERSION_MINOR = 22
//const LIBAVUTIL_VERSION_MICRO = 100
//const LIBAVUTIL_VERSION_INT =    AV_VERSION_INT(LIBAVUTIL_VERSION_MAJOR,                                                 LIBAVUTIL_VERSION_MINOR,                                                 LIBAVUTIL_VERSION_MICRO) 
//const LIBAVUTIL_VERSION =        AV_VERSION(LIBAVUTIL_VERSION_MAJOR,                                                 LIBAVUTIL_VERSION_MINOR,                                                 LIBAVUTIL_VERSION_MICRO) 
//const LIBAVUTIL_BUILD =          LIBAVUTIL_VERSION_INT 
//const LIBAVUTIL_IDENT =          "Lavu" AV_STRINGIFY(LIBAVUTIL_VERSION) 



/**
 * @file
 * @ingroup lavu
 * Libavutil version macros
 */

                        
                        

                   

/**
 * @addtogroup version_utils
 *
 * Useful to check and match library version in order to maintain
 * backward compatibility.
 *
 * The FFmpeg libraries follow a versioning sheme very similar to
 * Semantic Versioning (http://semver.org/)
 * The difference is that the component called PATCH is called MICRO in FFmpeg
 * and its value is reset to 100 instead of 0 to keep it above or equal to 100.
 * Also we do not increase MICRO for every bugfix or change in git master.
 *
 * Prior to FFmpeg 3.2 point releases did not change any lib version number to
 * avoid aliassing different git master checkouts.
 * Starting with FFmpeg 3.2, the released library versions will occupy
 * a separate MAJOR.MINOR that is not used on the master development branch.
 * That is if we branch a release of master 55.10.123 we will bump to 55.11.100
 * for the release and master will continue at 55.12.100 after it. Each new
 * point release will then bump the MICRO improving the usefulness of the lib
 * versions.
 *
 * @{
 */





/**
 * Extract version components from the full ::AV_VERSION_INT int as returned
 * by functions like ::avformat_version() and ::avcodec_version()
 */




/**
 * @}
 */

/**
 * @defgroup lavu_ver Version and Build diagnostics
 *
 * Macros and function useful to check at compiletime and at runtime
 * which version of libavutil is in use.
 *
 * @{
 */











/**
 * @defgroup lavu_depr_guards Deprecation Guards
 * FF_API_* defines may be placed below to indicate public API that will be
 * dropped at a future version bump. The defines themselves are not part of
 * the public API and may change, break or disappear at any time.
 *
 * @note, when bumping the major version it is recommended to manually
 * disable each FF_API_* in its own commit instead of disabling them all
 * at once through the bump. This improves the git bisect-ability of the change.
 *
 * @{
 */

                    
                                                                      
      
                       
                                                                      
      
                           
                                                                      
      
                          
                                                                      
      
                      
                                                                      
      
                            
                                                                      
      
                            
                                                                      
      
                        
                                                                      
      


/**
 * @}
 * @}
 */

                             
