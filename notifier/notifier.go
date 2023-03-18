package notifier

import (
	"github.com/CVEDB/notify/notifier/bark"
	"github.com/CVEDB/notify/notifier/chanify"
	"github.com/CVEDB/notify/notifier/dingtalk"
	"github.com/CVEDB/notify/notifier/discord"
	"github.com/CVEDB/notify/notifier/feishu"
	"github.com/CVEDB/notify/notifier/ftqq"
	"github.com/CVEDB/notify/notifier/gitter"
	"github.com/CVEDB/notify/notifier/googlechat"
	"github.com/CVEDB/notify/notifier/igot"
	"github.com/CVEDB/notify/notifier/mailgun"
	"github.com/CVEDB/notify/notifier/pushbullet"
	"github.com/CVEDB/notify/notifier/pushdeer"
	"github.com/CVEDB/notify/notifier/pushover"
	"github.com/CVEDB/notify/notifier/qpush"
	"github.com/CVEDB/notify/notifier/rocketchat"
	"github.com/CVEDB/notify/notifier/showdoc"
	"github.com/CVEDB/notify/notifier/slack"
	"github.com/CVEDB/notify/notifier/telegram"
	"github.com/CVEDB/notify/notifier/webhook"
	"github.com/CVEDB/notify/notifier/xz"
	"github.com/CVEDB/notify/notifier/zulip"
)

type NotifiesPackage struct {
	Bark       []*bark.Option       `yaml:"bark,omitempty"`
	Chanify    []*chanify.Option    `yaml:"chanify,omitempty"`
	Dingtalk   []*dingtalk.Option   `yaml:"dingtalk,omitempty"`
	Discord    []*discord.Option    `yaml:"discord,omitempty"`
	FeiShu     []*feishu.Option     `yaml:"feishu,omitempty"`
	FTQQ       []*ftqq.Option       `yaml:"ftqq,omitempty"`
	Gitter     []*gitter.Option     `yaml:"gitter,omitempty"`
	GoogleChat []*googlechat.Option `yaml:"googlechat,omitempty"`
	IGot       []*igot.Option       `yaml:"igot,omitempty"`
	Mailgun    []*mailgun.Option    `yaml:"mailgun,omitempty"`
	PushBullet []*pushbullet.Option `yaml:"pushbullet,omitempty"`
	PushDeer   []*pushdeer.Option   `yaml:"pushdeer,omitempty"`
	PushOver   []*pushover.Option   `yaml:"pushover,omitempty"`
	QPush      []*qpush.Option      `yaml:"qpush,omitempty"`
	RocketChat []*rocketchat.Option `yaml:"rocketchat,omitempty"`
	Slack      []*slack.Option      `yaml:"slack,omitempty"`
	ShowDoc    []*showdoc.Option    `yaml:"showdoc,omitempty"`
	Telegram   []*telegram.Option   `yaml:"telegram,omitempty"`
	XZ         []*xz.Option         `yaml:"xz,omitempty"`
	Webhook    []*webhook.Option    `yaml:"webhook,omitempty"`
	ZuLip      []*zulip.Option      `yaml:"zulip,omitempty"`

	// 计划中
	// Sendgraid  []*sendgraid.Option  `yaml:"sendgraid,omitempty"`
	// Mattermost []*mattermost.Option `yaml:"mattermost,omitempty" //https://api.mattermost.com/
}
