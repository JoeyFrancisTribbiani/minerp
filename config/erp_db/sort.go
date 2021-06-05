package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	Sort("product.sql", "ASIN.sql", "product.bak.sql")
	// Sort("回款记录.sql", "ASIN.sql", "回款.bak.sql")
	// Sort("批次成本.sql", "ASIN.sql", "批次.bak.sql")
	// Sort("业绩报表.sql", "ASIN.sql", "业绩.bak.sql")
	// Sort("店铺.sql", "店铺汇总.sql", "店铺.bak.sql")
	// Sort("店铺汇总.sql", "店铺.sql", "店铺汇总.bak.sql")
	// Sort("父ASIN.sql", "ASIN.sql", "ASIN.bak.sql")
	// Sort("ASIN.sql", "父ASIN.sql", "父ASIN.bak.sql")
	// Sort("ASIN.sql", "MSKU.sql", "MSKU.bak.sql")
	// Sort("ASIN.sql", "SKU.sql", "SKU.bak.sql")
}

func Sort(parentfile string, textfile string, resultfile string) error {
	file, err := os.Open(parentfile)
	if err != nil {
		log.Printf("Cannot open text file: %s, err:[%v]", textfile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		key := strings.Split(strings.Trim(line, " "), " ")[0]
		result, _ := MatchText(textfile, key)

		text := ""
		if result != "" {
			text = result
		} else {
			text = "-----" + line + "-----"
		}
		handler, _ := os.OpenFile(resultfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		defer handler.Close()

		write := bufio.NewWriter(handler)
		write.WriteString(text)
		write.WriteString("\n")
		write.Flush()
		handler.Close()
	}
	return err
}

func MatchText(textfile string, key string) (string, error) {
	file, err := os.Open(textfile)
	if err != nil {
		log.Printf("Cannot open text file: %s, err:[%v]", textfile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.Trim(line, " "), key) {
			return line, err
		}
	}
	return "", nil
}
