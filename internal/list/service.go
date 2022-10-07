package list

func ConvertCreateToList(cw CreateList) List {
	return List{
		Title: cw.Title,
	}
}
