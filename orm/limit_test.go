package orm

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
	"testing"
	"time"
)

/**
 * @Author: feiliang.wang
 * @Description: 分页查询测试
 * @File:  limit_test
 * @Version: 1.0.0
 * @Date: 2021/1/5 11:01
 */

func TestLimit(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			num  int
			size int
		}
		want string
	}{
		{
			name: "单个测试",
			args: struct {
				num  int
				size int
			}{
				num:  2,
				size: 100,
			},
			want: " LIMIT 100,100 ",
		},
	}
	for _, item := range tests {
		s := Limit(item.args.num, item.args.size)
		if s != item.want {
			t.Fatal(item.name)
		}
	}
}

func TestInsertInfoTx(t *testing.T) {
	url := `https://api.blockchair.com/dogecoin/outputs?q=recipient(DUFWBTPmwu5uozB6BZJDgdcvjzAbNVTQR9),is_spent(false),block_id(..3585045)&limit=100&offset=0`
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	type value struct {
		TxHash  string `json:"transaction_hash"`
		Index   int    `json:"index"`
		Value   int64  `json:"value"`
		IsSpent bool   `json:"is_spent"`
	}
	type datas struct {
		Data []value `json:"data"`
	}

	var values []value

	var data datas
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		t.Fatal(err)
	}
	values = append(values, data.Data...)
	url = `https://api.blockchair.com/dogecoin/outputs?q=recipient(DUFWBTPmwu5uozB6BZJDgdcvjzAbNVTQR9),is_spent(false),block_id(..3585045)&limit=100&offset=100`
	resp, err = http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		t.Fatal(err)
	}
	values = append(values, data.Data...)
	sb := strings.Builder{}
	for _, value := range values {
		sb.WriteString(fmt.Sprintf("%s\t%d\t%d\t%v\n", value.TxHash, value.Index, value.Value, value.IsSpent))
	}
	t.Log(sb.String())
}

var coins map[string]string = map[string]string{
	"bitcoin":      "okkong_wallet",
	"bitcoin-cash": "okkong_bch_wallet",
	"bitcoin-sv":   "okkong_bsv_wallet",
	"litecoin":     "okkong_ltc_wallet",
	"dogecoin":     "okkong_doge_wallet",
}

var url_head = `https://api.blockchair.com/`
var url_end = `/dashboards/transaction/`

var ctx = context.Background()

func test_coin(t *testing.T, coin, db_name string) {
	t.Logf("-------start coin %s -----------", coin)
	defer t.Logf("-------end coin %s -----------", coin)
	url := url_head + coin + url_end
	dsn := fmt.Sprintf("okminer:okminer@okni@tcp(rm-bp1cal5n6kwjz1ur6.mysql.rds.aliyuncs.com:3306)/%s?autocommit=true&loc=Local&charset=utf8&parseTime=true", db_name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	type Utxo struct {
		TxHash       string `json:"tx_hash"`
		Vout         int    `json:"vout"`
		Address      string `json:"address"`
		ScriptPubkey string `json:"scriptPubKey"`
		Amount       int64  `json:"amount"`
	}
	var utxos []Utxo
	_, err = FilterListInfo(db, logger, `utxos_info`, SqlFilterMap{"spend_status": &EnumFilter{0, 1}}, &utxos)
	if err != nil {
		t.Error(err)
	}
	type output struct {
		ScriptHex string `json:"script_hex"`
		IsSpent   bool   `json:"is_spent"`
		Value     int64  `json:"value"`
	}
	type transaction struct {
		Outputs []output `json:"outputs"`
	}

	type transactions map[string]transaction

	type datas struct {
		Data transactions `json:"data"`
	}
	t.Logf("coin %s utxos len %d", coin, len(utxos))
	for i, item := range utxos {
		time.Sleep(time.Second)
		resp, err := http.Get(url + item.TxHash)
		if err != nil {
			t.Logf("----- fail-----")
			t.Logf("index:%d tx:%s get utxo info fail.%+v", i, item.TxHash, err)
			continue
		}
		var data datas
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			t.Logf("----- fail-----")
			t.Logf("index:%d tx:%s get utxo info fail.%+v", i, item.TxHash, err)
			continue
		}
		if len(data.Data[item.TxHash].Outputs) == 0 {
			t.Logf("----- fail-----")
			t.Logf("index:%d tx:%s get utxo info fail.outputs len is 0", i, item.TxHash)
			continue
		}
		tn, ok := data.Data[item.TxHash]
		if !ok {
			t.Logf("----- fail-----")
			t.Logf("index:%d tx:%s get utxo info fail.outputs tx hash can't find", i, item.TxHash)
			continue
		}
		if len(tn.Outputs) <= item.Vout {
			t.Logf("----- fail-----")
			t.Logf("index:%d tx:%s get utxo info fail.outputs len %d <= vout %d", i, item.TxHash, len(tn.Outputs), item.Vout)
			continue
		}
		out := tn.Outputs[item.Vout]
		if out.IsSpent || out.ScriptHex != item.ScriptPubkey || out.Value != item.Amount {
			t.Logf("----- fail-----")
			t.Logf("index:%d tx:%s vout:%d can't compare,read:%+v db:%+v", i, item.TxHash, item.Vout, out, item)
			continue
		}
		//t.Logf("----- ok-----")
		//t.Logf("index:%d tx:%s vout:%d address:%s value:%d ", i, item.TxHash, item.Vout, item.Address, item.Amount)
	}
}

func TestDeleteInfoTx(t *testing.T) {
	for coin, db_name := range coins {
		test_coin(t, coin, db_name)
	}
}
