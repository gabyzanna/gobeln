package gobeln

import (
	"context"
	"fmt"
	"testing"

	pasproj "github.com/HRMonitorr/PasetoprojectBackend"
	"github.com/whatsauth/watoken"
)

//func TestGCHandlerFunc(t *testing.T) {
//	data := GCHandlerFunc("string", "GIS", "geogis")
//
//	fmt.Printf("%+v", data)
//}

func TestGeneratePaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("gisname", "PrivateKey")
	fmt.Println(hasil, err)
}

func TestUpdateData(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "gisname",
		Volume: "1",
	}
	up := UpdateNameGeo("MONGOSTRING", "dbsw", context.Background(), data)
	fmt.Println(up)
}

func TestDeleteDataGeo(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "gisname",
		Volume: "1",
	}
	up := DeleteDataGeo("MONGOSTRING", "dbsw", context.Background(), data)
	fmt.Println(up)
}

func TestInsertUser(t *testing.T) {
	conn := GetConnectionMongo("MONGOSTRING", "dbsw")
	pass, _ := pasproj.HashPass("testpass")
	data := RegisterStruct{
		Username: "cobacoba",
		Password: pass,
	}
	ins := InsertUserdata(conn, data.Username, data.Password)
	fmt.Println(ins)
}

func TestIsexist(t *testing.T) {
	data := IsExist("token", "PublicKey")
	fmt.Println(data)
}
