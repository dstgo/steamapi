package steam

// PlatformSupportedLanguages see https://partner.steamgames.com/doc/store/localization/languages
type PlatformSupportedLanguages string

const (
	arabic     PlatformSupportedLanguages = "ar"
	bulgarian                             = "bg"
	schinese                              = "zh-CN"
	tchinese                              = "zh-TW"
	czech                                 = "cs"
	danish                                = "da"
	dutch                                 = "nl"
	english                               = "en"
	finnish                               = "fi"
	french                                = "fr"
	german                                = "de"
	greek                                 = "el"
	hungarian                             = "hu"
	italian                               = "it"
	japanese                              = "ja"
	koreana                               = "ko"
	norwegian                             = "no"
	polish                                = "pl"
	portuguese                            = "pt"
	brazilian                             = "pt-BR"
	romanian                              = "ro"
	russian                               = "ru"
	spanish                               = "es"
	latam                                 = "es-419"
	swedish                               = "sv"
	thai                                  = "th"
	turkish                               = "tr"
	ukrainian                             = "uk"
	vietnamese                            = "vn"
)

type LanguageCode = uint

const (
	LangEnglish LanguageCode = iota
	LangGerman
	LangFrench
	LangItalian
	LangKorea
	LangSpanish
	LangSimplifiedChinese
	LangTraditionalChinese
	LangRussian
	LangThai
	LangJapanese
	LangPortugal
	LangPoland
	LangDenmark
	LangNetherlands
	LangFinland
	LangNorway
	LangSweden
	LangHungary
	LangCzech
	LangRomania
	LangTurkey
	LangPortugal1
	LangBulgaria
	LangGreek
	LangUkrainian
	LangUnknown
	LangSpanish1
	LangVietnamese
)
