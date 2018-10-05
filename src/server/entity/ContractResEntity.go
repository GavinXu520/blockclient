package entity


type DeployEntity struct {
	Decimals 	uint8 		`json:"decimals"`
	Total		int64		`json:"total"`
	Name 		string		`json:"name"`
	Symbol 		string 		`json:"symbol"`

	// passphrase
	Pwd 		string 		`json:"pwd"`


	Key 		KeyStorage	`json:"key"`
}

type ToekenTransferEntity struct {
	To 	string 		`json:"to"`
	Pwd 	string 		`json:"pwd"`
	Address string 		`json:"address"`
	Value 	int 		`json:"value"`
	Key 	KeyStorage 	`json:"key"`
}

// Keystorage context
type KeyStorage struct{
	Address		string		`json:"address"`
	Crypto		struct{
		Cipher		string		`json:"cipher"`
		Ciphertext	string 		`json:"ciphertext"`
		Cipherparams	struct{
			Iv 	string 		`json:"iv"`
		}		`json:"cipherparams"`
		Kdf 		string		`json:"kdf"`
		Kdfparams 	struct{
			Dklen		int	`json:"dklen"`
			N 		int 	`json:"n"`
			P 		int 	`json:"p"`
			R 		int 	`json:"r"`
			Salt 		string 	`json:"salt"`

		}		`json:"kdfparams"`
		Mac 		string 		`json:"mac"`

	}		`json:"crypto"`

	Id 		string 		`json:"id"`
	Version 	int 		`json:"version"`

}
