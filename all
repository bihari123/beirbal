package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var CreditCardLabel = "PFI_CREDIT_CARD"

type CreditCardEntry = utils.ProdigyOutput

var sampleTextCreditCard = []string{
	"card verification %v ",
	"card identification number %v ",
	"cvn %v ",
	"cid %v ",
	"cvc2 %v ",
	"cvv2 %v ",
	"pin block %v ",
	"security code %v ",
	"security number %v ",
	"security no %v ",
	"issue number %v ",
	"issue no %v ",
	/*
	   "cryptogramme %v ",
	   "numéro de sécurité %v ",
	   "numero de securite %v ",
	   "kreditkartenprüfnummer %v ",
	   "kreditkartenprufnummer %v ",
	   "prüfziffer %v ",
	   "prufziffer %v ",
	   "sicherheits Kode %v ",
	   "sicherheitscode %v ",
	   "sicherheitsnummer %v ",
	   "verfalldatum %v ",
	   "codice di verifica %v ",
	   "cod. sicurezza %v ",
	   "cod sicurezza %v ",
	   "n autorizzazione %v ",
	   "código %v ",
	   "codigo %v ",
	   "cod. seg %v ",
	   "cod seg %v ",
	   "código de segurança %v ",
	   "codigo de seguranca %v ",
	   "codigo de segurança %v ",
	   "código de seguranca %v ",
	   "cód. segurança %v ",
	   "cod. seguranca %v ",
	   "cod. segurança %v ",
	   "cód. seguranca %v ",
	   "cód segurança %v ",
	   "cod seguranca %v ",
	   "cod segurança %v ",
	   "cód seguranca %v ",
	   "número de verificação %v ",
	   "numero de verificacao %v ",
	   "ablauf %v ",
	   "gültig bis %v ",
	   "gültigkeitsdatum %v ",
	   "gultig bis %v ",
	   "gultigkeitsdatum %v ",
	   "scadenza %v ",
	   "data scad %v ",
	   "fecha de expiracion %v ",
	   "fecha de venc %v ",
	   "vencimiento %v ",
	   "válido hasta %v ",
	   "valido hasta %v ",
	   "vto %v ",
	   "data de expiração %v ",
	   "data de expiracao %v ",
	   "data em que expira %v ",
	   "validade %v ",
	   "valor %v ",
	   "vencimento %v ",
	   "transaction %v ",
	   "transaction number %v ",
	   "reference number %v ",
	*/

}

var training_statements_credit_card = utils.GenTrainingStatement(
	sampleTextCreditCard,
	CreditCardLabel,
)

func GetTextForFieldCreditCard(creditCardNumber string) CreditCardEntry {
	return utils.GetRandomEntry(training_statements_credit_card, creditCardNumber)
}
