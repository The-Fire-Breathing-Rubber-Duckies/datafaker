package pkg

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	dbI "github.com/the-fire-breathing-duckies/datafaker/pkg/db"
	"github.com/the-fire-breathing-duckies/datafaker/pkg/entities"
	"gopkg.in/yaml.v2"
)

type Cookery struct {
	Name    string   `yaml:"name"`
	Recipes []Recipe `yaml:"recipes"`
}

type Recipe struct {
	Name               string             `yaml:"name"`
	Quantity           int                `yaml:"quantity"`
	QuantityByRelation QuantityByRelation `yaml:"quantityByRelation,omitempty"`
	Table              string             `yaml:"table"`
	Fields             []Field            `yaml:"fields"`
}

type QuantityByRelation struct {
	Field    string `yaml:"field"`
	quantity int    `yaml:"quantity"`
}

type Field struct {
	Name          string        `yaml:"name"`
	Type          string        `yaml:"type"`
	Index         string        `yaml:"index,omitempty"`
	AutoIncrement bool          `yaml:"autoincrement,omitempty"`
	Entity        string        `yaml:"entity,omitempty"`
	EntityOpts    interface{}   `yaml:"entityOpts,omitempty"`
	Relation      FieldRelation `yaml:"relation,omitempty"`
}

type FieldRelation struct {
	Table string `yaml:"table"`
	Field string `yaml:"field"`
}

func ParseCookery(dbHandle *sql.DB, file string) (cookery Cookery) {
	cookery = openFile(file)

	runCookery(cookery, dbHandle)

	return cookery
}

func openFile(file string) (cookery Cookery) {
	// Open file
	yamlFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our yamlFile so that we can parse it later on
	defer yamlFile.Close()

	data, _ := ioutil.ReadAll(yamlFile)

	if err := yaml.Unmarshal(data, &cookery); err != nil {
		log.Fatalf("error: %v", err)
	}

	return cookery
}

func runCookery(cookery Cookery, dbHandle *sql.DB) {
	fmt.Printf("Cookery: %s\n", cookery.Name)
	fmt.Printf("%s\n", strings.Repeat("=", len(cookery.Name)))

	for _, recipe := range cookery.Recipes {
		runRecipe(recipe, dbHandle)
	}
}

func runRecipe(recipe Recipe, dbHandle *sql.DB) (fields map[string]interface{}) {
	fmt.Printf("\n\nRecipe: %s\n", recipe.Name)
	fmt.Printf("%s\n", strings.Repeat("-", len(recipe.Name)))

	for i := 0; i < recipe.Quantity; i++ {
		fields = make(map[string]interface{})
		for _, field := range recipe.Fields {
			// Skip fields with no name
			if field.Name == "" {
				continue
			}

			if result := processField(field); result != nil {
				fmt.Printf("%v %v\n", field.Name, result)
				fields[field.Name] = result
			}
		}

		// Insert
		dbI.Insert(dbHandle, recipe.Table, fields)
	}

	return fields
}

func processField(field Field) interface{} {
	if field.AutoIncrement {
		return nil
	}

	result := processEntity(field.Entity, field.EntityOpts)
	if result == nil {
		result = processGeneric(field.Type)
	}

	return result
}

func processEntity(entity string, opts interface{}) interface{} {
	switch entity {
	case "name":
		return entities.GetNames(1)[0]
	case "email":
		domain := ""
		if v := reflect.ValueOf(opts).MapIndex(reflect.ValueOf("domain")); v.IsValid() {
			domain = v.Interface().(string)
		}
		return entities.GetEmails(1, domain, "")[0]
	case "sentence":
		return gofakeit.HipsterSentence(10)
	case "paragraph":
		return gofakeit.HipsterParagraph(10, 10, 10, " ")
	default:
		return nil
	}
}

func processGeneric(t string) interface{} {
	switch t {
	case "text":
		return gofakeit.Word()
	case "int":
		return gofakeit.Int16()
	default:
		return nil
	}
}
