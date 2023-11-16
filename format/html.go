package format

import (
	"bytes"                           //
	"errors"                          //
	"github..com/x86kernel/htmlcolor" // Html syntax highlighter
	"github.com/PuerkitoBio/goquery"  // Html parser
	"io"                              //
)

// HtmlFormatter is a formatter for html markup language
type HtmlFormatter struct {
	parsedBody goquery.Document
	TextFormatter
}

// Format formats the html body
func (f *HtmlFormatter) Format(writer io.Writer, data []byte) error {
	htmlFormatter := htmlcolor.NewFormatter()
	buf := bytes.NewBuffer(make([]byte, 0, len(data)))
	err := htmlFormatter.Format(buf, data)

	if err == io.EOF {
		writer.Write(buf.Bytes())
		return nil
	}

	return errors.New("html formating error")
}

// Title returns the title of the formatter
func (f *HtmlFormatter) Title() string {
	return "[html]"
}

// Search searches for the query in the html body
func (f *HtmlFormatter) Search(q string, body []byte) ([]string, error) {
	if q == "" {
		buf := bytes.NewBuffer(make([]byte, 0, len(body)))
		err := f.Format(buf, body)
		return []string{buf.String()}, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	results := make([]string, 0, 8)
	doc.Find(q).Each(func(_ int, s *goquery.Selection) {
		htmlResult, err := goquery.OuterHtml(s)
		if err == nil {
			results = append(results, htmlResult)
		}
	})

	return results, nil
}
