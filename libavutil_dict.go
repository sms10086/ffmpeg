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

//#cgo pkg-config: libavutil
//#include <stdint.h>
//#include "libavutil/version.h"
//#include "libavutil/dict.h"
import "C"
import (
    "unsafe"
)

const AV_DICT_MATCH_CASE = 1
const AV_DICT_IGNORE_SUFFIX = 2
const AV_DICT_DONT_STRDUP_KEY = 4
const AV_DICT_DONT_STRDUP_VAL = 8
const AV_DICT_DONT_OVERWRITE = 16
const AV_DICT_APPEND = 32
const AV_DICT_MULTIKEY = 64



/**
 * @file
 * Public dictionary API.
 * @deprecated
 *  AVDictionary is provided for compatibility with libav. It is both in
 *  implementation as well as API inefficient. It does not scale and is
 *  extremely slow with large dictionaries.
 *  It is recommended that new code uses our tree container from tree.c/h
 *  where applicable, which uses AVL trees to achieve O(log n) performance.
 */

                     
                     

                   

                    

/**
 * @addtogroup lavu_dict AVDictionary
 * @ingroup lavu_data
 *
 * @brief Simple key:value store
 *
 * @{
 * Dictionaries are used for storing key:value pairs. To create
 * an AVDictionary, simply pass an address of a NULL pointer to
 * av_dict_set(). NULL can be used as an empty dictionary wherever
 * a pointer to an AVDictionary is required.
 * Use av_dict_get() to retrieve an entry or iterate over all
 * entries and finally av_dict_free() to free the dictionary
 * and all its contents.
 *
 @code
   AVDictionary *d = NULL;           // "create" an empty dictionary
   AVDictionaryEntry *t = NULL;

   av_dict_set(&d, "foo", "bar", 0); // add an entry

   char *k = av_strdup("key");       // if your strings are already allocated,
   char *v = av_strdup("value");     // you can avoid copying them like this
   av_dict_set(&d, k, v, AV_DICT_DONT_STRDUP_KEY | AV_DICT_DONT_STRDUP_VAL);

   while (t = av_dict_get(d, "", t, AV_DICT_IGNORE_SUFFIX)) {
       <....>                             // iterate over all entries in d
   }
   av_dict_free(&d);
 @endcode
 */

/**< Only get an entry with exact-case key match. Only relevant in av_dict_get(). */
/**< Return first entry in a dictionary whose first part corresponds to the search key,
                                         ignoring the suffix of the found key string. Only relevant in av_dict_get(). */
/**< Take ownership of a key that's been
                                         allocated with av_malloc() or another memory allocation function. */
/**< Take ownership of a value that's been
                                         allocated with av_malloc() or another memory allocation function. */
///< Don't overwrite existing entries.
/**< If the entry already exists, append to it.  Note that no
                                      delimiter is added, the strings are simply concatenated. */
/**< Allow to store several equal keys in the dictionary */

type AVDictionaryEntry struct {
    Key *byte
    Value *byte
}


type AVDictionary struct {
}


/**
 * Get a dictionary entry with matching key.
 *
 * The returned entry key or value must not be changed, or it will
 * cause undefined behavior.
 *
 * To iterate through all the dictionary entries, you can set the matching key
 * to the null string "" and set the AV_DICT_IGNORE_SUFFIX flag.
 *
 * @param prev Set to the previous matching element to find the next.
 *             If set to NULL the first matching element is returned.
 * @param key matching key
 * @param flags a collection of AV_DICT_* flags controlling how the entry is retrieved
 * @return found entry or NULL in case no matching entry was found in the dictionary
 */
func Av_dict_get(m *AVDictionary, key *byte,
                               prev *AVDictionaryEntry, flags int32) *AVDictionaryEntry {
    return (*AVDictionaryEntry)(unsafe.Pointer(C.av_dict_get(
        (*C.struct_AVDictionary)(unsafe.Pointer(m)), 
        (*C.char)(unsafe.Pointer(key)), 
        (*C.struct_AVDictionaryEntry)(unsafe.Pointer(prev)), C.int(flags))))
}

/**
 * Get number of entries in dictionary.
 *
 * @param m dictionary
 * @return  number of entries in dictionary
 */
func Av_dict_count(m *AVDictionary) int32 {
    return int32(C.av_dict_count((*C.struct_AVDictionary)(unsafe.Pointer(m))))
}

/**
 * Set the given entry in *pm, overwriting an existing entry.
 *
 * Note: If AV_DICT_DONT_STRDUP_KEY or AV_DICT_DONT_STRDUP_VAL is set,
 * these arguments will be freed on error.
 *
 * Warning: Adding a new entry to a dictionary invalidates all existing entries
 * previously returned with av_dict_get.
 *
 * @param pm pointer to a pointer to a dictionary struct. If *pm is NULL
 * a dictionary struct is allocated and put in *pm.
 * @param key entry key to add to *pm (will either be av_strduped or added as a new key depending on flags)
 * @param value entry value to add to *pm (will be av_strduped or added as a new key depending on flags).
 *        Passing a NULL value will cause an existing entry to be deleted.
 * @return >= 0 on success otherwise an error code <0
 */
func Av_dict_set(pm **AVDictionary, key *byte, value *byte, flags int32) int32 {
    return int32(C.av_dict_set((**C.struct_AVDictionary)(unsafe.Pointer(pm)), 
        (*C.char)(unsafe.Pointer(key)), (*C.char)(unsafe.Pointer(value)), 
        C.int(flags)))
}

/**
 * Convenience wrapper for av_dict_set that converts the value to a string
 * and stores it.
 *
 * Note: If AV_DICT_DONT_STRDUP_KEY is set, key will be freed on error.
 */
func Av_dict_set_int(pm **AVDictionary, key *byte, value int64, flags int32) int32 {
    return int32(C.av_dict_set_int((**C.struct_AVDictionary)(unsafe.Pointer(pm)), 
        (*C.char)(unsafe.Pointer(key)), C.longlong(value), C.int(flags)))
}

/**
 * Parse the key/value pairs list and add the parsed entries to a dictionary.
 *
 * In case of failure, all the successfully set entries are stored in
 * *pm. You may need to manually free the created dictionary.
 *
 * @param key_val_sep  a 0-terminated list of characters used to separate
 *                     key from value
 * @param pairs_sep    a 0-terminated list of characters used to separate
 *                     two pairs from each other
 * @param flags        flags to use when adding to dictionary.
 *                     AV_DICT_DONT_STRDUP_KEY and AV_DICT_DONT_STRDUP_VAL
 *                     are ignored since the key/value tokens will always
 *                     be duplicated.
 * @return             0 on success, negative AVERROR code on failure
 */
func Av_dict_parse_string(pm **AVDictionary, str *byte,
                         key_val_sep *byte, pairs_sep *byte,
                         flags int32) int32 {
    return int32(C.av_dict_parse_string(
        (**C.struct_AVDictionary)(unsafe.Pointer(pm)), 
        (*C.char)(unsafe.Pointer(str)), (*C.char)(unsafe.Pointer(key_val_sep)), 
        (*C.char)(unsafe.Pointer(pairs_sep)), C.int(flags)))
}

/**
 * Copy entries from one AVDictionary struct into another.
 * @param dst pointer to a pointer to a AVDictionary struct. If *dst is NULL,
 *            this function will allocate a struct for you and put it in *dst
 * @param src pointer to source AVDictionary struct
 * @param flags flags to use when setting entries in *dst
 * @note metadata is read using the AV_DICT_IGNORE_SUFFIX flag
 * @return 0 on success, negative AVERROR code on failure. If dst was allocated
 *           by this function, callers should free the associated memory.
 */
func Av_dict_copy(dst **AVDictionary, src *AVDictionary, flags int32) int32 {
    return int32(C.av_dict_copy((**C.struct_AVDictionary)(unsafe.Pointer(dst)), 
        (*C.struct_AVDictionary)(unsafe.Pointer(src)), C.int(flags)))
}

/**
 * Free all the memory allocated for an AVDictionary struct
 * and all keys and values.
 */
func Av_dict_free(m **AVDictionary)  {
    C.av_dict_free((**C.struct_AVDictionary)(unsafe.Pointer(m)))
}

/**
 * Get dictionary entries as a string.
 *
 * Create a string containing dictionary's entries.
 * Such string may be passed back to av_dict_parse_string().
 * @note String is escaped with backslashes ('\').
 *
 * @param[in]  m             dictionary
 * @param[out] buffer        Pointer to buffer that will be allocated with string containg entries.
 *                           Buffer must be freed by the caller when is no longer needed.
 * @param[in]  key_val_sep   character used to separate key from value
 * @param[in]  pairs_sep     character used to separate two pairs from each other
 * @return                   >= 0 on success, negative on error
 * @warning Separators cannot be neither '\\' nor '\0'. They also cannot be the same.
 */
func Av_dict_get_string(m *AVDictionary, buffer **byte,
                       key_val_sep byte, pairs_sep byte) int32 {
    return int32(C.av_dict_get_string((*C.struct_AVDictionary)(unsafe.Pointer(m)), 
        (**C.char)(unsafe.Pointer(buffer)), C.char(key_val_sep), 
        C.char(pairs_sep)))
}

/**
 * @}
 */

                          
