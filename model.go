package gofab

type Block struct {
	Number       uint64
	PreviousHash string
	DataHash     string
	CreateTime   string
	TxList       []Transaction
}

type Transaction struct {
	TransactionId string
	CreateTime    string
	Args          []string
}