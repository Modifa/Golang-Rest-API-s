package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"whatapp/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

//
/*DB ... */
type DB struct {
}

//
func (db *DB) SaveOnDB(functionnamewithschema string, m interface{}) ([]models.UserRegisted, error) {
	userReg := []models.UserRegisted{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("AuthDBNURL"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

//ReturnList
func (db *DB) ReturnContactList(functionnamewithschema string, m interface{}) ([]models.ListContacts, error) {
	ListContacts := []models.ListContacts{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("AuthDBNURL"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &ListContacts, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return ListContacts, nil
}

func (db *DB) ReturnUser(functionnamewithschema string, m interface{}) ([]models.LoggedUser, error) {
	LoggedUser := []models.LoggedUser{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("AuthDBNURL"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &LoggedUser, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return LoggedUser, nil
}

//
func (db *DB) ReturnGroups(functionnamewithschema string, m interface{}) ([]models.ListGroups, error) {
	LoggedUser := []models.ListGroups{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("AuthDBNURL"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &LoggedUser, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return LoggedUser, nil
}

//Convert Interface and return Query string
func ConVertInterface(funcstr string, m interface{}) string {
	q := "select * from " + funcstr + "("

	if m != nil {

		v := reflect.ValueOf(m)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {

			switch typeOfS.Field(i).Type.Name() {
			case "int", "int16", "int32", "int64", "int8":
				str := v.Field(i).Interface().(int64)
				strInt64 := strconv.FormatInt(str, 10)
				q += strInt64 + ","
			case "float64":
				str := v.Field(i).Interface().(float64)
				s := fmt.Sprintf("%f", str)
				q += s + ","
			case "bool":
				q += "'" + strconv.FormatBool(v.Field(i).Interface().(bool)) + "',"
			default:
				if v.Field(i).Interface().(string) == "" {
					q += "null,"
				} else {
					q += "'" + v.Field(i).Interface().(string) + "',"
				}
			}
		}

		q = q[0 : len(q)-len(",")]
	}

	q += ")"

	return q
}
