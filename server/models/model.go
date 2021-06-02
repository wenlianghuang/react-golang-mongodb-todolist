package models

/*type Subtask struct {
	Subtask string `json:"subtask"`
}*/

type SubNumber struct {
	Numberone   int    `json:"numberone"`
	Numbertwo   int    `json:"numbertwo"`
	Numberthree int    `json:"numberthree"`
	Substring   string `json:"substring"`
}
type ToDoList struct {
	//ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Test   string `json:"task,omitempty"`
	Status bool   `json:"status,omitempty"`
	TTT    string `json:"task2,omitempty"`
	Number string `json:"number,omitempty"`
	//SubTest []SubNumber `json:"subtest,omitempty"`

	//MyNumber SubNumber `json:"mynumber,omitempty"`
	//Subtask Subtask `json:"task,omitempty"`
	/*Task struct {
		Name string `json:"task,omitempty"`
	}*/
}
