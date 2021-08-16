package scam

import (
	"regexp"

	"github.com/otiai10/gosseract/v2"
)

var (
	rex = []*regexp.Regexp{
		regexp.MustCompile(`(?im)[\S\s]*believe[\S\s]*that[\S\s]*blockchain[\S\s]*and[\S\s]*bitcoin[\S\s]*wallet[\S\s]*will[\S\s]*make[\S\s]*the[\S\s]*world[\S\s]*more[\S\s]*advanced[\S\s]*and[\S\s]*more[\S\s]*comfortable[\S\s]*for[\S\s]*peoples[\S\s]*lives`),
		regexp.MustCompile(`(?im)[\S\s]*believe[\S\s]*that[\S\s]*blockchain[\S\s]*and[\S\s]*bitcoin[\S\s]*wallet[\S\s]*`),
		regexp.MustCompile(`(?im)[\S\s]*To[\S\s]*speed[\S\s]*up[\S\s]*the[\S\s]*process[\S\s]*of[\S\s]*mass`),
		regexp.MustCompile(`(?im)[\S\s]*Musk[\S\s]*`),
	}
)

// IsScam returns whether an image is detected as a scam.
// The function mainly matches against the twitter scam with the fake elon musk account giving away bitcoins
func IsScam(image []byte) (bool, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImageFromBytes(image)

	text, err := client.Text()
	if err != nil {
		return false, err
	}

	var matches int
	for _, re := range rex {
		if re.MatchString(text) {
			matches++

		}
	}

	if matches >= 2 {
		return true, nil
	}

	return false, nil
}
