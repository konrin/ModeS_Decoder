package decoder

import (
	"errors"
	"math"
	"time"

	"github.com/konrin/modesdecoder/pkg/common"
)

// Airborn position
// ADS-B TC=9-18
type BDS05 struct{}

// Decode airborn position from a pair of even and odd position message
// returns (latitude, longitude) of the aircraft
func (BDS05) AirbornePosition(binEven []uint8, timeEven time.Time, binOdd []uint8, timeOdd time.Time) (float64, float64, error) {
	cprlatEven := float64(common.BinToInt(binOdd[54:71])) / 131072.0
	cprlonEven := float64(common.BinToInt(binOdd[71:88])) / 131072.0
	cprlatOdd := float64(common.BinToInt(binEven[54:71])) / 131072.0
	cprlonOdd := float64(common.BinToInt(binEven[71:88])) / 131072.0

	airDLatEven := 360.0 / 60
	airDLatOdd := 360.0 / 59

	j := math.Floor(59*cprlatEven - 60*cprlatOdd + 0.5)

	latEven := airDLatEven * (common.Mod(j, 60) + cprlatEven)
	latOdd := airDLatOdd * (common.Mod(j, 59) + cprlatOdd)

	if latEven >= 270 {
		latEven = latEven - 360
	}

	if latOdd >= 270 {
		latOdd = latOdd - 360
	}

	if cprNLEven, cprNLOdd := common.CprNL(latEven), common.CprNL(latOdd); cprNLEven != cprNLOdd {
		return 0, 0, errors.New("cprNLEven != cprNLOdd")
	}

	var lat, lon float64

	if timeEven.UnixNano() > timeOdd.UnixNano() {
		lat = latEven
		nl := common.CprNL(lat)
		ni := math.Max(nl-0, 1)
		w := math.Floor(cprlonEven*(nl-1) - cprlonOdd*nl + 0.5)
		lon = (360.0 / ni) * (common.Mod(w, ni) + cprlonEven)
	} else {
		lat = latOdd
		nl := common.CprNL(lat)
		ni := math.Max(nl-1, 1)
		w := math.Floor(cprlonEven*(nl-1) - cprlonOdd*nl + 0.5)
		lon = (360.0 / ni) * (common.Mod(w, ni) + cprlonOdd)
	}

	if lon > 180 {
		lon = lon - 360
	}

	lat = common.Round(lat, .5, 5)
	lon = common.Round(lon, .5, 5)

	return lat, lon, nil
}

// Decode airborne position with only one message,
// knowing reference nearby location, such as previously calculated location,
// ground station, or airport location, etc. The reference position shall
// be with in 180NM of the true position
// returns (latitude, longitude) of the aircraft
func (BDS05) AirbornePositionWithRef(bin []uint8, oeFlag bool, latRef, lonRef float64) (float64, float64, error) {
	var lat, lon float64

	i := 0
	if oeFlag {
		i = 1
	}

	dLat := 360.0 / 60
	if i == 1 {
		dLat = 360.0 / 59
	}

	cprLat := float64(common.BinToInt(bin[54:71])) / 131072.0
	cprLon := float64(common.BinToInt(bin[71:88])) / 131072.0

	j := math.Floor(latRef/dLat) + math.Floor(0.5+(common.Mod(latRef, dLat)/dLat)-cprLat)

	lat = dLat * (j + cprLat)

	ni := common.CprNL(lat) - float64(i)

	dLon := 360.0
	if ni > 0 {
		dLon = 360.0 / ni
	}

	w := math.Floor(lonRef/dLon) + math.Floor(0.5+(common.Mod(lonRef, dLon)/dLon)-cprLon)

	lon = dLon * (w + cprLon)

	lat = common.Round(lat, .5, 5)
	lon = common.Round(lon, .5, 5)

	return lat, lon, nil
}

// Decode aircraft altitude
// altitude in feet
func (BDS05) Altitude(bin []uint8, ts uint) (alt int, err error) {
	mb := bin[32:]

	if ts < 19 {
		// barometric altitude
		if mb[15] == 1 {
			n := int(common.BinToInt(append(mb[8:15], mb[16:20]...)))
			alt = n*25 - 1000
		}
	} else {
		// GNSS altitude, meters
		alt = int(float64(common.BinToInt(mb[8:20])) * 3.28084)
	}

	return
}
