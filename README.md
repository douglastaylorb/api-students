# api-students
API to manage 'Curso golang' course students

Routes
- GET /students - List all students
- POST /students - Create student
- GET /students/:id - Get infos from a specific student
- PUT /students/:id - Update student
- DELETE /students/:id - Delete Student
- GET /students?active=<true/false> - List all active/non-active students

Struct Student:
- Name
- CPF
- Email
- Age
- Active