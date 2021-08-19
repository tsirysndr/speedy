package speedy

import "fmt"

type Provider string

const (
	OOKLA   Provider = "ookla"
	NETFLIX Provider = "netflix"
)

func Start(p Provider) (*Result, error) {
	switch p {
	case OOKLA:
		return StartOoklaTest()
	case NETFLIX:
		return StartNetflixTest()
	default:
		return nil, fmt.Errorf("unknown provider %s", p)
	}
}
