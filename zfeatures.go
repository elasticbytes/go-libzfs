package zfs

// #include <stdlib.h>
// #include <libzfs.h>
// #include "zfeatures.h"
// #include "common.h"
// #include "zpool.h"
// #include "zfs.h"
import "C"

const (
	ZFEATURE_STR_ASYNC_DESTROY string = "async_destroy"
	ZFEATURE_STR_EMPTY_BPOBJ = "empty_bpobj"
	ZFEATURE_STR_LZ4_COMPRESS = "lz4_compress"
	ZFEATURE_STR_BOOKMARKS = "bookmarks" // 0.6.4
	ZFEATURE_STR_EMBEDDED_DATA = "embedded_data"
	ZFEATURE_STR_ENABLED_TXG = "enabled_txg"
	ZFEATURE_STR_EXTENSIBLE_DATASET = "extensible_dataset"
	ZFEATURE_STR_HOLE_BIRTH = "hole_birth"
	ZFEATURE_STR_SPACEMAP_HISTOGRAM = "spacemap_histogram"
	ZFEATURE_STR_FILESYSTEM_LIMITS = "filesystem_limits" //0.6.5
	ZFEATURE_STR_LARGE_BLOCKS = "large_blocks"
	ZFEATURE_STR_EDONR = "edonr" //0.7.0
	ZFEATURE_STR_LARGE_DNODE = "large_dnode"
	ZFEATURE_STR_MULTI_VDEV_CRASH_DUMP = "multi_vdev_crash_dump"
	ZFEATURE_STR_SHA512 = "sha512"
	ZFEATURE_STR_SKEIN = "skein"
	ZFEATURE_STR_USEROBJ_ACCOUNTING = "userobj_accounting"
	ZFEATURE_STR_ALLOCATION_CLASSES = "allocation_classes" // 0.8.0
	ZFEATURE_STR_BOOKMARK_V2 = "bookmark_v2"
	ZFEATURE_STR_DEVICE_REMOVAL = "device_removal"
	ZFEATURE_STR_ENCRYPTION = "encryption"
	ZFEATURE_STR_OBSOLETE_COUNTS = "obsolete_counts"
	ZFEATURE_STR_PROJECT_QUOTA = "project_quota"
	ZFEATURE_STR_RESILVER_DEFER = "resilver_defer"
	ZFEATURE_STR_SPACEMAP_V2 = "spacemap_v2"
	ZFEATURE_STR_ZPOOL_CHECKPOINT = "zpool_checkpoint"
)

var featuresTable []string

func init() {
	featuresTable = make([]string, C.SPA_FEATURES)
	for i := 0; i < int(C.SPA_FEATURES); i++ {
		name := C.GoString(C.zfeatures_get_name(C.int(i)))
		featuresTable[i] = name
	}
}
