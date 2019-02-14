package response

import "encoding/json"

//taobao.tbk.dg.item.coupon.get( 好券清单API【导购】 )
type TbkDgItemCouponGetResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkDgItemCouponGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
