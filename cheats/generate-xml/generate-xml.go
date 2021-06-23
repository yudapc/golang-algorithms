package generatexml

type Document struct {
	Title   string
	URL     string
	Content struct {
		Articles []struct {
			Title      string
			URL        string
			Categories []string
			Info       string
		}
	}
}
