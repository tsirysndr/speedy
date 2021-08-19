package speedy

import (
	"github.com/showwin/speedtest-go/speedtest"
	"go.uber.org/multierr"
)

func StartOoklaTest() (*Result, error) {
	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})
	s := targets[0]
	err := multierr.Combine(
		s.DownloadTest(false),
		s.UploadTest(false),
	)
	if err != nil {
		return nil, err
	}
	return &Result{
		Download: s.DLSpeed,
		Upload:   s.ULSpeed,
	}, nil
}
