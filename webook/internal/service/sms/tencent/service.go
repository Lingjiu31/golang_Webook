package tencent

import (
	"context"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Service struct {
	appId    string
	signName string
	client   *sms.Client
}

func NewService(appId string, signName string, client *sms.Client) *Service {
	return &Service{
		appId:    appId,
		signName: signName,
		client:   client,
	}
}

func (s *Service) Send(ctx context.Context, tpl string, args []string, number []string) error {
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = common.StringPtr(s.appId)
	req.SignName = common.StringPtr(s.signName)
	req.TemplateId = common.StringPtr(tpl)
	req.TemplateParamSet = common.StringPtrs(args)
	req.PhoneNumberSet = common.StringPtrs(number)

	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "OK" {
			return fmt.Errorf("短信发送失败 %s, %s", *status.Code, *status.Message)
		}
	}
	return nil
}
