package jobs

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	sp "gopkg.me/selling-partner-api-sdk/pkg/selling-partner"
	"gopkg.me/selling-partner-api-sdk/sellers"
)

func main() {
	// Fuck()
	s := &SyncReport{}
	s.Sync()
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

func (s SyncReport) Sync() error {

	sellingPartner, err := sp.NewSellingPartner(&sp.Config{
		ClientID:     "amzn1.application-oa2-client.0c73f7463a85441f9ea68bbcb59702af",
		ClientSecret: "2bc5c7bed95d3891bf7b7562c7eb8497c307f8beebecc628614bfcde0d81a3b5",
		RefreshToken: "Atzr|IwEBIFgMrY3c7wGNme00P-GXaiH1LEIkKrjl5pYKACuKjsC7qRqQefKlnv5gaTv9sXfM_-eXpIQBIZiqgh8f7wer0cZVPxqcGyrpAJVPj5Zi_WU4PSp3vVjNks9C47HIxGES94YDGrOTVXev-KtKkEsE-laUBsswyv3oFAO02G017HUHVf2V_jQo_Vw_fibbS4XL0FfzTlvKUKZzOe34A28MxyL5FVuoZEiOftrPLMs0BMyqD3bp03Q3Ex8zG1JamggJJRH-kFAz5VfCBQnTR93ifsYRiE2KFCDquau3T4K_OYrAX5wW_BJBjU01IvjqE_2A2rk",
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

	seller, err := sellers.NewClientWithResponses(endpoint,
		sellers.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
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

		sellers.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
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
	_, err = seller.GetMarketplaceParticipationsWithResponse(ctx)

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	return nil
}

func Fuck() error {
	// func (t SyncProduct) Exec() error {
	// func (t SyncProduct) Exec(arg interface{}) error {
	// str := time.Now().Format(timeFormat) + " [INFO] JobCore ExamplesOne exec success"
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；
	// switch arg.(type) {

	// case string:
	// 	if arg.(string) != "" {
	// 		fmt.Println("string", arg.(string))
	// 		fmt.Println(str, arg.(string))
	// 	} else {
	// 		fmt.Println("arg is nil")
	// 		fmt.Println(str, "arg is nil")
	// 	}
	// 	break
	// }
	// return nil

	sellingPartner, err := sp.NewSellingPartner(&sp.Config{
		ClientID:     "amzn1.application-oa2-client.0c73f7463a85441f9ea68bbcb59702af",
		ClientSecret: "2bc5c7bed95d3891bf7b7562c7eb8497c307f8beebecc628614bfcde0d81a3b5",
		RefreshToken: "Atzr|IwEBIFgMrY3c7wGNme00P-GXaiH1LEIkKrjl5pYKACuKjsC7qRqQefKlnv5gaTv9sXfM_-eXpIQBIZiqgh8f7wer0cZVPxqcGyrpAJVPj5Zi_WU4PSp3vVjNks9C47HIxGES94YDGrOTVXev-KtKkEsE-laUBsswyv3oFAO02G017HUHVf2V_jQo_Vw_fibbS4XL0FfzTlvKUKZzOe34A28MxyL5FVuoZEiOftrPLMs0BMyqD3bp03Q3Ex8zG1JamggJJRH-kFAz5VfCBQnTR93ifsYRiE2KFCDquau3T4K_OYrAX5wW_BJBjU01IvjqE_2A2rk",
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

	seller, err := sellers.NewClientWithResponses(endpoint,
		sellers.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
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

		sellers.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
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
	_, err = seller.GetMarketplaceParticipationsWithResponse(ctx)

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	return nil
}
