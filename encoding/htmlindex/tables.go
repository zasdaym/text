// Code generated by running "go generate" in github.com/zasdaym/text. DO NOT EDIT.

package htmlindex

type htmlEncoding byte

const (
	utf8 htmlEncoding = iota
	ibm866
	iso8859_2
	iso8859_3
	iso8859_4
	iso8859_5
	iso8859_6
	iso8859_7
	iso8859_8
	iso8859_8I
	iso8859_10
	iso8859_13
	iso8859_14
	iso8859_15
	iso8859_16
	koi8r
	koi8u
	macintosh
	windows874
	windows1250
	windows1251
	windows1252
	windows1253
	windows1254
	windows1255
	windows1256
	windows1257
	windows1258
	macintoshCyrillic
	gbk
	gb18030
	big5
	eucjp
	iso2022jp
	shiftJIS
	euckr
	replacement
	utf16be
	utf16le
	xUserDefined
	numEncodings
)

var canonical = [numEncodings]string{
	"utf-8",
	"ibm866",
	"iso-8859-2",
	"iso-8859-3",
	"iso-8859-4",
	"iso-8859-5",
	"iso-8859-6",
	"iso-8859-7",
	"iso-8859-8",
	"iso-8859-8-i",
	"iso-8859-10",
	"iso-8859-13",
	"iso-8859-14",
	"iso-8859-15",
	"iso-8859-16",
	"koi8-r",
	"koi8-u",
	"macintosh",
	"windows-874",
	"windows-1250",
	"windows-1251",
	"windows-1252",
	"windows-1253",
	"windows-1254",
	"windows-1255",
	"windows-1256",
	"windows-1257",
	"windows-1258",
	"x-mac-cyrillic",
	"gbk",
	"gb18030",
	"big5",
	"euc-jp",
	"iso-2022-jp",
	"shift_jis",
	"euc-kr",
	"replacement",
	"utf-16be",
	"utf-16le",
	"x-user-defined",
}

var nameMap = map[string]htmlEncoding{
	"unicode-1-1-utf-8":   utf8,
	"utf-8":               utf8,
	"utf8":                utf8,
	"866":                 ibm866,
	"cp866":               ibm866,
	"csibm866":            ibm866,
	"ibm866":              ibm866,
	"csisolatin2":         iso8859_2,
	"iso-8859-2":          iso8859_2,
	"iso-ir-101":          iso8859_2,
	"iso8859-2":           iso8859_2,
	"iso88592":            iso8859_2,
	"iso_8859-2":          iso8859_2,
	"iso_8859-2:1987":     iso8859_2,
	"l2":                  iso8859_2,
	"latin2":              iso8859_2,
	"csisolatin3":         iso8859_3,
	"iso-8859-3":          iso8859_3,
	"iso-ir-109":          iso8859_3,
	"iso8859-3":           iso8859_3,
	"iso88593":            iso8859_3,
	"iso_8859-3":          iso8859_3,
	"iso_8859-3:1988":     iso8859_3,
	"l3":                  iso8859_3,
	"latin3":              iso8859_3,
	"csisolatin4":         iso8859_4,
	"iso-8859-4":          iso8859_4,
	"iso-ir-110":          iso8859_4,
	"iso8859-4":           iso8859_4,
	"iso88594":            iso8859_4,
	"iso_8859-4":          iso8859_4,
	"iso_8859-4:1988":     iso8859_4,
	"l4":                  iso8859_4,
	"latin4":              iso8859_4,
	"csisolatincyrillic":  iso8859_5,
	"cyrillic":            iso8859_5,
	"iso-8859-5":          iso8859_5,
	"iso-ir-144":          iso8859_5,
	"iso8859-5":           iso8859_5,
	"iso88595":            iso8859_5,
	"iso_8859-5":          iso8859_5,
	"iso_8859-5:1988":     iso8859_5,
	"arabic":              iso8859_6,
	"asmo-708":            iso8859_6,
	"csiso88596e":         iso8859_6,
	"csiso88596i":         iso8859_6,
	"csisolatinarabic":    iso8859_6,
	"ecma-114":            iso8859_6,
	"iso-8859-6":          iso8859_6,
	"iso-8859-6-e":        iso8859_6,
	"iso-8859-6-i":        iso8859_6,
	"iso-ir-127":          iso8859_6,
	"iso8859-6":           iso8859_6,
	"iso88596":            iso8859_6,
	"iso_8859-6":          iso8859_6,
	"iso_8859-6:1987":     iso8859_6,
	"csisolatingreek":     iso8859_7,
	"ecma-118":            iso8859_7,
	"elot_928":            iso8859_7,
	"greek":               iso8859_7,
	"greek8":              iso8859_7,
	"iso-8859-7":          iso8859_7,
	"iso-ir-126":          iso8859_7,
	"iso8859-7":           iso8859_7,
	"iso88597":            iso8859_7,
	"iso_8859-7":          iso8859_7,
	"iso_8859-7:1987":     iso8859_7,
	"sun_eu_greek":        iso8859_7,
	"csiso88598e":         iso8859_8,
	"csisolatinhebrew":    iso8859_8,
	"hebrew":              iso8859_8,
	"iso-8859-8":          iso8859_8,
	"iso-8859-8-e":        iso8859_8,
	"iso-ir-138":          iso8859_8,
	"iso8859-8":           iso8859_8,
	"iso88598":            iso8859_8,
	"iso_8859-8":          iso8859_8,
	"iso_8859-8:1988":     iso8859_8,
	"visual":              iso8859_8,
	"csiso88598i":         iso8859_8I,
	"iso-8859-8-i":        iso8859_8I,
	"logical":             iso8859_8I,
	"csisolatin6":         iso8859_10,
	"iso-8859-10":         iso8859_10,
	"iso-ir-157":          iso8859_10,
	"iso8859-10":          iso8859_10,
	"iso885910":           iso8859_10,
	"l6":                  iso8859_10,
	"latin6":              iso8859_10,
	"iso-8859-13":         iso8859_13,
	"iso8859-13":          iso8859_13,
	"iso885913":           iso8859_13,
	"iso-8859-14":         iso8859_14,
	"iso8859-14":          iso8859_14,
	"iso885914":           iso8859_14,
	"csisolatin9":         iso8859_15,
	"iso-8859-15":         iso8859_15,
	"iso8859-15":          iso8859_15,
	"iso885915":           iso8859_15,
	"iso_8859-15":         iso8859_15,
	"l9":                  iso8859_15,
	"iso-8859-16":         iso8859_16,
	"cskoi8r":             koi8r,
	"koi":                 koi8r,
	"koi8":                koi8r,
	"koi8-r":              koi8r,
	"koi8_r":              koi8r,
	"koi8-ru":             koi8u,
	"koi8-u":              koi8u,
	"csmacintosh":         macintosh,
	"mac":                 macintosh,
	"macintosh":           macintosh,
	"x-mac-roman":         macintosh,
	"dos-874":             windows874,
	"iso-8859-11":         windows874,
	"iso8859-11":          windows874,
	"iso885911":           windows874,
	"tis-620":             windows874,
	"windows-874":         windows874,
	"cp1250":              windows1250,
	"windows-1250":        windows1250,
	"x-cp1250":            windows1250,
	"cp1251":              windows1251,
	"windows-1251":        windows1251,
	"x-cp1251":            windows1251,
	"ansi_x3.4-1968":      windows1252,
	"ascii":               windows1252,
	"cp1252":              windows1252,
	"cp819":               windows1252,
	"csisolatin1":         windows1252,
	"ibm819":              windows1252,
	"iso-8859-1":          windows1252,
	"iso-ir-100":          windows1252,
	"iso8859-1":           windows1252,
	"iso88591":            windows1252,
	"iso_8859-1":          windows1252,
	"iso_8859-1:1987":     windows1252,
	"l1":                  windows1252,
	"latin1":              windows1252,
	"us-ascii":            windows1252,
	"windows-1252":        windows1252,
	"x-cp1252":            windows1252,
	"cp1253":              windows1253,
	"windows-1253":        windows1253,
	"x-cp1253":            windows1253,
	"cp1254":              windows1254,
	"csisolatin5":         windows1254,
	"iso-8859-9":          windows1254,
	"iso-ir-148":          windows1254,
	"iso8859-9":           windows1254,
	"iso88599":            windows1254,
	"iso_8859-9":          windows1254,
	"iso_8859-9:1989":     windows1254,
	"l5":                  windows1254,
	"latin5":              windows1254,
	"windows-1254":        windows1254,
	"x-cp1254":            windows1254,
	"cp1255":              windows1255,
	"windows-1255":        windows1255,
	"x-cp1255":            windows1255,
	"cp1256":              windows1256,
	"windows-1256":        windows1256,
	"x-cp1256":            windows1256,
	"cp1257":              windows1257,
	"windows-1257":        windows1257,
	"x-cp1257":            windows1257,
	"cp1258":              windows1258,
	"windows-1258":        windows1258,
	"x-cp1258":            windows1258,
	"x-mac-cyrillic":      macintoshCyrillic,
	"x-mac-ukrainian":     macintoshCyrillic,
	"chinese":             gbk,
	"csgb2312":            gbk,
	"csiso58gb231280":     gbk,
	"gb2312":              gbk,
	"gb_2312":             gbk,
	"gb_2312-80":          gbk,
	"gbk":                 gbk,
	"iso-ir-58":           gbk,
	"x-gbk":               gbk,
	"gb18030":             gb18030,
	"big5":                big5,
	"big5-hkscs":          big5,
	"cn-big5":             big5,
	"csbig5":              big5,
	"x-x-big5":            big5,
	"cseucpkdfmtjapanese": eucjp,
	"euc-jp":              eucjp,
	"x-euc-jp":            eucjp,
	"csiso2022jp":         iso2022jp,
	"iso-2022-jp":         iso2022jp,
	"csshiftjis":          shiftJIS,
	"ms932":               shiftJIS,
	"ms_kanji":            shiftJIS,
	"shift-jis":           shiftJIS,
	"shift_jis":           shiftJIS,
	"sjis":                shiftJIS,
	"windows-31j":         shiftJIS,
	"x-sjis":              shiftJIS,
	"cseuckr":             euckr,
	"csksc56011987":       euckr,
	"euc-kr":              euckr,
	"iso-ir-149":          euckr,
	"korean":              euckr,
	"ks_c_5601-1987":      euckr,
	"ks_c_5601-1989":      euckr,
	"ksc5601":             euckr,
	"ksc_5601":            euckr,
	"windows-949":         euckr,
	"csiso2022kr":         replacement,
	"hz-gb-2312":          replacement,
	"iso-2022-cn":         replacement,
	"iso-2022-cn-ext":     replacement,
	"iso-2022-kr":         replacement,
	"replacement":         replacement,
	"utf-16be":            utf16be,
	"utf-16":              utf16le,
	"utf-16le":            utf16le,
	"x-user-defined":      xUserDefined,
}

var localeMap = []htmlEncoding{
	windows1252, // und_Latn
	windows1256, // ar
	windows1251, // ba
	windows1251, // be
	windows1251, // bg
	windows1250, // cs
	iso8859_7,   // el
	windows1257, // et
	windows1256, // fa
	windows1255, // he
	windows1250, // hr
	iso8859_2,   // hu
	shiftJIS,    // ja
	windows1251, // kk
	euckr,       // ko
	windows1254, // ku
	windows1251, // ky
	windows1257, // lt
	windows1257, // lv
	windows1251, // mk
	iso8859_2,   // pl
	windows1251, // ru
	windows1251, // sah
	windows1250, // sk
	iso8859_2,   // sl
	windows1251, // sr
	windows1251, // tg
	windows874,  // th
	windows1254, // tr
	windows1251, // tt
	windows1251, // uk
	windows1258, // vi
	gb18030,     // zh-hans
	big5,        // zh-hant
}

const locales = "und_Latn ar ba be bg cs el et fa he hr hu ja kk ko ku ky lt lv mk pl ru sah sk sl sr tg th tr tt uk vi zh-hans zh-hant"