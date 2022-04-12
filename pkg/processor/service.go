package processor

type Processor interface {
	CreateUserTransaction(createUserTransactionRequest CreateUserTransaction) (int, error)
}
