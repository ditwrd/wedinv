package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1727635671")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(0, []byte(`{
			"autogeneratePattern": "[A-Za-z0-9]{4}",
			"hidden": false,
			"id": "text3208210256",
			"max": 4,
			"min": 4,
			"name": "id",
			"pattern": "^[A-Za-z0-9]+$",
			"presentable": false,
			"primaryKey": true,
			"required": true,
			"system": true,
			"type": "text"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1727635671")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(0, []byte(`{
			"autogeneratePattern": "[A-za-z0-9]{11}",
			"hidden": false,
			"id": "text3208210256",
			"max": 11,
			"min": 11,
			"name": "id",
			"pattern": "^[A-za-z0-9]+$",
			"presentable": false,
			"primaryKey": true,
			"required": true,
			"system": true,
			"type": "text"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
