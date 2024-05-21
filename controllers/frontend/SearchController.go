/*
   使用 Beego 框架和 Elasticsearch 客户端实现商品数据的增删改查功能，包括基本的添加、更新、删除、单条查询、多条查询、
条件查询和分页查询。每个方法都涉及到从数据库获取数据、与 Elasticsearch 交互、处理查询结果，并将结果返回给客户端或输出到控制台。
*/

package frontend

import (
	"B2CProject/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/olivere/elastic/v7"
	"math"
	"reflect"
	"strconv"
)

type SearchController struct {
	BaseController
}

func init() {
	//exists, err := models.EsClient.IndexExists("product").Do(context.Background())
	//if err != nil {
	//	beego.Error(err)
	//}
	//if !exists {
	//	// Create a new index.
	//	mapping := `
	//	{
	//		"settings": {
	//		  "number_of_shards": 1,
	//		  "number_of_replicas": 0
	//		},
	//		"mappings": {
	//		  "properties": {
	//			"content": {
	//			  "type": "text",
	//			  "analyzer": "ik_max_word",
	//			  "search_analyzer": "ik_max_word"
	//			},
	//			"title": {
	//			  "type": "text",
	//			  "analyzer": "ik_max_word",
	//			  "search_analyzer": "ik_max_word"
	//			}
	//		  }
	//		}
	//	  }
	//	`
	//	_, err := models.EsClient.CreateIndex("product").Body(mapping).Do(context.Background())
	//	if err != nil {
	//		// Handle error
	//		beego.Error(err)
	//	}
	//
	//}
}

// 增加商品数据
func (c *SearchController) AddProduct() {
	product := []models.Product{}
	models.DB.Find(&product)

	for i := 0; i < len(product); i++ {
		_, err := models.EsClient.Index().
			Index("product").
			Id(strconv.Itoa(product[i].Id)).
			BodyJson(product[i]).
			Do(context.Background())
		if err != nil {
			// Handle error
			beego.Error(err)
		}
	}

	c.Ctx.WriteString("AddProduct success")

}

// 更新数据
func (c *SearchController) Update() {
	//从数据库获取修改

	product := models.Product{}
	models.DB.Where("id=20").Find(&product)
	product.Title = "苹果电脑"
	product.SubTitle = "苹果电脑"
	res, err := models.EsClient.Update().
		Index("product").
		Type("_doc").
		Id("20").
		Doc(product).
		Do(context.Background())
	if err != nil {
		beego.Error(err)
	}
	fmt.Printf("update %s\n", res.Result)

	c.Ctx.WriteString("修改数据")
}

// 删除
func (c *SearchController) Delete() {
	res, err := models.EsClient.Delete().
		Index("product").
		Type("_doc").
		Id("20").
		Do(context.Background())

	if err != nil {
		beego.Error(err)
	}
	fmt.Printf("Delete %s\n", res.Result)

	c.Ctx.WriteString("删除成功")
}

// 查询一条数据
func (c *SearchController) GetOne() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.Ctx.WriteString("GetOne")
		}
	}()

	result, _ := models.EsClient.Get().
		Index("product").
		Id("19").
		Do(context.Background())

	fmt.Println(result.Source)

	product := models.Product{}
	json.Unmarshal(result.Source, &product)
	c.Data["json"] = product
	c.ServeJSON()

}

// 查询多条数据
func (c *SearchController) Query() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.Ctx.WriteString("Query")
		}
	}()

	query := elastic.NewMatchQuery("Title", "旗舰")
	searchResult, err := models.EsClient.Search().
		Index("product").        // search in index "twitter"
		Query(query).            // specify the query
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	productList := []models.Product{}
	var product models.Product
	for _, item := range searchResult.Each(reflect.TypeOf(product)) {
		g := item.(models.Product)
		fmt.Printf("标题： %v\n", g.Title)
		productList = append(productList, g)
	}

	c.Data["json"] = productList
	c.ServeJSON()

}

// 条件筛选查询
func (c *SearchController) FilterQuery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.Ctx.WriteString("Query")
		}
	}()

	//筛选
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("Title", "小米"))
	boolQ.Filter(elastic.NewRangeQuery("Id").Gt(19))
	boolQ.Filter(elastic.NewRangeQuery("Id").Lt(31))
	searchResult, err := models.EsClient.Search().Index("product").Type("_doc").Query(boolQ).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	var product models.Product
	for _, item := range searchResult.Each(reflect.TypeOf(product)) {
		t := item.(models.Product)
		fmt.Printf("Id:%v 标题：%v\n", t.Id, t.Title)
	}

	c.Ctx.WriteString("filter Query")
}

// 分页查询
func (c *SearchController) ProductList() {
	c.BaseInit()
	keyword := c.GetString("keyword")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.Ctx.WriteString("ProductList")
		}
	}()

	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	pageSize := 5

	query := elastic.NewMatchQuery("Title", keyword)
	searchResult, err := models.EsClient.Search().
		Index("product").
		Query(query).
		Sort("Price", true). //true 升序
		Sort("Id", false).   //false 降序
		From((page - 1) * pageSize).Size(pageSize).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	//查询符合条件的商品的总数
	searchResult2, _ := models.EsClient.Search().
		Index("product").        // search in index "twitter"
		Query(query).            // specify the query
		Do(context.Background()) // execute

	productList := []models.Product{}
	var product models.Product
	for _, item := range searchResult.Each(reflect.TypeOf(product)) {
		g := item.(models.Product)
		fmt.Printf("标题： %v\n", g.Title)
		productList = append(productList, g)
	}
	c.Data["productList"] = productList
	c.Data["totalPages"] = math.Ceil(float64(len(searchResult2.Each(reflect.TypeOf(product)))) / float64(pageSize))
	c.Data["page"] = page
	c.Data["keyword"] = keyword
	// c.ServeJSON()
	c.TplName = "frontend/elasticsearch/list.html"
}
