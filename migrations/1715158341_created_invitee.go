package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "xuiagwc0s3jkqd4",
			"created": "2024-05-08 08:52:21.462Z",
			"updated": "2024-05-08 08:52:21.462Z",
			"name": "invitee",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "kf7h2tq6",
					"name": "name",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "epdheg19",
					"name": "status",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"waiting",
							"accepted",
							"declined"
						]
					}
				},
				{
					"system": false,
					"id": "r9fzhzsl",
					"name": "invited_by",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"groom_bride",
							"groom",
							"bride"
						]
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `name_idx` + "`" + ` ON ` + "`" + `invitee` + "`" + ` (` + "`" + `name` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("xuiagwc0s3jkqd4")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
