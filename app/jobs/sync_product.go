package jobs

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.me/selling-partner-api-sdk/finances"
	"gopkg.me/selling-partner-api-sdk/ordersV0"
	sp "gopkg.me/selling-partner-api-sdk/pkg/selling-partner"
	"gopkg.me/selling-partner-api-sdk/reports"
)

func main() {
	// Fuck()
	s := &SyncReport{}
	s.SyncFinance()
	// s.SyncOrders()
}

// 需要将定义的struct 添加到字典中；
// 字典 key 可以配置到 自动任务 调用目标 中；
// func InitJob() {
// 	jobList = map[string]JobsExec{
// 		"ExamplesOne": ExamplesOne{},
// 		// ...
// 	}
// }

// 新添加的job 必须按照以下格式定义，并实现Exec函数
type SyncReport struct {
}

func (s SyncReport) SyncFinance() error {

	sellingPartner, err := sp.NewSellingPartner(&sp.Config{
		ClientID:     "amzn1.application-oa2-client.0c73f7463a85441f9ea68bbcb59702af",
		ClientSecret: "2bc5c7bed95d3891bf7b7562c7eb8497c307f8beebecc628614bfcde0d81a3b5",
		// RefreshToken: "Atzr|IwEBIFgMrY3c7wGNme00P-GXaiH1LEIkKrjl5pYKACuKjsC7qRqQefKlnv5gaTv9sXfM_-eXpIQBIZiqgh8f7wer0cZVPxqcGyrpAJVPj5Zi_WU4PSp3vVjNks9C47HIxGES94YDGrOTVXev-KtKkEsE-laUBsswyv3oFAO02G017HUHVf2V_jQo_Vw_fibbS4XL0FfzTlvKUKZzOe34A28MxyL5FVuoZEiOftrPLMs0BMyqD3bp03Q3Ex8zG1JamggJJRH-kFAz5VfCBQnTR93ifsYRiE2KFCDquau3T4K_OYrAX5wW_BJBjU01IvjqE_2A2rk",
		RefreshToken: "Atzr|IwEBIDLdczknrBmE9Bh4TLUARxlMcdFCR1zW79YtDPt006uP7a9SXCTfZa8M80OXoxxGaIp0hn-BjN5oC09ZDlg-SweQ9RnsUOf1nqqsgrVPHFZYD1esdrfYd6ymMzkXv4XaCNI97C0MH_ZiuUV03FdjGxrR85gLNYB_cW8_K8AknXE-n1Z9fKAQnXeReMzFSgbTZV8GBT2IlcHUyOsJjRDSpkX_CZb_GdoHj1FdBZC6HtWrBJd36UVdgT-eTNFpmi9-dYFsXQtPlkGkElGaFgdm8viAXXaZoUA8FGBh-1a_3gEMk1L-AbgBTqjfV6grZI9t-5w",
		AccessKeyID:  "AKIAVCZHDCZBJL6PSYNE",
		SecretKey:    "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",
		Region:       "us-east-1",
		RoleArn:      "arn:aws:iam::349584954946:role/erp_role",
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	endpoint := "https://sellingpartnerapi-na.amazon.com"

	finance, err := finances.NewClientWithResponses(endpoint,
		finances.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
			req.Header.Add("X-Amzn-Requestid", uuid.New().String())
			err = sellingPartner.SignRequest(req)
			if err != nil {
				return errors.Wrap(err, "sign error")
			}

			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				return errors.Wrap(err, "DumpRequest Error")
			}
			log.Printf("DumpRequest = %s", dump)
			return nil
		}),

		finances.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
			dump, err := httputil.DumpResponse(rsp, true)
			if err != nil {
				return errors.Wrap(err, "DumpResponse Error")
			}
			log.Printf("DumpResponse = %s", dump)

			return nil
		}),
	)

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()
	var size int32 = 10
	// var types []string = []string{"GET_MERCHANT_LISTINGS_ALL_DATA"}
	// var status []string = []string{"DONE"}
	// var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	// m, _ := time.ParseDuration("-1M")
	var since = time.Now().AddDate(0, -1, -21)
	var until = time.Now().AddDate(0, 0, -1)

	// var nextToken = ""

	// finance.GetOrdersWithResponse(ctx, &ordersV0.GetOrdersParams{
	finance.ListFinancialEventGroupsWithResponse(ctx, &finances.ListFinancialEventGroupsParams{
		&size,
		&until,
		&since,
		nil,
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	return nil
}
func (s SyncReport) SyncOrders() error {

	sellingPartner, err := sp.NewSellingPartner(&sp.Config{
		ClientID:     "amzn1.application-oa2-client.0c73f7463a85441f9ea68bbcb59702af",
		ClientSecret: "2bc5c7bed95d3891bf7b7562c7eb8497c307f8beebecc628614bfcde0d81a3b5",
		// RefreshToken: "Atzr|IwEBIFgMrY3c7wGNme00P-GXaiH1LEIkKrjl5pYKACuKjsC7qRqQefKlnv5gaTv9sXfM_-eXpIQBIZiqgh8f7wer0cZVPxqcGyrpAJVPj5Zi_WU4PSp3vVjNks9C47HIxGES94YDGrOTVXev-KtKkEsE-laUBsswyv3oFAO02G017HUHVf2V_jQo_Vw_fibbS4XL0FfzTlvKUKZzOe34A28MxyL5FVuoZEiOftrPLMs0BMyqD3bp03Q3Ex8zG1JamggJJRH-kFAz5VfCBQnTR93ifsYRiE2KFCDquau3T4K_OYrAX5wW_BJBjU01IvjqE_2A2rk",
		RefreshToken: "Atzr|IwEBIDLdczknrBmE9Bh4TLUARxlMcdFCR1zW79YtDPt006uP7a9SXCTfZa8M80OXoxxGaIp0hn-BjN5oC09ZDlg-SweQ9RnsUOf1nqqsgrVPHFZYD1esdrfYd6ymMzkXv4XaCNI97C0MH_ZiuUV03FdjGxrR85gLNYB_cW8_K8AknXE-n1Z9fKAQnXeReMzFSgbTZV8GBT2IlcHUyOsJjRDSpkX_CZb_GdoHj1FdBZC6HtWrBJd36UVdgT-eTNFpmi9-dYFsXQtPlkGkElGaFgdm8viAXXaZoUA8FGBh-1a_3gEMk1L-AbgBTqjfV6grZI9t-5w",
		AccessKeyID:  "AKIAVCZHDCZBJL6PSYNE",
		SecretKey:    "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",
		Region:       "us-east-1",
		RoleArn:      "arn:aws:iam::349584954946:role/erp_role",
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	endpoint := "https://sellingpartnerapi-na.amazon.com"

	order, err := ordersV0.NewClientWithResponses(endpoint,
		ordersV0.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
			req.Header.Add("X-Amzn-Requestid", uuid.New().String())
			err = sellingPartner.SignRequest(req)
			if err != nil {
				return errors.Wrap(err, "sign error")
			}

			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				return errors.Wrap(err, "DumpRequest Error")
			}
			log.Printf("DumpRequest = %s", dump)
			return nil
		}),

		ordersV0.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
			dump, err := httputil.DumpResponse(rsp, true)
			if err != nil {
				return errors.Wrap(err, "DumpResponse Error")
			}
			log.Printf("DumpResponse = %s", dump)

			return nil
		}),
	)

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()
	var size int = 10
	// var types []string = []string{"GET_MERCHANT_LISTINGS_ALL_DATA"}
	// var status []string = []string{"DONE"}
	var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	// m, _ := time.ParseDuration("-1M")
	var since = time.Now().AddDate(0, -2, -21).UTC().Format(time.RFC3339)
	var until = time.Now().AddDate(0, -1, 0).UTC().Format(time.RFC3339)

	// var nextToken = ""

	order.GetOrdersWithResponse(ctx, &ordersV0.GetOrdersParams{
		&since,
		&until,
		nil,
		nil,
		nil,
		marketplaceIds,
		nil,
		nil,
		nil,
		nil,
		&size,
		nil,
		nil,
		nil,
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	return nil
}

func (s SyncReport) Sync() error {

	sellingPartner, err := sp.NewSellingPartner(&sp.Config{
		ClientID:     "amzn1.application-oa2-client.0c73f7463a85441f9ea68bbcb59702af",
		ClientSecret: "2bc5c7bed95d3891bf7b7562c7eb8497c307f8beebecc628614bfcde0d81a3b5",
		// RefreshToken: "Atzr|IwEBIFgMrY3c7wGNme00P-GXaiH1LEIkKrjl5pYKACuKjsC7qRqQefKlnv5gaTv9sXfM_-eXpIQBIZiqgh8f7wer0cZVPxqcGyrpAJVPj5Zi_WU4PSp3vVjNks9C47HIxGES94YDGrOTVXev-KtKkEsE-laUBsswyv3oFAO02G017HUHVf2V_jQo_Vw_fibbS4XL0FfzTlvKUKZzOe34A28MxyL5FVuoZEiOftrPLMs0BMyqD3bp03Q3Ex8zG1JamggJJRH-kFAz5VfCBQnTR93ifsYRiE2KFCDquau3T4K_OYrAX5wW_BJBjU01IvjqE_2A2rk",
		RefreshToken: "Atzr|IwEBIDLdczknrBmE9Bh4TLUARxlMcdFCR1zW79YtDPt006uP7a9SXCTfZa8M80OXoxxGaIp0hn-BjN5oC09ZDlg-SweQ9RnsUOf1nqqsgrVPHFZYD1esdrfYd6ymMzkXv4XaCNI97C0MH_ZiuUV03FdjGxrR85gLNYB_cW8_K8AknXE-n1Z9fKAQnXeReMzFSgbTZV8GBT2IlcHUyOsJjRDSpkX_CZb_GdoHj1FdBZC6HtWrBJd36UVdgT-eTNFpmi9-dYFsXQtPlkGkElGaFgdm8viAXXaZoUA8FGBh-1a_3gEMk1L-AbgBTqjfV6grZI9t-5w",
		AccessKeyID:  "AKIAVCZHDCZBJL6PSYNE",
		SecretKey:    "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",
		Region:       "us-east-1",
		RoleArn:      "arn:aws:iam::349584954946:role/erp_role",
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	endpoint := "https://sellingpartnerapi-na.amazon.com"

	reporter, err := reports.NewClientWithResponses(endpoint,
		reports.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
			req.Header.Add("X-Amzn-Requestid", uuid.New().String())
			err = sellingPartner.SignRequest(req)
			if err != nil {
				return errors.Wrap(err, "sign error")
			}

			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				return errors.Wrap(err, "DumpRequest Error")
			}
			log.Printf("DumpRequest = %s", dump)
			return nil
		}),

		reports.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
			dump, err := httputil.DumpResponse(rsp, true)
			if err != nil {
				return errors.Wrap(err, "DumpResponse Error")
			}
			log.Printf("DumpResponse = %s", dump)

			return nil
		}),
	)

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()
	var size int = 10
	var types []string = []string{"GET_ORDER_REPORT_DATA_SHIPPING"}
	var status []string = []string{""}
	var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	// var marketplaceIds []string = []string{"US1"}
	// var since = time.Date(2021, time.March, 16, 12, 0, 0, 0, time.UTC)
	// var until = time.Date(2021, time.May, 1, 12, 0, 0, 0, time.UTC)
	var nextToken = ""
	_, err = reporter.GetReportsWithResponse(ctx, &reports.GetReportsParams{
		&types,
		&status,
		&marketplaceIds,
		&size,
		// &since,
		// &until,
		nil,
		nil,
		&nextToken,
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	return nil
}
