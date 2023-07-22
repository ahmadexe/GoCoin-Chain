package transaction

type TransactionResponse struct {
	SenderPrivateKey           *string `json:"senderPrivateKey"`
	SenderPublicKey            *string `json:"senderPublicKey"`
	SenderBlockchainAddress    *string `json:"senderBlockchainAddress"`
	RecipientBlockchainAddress *string `json:"recipientBlockchainAddress"`
	Value                      *string `json:"value"`
}

func (tr *TransactionResponse) Validate() bool {
	return tr.SenderPrivateKey != nil && tr.SenderPublicKey != nil && tr.SenderBlockchainAddress != nil && tr.RecipientBlockchainAddress != nil && tr.Value != nil
}
