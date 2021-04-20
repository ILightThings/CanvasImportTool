# Canvas LMS User Import Tool

This tool was designed to make it easier to import users into courses on the Canvas LMS Platform.



The template CSV has 4 columns.

| course_id | user_id | role    | status |
| --------- | ------- | ------- | ------ |
| CourseA   | UserABC | student | active |
| CourseB   | UserDEF | faculty | active |
| CourseC   |         |         |        |
| CourseD   |         |         |        |



The program will combine column (user_id,role,status) and apply them to every row in the course_id column.



A new file named `importme.csv` will have the following result.

| course_id | user_id | role    | status |
| --------- | ------- | ------- | ------ |
| CourseA   | UserABC | student | active |
| CourseB   | UserABC | student | active |
| CourseC   | UserABC | student | active |
| CourseD   | UserABC | student | active |
| CourseA   | UserDEF | faculty | active |
| CourseB   | UserDEF | faculty | active |
| CourseC   | UserDEF | faculty | active |
| CourseD   | UserDEF | faculty | active |



You can have more students then courses or more courses then students.



## Usage

Have a template file of your students and courses in a csv file name `users.csv` with the following format:

| course_id | user_id | role | status |
| --------- | ------- | ---- | ------ |
|           |         |      |        |



Run the CanvasImportTool 

Upload the newly created importme.csv file.



## Note

Please ensure that none of the data entered have any commas as it is used as the delimiter.



## Build

```go
go build CanvasImportTool.go
```