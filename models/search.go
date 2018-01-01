package models

func GetSearchRresult(serarchKey string) (index *Index, err error) {
	index = new(Index)
	index.Title = `搜索结果`
	IndexItem, err := getSearchRresult(serarchKey)
	if err != nil {
		return index, err
	}
	index.IndexItems = append(index.IndexItems, *IndexItem)
	return
}
