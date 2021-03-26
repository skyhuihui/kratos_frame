package model

type InsertTestParams struct {
}

type FindTestParams struct {
	PageParams
}

type DeleteTestParams struct {
	TestId uint
}

type UpdateTestParams struct {
	TestId uint
}
