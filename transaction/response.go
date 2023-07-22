package transaction

type TransactionResponse struct {
	SenderPrivateKey           *string `json:"sender_private_key"`
	SenderPublicKey            *string `json:"sender_public_key"`
	SenderBlockchainAddress    *string `json:"sender_blockchain_address"`
	RecipientBlockchainAddress *string `json:"recipient_blockchain_address"`
	Value                      *string `json:"value"`
}

func (tr *TransactionResponse) Validate() bool {
	return tr.SenderPrivateKey != nil && tr.SenderPublicKey != nil && tr.SenderBlockchainAddress != nil && tr.RecipientBlockchainAddress != nil && tr.Value != nil
}
