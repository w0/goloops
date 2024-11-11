package audiocontent

import (
	"fmt"
	"io"

	"github.com/groob/plist"
)

const (
	acBaseUrl  = "https://audiocontentdownload.apple.com"
	acFilePath = "/lp10_ms3_content_2016/"
)

type AudioContent struct {
	ConfigVersion string `plist:"ConfigVersion"`
	/*Content       map[string][]struct {
	Name     string   `plist:"Name"`
	Packages []string `plist:"Packages"`
	Locale   []struct {
		Description string `plist:"Description"`
		DisplayName string `plist:"DisplayName"`
		} `plist:"_LOCALIZABLE_"`
		} `plist:"Content"`*/
	Packages map[string]struct {
		ContainsAppleLoops                  bool   `plist:"ContainsAppleLoops"`
		ContainsGarageBandLegacyInstruments bool   `plist:"ContainsGarageBandLegacyInstruments"`
		DownloadName                        string `plist:"DownloadName"`
		DownloadSize                        any    `plist:"DownloadSize"`
		FileCheck                           any    `plist:"FileCheck"`
		InstalledSize                       any    `plist:"InstalledSize"`
		IsMandatory                         bool   `plist:"IsMandatory"`
		PackageID                           string `plist:"PackageID"`
	} `plist:"Packages"`
}

func (ac *AudioContent) GetDownloadUrl(PackageID string) string {
	return fmt.Sprintf("%s%s%s", acBaseUrl, acFilePath, ac.Packages[PackageID].DownloadName)
}

func NewAudioContent(PlistPath io.Reader) (AudioContent, error) {

	d := plist.NewDecoder(PlistPath)

	var ac AudioContent

	err := d.Decode(&ac)

	if err != nil {
		return AudioContent{}, err
	}

	return ac, nil
}

func (ac *AudioContent) ListMandatory() {
	fmt.Printf("Mandatory Audio Content:\n")
	count := 1
	for k, v := range ac.Packages {
		if v.IsMandatory == true {
			fmt.Printf("\t%d. %s\n", count, k)
			count++
		}
	}
}

func (ac *AudioContent) GetMandatory() {
	for k, v := range ac.Packages {
		if v.IsMandatory == true {
			fmt.Println(ac.GetDownloadUrl(k))
		}
	}
}