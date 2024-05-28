package format

import (
	"test-func/pkg/vdk/av/avutil"
	"test-func/pkg/vdk/format/aac"
	"test-func/pkg/vdk/format/flv"
	"test-func/pkg/vdk/format/mp4"
	"test-func/pkg/vdk/format/rtmp"
	"test-func/pkg/vdk/format/rtsp"
	"test-func/pkg/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
