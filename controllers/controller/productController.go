package controller

import (
	// 文字列と基本データ型の変換パッケージ
	strconv "strconv"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	// entity "../../models/entity"
	entity "mvc/models/entity"

	// DBアクセス用モジュール
	// db "../../models/db"
	db "mvc/models/db"
)

// 商品の購入状態を定義
const (
	// NotPurchased は 未購入
	NotPurchased = 0

	// Purchased は 購入済
	Purchased = 1
)

// FetchAllProducts は 全ての商品情報を取得する
func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}

// FindProduct は 指定したIDの商品情報を取得する
func FindProduct(c *gin.Context) {
	// クエリ文字列の取得
	productIDStr := c.Query("productID")

	// 文字列を10進数と解釈してintで返す
	productID, _ := strconv.Atoi(productIDStr)

	resultProduct := db.FindProduct(productID)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProduct)
}

// AddProduct は 商品をDBへ登録する
func AddProduct(c *gin.Context) {
	// c.PostFormでフォームの値を取得 str
	productName := c.PostForm("productName")
	productMemo := c.PostForm("productMemo")

	// productの定義してinsertする
	var product = entity.Product{
		Name:    productName,
		Memo:    productMemo,
		Default: 0,
		State:   NotPurchased,
	}

	db.InsertProduct(&product)
}

// AddDefaultProduct は デフォルト商品をDBへ登録する
// デフォルト商品：定期的に買う商品
func AddDefaultProduct(c *gin.Context) {
	// c.PostFormでフォームの値を取得 str
	productName := "納豆"
	productMemo := "定期的に買う商品"
	productDefault := 1

	// productの定義してinsertする
	var product = entity.Product{
		Name:    productName,
		Memo:    productMemo,
		Default: productDefault,
		State:   NotPurchased,
	}

	db.InsertProduct(&product)
}

// ChangeStateProduct は 商品情報の状態を変更する
func ChangeStateProduct(c *gin.Context) {
	reqProductID := c.PostForm("productID")
	reqProductState := c.PostForm("productState")

	productID, _ := strconv.Atoi(reqProductID)
	productState, _ := strconv.Atoi(reqProductState)
	changeState := NotPurchased

	// 商品状態が未購入の場合
	if productState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}

	db.UpdateStateProduct(productID, changeState)
}

// DeleteProduct は 商品情報をDBから削除する
func DeleteProduct(c *gin.Context) {
	productIDStr := c.PostForm("productID")
	productDefaultstr := c.PostForm("productDefault")
	productStatestr := c.PostForm("productState")

	productID, _ := strconv.Atoi(productIDStr)
	productDefault, _ := strconv.Atoi(productDefaultstr)
	productState, _ := strconv.Atoi(productStatestr)

	// デフォルト商品でないなら削除 デフォルトなら商品状態を変更
	if productDefault == 1 {
		productState = NotPurchased
		db.UpdateStateProduct(productID, productState)
	} else {
		db.DeleteProduct(productID)
	}
}
