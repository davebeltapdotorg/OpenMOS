package service

import (
	"math"
	"strconv"
	"strings"

	"airshift/openmos/internal/model"
	"airshift/openmos/internal/xml"
)

// MOS item durations are sample counts. A missing or invalid rate leaves seconds unknown.
func resolveItemTiming(itemEdDur, objDur, objTB string) (samples int, haveSamples bool, seconds int, timeBase int, haveRate bool) {
	samples, haveSamples = parseSamples(itemEdDur)
	if !haveSamples {
		samples, haveSamples = parseSamples(objDur)
	}
	rate, err := strconv.ParseFloat(strings.TrimSpace(objTB), 64)
	maxInt := float64(int(^uint(0) >> 1))
	if err != nil || rate <= 0 || math.IsNaN(rate) || math.IsInf(rate, 0) || rate > maxInt {
		return samples, haveSamples, 0, 0, false
	}
	timeBase = int(math.Round(rate))
	if !haveSamples {
		return samples, false, 0, timeBase, true
	}
	wholeSeconds := math.Round(float64(samples) / rate)
	if wholeSeconds > maxInt {
		return samples, true, 0, timeBase, false
	}
	return samples, true, int(wholeSeconds), timeBase, true
}

func parseSamples(raw string) (int, bool) {
	raw = strings.TrimSpace(raw)
	base := 10
	if strings.HasPrefix(strings.ToLower(raw), "0x") {
		raw, base = raw[2:], 16
	} else if strings.HasPrefix(raw, "x") {
		raw, base = raw[1:], 16
	}
	n, err := strconv.ParseInt(raw, base, strconv.IntSize)
	return int(n), err == nil && n >= 0
}

func mediaPathsFrom(paths *xml.ObjPaths, bare string) *model.MediaPaths {
	out := &model.MediaPaths{}
	if paths != nil {
		out.Essence = convertPaths(paths.Essence)
		out.Proxy = convertPaths(paths.Proxy)
		out.Metadata = convertPaths(paths.Metadata)
	}
	if bare = strings.TrimSpace(bare); bare != "" {
		out.Essence = append(out.Essence, model.MediaPath{URL: bare})
	}
	if len(out.Essence)+len(out.Proxy)+len(out.Metadata) == 0 {
		return nil
	}
	return out
}

func convertPaths(paths []xml.ObjPath) []model.MediaPath {
	var out []model.MediaPath
	for _, path := range paths {
		if value := strings.TrimSpace(path.Value); value != "" {
			out = append(out, model.MediaPath{URL: value, TechDescription: strings.TrimSpace(path.TechDescription)})
		}
	}
	return out
}
