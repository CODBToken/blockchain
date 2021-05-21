package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/yunhailanuxgk/go-jinbao/common"
	"github.com/yunhailanuxgk/go-jinbao/core/types"
	"github.com/yunhailanuxgk/go-jinbao/ethclient"
	"io/ioutil"
	"math/big"
	"path"
	"sort"
	"time"
)

type item struct {
	Addr           common.Address
	Total, Counter *big.Int
}

type itemlist []item

func (o itemlist) Len() int {
	return len(o)
}

func (o itemlist) Less(i, j int) bool {
	return o[i].Total.Cmp(o[j].Total) > 0
}

func (o itemlist) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o itemlist) toCSV(fp string) {
	sort.Sort(o)
	buf := new(bytes.Buffer)
	buf.Write([]byte("Idx,Address,Times,Amount\r\n"))
	temp := "%d,%s,%d,%d\r\n"
	base := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	for i, item := range o {
		r := fmt.Sprintf(temp, i, item.Addr.Hex(), item.Counter, new(big.Int).Div(item.Total, base))
		buf.Write([]byte(r))
	}
	err := ioutil.WriteFile(fp, buf.Bytes(), 0755)
	fmt.Println("output", fp, "err", err)
}

func main() {
	ctx := context.Background()
	signer := types.NewEIP155Signer(big.NewInt(111))
	to := common.HexToAddress("0x830cfaaa057b3664773ec758eccf14aa24309c90") // coinsuper
	client, err := ethclient.Dial("/Users/liangc/Library/jinbao/jinbao.ipc")
	if err != nil {
		panic(err)
		return
	}
	current, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		return
	}
	fmt.Println(current.Number())
	s := time.Now()
	fmt.Println("start.")
	total := 0
	fc := make(map[common.Address]item)

	s2w := time.Now()
	for i := current.Number(); i.Cmp(big.NewInt(0)) > 0; i = new(big.Int).Sub(i, big.NewInt(1)) {
		blk, _ := client.BlockByNumber(ctx, i)
		for _, tx := range blk.Transactions() {
			if tx.To() != nil && *tx.To() == to {
				from, _ := types.Sender(signer, tx)
				v, ok := fc[from]
				if !ok {
					v = item{
						Addr:    from,
						Total:   big.NewInt(0),
						Counter: big.NewInt(0),
					}
				}
				v.Total = new(big.Int).Add(v.Total, tx.Value())
				v.Counter = new(big.Int).Add(v.Counter, big.NewInt(1))
				fc[from] = v
				total++
			}
		}
		if i.Int64()%20000 == 0 {
			fmt.Println("->", i.Int64(), time.Since(s2w))
			s2w = time.Now()
		}
	}

	il := make([]item, 0)
	for _, v := range fc {
		il = append(il, v)
	}

	itemlist(il).toCSV(path.Join(
		"/Users/liangc/Documents/JINBAO/coinsuper/",
		fmt.Sprintf("%s_%d.csv", time.Now().Format("2006-01-02"), current.Number())))
	//ret, _ := json.Marshal(fc)
	//buf := new(bytes.Buffer)
	//json.Indent(buf, ret, "", "\t")
	//fmt.Println(buf.String())
	fmt.Println("timeused", time.Since(s))
	fmt.Println("total tx", total, "total account", len(fc))
}
