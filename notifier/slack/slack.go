package slack

import (
	"fmt"
	"github.com/CVEDB/notify/types"
	"github.com/CVEDB/notify/utils"
	"github.com/CVEDB/requests"
	"github.com/CVEDB/requests/ext"
)

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Text string `yaml:"text" json:"text"`
}

type notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *notifier {
	noticer := &notifier{}
	noticer.Option = opt
	return noticer
}

func (n *notifier) format(messages []string) (string, ext.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	json := utils.StructToJson(n.MessageParams)
	return n.Webhook, ext.Json(json)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Content == "ok" {
		return nil
	}
	return fmt.Errorf("[Slack] [%v] %s", resp.StatusCode, resp.Content)
}
