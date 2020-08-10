package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Student struct {
	studenid  string    `json:"studenid"`
	firstname string    `json:"firstname"`
	lastname  string    `json:"lastname"`
	age       string    `json:"age"`
	phone     string    `json:"phone"`
	city      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DATABASE INSTANCE
var stu_collection *mongo.Collection

func StudentCollection(c *mongo.Database) {
	stu_collection = c.Collection("student")
}

func GetAllStudents(c *gin.Context) {
	all_student := []Student{}
	cursor, err := stu_collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all student, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var student Student
		cursor.Decode(&student)
		all_student = append(all_student, student)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Students",
		"data":    all_student,
	})
	return
}

func CreateStudent(c *gin.Context) {
	var student Student
	c.BindJSON(&student)

	// id := guuid.New().String()

	newStudent := Student{
		studenid:  student.studenid,
		firstname: student.firstname,
		lastname:  student.lastname,
		age:       student.age,
		phone:     student.phone,
		city:      student.city,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := stu_collection.InsertOne(context.TODO(), newStudent)

	if err != nil {
		log.Printf("Error while inserting new student into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Student created Successfully",
	})
	return
}

func GetSingleStudent(c *gin.Context) {
	studentid := c.Param("studentid")

	student := Student{}
	err := stu_collection.FindOne(context.TODO(), bson.M{"studentid": studentid}).Decode(&student)
	if err != nil {
		log.Printf("Error while getting a single student, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Student",
		"data":    student,
	})
	return
}

func EditStudent(c *gin.Context) {
	studentid := c.Param("studentid")
	var student Student
	c.BindJSON(&student)

	newData := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := stu_collection.UpdateOne(context.TODO(), bson.M{"studentid": studentid}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Student Edited Successfully",
	})
	return
}

func DeleteStudent(c *gin.Context) {
	studentid := c.Param("studentid")

	_, err := stu_collection.DeleteOne(context.TODO(), bson.M{"studentid": studentid})
	if err != nil {
		log.Printf("Error while deleting a single student, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Student deleted successfully",
	})
	return
}
