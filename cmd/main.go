package main

import (
	"fmt"
	"log"
	"simple-clinic/config"
	"simple-clinic/controller"
	"simple-clinic/models"
	"simple-clinic/storage/jsondb"
)

func main() {

	var cfg = config.Load()

	jsondb, err := jsondb.NewJsonDbConnection(&cfg)
	if err != nil {
		panic(err)
	}

	con := controller.NewController(&cfg, jsondb)

	// Clinic(con)
	Branch(con)
}

func Branch(con *controller.Conn) {

	resp, err := con.CreateBranch(&models.CreateBranch{
		ClinicID: "90384823904",
		Address:  "Tashkent",
		Phone:    "883-349-435-45",
	})

	if err != nil {
		log.Println("Create branch:", err.Error())
		return
	}

	fmt.Println(resp)
}

func Clinic(con *controller.Conn) {

	// Create
	// {
	// 	for i := 0; i < 100; i++ {
	// 		resp, err := con.CreateClinic(
	// 			&models.CreateClinic{
	// 				Name: faker.FirstName(),
	// 				Logo: faker.LastName(),
	// 				City: faker.DomainName(),
	// 			},
	// 		)
	// 		if err != nil {
	// 			log.Println("Create clinic:", err.Error())
	// 			return
	// 		}

	// 		fmt.Println(resp)
	// 	}
	// }

	// GetByID
	// {
	// 	resp, err := con.GetByIDClinic(&models.ClinicPrimaryKey{ID: "bd82bb75-ce28-479d-88a8-491c50f5df5e"})
	// 	if err != nil {
	// 		log.Println(config.Error, "GetByID clinic:", err.Error())
	// 		return
	// 	}

	// 	fmt.Println(resp)
	// }

	// GetList
	{
		// resp, err := con.GetListClinic(&models.GetListClinicRequest{
		// 	Offset: 0,
		// 	Limit:  10,
		// })
		// if err != nil {
		// 	log.Println(config.Error, "GetByID clinic:", err.Error())
		// 	return
		// }

		// fmt.Println("Clinic count:", resp.Count)
		// for ind, clinic := range resp.Clinics {
		// 	fmt.Println(ind+1, clinic)
		// }
	}

	// Update
	// {
	// 	resp, err := con.UpdateClinic(
	// 		&models.UpdateClinic{
	// 			ClinicID: "bd82bb75-ce28-479d-88a8-491c50f5df5e",
	// 			Name:     "Nuravshon",
	// 			Logo:     "logo",
	// 			City:     "Jizzah",
	// 		},
	// 	)
	// 	if err != nil {
	// 		log.Println(config.Error, "Update clinic:", err.Error())
	// 		return
	// 	}

	// 	fmt.Println(resp)
	// }

	// Delete
	{
		// err := con.DeleteClinic(&models.ClinicPrimaryKey{ID: "bd82bb75-ce28-479d-88a8-491c50f5df5e"})
		// if err != nil {
		// 	log.Println(config.Error, "Delete clinic:", err.Error())
		// 	return
		// }
	}
}
