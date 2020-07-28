# Mero Edu 

## Running development
`make run-dev`

## Stoping development
`make stop-dev`

## Code Folder Structure

```
domain // Entity
├── course.go
├── category.go
└── author.go 

course
├── delivery
│   └── http
│       ├── course_handler.go
│       └── course_test.go
├── mocks
│   ├── courseRepository.go
│   └── courseUsecase.go
├── repository // implementation
│   ├── mysql_course.go
│   └── mysql_course_test.go
├── repository.go 
├── usecase // Implementation
│   ├── courseu_usecase_test.go
│   └── course_usecase.go
└── usecase.go // Usecase Interface.
```