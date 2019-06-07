package webhook_helper

func deepcopy_webhook(wh *Webhook) *Webhook {
	return &Webhook{
		Id:  wh.Id,
		Url: wh.Url,
	}
}
