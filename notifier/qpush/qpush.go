package qpush

import (
	"fmt"
	"github.com/CVEDB/notify/types"
	"github.com/CVEDB/notify/utils"
	"github.com/CVEDB/requests"
	"github.com/CVEDB/requests/ext"
)

const DefaultWebhook = "http://qpush.me/pusher/push_site/"

type Option struct {
	types.BaseOption `yaml:",inline"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Name string  `yaml:"name" dict:"name"`
	Code string  `yaml:"code" dict:"code"`
	Msg  *string `yaml:"msg,omitempty" dict:"msg[text],omitempty"`
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
	data := utils.StructToDict(n.MessageParams)
	return DefaultWebhook, ext.Data(data)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok {
		return nil
	}
	return fmt.Errorf("[QPush] [%v] %s", resp.StatusCode, resp.Content)
}
