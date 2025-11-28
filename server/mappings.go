package server

const (
	UnlabeledRegion            = "Unlabeled Region"
	UndefinedDescription       = "This dive site is missing a description."
	PrefixForTagsInDescription = "tags:"
	RegionTagPrefix            = "_region_"
)

var CylinderTypeMappings = map[string]string{
	"AL100": "aluminium",
	"HP100": "steel",
	"HP130": "steel",
}

var SpecialTagValueMappings = map[string]string{
	"europe":        "Europe",
	"asia":          "Asia",
	"north-america": "North America",
	"atlantic":      "Atlantic Ocean",
	"indian":        "Indian Ocean",
	"pacific":       "Pacific Ocean",
	"mediterranean": "Mediterranean Sea",
	"red-sea":       "Red Sea",
}

var AwardMappings = map[string]string{
	"1st-dive":              "First dive!",
	"1st-seawater-dive":     "First seawater dive!",
	"1st-shark-encounter":   "First shark encounter!",
	"1st-night-dive":        "First night dive!",
	"1st-30m-dive":          "First 30m dive!",
	"1st-40m-dive":          "First 40m dive!",
	"1st-wreck-dive":        "First wreck dive!",
	"1st-wreck-penetration": "First wreck penetration dive!",
	"cert-owd":              "OWD diver! (CMAS)",
	"cert-aowd-nitrox":      "AOWD diver! Nitrox specialty diver! (SSI)",
	"cert-navigation":       "Navigation specialty diver! (SSI)",
	"cert-dry":              "Dry suit specialty diver! (SSI)",
	"cert-deep":             "Deep specialty diver! (PADI)",
	"cert-wreck":            "Wreck specialty diver! (PADI)",
	"100th-dive":            "100th dive!",
}
