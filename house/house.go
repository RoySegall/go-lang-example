package house

// this is a publis function because it's start with an uppercase.
func Address() string {
	return "Ramat gan"
}

// Not accesable out side.
func size() string {
	return "4 rooms"
}

func GetSize() string {
	return size();
}
