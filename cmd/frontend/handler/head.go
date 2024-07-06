package handler

import (
	"errors"
	"io"

	"github.com/PuerkitoBio/goquery"
)

/*
Merges the <head> found inside <body> into document <head>.
  - document is parsed from r using goquery and the result is written to w
*/
func mergeBodyHeadTags(r io.Reader, w io.Writer) error {
	headTags := []string{"meta", "link", "title", "style"}

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return errors.Join(errors.New("goquery.NewDocumentFromReader"), err)
	}

	var headTagNodes []*goquery.Selection
	for _, tag := range headTags {
		doc.Find("body").Find(tag).Each(func(i int, s *goquery.Selection) {
			headTagNodes = append(headTagNodes, s.Remove())
		})
	}

	for _, bodyNode := range headTagNodes {
		name := bodyNode.Get(0).Data
		head := doc.Find("head")
		if head.Length() == 0 {
			doc.Add("head").AppendSelection(bodyNode)
			continue
		}

		currentHeadTags := head.Find(name)
		currentHeadTags.Each(func(i int, s *goquery.Selection) {
			if shouldMergeTags(s, bodyNode) {
				s.Remove()
				return
			}
		})
		doc.Find("head").AppendSelection(bodyNode)
	}

	rendered, err := doc.Html()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(rendered))
	return err
}

func shouldMergeTags(head, body *goquery.Selection) bool {
	return body.Get(0).Data == "title" || (body.Get(0).Data == "meta" && (head.AttrOr("property", "#h") == body.AttrOr("property", "#b") || head.AttrOr("name", "#h") == body.AttrOr("name", "#b")))
}
