package vm

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cc14514/go-lib"
	"github.com/yunhailanuxgk/go-jinbao/accounts/keystore"
	"github.com/yunhailanuxgk/go-jinbao/common"
	"github.com/yunhailanuxgk/go-jinbao/core/types"
	"github.com/yunhailanuxgk/go-jinbao/crypto"
	"github.com/yunhailanuxgk/go-jinbao/ethclient"
	"math/big"
	"testing"
	"time"
)

/*
0204 : 400
0205 : 200
0206 : 200
0218 : 5200
*/

/*
0 9ba390bbb0021693f30e2ba4e515feafbdefac3a52bcecd98282edbbb05f4192 0xCFE65DD56Ea69d9f6542103db5Db3d1F64D7057A
1 869db7d9a6dc34e3aabf1deae3c1cd100aac8e5fee5331bfed094b7b4c2a8b95 0x8eb9238a7A34779A939bDfB39ba389848bb8658E
2 0215f23d2390af65f588eedc132bc4f83dbc07f515dd5f6b476182c26ca3d1db 0xca04D510109B16a4173CF52E30ecBc5d0D22B0Fe
3 3d24e16d167cb77b52ab5ee955bff771c1eceed9a81a92d8bf9e5cd283aa0a1b 0x468B2226e261DcCCB28b9883675880c8af2417dA
4 930b5b6f4c41029615ec19aae7162f3161cb782dffb6261587eb2fecbe8d28f4 0xFba3f30F82081c528301EFdaa3d0Ff28a12Ae106
5 132e970bde064bc9fe5300a7264b44e4255d2ca0b1aa5ca12478c0ef9615c413 0x905f5c896FE97589DDfd44BE6E03BE8E1561c271

---------------------
0 d84cdf13fe1e18a039772e88669c48271ac561e2daf373a08b6970276a4faaa0 0x323616CA7316ca6Ae61136011C1a9ed47544D772
1 6697eb6b59cce7fd288a65638c470c99b7e3673e861209052b1bb90c735aeb3b 0x96e159ec1e6e9cEC2ad82bdAb6B7cebb2Ab0fE2D
2 5b29e2d6342b53c629de3930546a3f91a0b2ea20fd3f4c160398ddc84100f435 0x559b697AE6C68Ce03079dC3Cc051611777200d8f
3 9713dca7e3050157bddbc13eb5ecee83c23af9f9dcc87c9c5cd0d790a837bc5e 0x32D254AbC39daEc78a18017c50ea18107B26A9c2
4 ed866a2ed9a9a238e0fca9e54769702f63f87460a157e6446c99c3bcb43add62 0xcDbB3523E45935785cc54a39E8d1709B9fee8Ec2
5 8f7afb17daa783662a20af12cc84995ff3d7bcafcf37f3ed572eb3953fe9e32c 0xd303A465A4723F1EA2AE8D0443766d34B183fA5e
6 257d83e2e6a0f8c5133840cddbb3664cff3713c9465e19b75f420c9801ac8296 0xd1B637986001a40EE1067b9322f958021267Fbfd
7 3a9a3bf19fb12fbd80b1beba061450389c4ee49f2bd9ce23d993db0359d1afb8 0xE5a925F8587117EE6E60aF3Ff7B3d13ecb0A174B
8 d8e326b485d9db996bcf8de59fe9efebf9049eccf3d83da791a263cfe0bc376f 0x1B59D90C699950d54C788cb1261ba71b74Df88Ce
9 7ac94c9e3a8d0ff3d20c21abb21f0f26d53e3b7cfb1063145149f85b423260f4 0xf07abc2c643B144BCd2C9AB5b5ECCFA9FF86cdE6

*/
var (
	// 老鼠仓 : 0xe35d492cba7cb79bbea670ea1a546456b6d63c42
	mouse  = `{"address":"e35d492cba7cb79bbea670ea1a546456b6d63c42","crypto":{"cipher":"aes-128-ctr","ciphertext":"9a73e5839869efb0ed144c3900c714f3957aac6cf2cf110839114ec1698b9438","cipherparams":{"iv":"6d44c9a0bd6e0832375c7a235a042362"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"9af2dc37bec3199c2241807dd1c3a745d7aa7249ef5ad3bbd470a173ef2b1a99"},"mac":"d0ed5370fddc31adfc87f8660d9fee4003368f39c7e7cba78a4f005b9fa25534"},"id":"d1f91170-7e2b-45d2-a2a4-3c220320e57a","version":3}`
	amount = "200000000000000000000" // 200
	//0x8472bFaFCf95a26FDBa5a8E2a61097F286756715
	//contractAddr = common.HexToAddress("0x8472bFaFCf95a26FDBa5a8E2a61097F286756715")
	lockLedgerAddr2 = common.HexToAddress("0x8472bFaFCf95a26FDBa5a8E2a61097F286756715")
	prvs            = []string{
		"9ba390bbb0021693f30e2ba4e515feafbdefac3a52bcecd98282edbbb05f4192", // 0204~0206 600
		"869db7d9a6dc34e3aabf1deae3c1cd100aac8e5fee5331bfed094b7b4c2a8b95",
		"0215f23d2390af65f588eedc132bc4f83dbc07f515dd5f6b476182c26ca3d1db",
		"3d24e16d167cb77b52ab5ee955bff771c1eceed9a81a92d8bf9e5cd283aa0a1b",
		"930b5b6f4c41029615ec19aae7162f3161cb782dffb6261587eb2fecbe8d28f4",
		"132e970bde064bc9fe5300a7264b44e4255d2ca0b1aa5ca12478c0ef9615c413",
		// ----------
		"d84cdf13fe1e18a039772e88669c48271ac561e2daf373a08b6970276a4faaa0",
		"6697eb6b59cce7fd288a65638c470c99b7e3673e861209052b1bb90c735aeb3b",
		"5b29e2d6342b53c629de3930546a3f91a0b2ea20fd3f4c160398ddc84100f435",
		"9713dca7e3050157bddbc13eb5ecee83c23af9f9dcc87c9c5cd0d790a837bc5e",
		"ed866a2ed9a9a238e0fca9e54769702f63f87460a157e6446c99c3bcb43add62",
		"8f7afb17daa783662a20af12cc84995ff3d7bcafcf37f3ed572eb3953fe9e32c",
		"257d83e2e6a0f8c5133840cddbb3664cff3713c9465e19b75f420c9801ac8296",
		"3a9a3bf19fb12fbd80b1beba061450389c4ee49f2bd9ce23d993db0359d1afb8",
		"d8e326b485d9db996bcf8de59fe9efebf9049eccf3d83da791a263cfe0bc376f",
		"7ac94c9e3a8d0ff3d20c21abb21f0f26d53e3b7cfb1063145149f85b423260f4",
		// -----------
	}
	hexToPrv = func(h string) *ecdsa.PrivateKey {
		b, _ := hex.DecodeString(h)
		p, _ := crypto.ToECDSA(b)
		return p
	}
	prvToAddr = func(prv *ecdsa.PrivateKey) common.Address {
		return crypto.PubkeyToAddress(prv.PublicKey)
	}
	prvRandom = func() string {
		p, _ := crypto.GenerateKey()
		i := new(big.Int).Mod(p.X, big.NewInt(int64(len(prvs))))
		return prvs[i.Int64()]
	}
	getAddr = func() (*ecdsa.PrivateKey, common.Address) {
		var (
			prvHex = prvRandom()
			prv    = hexToPrv(prvHex)
		)
		addr := prvToAddr(prv)
		return prv, addr
	}
)

func TestShowaddr(t *testing.T) {
	lib.Start()
	for i, prvH := range prvs {
		prv := hexToPrv(prvH)
		addr := prvToAddr(prv)
		_, ok := lib.Mdb[addr.Hash()]
		t.Log(i, prvH, addr.Hex(), ok)
	}
}

/*
0206 : from 0x905f5c896FE97589DDfd44BE6E03BE8E1561c271 tx 0x17c86be3e077a0d0cb0a93953e0c2353a252efa15a5d91c59e7cdd4e0dff4ffc amount 200
0218 : from 0xcfe65dd56ea69d9f6542103db5db3d1f64d7057a tx 0x52d62c737468405937d25414106dd20eb7cbe6d915776196287bf57d4ef92d9d amount 200
	 : from 0x8eb9238a7A34779A939bDfB39ba389848bb8658E tx 0xe041da3ea27eb02dbdd09bfc9e16883ecf976db0ace62d409dd27f94b960e866 amount 5000

0227 : from 0xE5a925F8587117EE6E60aF3Ff7B3d13ecb0A174B tx 0x50f8d6a8b260da221a6546003c5e13b83761f15a0d4fc652e79a90dd44978be5 amount 800
0228 : from 0xFba3f30F82081c528301EFdaa3d0Ff28a12Ae106 tx 0x89e28d13ddd07de6006480f2fadb0a913cbbdf3c6ff7fc52cf920c3ffd60d01c amount 800
0301 : from 0xd303A465A4723F1EA2AE8D0443766d34B183fA5e tx 0x2f0f2a8064d942d48eeaabde72af7af53f15975868637a2cf65251a7f06aac28 amount 2000
0302 : from 0x96e159ec1e6e9cEC2ad82bdAb6B7cebb2Ab0fE2D tx 0x6b9e9108b110749d3f063c0a918e592ca7b4e6a1945b57a10ffc7fd1ff250ba8 amount 1000
0304 : from 0xFba3f30F82081c528301EFdaa3d0Ff28a12Ae106 tx 0xce1245033d3f972d35ecfe320de674eb85a0664605748caee906381e587f8676 amount 2000

0306 : from 0xcDbB3523E45935785cc54a39E8d1709B9fee8Ec2 tx 0x677bfbe98b424980529650f51b2762c64d0ad2bb393c920f0b72acec53f15d4b amount 30000
0308 : from 0x323616CA7316ca6Ae61136011C1a9ed47544D772 tx 0x3b2701e63b1da1f5f8a2ce5ec119b75c08c2b37d17b1097aa1ad3400a02c62e0 amount 4000

*/
func argsFn(to, a, n string, laddt int) (_to string, _a string, _c *big.Int, _p string, _ca common.Address) {
	_to, _a = to, a
	switch n {
	case "test":
		_c = big.NewInt(3)
		_p = "/Users/liangc/Library/jinbao/testnet/jinbao.ipc"
	case "dev":
		_c = big.NewInt(4)
		_p = "/Users/liangc/Library/jinbao/devnet/jinbao.ipc"
	default:
		_c = big.NewInt(111)
		_p = "/Users/liangc/Library/jinbao/jinbao.ipc"
	}

	switch laddt {
	case 2:
		_ca = lockLedgerAddr2
	default:
		_ca = lockLedgerAddr
	}
	return
}

func TestClient_SendTransaction(t *testing.T) {
	var (
		// 主账户 : 200
		//to, amount, cid, ipcpath ,contractAddr = argsFn( "0xe20ba01bcb4b693458dab1544ee6673ea6aaea9d", "200000000000000000000", "main",1) // 200

		// 老鼠仓 :
		/*
			800   ==> 800000000000000000000
			1000  ==> 1000000000000000000000
			2000  ==> 2000000000000000000000
			4000  ==> 4000000000000000000000
			5000  ==> 5000000000000000000000
			30000 ==> 30000000000000000000000
		*/
		//loc1 = "0xb0b3c7d80866cd70691954bead9806bf594008f5"
		loc2                                   = "0xa4dc639179ef3cc2677fb1029b385e7bb81bc21f"
		loc                                    = loc2
		to, amount, cid, ipcpath, contractAddr = argsFn(loc, "4000000000000000000000", "main", 2)

		// 测试
		//to, amount, cid, ipcpath = argsFn( "0x75f49085b8d187e8747ffa63f95a22415d37e3e2", "5000000000000000000000", "test") // 5000
		ctx      = context.Background()
		gas      = big.NewInt(150000)
		price, _ = new(big.Int).SetString("18000000000", 10)
	)
	prv, addr := getAddr()
	client, err := ethclient.Dial(ipcpath)
	if err != nil {
		panic(err)
	}
	n, _ := client.PendingNonceAt(ctx, addr)
	tx := types.NewTransaction(n, contractAddr,
		new(big.Int), gas, price,
		[]byte(fmt.Sprintf("unlock,%s,%s", to, amount)))
	signer := types.NewEIP155Signer(cid)
	tx, _ = types.SignTx(tx, signer, prv)
	err = client.SendTransaction(ctx, tx)
	fmt.Println("from", addr.Hex(), "tx", tx.Hash().Hex(), "amount", amount)
}

func TestFindTx(t *testing.T) {
	ctx := context.Background()
	signer := types.NewEIP155Signer(big.NewInt(2))
	to := common.HexToAddress("0x830cfaaa057b3664773ec758eccf14aa24309c90") // coinsuper
	client, err := ethclient.Dial("/Users/liangc/Library/jinbao/testnet/jinbao.ipc")
	if err != nil {
		t.Error(err)
		return
	}
	current, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		return
	}
	t.Log(current.Number())
	s := time.Now()
	fmt.Println("start.")
	total := 0
	fc := make(map[common.Address]*big.Int)
	for i := current.Number(); i.Cmp(big.NewInt(0)) > 0; i = new(big.Int).Sub(i, big.NewInt(1)) {
		blk, _ := client.BlockByNumber(ctx, i)
		for _, tx := range blk.Transactions() {
			if tx.To() != nil && *tx.To() == to {
				from, _ := types.Sender(signer, tx)
				v, ok := fc[from]
				if !ok {
					v = big.NewInt(0)
				}
				v = new(big.Int).Add(v, tx.Value())
				fc[from] = v
				total++
			}
		}
	}

	ret, _ := json.Marshal(fc)
	buf := new(bytes.Buffer)
	json.Indent(buf, ret, "", "\t")
	fmt.Println(buf.String())
	fmt.Println("timeused", time.Since(s))
	fmt.Println("total tx", total)

}

func TestLockAddr(t *testing.T) {
	addr := common.BytesToAddress(crypto.Keccak256(lockLedgerAddr.Bytes()))
	t.Log(addr.Hex())
	t.Log("lockledgerAddr", common.BytesToAddress(crypto.Keccak256(lockLedgerAddr.Bytes())).Hex())
}

// 老鼠仓
func TestMouse(t *testing.T) {
	var (
		key, _ = keystore.DecryptKey([]byte(mouse), "123456")
		prv    = key.PrivateKey
		addr   = prvToAddr(prv)

		ipcpath = "/Users/liangc/Library/jinbao/jinbao.ipc"
		cid     = big.NewInt(111)

		ctx      = context.Background()
		gas      = big.NewInt(21000)
		price, _ = new(big.Int).SetString("18000000000", 10)
	)

	client, err := ethclient.Dial(ipcpath)
	if err != nil {
		panic(err)
	}

	n, _ := client.PendingNonceAt(ctx, addr)
	a := big.NewInt(100000000000000000) //0.1

	for _, _prv := range prvs {
		tprv := hexToPrv(_prv)
		to := prvToAddr(tprv)
		//TODO loop
		tx := types.NewTransaction(n, to, a, gas, price, nil)
		signer := types.NewEIP155Signer(cid)
		tx, _ = types.SignTx(tx, signer, prv)
		err = client.SendTransaction(ctx, tx)
		fmt.Println("nonce", n, "from", addr.Hex(), "tx", tx.Hash().Hex(), "amount", amount)
		n = n + 1
	}
}
