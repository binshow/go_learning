package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strconv"
)

/**
使用go操作es实现三种分页方式：

*/

var (
	client *elastic.Client
	host   = "http://localhost:9200/"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}

	_, _, err = client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}

	version, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", version)
}

func main() {
	//create()
	//getById("1")
	//getById("2")
	//delete("5")
	//update()
	//query()
	//search(3,2)

	for i := 100; i < 1000; i++ {
		create(strconv.Itoa(i))
	}
}

// 创建
func create(id string) {
	//1. 插入结构体
	name := "binshow" + id
	p1 := Person{name, 22, "niubi", []string{"wangzherongyao"}}
	put1, err := client.Index().Index("ustc").Type("person").Id(id).BodyJson(p1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	////2. 插入字符串
	//p2 := `{"name":"hhj","age":23,"about":"smart","interests":["guangjie"]}`
	//put2, err := client.Index().Index("ustc").Type("person").Id("5").BodyJson(p2).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

}

// 查询
func getById(id string) {
	get1, err := client.Get().Index("ustc").Type("person").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		var p Person
		err := json.Unmarshal(get1.Source, &p)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(p.Name)
		fmt.Println(p.Age)
		fmt.Println(string(get1.Source))
	}
}

// 删除
func delete(id string) {
	res, err := client.Delete().Index("ustc").
		Type("person").
		Id(id).
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)

}

// 更新
func update() {
	res, err := client.Update().
		Index("ustc").
		Type("person").
		Id("1").
		Doc(map[string]interface{}{"age": 88}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)
}

// 搜索
func query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = client.Search("ustc").Type("person").Do(context.Background())
	printPerson(res, err)

	//字段相等
	q := elastic.NewQueryStringQuery("name:binshow")
	res, err = client.Search("ustc").Type("person").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printPerson(res, err)

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("ustc").Type("employee").Query(q).Do(context.Background())
	printPerson(res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "niubi")
	res, err = client.Search("ustc").Type("person").Query(matchPhraseQuery).Do(context.Background())
	printPerson(res, err)

	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err = client.Search("ustc").Type("person").Aggregation("all_interests", aggs).Do(context.Background())
	printPerson(res, err)
}

//打印查询到的Person
func printPerson(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Person
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Person)
		fmt.Printf("%#v\n", t)
	}
}

// 简单分页查询
func search(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("ustc").
		Type("person").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printPerson(res, err)
}
