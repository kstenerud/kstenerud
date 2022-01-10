package main

import (
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

// ============================================================================

// https://shields.io/
// https://simpleicons.org/
// https://github.com/simple-icons/simple-icons/blob/develop/slugs.md

var color_language = "blue"
var color_format = "239120"
var color_technology = "2c2d72"
var color_platform = "black"

var languages []Badge = []Badge{
	NewNoLogoBadge("Assembler", "ASM", color_language),
	NewLogoBadge("Bash", "gnu-bash", color_language),
	NewLogoBadge("C", "c", color_language),
	NewLogoBadge("C++", "c++", color_language),
	NewLogoBadge("C#", "c-sharp", color_language),
	NewNoLogoBadge("COBOL", "CB", color_language),
	NewNoLogoBadge("Digital Logic", "DL", color_language),
	NewLogoBadge("Flutter", "flutter", color_language),
	NewLogoBadge("Go", "go", color_language),
	NewLogoBadge("Java", "java", color_language),
	NewLogoBadge("Kotlin", "kotlin", color_language),
	NewLogoBadge("LUA", "lua", color_language),
	NewNoLogoBadge("Objective-C", "OC", color_language),
	NewLogoBadge("Perl", "perl", color_language),
	NewLogoBadge("PHP", "php", color_language),
	NewLogoBadge("Python", "python", color_language),
	NewNoLogoBadge("SQL", "SQL", color_language),
	NewLogoBadge("Swift", "swift", color_language),
}

var formats []Badge = []Badge{
	NewLogoBadge("CSS", "css3", color_format),
	NewLogoBadge("HTML", "html5", color_format),
	NewLogoBadge("JSON", "json", color_format),
	NewLogoBadge("Markdown", "markdown", color_format),
	NewNoLogoBadge("XML", "<>", color_format),
	NewNoLogoBadge("YAML", "YML", color_format),
}

var technologies []Badge = []Badge{
	NewLogoBadge("Cocos", "cocos", color_technology),
	NewLogoBadge("Docker", "docker", color_technology),
	NewLogoBadge("Django", "django", color_technology),
	NewLogoBadge("Hibernate", "hibernate", color_technology),
	NewLogoBadge("MySQL", "mysql", color_technology),
	NewLogoBadge("NET", ".net", color_technology),
	NewLogoBadge("PostgreSQL", "postgresql", color_technology),
	NewLogoBadge("Spring", "spring", color_technology),
	NewLogoBadge("SQLite", "sqlite", color_technology),
	NewLogoBadge("Unity", "unity", color_technology),
}

var platforms []Badge = []Badge{
	NewLogoBadge("Android", "android", color_platform),
	NewLogoBadge("BSD", "freebsd", color_platform),
	NewLogoBadge("iOS", "ios", color_platform),
	NewLogoBadge("Linux", "linux", color_platform),
	NewLogoBadge("macOS", "macos", color_platform),
	NewLogoBadge("NixOS", "nixos", color_platform),
	NewLogoBadge("Windows", "windows", color_platform),
}

// ============================================================================

type Badge interface {
	String() string
	GetMessage() string
}

type NoLogoBadge struct {
	Message string
	Label   string
	Color   string
}

func NewNoLogoBadge(message string, label string, color string) *NoLogoBadge {
	return &NoLogoBadge{
		Message: message,
		Label:   label,
		Color:   color,
	}
}

func (_this *NoLogoBadge) String() string {
	return fmt.Sprintf("![%v](https://img.shields.io/static/v1?label=%v&message=%v&color=%v)",
		strings.ToUpper(_this.Message),
		url.QueryEscape(_this.Label),
		url.QueryEscape(strings.ToUpper(_this.Message)),
		url.QueryEscape(_this.Color))
}

func (_this *NoLogoBadge) GetMessage() string {
	return _this.Message
}

type LogoBadge struct {
	Message string
	Logo    string
	Color   string
}

func NewLogoBadge(message string, logo string, color string) *LogoBadge {
	return &LogoBadge{
		Message: message,
		Logo:    logo,
		Color:   color,
	}
}

func (_this *LogoBadge) GetMessage() string {
	return _this.Message
}

func (_this *LogoBadge) String() string {
	return fmt.Sprintf("![%v](https://img.shields.io/static/v1?label=%%7F&message=%v&color=%v&logo=%v&logoColor=white)",
		strings.ToUpper(_this.Message),
		url.QueryEscape(strings.ToUpper(_this.Message)),
		url.QueryEscape(_this.Color),
		url.QueryEscape(_this.Logo))
}

func sortBadges(badges []Badge) {
	sort.Slice(badges, func(i, j int) bool {
		return strings.Compare(strings.ToLower(badges[i].GetMessage()),
			strings.ToLower(badges[j].GetMessage())) < 0
	})
}

func getBadgeString(badges []Badge, index int) string {
	if index >= len(badges) {
		return ""
	}
	return badges[index].String()
}

func getHalfIndex(badges []Badge) int {
	length := len(badges)
	if length&1 == 1 {
		return length/2 + 1
	}
	return length / 2
}

func maxLength(slices ...interface{}) int {
	max := 0
	for _, s := range slices {
		length := reflect.ValueOf(s).Len()
		if length > max {
			max = length
		}
	}
	return max
}

// ============================================================================

func main() {
	sortBadges(languages)
	sortBadges(formats)
	sortBadges(technologies)
	sortBadges(platforms)

	lang1 := languages[:getHalfIndex(languages)]
	lang2 := languages[len(lang1):]

	tableLength := maxLength(lang1, lang2, formats, technologies, platforms)

	lang1 = languages[:tableLength]
	lang2 = languages[len(lang1):]

	fmt.Printf("| Languages | Languages | Technologies | Platforms | Formats |\n")
	fmt.Printf("| - | - | - | - | - |\n")

	for i := 0; i < tableLength; i++ {
		fmt.Printf("| %v | %v | %v | %v | %v |\n",
			getBadgeString(lang1, i),
			getBadgeString(lang2, i),
			getBadgeString(technologies, i),
			getBadgeString(platforms, i),
			getBadgeString(formats, i))
	}
}
