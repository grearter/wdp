package food

func (food *Food) Add(foodModel *Food) error {
	c := food.GetC()
	defer c.Database.Session.Close()

	return c.Insert(foodModel)
}
