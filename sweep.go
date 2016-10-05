package ysok

import (
	"os"
	"time"

	"github.com/nlopes/slack"
)

var (
	// CmdSweep is definition of sweep subcommand.
	CmdSweep = &Command{
		Run:       runSweep,
		UsageLine: "sweep ",
		Short:     "過去のファイル削除",
		Long: `過去のファイルを削除します。

Options:
    -t YOUR_SLACK_TOKEN, --token YOUR_SLACK_TOKEN
        Slack のトークンを指定します。指定しない場合、環境変数 YSOK_TOKEN が使用されます。

    -u YOUR_SLACK_USER_ID, --user YOUR_SLACK_USER_ID
        Slack のユーザーIDを指定します。指定しない場合、環境変数 YSOK_USER が使用されます。

    -d DAYS_AGO, --days DAYS_AGO
        指定された日数以前のファイルを削除します。デフォルトは 30。7 以上の数字を指定して下さい。
	`,
	}
	token string
	days  int
	user  string
)

func init() {
	CmdSweep.Flag.StringVar(&token, "token", "", "Your slack token")
	CmdSweep.Flag.StringVar(&token, "t", "", "Your slack token")
	CmdSweep.Flag.StringVar(&user, "user", "", "Your slack user id")
	CmdSweep.Flag.StringVar(&user, "u", "", "Your slack user id")
	CmdSweep.Flag.IntVar(&days, "days", 30, "Days ago")
	CmdSweep.Flag.IntVar(&days, "d", 30, "Days ago")
}

// runSweep executes sweep command and return exit code.
func runSweep(args []string) int {
	u := getUser()
	if u == "" {
		errf("ユーザーIDを読み込むことが出来ませんでした")
		return ErrNoUser
	}
	t := getToken()
	if t == "" {
		errf("トークンを読み込むことが出来ませんでした")
		return ErrNoToken
	}
	if days < 7 {
		errf("日数は7以上を指定して下さい")
		return ErrInvalidDays
	}
	s := slack.New(getToken())
	cnt, err := getFileCount(s)
	if err != nil {
		errf(err.Error())
		return ErrGetFileCount
	}

	files, err := getFiles(s, cnt)
	if err != nil {
		errf(err.Error())
		return ErrGetFiles
	}

	threshold := time.Now().AddDate(0, 0, days*-1)
	for _, f := range files {
		if f.Created.Time().Before(threshold) {
			rmFile(s, f)
		}
	}

	return 0
}

func getUser() string {
	if user != "" {
		return user
	}

	return os.Getenv("YSOK_USER")
}

func getToken() string {
	if token != "" {
		return token
	}

	return os.Getenv("YSOK_TOKEN")
}

func getFileCount(s *slack.Client) (int, error) {
	p := slack.GetFilesParameters{User: getUser()}
	_, paging, err := s.GetFiles(p)
	if err != nil {
		return 0, err
	}
	return paging.Total, nil
}

func getFiles(s *slack.Client, cnt int) ([]slack.File, error) {
	p := slack.GetFilesParameters{User: getUser(), Count: cnt}
	files, _, err := s.GetFiles(p)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func rmFile(s *slack.Client, f slack.File) {
	err := s.DeleteFile(f.ID)
	result := "success"
	if err != nil {
		result = "failure"
		errf(err.Error())
	}
	outf("Delete %v[%v](%v) -----> %v", f.Name, f.ID, f.Created.Time(), result)
}
