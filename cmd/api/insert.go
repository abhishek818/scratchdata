package api

import (
	"errors"
	"scratchdata/models"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

func (a *API) Insert(c *fiber.Ctx) error {
	if c.QueryBool("debug", false) {
		rid := ulid.Make().String()
		log.Debug().
			Str("request_id", rid).
			Interface("headers", c.GetReqHeaders()).
			Str("body", string(c.Body())).
			Interface("queryParams", c.Queries()).
			Msg("Incoming request")
	}

	// TODO: this block can be abstracted as we also use it for query
	/////
	apiKey := c.Locals("apiKey").(models.APIKey)

	// TODO: read-only vs read-write connections
	connectionSetting := a.db.GetDatabaseConnection(apiKey.DestinationID)

	if connectionSetting.ID == "" {
		return errors.New("No DB Connections set up")
	}
	/////

	input := c.Body()
	if !gjson.ValidBytes(input) {
		return errors.New("Invalid JSON")
	}

	// Get table (from header, query, or body)
	// Get flatten type
	// Get data location (body, data key)
	// Get rows if array
	// Flatten per algorithm
	// Create rowid and write
	a.dataTransport.Write(connectionSetting.ID, []byte{})

	// table_name, table_location := i.getField("X-SCRATCHDB-TABLE", "table", "table", c)
	// if table_name == "" {
	// 	return fiber.NewError(fiber.StatusBadRequest, "You must specify a table name")
	// }

	// flattenAlgorithm, _ := i.getField("X-SCRATCHDB-FLATTEN", "flatten", "flatten", c)

	// data_path := "$"
	// if table_location == "body" {
	// 	data_path = "$.data"
	// }

	// root, err := ajson.Unmarshal(input)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, err.Error())
	// }

	// x, err := root.JSONPath(data_path)
	// if err != nil {
	// 	return err
	// }

	// dir := filepath.Join(i.Config.Ingest.DataDir, api_key, table_name)

	// // TODO: make sure this is atomic!
	// writer, ok := i.writers[dir]
	// if !ok {
	// 	writer = NewFileWriter(
	// 		dir,
	// 		i.Config,
	// 		filepath.Join("data", api_key, table_name),
	// 		api_key, table_name,
	// 	)
	// 	i.writers[dir] = writer
	// }

	// if x[0].Type() == ajson.Array {
	// 	objects, err := x[0].GetArray()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	for _, o := range objects {

	// 		if flattenAlgorithm == "explode" {
	// 			flats, err := FlattenJSON(o.String(), nil, false)
	// 			if err != nil {
	// 				return err
	// 			}

	// 			for _, flat := range flats {
	// 				err = writer.Write(flat)
	// 				if err != nil {
	// 					log.Error().Err(err).Str("flat", flat).Msg("Unable to write object")
	// 				}

	// 			}

	// 		} else {
	// 			flat, err := flatten.FlattenString(o.String(), "", flatten.UnderscoreStyle)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			err = writer.Write(flat)
	// 			if err != nil {
	// 				log.Error().Err(err).Str("flat", flat).Msg("Unable to write object")
	// 			}
	// 		}
	// 	}

	// } else if x[0].Type() == ajson.Object {
	// 	if flattenAlgorithm == "explode" {
	// 		flats, err := FlattenJSON(x[0].String(), nil, false)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		for _, flat := range flats {
	// 			err = writer.Write(flat)
	// 			if err != nil {
	// 				log.Error().Err(err).Str("flat", flat).Msg("Unable to write object")
	// 			}

	// 		}

	// 	} else {
	// 		flat, err := flatten.FlattenString(x[0].String(), "", flatten.UnderscoreStyle)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		err = writer.Write(flat)
	// 		if err != nil {
	// 			return fiber.NewError(fiber.StatusBadRequest, err.Error())
	// 		}
	// 	}
	// }

	return c.SendString("ok")
}
