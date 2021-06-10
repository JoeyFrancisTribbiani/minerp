package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	sp "gopkg.me/selling-partner-api-sdk/pkg/selling-partner"
	"gopkg.me/selling-partner-api-sdk/reports"
)

func main() {
	// Fuck()
	// s := &SyncReport{}
	// s.SyncReport()
	// Fuck1()
	// ReadExcel()
	// GetAllReportTypes()
	CreateAllReports()
}

func CreateAllReports() {
	resultFile := "ReportTypes.txt"
	file, err := os.Open(resultFile)
	if err != nil {
		log.Printf("Cannot open text file: %s, err:[%v]", resultFile, err)
	}
	defer file.Close()

	var reportTypeList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		reportTypeList = append(reportTypeList, line)
	}

	reporter, err := GetReporter()

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()

	var marketplaceIds []string = []string{"ATVPDKIKX0DER"}

	for _, v := range reportTypeList {
		reporter.CreateReportWithResponse(ctx, reports.CreateReportJSONRequestBody{
			DataEndTime:    nil,
			DataStartTime:  nil,
			MarketplaceIds: marketplaceIds,
			ReportOptions:  nil,
			ReportType:     v,
		})
	}
}

func GetAllReportTypes() {
	mdFilePath := "/root/git/go-admin/selling-partner-api-docs/references/reports-api/reportType_string_array_values.md"
	resultFile := "ReportTypes.txt"
	file, err := os.Open(mdFilePath)
	if err != nil {
		log.Printf("Cannot open text file: %s, err:[%v]", mdFilePath, err)
	}
	defer file.Close()

	pattenStr1 := "<p><strong>reportType</strong> value: "
	pattenStr2 := "<p><strong>reportType</strong> value:</p>"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var reportType = ""
		if strings.HasPrefix(line, pattenStr1) {
			reportType = strings.Replace(line, pattenStr1, "", 1)
			reportType = strings.Trim(reportType, "</p></td>")
			reportType = strings.Trim(reportType, "</p>")
		}

		if strings.HasPrefix(line, pattenStr2) {
			scanner.Scan()
			line = scanner.Text()
			reportType = strings.Trim(line, "</p></td>")
			// reportType = strings.Trim(reportType, "</p>")
		}
		if reportType == "" {
			continue
		}
		handler, _ := os.OpenFile(resultFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		defer handler.Close()

		write := bufio.NewWriter(handler)
		write.WriteString(reportType)
		write.WriteString("\n")
		write.Flush()
		handler.Close()
	}
}

func GetReporter() (*reports.ClientWithResponses, error) {
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

	return reporter, err
}
func ReadCsv() {

}

// 新添加的job 必须按照以下格式定义，并实现Exec函数
type SyncReport struct {
}

func (s SyncReport) SyncReport() error {

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
	var reportDocumentId = "amzn1.spdoc.1.3.e8bb480b-e1e2-4677-aee1-df235ea124f0.TNKPBOO8KOZLO.47700"
	reportRsp, err := reporter.GetReportDocumentWithResponse(ctx, reportDocumentId)
	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	var url = reportRsp.Model.Payload.Url
	// var aes = reportRsp.Model.Payload.EncryptionDetails.Key
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("report.txt", data, 0644)

	// fmt.Printf(string(reportRsp.Body))

	return nil
}

func Fuck1() {
	var data, _ = ioutil.ReadFile("report.txt")
	var key, _ = base64.StdEncoding.DecodeString("9Cn5K4TgMPXcmq6Qq5E3UpG7+DWFxfoZqAQ6202/g1g=")
	var iv, _ = base64.StdEncoding.DecodeString("0flfi6vuapTGn7zzvxyRaA==")

	var result, _ = Decrypt(key, iv, data)
	// var b, _ = base64.StdEncoding.DecodeString(string(result))
	// var str = string(b)

	// fmt.Println(str)

	ioutil.WriteFile("fuck.txt", result, 0644)
}

func pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Decrypt(key []byte, iv []byte, data []byte) ([]byte, error) {
	// func Decrypt(text string) (string, error) {
	// decode_data, err := base64.StdEncoding.DecodeString(text)
	// if err != nil {
	// 	return nil, nil
	// }
	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher(key)
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//输出到[]byte数组
	origin_data := make([]byte, len(data))
	blockMode.CryptBlocks(origin_data, data)
	//去除填充,并返回
	return unpad(origin_data), nil
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
