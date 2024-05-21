package alipay

// BillDownloadURLQuery 查询对账单下载地址接口 https://docs.open.alipay.com/api_15/alipay.data.dataservice.bill.downloadurl.query
func (c *Client) BillDownloadURLQuery(param BillDownloadURLQuery) (result *BillDownloadURLQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// BillBalanceQuery 支付宝商家账户当前余额查询接口 https://opendocs.alipay.com/apis/api_15/alipay.data.bill.balance.query
func (c *Client) BillBalanceQuery(param BillBalanceQuery) (result *BillBalanceQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// BillAccountLogQuery 查询账户账务明细接口请求参数 https://opendocs.alipay.com/apis/api_15/alipay.data.bill.accountlog.query
func (c *Client) BillAccountLogQuery(param BillAccountLogQuery) (result *BillAccountLogQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}
