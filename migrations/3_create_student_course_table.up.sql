
CREATE TABLE StudentCourse
(
    id serial PRIMARY KEY,
    studentId int,
    courseId int,
    CONSTRAINT fk_studentId
      FOREIGN KEY(studentId) 
	  REFERENCES Student(id),
    CONSTRAINT fk_courseId
      FOREIGN KEY(courseId) 
	  REFERENCES Course(id)
);