package utils

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func ReadNodeByFile(path string) *html.Node {
	file, e := os.Open(path)
	defer file.Close()
	if e != nil {
		fmt.Println("html文件打开失败", e)
	}
	node, err := html.Parse(file)
	if err != nil {
		fmt.Println("读取html内容失败", err)
	}
	return node
}

func ReadNodeByHttp(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取http连接失败", resp)
	}
	defer resp.Body.Close()
	node, err1 := html.Parse(resp.Body)
	if err1 != nil {
		fmt.Println("读取html内容失败", err)
	}
	return node
}
func FindNodeByAtrr(atrrName, attrValue string, node *html.Node) *html.Node {
	if node == nil {
		return nil
	}
	attributes := node.Attr
	for _, attr := range attributes {
		if attr.Key == atrrName && attr.Val == attrValue {
			return node
		}
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if n.Type != html.ElementNode {
			continue
		}
		destNode := FindNodeByAtrr(atrrName, attrValue, n)
		if destNode != nil {
			return destNode
		}
	}
	return nil
}

func recursion(node *html.Node, Data string) *html.Node {
	//fmt.Println(node.Data,",",node.Data == "body")
	if node == nil {
		return nil
	}
	if node.Data == Data {
		return node
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		recursion := recursion(c, Data)
		if recursion != nil {
			return recursion
		}
	}
	return nil
}
func GetHeadHtml(node *html.Node) *html.Node {
	return node.FirstChild.NextSibling
}

func GetHeadNode(node *html.Node) *html.Node {
	return recursion(node, "head")
}
func GetBodyNode(node *html.Node) *html.Node {
	return recursion(node, "body")
}

type Parse interface {
	HtmlParse()
}
