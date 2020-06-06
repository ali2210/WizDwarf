package db



import (
	wallet "./cloudwalletclass"
	walletAcc "../structs"
	"fmt"
	"context"
	firebase "firebase.google.com/go"
)

// constants 

const collectionName string = "EthereumPrivateLedger"




// Interface 
type PublicLedger interface{

	// Public Ledger 
	CreatePublicAddress(w *wallet.EthereumWalletAcc , clientID *firebase.App)(*wallet.EthereumWalletAcc, error);
	FindMyPublicAddress(w *walletAcc.Acc, clientID *firebase.App)(*wallet.EthereumWalletAcc,error);
	FindMyAddressByEmail(w *walletAcc.Acc, clientID *firebase.App)(*wallet.EthereumWalletAcc,error);
}



type ledgerPublic struct{}



func NewCollectionInstance() PublicLedger{
	return &ledgerPublic{}
}


func (*ledgerPublic)CreatePublicAddress(w *wallet.EthereumWalletAcc, clientID *firebase.App)(*wallet.EthereumWalletAcc, error){

		ctx := context.Background()

		client , err := clientID.Firestore(ctx); if err != nil{
			fmt.Println("Error", err)
			return nil, err
		}
		defer client.Close()
		fmt.Println("clientID:", clientID)

		_, _ ,err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
			"Email" : w.Email,
			"Password" : w.Password,
			"EthAddress": w.EthAddress,
			"Terms":w.Terms,
		}); if err != nil{
			fmt.Println("Error", err)
			return nil , err	
		}
		return w,nil

}


func (*ledgerPublic)FindMyPublicAddress(w *walletAcc.Acc, clientID *firebase.App)(*wallet.EthereumWalletAcc, error){


	ctx := context.Background()

	client , err := clientID.Firestore(ctx); if err != nil {
		fmt.Println("Error:", err)
		return nil, err 
	}
	defer client.Close()
	var ethereumDetials wallet.EthereumWalletAcc

	iterator := client.Collection(collectionName).Where("Email", "==" , w.Email ).Where("Password" , "==" , w.Password).Where("EthAddress", "==" , w.EthAddress).Documents(ctx)
	defer iterator.Stop()

	for{
		doc , err := iterator.Next(); if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		ethereumDetials = wallet.EthereumWalletAcc{
			Email : doc.Data()["Email"].(string),
			Password : doc.Data()["Password"].(string),
			EthAddress : doc.Data()["EthAddress"].(string),
			Terms : doc.Data()["Terms"].(bool),
		}
		break
	}
	return &ethereumDetials, nil
}

func (*ledgerPublic)FindMyAddressByEmail(w *walletAcc.Acc, clientID *firebase.App)(*wallet.EthereumWalletAcc,error){
	ctx := context.Background()

	client , err := clientID.Firestore(ctx); if err != nil {
		fmt.Println("Error:", err)
		return nil, err 
	}
	defer client.Close()
	var ethereumDetials wallet.EthereumWalletAcc

	iterator := client.Collection(collectionName).Where("Email", "==" , w.Email ).Where("Password" , "==" , w.Password).Documents(ctx)
	defer iterator.Stop()

	for{
		doc , err := iterator.Next(); if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		ethereumDetials = wallet.EthereumWalletAcc{
			Email : doc.Data()["Email"].(string),
			Password : doc.Data()["Password"].(string),
			EthAddress : doc.Data()["EthAddress"].(string),
			Terms : doc.Data()["Terms"].(bool),
		}
		break
	}
	return &ethereumDetials, nil	
}

